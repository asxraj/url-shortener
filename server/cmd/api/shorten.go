package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asxraj/url-shortener/internal/database"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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

// try out giving localhost as name to see if the loop will happen
func (app *application) shortenURL(w http.ResponseWriter, r *http.Request) {
	body := &request{}

	err := app.readJSON(w, r, &body)
	if err != nil {
		app.errorResponseJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	rdb := database.CreateClient()
	defer rdb.Close()

	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	val, _ := rdb.Get(database.Ctx, id).Result()
	if val != "" {
		app.errorResponseJSON(w, r, http.StatusForbidden, "custom short URL is already in use")
		return
	}

	val, err = rdb.Get(database.Ctx, r.RemoteAddr).Result()
	if err == redis.Nil {
		_ = rdb.Set(database.Ctx, r.RemoteAddr, os.Getenv("API_QUOTA"), 30*time.Minute)
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			ttl, _ := rdb.TTL(database.Ctx, r.RemoteAddr).Result()
			app.errorResponseJSON(w, r, http.StatusServiceUnavailable, fmt.Sprintf("Rate limit exceeded, try again in %d", ttl/time.Nanosecond/time.Minute))
		}
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = rdb.Set(database.Ctx, id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		app.errorResponseJSON(w, r, http.StatusInternalServerError, err.Error())
	}

	resp := response{URL: body.URL, Expiry: body.Expiry, XRateLimitReset: 30}

	rdb.Decr(database.Ctx, r.RemoteAddr)

	val, _ = rdb.Get(database.Ctx, r.RemoteAddr).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := rdb.TTL(database.Ctx, r.RemoteAddr).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = fmt.Sprintf("%v/%v", os.Getenv("DOMAIN"), id)

	app.writeJSON(w, http.StatusOK, resp, nil)
}
