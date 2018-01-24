package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func startREST(addr string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/lookup/{lookupId}", LookupHandler)
	router.HandleFunc("/set", SetHandler)
	log.Fatal(http.ListenAndServe(addr, router))
}

// LookupHandler implements the REST GET handler
func LookupHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	lookupId, err := strconv.ParseInt(vars["lookupId"], 10, 64)

	if err != nil {
		response := &LookupResponse{
			Found:  false,
			Reason: err.Error(),
		}
		resp, _ := json.Marshal(response)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(resp)
		return
	}

	response := &LookupResponse{
		Entity: &Entity{Id: lookupId, Name: "Hello"},
		Found:  true,
	}

	resp, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// SetHandler implements the REST POST handler
func SetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	req := &SetRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		resp, _ := json.Marshal(&SetResponse{Success: false, Reason: err.Error()})
		w.Write(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validate(req) {
		resp, _ := json.Marshal(&SetResponse{Success: false, Reason: "Unable to validate request"})
		w.Write(resp)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(&SetResponse{Entity: req.Entity, Success: true})
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}
