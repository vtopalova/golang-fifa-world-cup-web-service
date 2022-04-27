package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)

}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	year := req.URL.Query().Get("year")
	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
		res.Write(winners)
	} else {
		winnerYear, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
		}
		res.Write(winnerYear)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	if !data.IsAccessTokenValid(accessToken) {
		res.WriteHeader(http.StatusUnauthorized)
	} else {
		err := data.AddNewWinner(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		res.WriteHeader(http.StatusCreated)
	}

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ListWinners(res, req)
		break
	case http.MethodPost:
		AddNewWinner(res, req)
		break
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
