package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/adityasingh-ops/quizzy_backend/config"
	"github.com/adityasingh-ops/quizzy_backend/internal/handlers"
	"github.com/adityasingh-ops/quizzy_backend/internal/repositories"
	"github.com/adityasingh-ops/quizzy_backend/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	// Initialize repositories
	db := client.Database(cfg.DBName)
	quizRepo := repositories.NewQuizRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)
	scoreRepo := repositories.NewScoreRepository(db)

	// Initialize services
	quizService := services.NewQuizService(quizRepo)
	questionService := services.NewQuestionService(questionRepo)
	scoreService := services.NewScoreService(scoreRepo)

	// Initialize handlers
	quizHandler := handlers.NewQuizHandler(quizService)
	questionHandler := handlers.NewQuestionHandler(questionService)
	scoreHandler := handlers.NewScoreHandler(scoreService)

	// Set up Gin router
	router := gin.Default()

	// Register routes
	api := router.Group("/api")
	{
		quizzes := api.Group("/quizzes")
		{
			quizzes.POST("/", quizHandler.CreateQuiz)
			quizzes.GET("/", quizHandler.ListQuizzes)
			quizzes.GET("/:id", quizHandler.GetQuiz)
			quizzes.PUT("/:id", quizHandler.UpdateQuiz)
			quizzes.DELETE("/:id", quizHandler.DeleteQuiz)
			quizzes.GET("/random", quizHandler.GetRandomQuiz)
		}

		questions := api.Group("/questions")
		{
			questions.POST("/", questionHandler.CreateQuestion)
			questions.GET("/:id", questionHandler.GetQuestion)
			questions.PUT("/:id", questionHandler.UpdateQuestion)
			questions.DELETE("/:id", questionHandler.DeleteQuestion)
			questions.GET("/quiz/:quizId", questionHandler.ListQuestionsByQuiz)
		}

		scores := api.Group("/scores")
		{
			scores.POST("/", scoreHandler.SubmitScore)
			scores.GET("/user/:userId", scoreHandler.GetUserScores)
			scores.GET("/quiz/:quizId", scoreHandler.GetQuizScores)
			scores.GET("/average/user/:userId", scoreHandler.GetUserAverageScore)
		}
	}

	// Start the server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}