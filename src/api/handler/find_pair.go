package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	uc "pairs/src/api/usecases"
	entity "pairs/src/entities"
	lib "pairs/src/libs"
)

// FindPairs handler of the find-pair
func FindPairs(res http.ResponseWriter, req *http.Request) {

	appReq := lib.NewRequest(res, req)
	log.Println("API: Get find pair....")

	// Request parse
	var payload entity.Payload
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		appReq.AddError(http.StatusBadRequest, "internal", "body", "Invalid request!")
		return
	}

	// NewFindPair usecase return target indices
	uc := uc.NewFindPair(appReq, payload)
	uc.GetPairs()
	if appReq.HasErrors() {
		appReq.WriteErrorResponse()
		return
	}
	appReq.WriteSuccessResponse()

}
