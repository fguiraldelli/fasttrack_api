package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// question represents data about a question and its answer
type question struct {
	ID             string   `json:"id"`
	Question       string   `json:"question"`
	Answers        []string `json:"answers"`
	Correct_answer string   `json:"correct_answer"`
	Answered       bool     `json:"answered"`
	Is_corrected   bool     `json:"is_corrected"`
}

var questions = []question{
	{ID: "1", Question: "Question 1", Answers: []string{"A1", "A2", "A3", "A4", "A5"}, Correct_answer: "A4", Answered: false, Is_corrected: false},
	{ID: "2", Question: "Question 2", Answers: []string{"B1", "B2", "B3", "B4", "B5"}, Correct_answer: "B1", Answered: false, Is_corrected: false},
	{ID: "3", Question: "Question 3", Answers: []string{"C1", "C2", "C3", "C4", "C5"}, Correct_answer: "C5", Answered: false, Is_corrected: false},
	{ID: "4", Question: "Question 4", Answers: []string{"D1", "D2", "D3", "D4", "D5"}, Correct_answer: "D2", Answered: false, Is_corrected: false},
	{ID: "5", Question: "Question 5", Answers: []string{"E1", "E2", "E3", "E4", "E5"}, Correct_answer: "E3", Answered: false, Is_corrected: false},
}

//getQuestions responds with the list of all questions as JSON.
func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.Run("localhost:8080")
}
