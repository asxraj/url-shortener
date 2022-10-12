package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/asxraj/url-shortener/internal/database"
	"github.com/asxraj/url-shortener/internal/models"
	"github.com/asxraj/url-shortener/internal/validator"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func (app *application) resolveURL(w http.ResponseWriter, r *http.Request) {
	url := httprouter.ParamsFromContext(r.Context()).ByName("url")

	rdb := database.CreateClient()
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		app.errorResponseJSON(w, r, http.StatusNotFound, map[string]any{"error": "short url not found in the database"})
		return
	} else if err != nil {
		app.errorResponseJSON(w, r, http.StatusInternalServerError, "cannot connect to DB")
		return
	}

	err = app.models.Users.ClickURL(url, r.RemoteAddr)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusMovedPermanently, map[string]any{
		"to": "https://" + val,
	}, nil)
}

// try out giving localhost as name to see if the loop will happen
func (app *application) shortenURL(w http.ResponseWriter, r *http.Request) {

	input := &request{}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponseJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	var identifier string
	var api_quota int

	user := app.contextGetUser(r)
	if user.ID != 0 {
		identifier = fmt.Sprint(user.ID)
		api_quota = 20 // Could set this from database
	} else {
		identifier = r.RemoteAddr
		api_quota, err = strconv.Atoi(os.Getenv("API_QUOTA"))
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	u, err := url.Parse(input.URL)
	if err != nil {
		panic(err)
	}
	v := validator.New()

	if u.Scheme == "" {
		u.Host = input.URL
	} else if u.Scheme == "http" {
		v.AddError("url", "http is not allowed")
	}
	input.URL = u.Host

	models.ValidateUrl(v, input.URL)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// enforcing https
	input.URL = u.Host

	if input.CustomShort == "" {
		input.CustomShort = uuid.New().String()[:6]
	}
	if input.Expiry == 0 {
		input.Expiry = 24
	}

	rdb := database.CreateClient()
	defer rdb.Close()

	// Checking if custom alias is already in use
	val, _ := rdb.Get(database.Ctx, input.CustomShort).Result()
	if val != "" {
		app.errorResponseJSON(w, r, http.StatusForbidden, "custom short URL is already in use")
		return
	}

	// Checking the ratelimiting (IP address basis for non-users and ID for users)
	val, err = rdb.Get(database.Ctx, identifier).Result()
	if err == redis.Nil {
		_ = rdb.Set(database.Ctx, identifier, api_quota, 30*time.Minute)
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			ttl, _ := rdb.TTL(database.Ctx, identifier).Result()
			app.errorResponseJSON(w, r, http.StatusServiceUnavailable, fmt.Sprintf("Rate limit exceeded, try again in %d", ttl/time.Nanosecond/time.Minute))
			return
		}
	}

	// Make this a transaction with psql  ??? perhaps with the help of passing the context
	err = rdb.Set(database.Ctx, input.CustomShort, input.URL, input.Expiry*time.Hour).Err()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if user.ID != 0 {
		err = app.models.Users.SaveURL(user, input.URL, input.CustomShort, time.Now().Add(input.Expiry*time.Hour))
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	resp := response{URL: input.URL, Expiry: input.Expiry, XRateLimitReset: 30}

	rdb.Decr(database.Ctx, identifier)

	val, _ = rdb.Get(database.Ctx, identifier).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := rdb.TTL(database.Ctx, identifier).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = fmt.Sprintf("%v/%v", os.Getenv("DOMAIN"), input.CustomShort)

	app.writeJSON(w, http.StatusOK, resp, nil)
}
