package urlcontroller

import (
	"fmt"
	"net/http"

	"github.com/codeshaine/url-shortner/db"
	"github.com/codeshaine/url-shortner/internal/response"
	"github.com/codeshaine/url-shortner/internal/utils"
	"github.com/go-chi/chi/v5"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	urlToShorten := r.URL.Query().Get("url")
	shortUrl := utils.GenerateUnqueUrl()
	urlData, err := db.InsertUrl(urlToShorten, shortUrl)

	if err != nil {
		res := response.ErrorResponse(fmt.Sprintf("db error: %v", err))
		response.Json(w, http.StatusBadRequest, res)
		return
	}
	res := response.SuccessResponse(urlData)
	response.Json(w, http.StatusCreated, res)
}

func HanldeRedirect(w http.ResponseWriter, r *http.Request) {
	urlToRedirect := chi.URLParam(r, "url")
	urlData, err := db.GetLongUrl(urlToRedirect)
	if err != nil {
		res := response.ErrorResponse(fmt.Sprintf("db error: %v", err))
		response.Json(w, http.StatusBadRequest, res)
		return
	}
	if urlData.ClickCount > 1000 { //not good either
		db.Db.Exec("DELETE FROM urls WHERE short_url=$1", urlData.ShortUrl) //not good
		resData := response.ErrorResponse("Error:click count excceded")
		response.Json(w, http.StatusBadRequest, resData)
		return
	}
	resData := response.SuccessResponse(urlData)
	response.Json(w, http.StatusOK, resData)
}
