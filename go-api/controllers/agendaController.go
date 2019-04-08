package controllers

import (
	"encoding/json"
	"go-api/models"
	"go-api/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetAgenda(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	directorios, err := models.GetAgenda()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, directorios)
}

func PostAgenda(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	body, _ := ioutil.ReadAll(r.Body)
	var dir models.Agenda
	err := json.Unmarshal(body, &dir)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewAgenda(dir)

	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{dir, http.StatusCreated})

}

func GetAgendaId(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	directorios, err := models.GetAgendaId(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, directorios)
}

func PutAgenda(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)
	body, _ := ioutil.ReadAll(r.Body)
	var directorios models.Agenda
	err := json.Unmarshal(body, &directorios)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	directorios.UID = uint32(uid)
	rows, err := models.UpdateAgenda(directorios)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}

func DeleteAgenda(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)
	_, err := models.DeleteAgenda(uint32(uid))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, http.StatusNoContent)
}
