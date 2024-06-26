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

	router.Use(h.logMiddleware)
	router.Use(h.CORSMiddleware)
	api := router.Group("/api")
	{
		api.GET("/admin-exists", h.adminExist)

		// Группы роутеров
		auth := api.Group("/auth")
		{
			// endpoints группы
			auth.POST("/register", h.register)
			auth.POST("/login", h.logIn)
		}

		users := api.Group("/user", h.userIdentity, h.userCheckApprove)
		{
			users.DELETE("/delete", h.deleteUser)
			users.PUT("/update", h.updateUser)
			users.GET("/info", h.infoUser)
			users.GET("/limit", h.limitUser)
			users.POST("/new-password", h.newPassword)
		}

		folders := api.Group("/folder", h.userIdentity, h.userCheckApprove)
		{
			folders.POST("/create", h.createFolder)
			folders.GET("/all", h.getAllFolders)
			// Все файл в директории
			folders.GET("/:folder_id", h.getFilesInFolder)
			folders.DELETE("/:folder_id", h.deleteFolder)
			folders.PUT("/update/:folder_id", h.updateFolder)
		}

		files := api.Group("/file", h.userIdentity, h.userCheckApprove)
		{
			// create file
			files.POST("/upload", h.maxFileLimit, h.uploadFile)
			files.GET("/download/:file_id", h.downloadFile)
			files.GET("/list", h.getFilesInRoot)
			files.POST("/download-many", h.downloadManyFiles)
			files.DELETE("/:file_id", h.deleteFile)
			files.PUT("update/:file_id", h.updateFile)
		}

		admin := api.Group("/admin", h.userIdentity, h.adminIdentify)
		{
			admin.GET("/user-list", h.userList)
			admin.GET("/user-list/search", h.userListSearch)
			admin.PUT("/toggle-approve/:user_id", h.toggleUserApprove)
		}

		fileUrl := api.Group("/url", h.userIdentity, h.userCheckApprove)
		{
			fileUrl.POST("/create", h.createUrl)
		}

		urlGet := api.Group("/url-get")
		{
			urlGet.GET("/files/:uuid", h.getFilesUUid)
			urlGet.GET("/download/:uuid", h.downloadFilesUUid)
		}
	}

	return router
}
