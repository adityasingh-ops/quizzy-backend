package services

import (
	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"github.com/adityasingh-ops/quizzy_backend/internal/repositories"
)

type ScoreService struct {
	scoreRepo *repositories.ScoreRepository
}

func NewScoreService(scoreRepo *repositories.ScoreRepository) *ScoreService {
	return &ScoreService{scoreRepo: scoreRepo}
}

func (s *ScoreService) SubmitScore(score *models.UserScore) (*models.UserScore, error) {
	return s.scoreRepo.Create(score)
}

func (s *ScoreService) GetUserScores(userID string) ([]*models.UserScore, error) {
	return s.scoreRepo.GetByUserID(userID)
}

func (s *ScoreService) GetQuizScores(quizID string) ([]*models.UserScore, error) {
	return s.scoreRepo.GetByQuizID(quizID)
}

func (s *ScoreService) GetUserAverageScore(userID string) (float64, error) {
	return s.scoreRepo.GetAverageScoreByUserID(userID)
}