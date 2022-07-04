package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user represents data about the users and theirs respective answers
type registred_user struct {
	ID                       string     `json:"id"`
	Name                     string     `json:"name"`
	Email                    string     `json:"email"`
	Quiz                     []question `json:"questions"`
	Number_corrected_answers int        `json:"number_corrected_answers"`
	User_rated               float64    `json:"user_rated"`
}

// question represents data about a question and its answer
type question struct {
	ID             string   `json:"id"`
	Question       string   `json:"question"`
	Answers        []string `json:"answers"`
	Correct_answer string   `json:"correct_answer"`
	Answered       bool     `json:"answered"`
	Is_corrected   bool     `json:"is_corrected"`
}

var users = []registred_user{
	{ID: "1", Name: "John Doe", Email: "doe.jonh@hotmail.com", Quiz: questions, Number_corrected_answers: 0, User_rated: 0.00},
	{ID: "2", Name: "Jane Doe", Email: "janedoe1989@gmail.com", Quiz: questions, Number_corrected_answers: 0, User_rated: 0.00},
}

var questions = []question{
	{ID: "1", Question: "Question 1", Answers: []string{"A1", "A2", "A3", "A4", "A5"}, Correct_answer: "A4", Answered: false, Is_corrected: false},
	{ID: "2", Question: "Question 2", Answers: []string{"B1", "B2", "B3", "B4", "B5"}, Correct_answer: "B1", Answered: false, Is_corrected: false},
	{ID: "3", Question: "Question 3", Answers: []string{"C1", "C2", "C3", "C4", "C5"}, Correct_answer: "C5", Answered: false, Is_corrected: false},
	{ID: "4", Question: "Question 4", Answers: []string{"D1", "D2", "D3", "D4", "D5"}, Correct_answer: "D2", Answered: false, Is_corrected: false},
	{ID: "5", Question: "Question 5", Answers: []string{"E1", "E2", "E3", "E4", "E5"}, Correct_answer: "E3", Answered: false, Is_corrected: false},
}

func verifyEmail(new_email string, existed_email string) error {
	if new_email == existed_email {
		return errors.New("this e-mail already exists")
	}
	return nil
}

//getQuestions responds with the list of all questions as JSON.
func getQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

//getUsers list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

//Post method to register a new user
func registerUser(c *gin.Context) {
	var newUser registred_user

	//Call a BindJson to bind  the received JSON to new user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	//add a new user to the users slice.
	for _, existed_user := range users {
		if err := verifyEmail(existed_user.Email, newUser.Email); err != nil {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
	}
	newUser.ID = strconv.Itoa(len(users) + 1)
	newUser.Quiz = questions
	newUser.Number_corrected_answers = 0
	newUser.User_rated = 0.00
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, "user created sucessfully")
}

func main() {
	router := gin.Default()
	router.GET("/questions", getQuestions)
	router.GET("/users", getUsers)
	router.POST("/user", registerUser)
	router.Run("localhost:8080")
}
