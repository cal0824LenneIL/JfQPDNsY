// 代码生成时间: 2025-10-12 03:36:22
package main

import (
    "net/http"
    "github.com/labstack/echo"
    "golang.org/x/crypto/bcrypt"
)

// Question represents a question model
type Question struct {
    ID    uint   `json:"id"`
    Title string `json:"title"`
    Answer string `json:"answer"`
}

// QuestionService handles operations related to questions
type QuestionService struct {
    // No additional fields needed for this service
}

// NewQuestionService creates a new QuestionService
func NewQuestionService() *QuestionService {
    return &QuestionService{}
}

// CreateQuestion adds a new question to the database
func (s *QuestionService) CreateQuestion(title, answer string) (*Question, error) {
    // Simulate creating a question in a database
    // In a real-world scenario, you would interact with an actual database
    // and handle potential errors (e.g., unique constraint violations)
    // For simplicity, let's assume the operation is successful
    
    // Hash the answer using bcrypt for security
    hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(answer), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Create a new Question
    question := &Question{
        Title: title,
        Answer: string(hashedAnswer),
    }

    // Return the created question
    return question, nil
}

// QuestionHandler handles HTTP requests for questions
func QuestionHandler(e *echo.Echo, s *QuestionService) {
    e.POST("/question", func(c echo.Context) error {
        title := c.QueryParam("title")
        answer := c.QueryParam("answer")

        if title == "" || answer == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Title and answer are required")
        }

        question, err := s.CreateQuestion(title, answer)
        if err != nil {
            return err
        }

        // Return the created question as JSON
        return c.JSON(http.StatusCreated, question)
    })
}

func main() {
    e := echo.New()
    s := NewQuestionService()
    QuestionHandler(e, s)

    e.Logger.Fatal(e.Start(":8080"))
}