package repositories

import (
	"context"

	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionRepository struct {
	collection *mongo.Collection
}

func NewQuestionRepository(db *mongo.Database) *QuestionRepository {
	return &QuestionRepository{
		collection: db.Collection("questions"),
	}
}

func (r *QuestionRepository) Create(question *models.Question) (*models.Question, error) {
	result, err := r.collection.InsertOne(context.Background(), question)
	if err != nil {
		return nil, err
	}

	question.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return question, nil
}

func (r *QuestionRepository) GetByID(id string) (*models.Question, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var question models.Question
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&question)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (r *QuestionRepository) Update(id string, question *models.Question) (*models.Question, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	_, err = r.collection.ReplaceOne(context.Background(), bson.M{"_id": objectID}, question)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuestionRepository) Delete(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}

func (r *QuestionRepository) ListByQuizID(quizID string) ([]*models.Question, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"quizId": quizID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var questions []*models.Question
	if err = cursor.All(context.Background(), &questions); err != nil {
		return nil, err
	}

	return questions, nil
}