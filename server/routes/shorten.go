package routes

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asxraj/url-shortener/database"
	"github.com/asxraj/url-shortener/helpers"
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
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	body := &request{}

	err := helpers.ReadJSON(w, r, &body)
	if err != nil {
		helpers.ErrorResponseJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	rdb2 := database.CreateClient(1)
	defer rdb2.Close()

	val, err := rdb2.Get(database.Ctx, r.RemoteAddr).Result()
	if err == redis.Nil {
		_ = rdb2.Set(database.Ctx, r.RemoteAddr, os.Getenv("API_QUOTA"), 30*time.Minute)
	} else {
		valInt, _ := strconv.Atoi(val)
		if valInt <= 0 {
			ttl, _ := rdb2.TTL(database.Ctx, r.RemoteAddr).Result()
			helpers.ErrorResponseJSON(w, r, http.StatusServiceUnavailable, fmt.Sprintf("Rate limit exceeded, try again in %d", ttl/time.Nanosecond/time.Minute))
		}
	}

	var id string

	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	rdb3 := database.CreateClient(3)
	defer rdb3.Close()

	val, _ = rdb3.Get(database.Ctx, id).Result()
	if val != "" {
		helpers.ErrorResponseJSON(w, r, http.StatusForbidden, "custom short URL is already in use")
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = rdb3.Set(database.Ctx, id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		helpers.ErrorResponseJSON(w, r, http.StatusInternalServerError, err.Error())

	}

	resp := response{URL: body.URL, Expiry: body.Expiry, XRateLimitReset: 30}

	rdb2.Decr(database.Ctx, r.RemoteAddr)

	val, _ = rdb2.Get(database.Ctx, r.RemoteAddr).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := rdb2.TTL(database.Ctx, r.RemoteAddr).Result()
	resp.XRateLimitReset = ttl / time.Nanosecond / time.Minute

	resp.CustomShort = fmt.Sprintf("%v/%v", os.Getenv("DOMAIN"), id)

	helpers.WriteJSON(w, http.StatusOK, resp, nil)
}
