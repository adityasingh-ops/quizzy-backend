package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"github.com/adityasingh-ops/quizzy_backend/internal/services"
)

type QuizHandler struct {
	quizService *services.QuizService
}

func NewQuizHandler(quizService *services.QuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdQuiz, err := h.quizService.CreateQuiz(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdQuiz)
}

func (h *QuizHandler) GetQuiz(c *gin.Context) {
	id := c.Param("id")
	quiz, err := h.quizService.GetQuiz(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (h *QuizHandler) ListQuizzes(c *gin.Context) {
	quizzes, err := h.quizService.ListQuizzes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

func (h *QuizHandler) UpdateQuiz(c *gin.Context) {
	id := c.Param("id")
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedQuiz, err := h.quizService.UpdateQuiz(id, &quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedQuiz)
}

func (h *QuizHandler) DeleteQuiz(c *gin.Context) {
	id := c.Param("id")
	err := h.quizService.DeleteQuiz(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz deleted successfully"})
}

func (h *QuizHandler) GetRandomQuiz(c *gin.Context) {
	randomQuiz, err := h.quizService.GetRandomQuiz()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, randomQuiz)
}
