package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	models "ydgo/src/models"
	service "ydgo/src/service"
)

type myHandler struct{}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":6868",
		Handler: &myHandler{},
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/question"] = gptQuestionAction
	fmt.Printf("...")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func gptQuestionAction(w http.ResponseWriter, r *http.Request) {
	r.Host = ""
	var questionRequest models.QuestionRequest
	err := json.NewDecoder(r.Body).Decode(&questionRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Message: err.Error()})
		return
	}
	questionService := service.QuestionService{questionRequest}
	questionResponse := questionService.FindAnswer()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(questionResponse)
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Host = ""
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	http.Error(w, "404 not found.", http.StatusNotFound)
	io.WriteString(w, "path: "+r.URL.String())
	return
}
