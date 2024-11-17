package hello

import (
	"backend/src/utils"
	"net/http"
)

func GetHelloController(w http.ResponseWriter, r *http.Request){
	response := GetHelloService()
	utils.RespondWithJSON(w, http.StatusOK, response)
}