package models
type Question struct {
    ID            string   `json:"id" bson:"_id,omitempty"`
    QuizID        string   `json:"quizId" bson:"quizId"`
    Text          string   `json:"text" bson:"text"`
    Options       []string `json:"options" bson:"options"`
    CorrectAnswer int      `json:"correctAnswer" bson:"correctAnswer"`
    Explanation   string   `json:"explanation" bson:"explanation"`
    TimerDuration int      `json:"timerDuration" bson:"timerDuration"` // in seconds
}