// internal/repositories/quiz_repository.go
package repositories

import (
	"context"

	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

type QuizRepository struct {
	collection *mongo.Collection
}

func NewQuizRepository(db *mongo.Database) *QuizRepository {
	return &QuizRepository{
		collection: db.Collection("quizzes"),
	}
}

func (r *QuizRepository) Create(quiz *models.Quiz) (*models.Quiz, error) {
	result, err := r.collection.InsertOne(context.Background(), quiz)
	if err != nil {
		return nil, err
	}

	quiz.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return quiz, nil
}

func (r *QuizRepository) GetByID(id string) (*models.Quiz, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var quiz models.Quiz
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&quiz)
	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

func (r *QuizRepository) List() ([]*models.Quiz, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var quizzes []*models.Quiz
	if err = cursor.All(context.Background(), &quizzes); err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (r *QuizRepository) Update(id string, quiz *models.Quiz) (*models.Quiz, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, quiz)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}

func (r *QuizRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

func (r *QuizRepository) GetRandom() (*models.Quiz, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}},
	}

	cursor, err := r.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var quizzes []models.Quiz
	if err = cursor.All(context.Background(), &quizzes); err != nil {
		return nil, err
	}

	if len(quizzes) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &quizzes[0], nil
}