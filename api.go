package main

import (
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	// APIS
	router.HandleFunc("/alert", makeHTTPHandleFunc(s.handleSaveAlert)).Methods("POST")
	router.HandleFunc("/alert", makeHTTPHandleFunc(s.GetAlertByType)).Methods("GET")
	router.HandleFunc("/filterAlert", makeHTTPHandleFunc(s.GetAlertByTypeAndTimeRange)).Methods("GET")

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// POST API to save all details for alerts
func (s *APIServer) handleSaveAlert(w http.ResponseWriter, r *http.Request) error {
	req := new(Alert)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	alert, err := NewAlert(req.AlertType, req.Time, req.Description)
	if err != nil {
		return err
	}
	if err := s.store.SaveAlert(*alert); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, alert)
}

// To get alert information based on Alert type
func (s *APIServer) GetAlertByType(w http.ResponseWriter, r *http.Request) error {
	alertType := r.URL.Query().Get("alertType")
	pageNo, err := toInt64(r.URL.Query().Get("pageNo"))
	if err != nil {
		pageNo = 0
	}
	pageSize, err := toInt64(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize > 10 {
		pageSize = 10
	}
	alerts, err := s.store.GetAlertByType(alertType, pageNo, pageSize)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, alerts)
}

func toInt64(varStr string) (int64, error) {
	vartoInt64, err := strconv.ParseInt(varStr, 10, 64)
	if err != nil {
		return vartoInt64, fmt.Errorf("invalid ref given %s", varStr)
	}
	return vartoInt64, nil
}

// To get alert information based on Alert type and from specific time frame(epoc time format)

func (s *APIServer) GetAlertByTypeAndTimeRange(w http.ResponseWriter, r *http.Request) error {
	alertType := r.URL.Query().Get("alertType")
	fromTime, err := toInt64(r.URL.Query().Get("from"))
	if err != nil {
		currentTime := time.Now()
		oneWeekAgo := currentTime.AddDate(0, 0, -7)
		fromTime = oneWeekAgo.Unix()
	}
	toTime, err := toInt64(r.URL.Query().Get("to"))
	if err != nil {

		toTime = time.Now().Unix()
	}

	pageNo, err := toInt64(r.URL.Query().Get("pageNo"))
	if err != nil {
		pageNo = 0
	}
	pageSize, err := toInt64(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize > 10 {
		pageSize = 10
	}
	alerts, err := s.store.GetAlertByTypeAndInRange(alertType, fromTime, toTime, pageNo, pageSize)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, alerts)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, err.Error())
		}
	}
}
