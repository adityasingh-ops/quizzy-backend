package repositories

import (
	"context"

	"github.com/adityasingh-ops/quizzy_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScoreRepository struct {
	collection *mongo.Collection
}

func NewScoreRepository(db *mongo.Database) *ScoreRepository {
	return &ScoreRepository{
		collection: db.Collection("scores"),
	}
}

func (r *ScoreRepository) Create(score *models.UserScore) (*models.UserScore, error) {
	result, err := r.collection.InsertOne(context.Background(), score)
	if err != nil {
		return nil, err
	}

	score.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return score, nil
}

func (r *ScoreRepository) GetByUserID(userID string) ([]*models.UserScore, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var scores []*models.UserScore
	if err = cursor.All(context.Background(), &scores); err != nil {
		return nil, err
	}

	return scores, nil
}

func (r *ScoreRepository) GetByQuizID(quizID string) ([]*models.UserScore, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"quizId": quizID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var scores []*models.UserScore
	if err = cursor.All(context.Background(), &scores); err != nil {
		return nil, err
	}

	return scores, nil
}

func (r *ScoreRepository) GetAverageScoreByUserID(userID string) (float64, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"userId": userID}},
		{"$group": bson.M{"_id": nil, "averageScore": bson.M{"$avg": "$score"}}},
	}

	cursor, err := r.collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(context.Background())

	var result struct {
		AverageScore float64 `bson:"averageScore"`
	}

	if cursor.Next(context.Background()) {
		err := cursor.Decode(&result)
		if err != nil {
			return 0, err
		}
		return result.AverageScore, nil
	}

	return 0, mongo.ErrNoDocuments
}