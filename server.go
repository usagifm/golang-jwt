package main

import (
	"github.com/gin-gonic/gin"
	"github.com/usagifm/golang-jwt/config"
	"github.com/usagifm/golang-jwt/controller"
	"github.com/usagifm/golang-jwt/middleware"
	"github.com/usagifm/golang-jwt/repository"
	"github.com/usagifm/golang-jwt/service"
	"gorm.io/gorm"
)

// variable global untuk controller2
var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	diaryRepository repository.DiaryRepository = repository.NewDiaryRepository(db)
	jwtService      service.JWTService         = service.NewJWTService()
	userService     service.UserService        = service.NewUserService(userRepository)
	diaryService    service.DiaryService       = service.NewDiaryService(diaryRepository)
	authService     service.AuthService        = service.NewAuthService(userRepository)
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	diaryController controller.DiaryController = controller.NewDiaryController(diaryService, jwtService)
)

func main() {
	// if failed to connect then throw error
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()

	// auth routes
	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// user routes
	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
		userRoutes.GET("/diaries/self", userController.MyDiaries)

		userRoutes.GET("/diaries", diaryController.All)
		userRoutes.POST("/diaries/diary/insert", diaryController.Insert)
		userRoutes.GET("/diaries/diary/:id", diaryController.FindById)
		userRoutes.PUT("/diaries/diary/:id/update", diaryController.Update)
		userRoutes.DELETE("/diaries/diary/:id/delete", diaryController.Delete)
	}

	server.Run() // listen to default port :8080
}
