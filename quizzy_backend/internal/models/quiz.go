package models

type Quiz struct {
    ID            string `json:"id" bson:"_id,omitempty"`
    Name          string `json:"name" bson:"name"`
    Description   string `json:"description" bson:"description"`
    ImageURL      string `json:"imageUrl" bson:"imageUrl"`
    QuestionCount int    `json:"questionCount" bson:"questionCount"`
    Difficulty    string `json:"difficulty" bson:"difficulty"`
}