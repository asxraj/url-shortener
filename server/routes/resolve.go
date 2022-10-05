package routes

import (
	"fmt"
	"net/http"

	"github.com/asxraj/url-shortener/database"
	"github.com/asxraj/url-shortener/helpers"
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
)

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	url := httprouter.ParamsFromContext(r.Context()).ByName("url")
	fmt.Println(url)

	rdb := database.CreateClient(0)
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		helpers.ErrorResponseJSON(w, r, http.StatusNotFound, "short url not found in the database")
	} else if err != nil {
		helpers.ErrorResponseJSON(w, r, http.StatusInternalServerError, "cannot connect to DB")
		return
	}

	http.Redirect(w, r, fmt.Sprintf("https://%v", val), http.StatusSeeOther)
}
