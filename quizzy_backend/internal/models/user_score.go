package models
type UserScore struct {
    ID     string  `json:"id" bson:"_id,omitempty"`
    UserID string  `json:"userId" bson:"userId"`
    QuizID string  `json:"quizId" bson:"quizId"`
    Score  int     `json:"score" bson:"score"`
    Date   string  `json:"date" bson:"date"`
}