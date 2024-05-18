package api

import (
	"github.com/KodokuOdius/SecureFileChanger/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// Уровень обработчиков (работа с http)
func (h *Handler) InitRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Группы роутеров
	auth := router.Group("/auth")
	{
		// endpoints группы
		// /auth/sign-up
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/upload", h.uploadFile)
		api.GET("/download", h.downloadFile)
		users := api.Group("/user")
		{
			users.POST("/create", h.userCreate)
		}
		folders := api.Group("folder")
		{
			folders.POST("/create", h.createFolder)
			folders.GET("/all", h.getAllFolders)
			folders.GET("/:folder_id", h.getFolderById)
			folders.DELETE("/:folder_id", h.deleteFolder)
		}
	}

	return router
}
