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

	// h.userApproved
	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/user")
		{
			users.POST("/disable/:user_id", h.disableUser)
			users.DELETE("/delete", h.deleteUser)
			users.PUT("/update", h.updateUser)
			users.POST("/new-password", h.newPassword)
		}

		folders := api.Group("folder")
		{
			folders.POST("/create", h.createFolder)
			folders.GET("/all", h.getAllFolders)
			// Все файл в директории
			folders.GET("/:folder_id", h.getFilesInFolder)
			folders.DELETE("/:folder_id", h.deleteFolder)
			folders.PUT("/update", h.updateFolder)
		}

		files := api.Group("/file")
		{
			files.POST("/create", h.createFile)
			files.DELETE("/:file_id", h.deleteFile)
			files.POST("/upload", h.uploadFile)
			files.GET("/download", h.downloadFile)
		}
	}

	return router
}
