package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"github.com/adityasingh-ops/quizzy_backend/internal/services"
)

type ScoreHandler struct {
	scoreService *services.ScoreService
}

func NewScoreHandler(scoreService *services.ScoreService) *ScoreHandler {
	return &ScoreHandler{scoreService: scoreService}
}

func (h *ScoreHandler) SubmitScore(c *gin.Context) {
	var score models.UserScore
	if err := c.ShouldBindJSON(&score); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submittedScore, err := h.scoreService.SubmitScore(&score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, submittedScore)
}

func (h *ScoreHandler) GetUserScores(c *gin.Context) {
	userID := c.Param("userId")
	scores, err := h.scoreService.GetUserScores(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, scores)
}

func (h *ScoreHandler) GetQuizScores(c *gin.Context) {
	quizID := c.Param("quizId")
	scores, err := h.scoreService.GetQuizScores(quizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, scores)
}

func (h *ScoreHandler) GetUserAverageScore(c *gin.Context) {
	userID := c.Param("userId")
	averageScore, err := h.scoreService.GetUserAverageScore(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"averageScore": averageScore})
}