package services

import (
	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"github.com/adityasingh-ops/quizzy_backend/internal/repositories"
)

type QuizService struct {
	quizRepo *repositories.QuizRepository
}

func NewQuizService(quizRepo *repositories.QuizRepository) *QuizService {
	return &QuizService{quizRepo: quizRepo}
}

func (s *QuizService) CreateQuiz(quiz *models.Quiz) (*models.Quiz, error) {
	return s.quizRepo.Create(quiz)
}

func (s *QuizService) GetQuiz(id string) (*models.Quiz, error) {
	return s.quizRepo.GetByID(id)
}

func (s *QuizService) ListQuizzes() ([]*models.Quiz, error) {
	return s.quizRepo.List()
}

func (s *QuizService) UpdateQuiz(id string, quiz *models.Quiz) (*models.Quiz, error) {
	return s.quizRepo.Update(id, quiz)
}

func (s *QuizService) DeleteQuiz(id string) error {
	return s.quizRepo.Delete(id)
}

func (s *QuizService) GetRandomQuiz() (*models.Quiz, error) {
	return s.quizRepo.GetRandom()
}