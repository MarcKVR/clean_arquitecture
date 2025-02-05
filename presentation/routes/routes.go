package routes

import (
	"github.com/MarcKVR/clean_arquitecture/application/service"
	"github.com/MarcKVR/clean_arquitecture/application/usecases"
	"github.com/MarcKVR/clean_arquitecture/cmd/config"
	"github.com/MarcKVR/clean_arquitecture/infrastructure/repository"
	"github.com/MarcKVR/clean_arquitecture/infrastructure/s3"
	"github.com/MarcKVR/clean_arquitecture/presentation/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// // func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {

// func RegisterRoutes(app *fiber.App, db *gorm.DB, logger *log.Logger) {
// 	// // Routes
// 	apiGroup := app.Group("/api")

// 	authRepo := repository.NewAuthRepository(db, logger)
// 	authService := service.NewAuthService(authRepo, logger)
// 	authHandler := handler.NewAuthHandler(authService)
// 	app.Post("/login", authHandler.Login)

// 	// Users routes
// 	userRepo := repository.NewUserRepository(db, logger)
// 	userService := service.NewUserService(userRepo, logger)
// 	userHandler := handler.NewUserHandler(userService)

// 	apiGroup.Post("/users", userHandler.Create)
// 	apiGroup.Get("/users/:id", userHandler.GetUser)

// }

func RegisterRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {
	// apiGroup := app.Group("/api")

	repo := repository.NewExchangeRateRepository(cfg.ExchangeAPIURL)
	usecase := usecases.NewExchangeRateUsecase(repo)
	handlers.NewExchangeRateHandler(app, usecase)

	s3Repo := s3.NewS3SClient(cfg)
	s3Service := service.NewFileService(s3Repo)
	s3Handler := handlers.NewFileHandler(s3Service)
	app.Post("/upload", s3Handler.UploadFile)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users/:id", userHandler.GetUserById)
	app.Get("/users", userHandler.GetAllUsers)
	app.Put("/users", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
