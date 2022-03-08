package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	models "ydgo/src/models"
)

type QuestionService struct {
	QuestionRequest models.QuestionRequest
}

type jsonResult struct {
	List []jsonVal `json:"list"`
}

type jsonVal struct {
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
}

type unRelevantJson struct {
	List []unRelevantJsonVal `json:"list"`
}

type unRelevantJsonVal struct {
	Id       int    `json:"id"`
	Response string `json:"response"`
}

func (qS QuestionService) FindAnswer() models.QuestionResponse {
	// Open our jsonFile
	jsonFile, err := os.Open("src/relevant.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result jsonResult
	json.Unmarshal([]byte(byteValue), &result)
	var relevantResponse string = ""
	for _, s := range result.List {
		if strings.Contains(strings.ToLower(s.Question), strings.ToLower(qS.QuestionRequest.Question)) || strings.Contains(strings.ToLower(qS.QuestionRequest.Question), strings.ToLower(s.Question)) {
			relevantResponse = s.Answer
		}
	}

	unRelevantResponse := ""
	if relevantResponse == "" {
		jsonFile, err := os.Open("src/unrelevant.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var resultNew unRelevantJson
		json.Unmarshal([]byte(byteValue), &resultNew)
		var randInt = rand.Intn(22)
		for _, s := range resultNew.List {
			if s.Id == randInt {
				unRelevantResponse += s.Response
			}
		}
	}

	return models.QuestionResponse{Answer: relevantResponse, UnrelevantAnswer: unRelevantResponse}
}
