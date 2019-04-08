package controllers

import (
	"go-api/utils"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, struct {
		Message string `json:"message"`
	}{
		Message: "Go RestFul Api",
	})
}
