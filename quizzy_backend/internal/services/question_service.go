package services

import (
	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"github.com/adityasingh-ops/quizzy_backend/internal/repositories"
)

type QuestionService struct {
	questionRepo *repositories.QuestionRepository
}

func NewQuestionService(questionRepo *repositories.QuestionRepository) *QuestionService {
	return &QuestionService{questionRepo: questionRepo}
}

func (s *QuestionService) CreateQuestion(question *models.Question) (*models.Question, error) {
	return s.questionRepo.Create(question)
}

func (s *QuestionService) GetQuestion(id string) (*models.Question, error) {
	return s.questionRepo.GetByID(id)
}

func (s *QuestionService) UpdateQuestion(id string, question *models.Question) (*models.Question, error) {
	return s.questionRepo.Update(id, question)
}

func (s *QuestionService) DeleteQuestion(id string) error {
	return s.questionRepo.Delete(id)
}

func (s *QuestionService) ListQuestionsByQuiz(quizID string) ([]*models.Question, error) {
	return s.questionRepo.ListByQuizID(quizID)
}