package api

import (
	"imap-sync/config"
	"imap-sync/controller"
	"imap-sync/internal"
	"imap-sync/logger"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

var log = logger.Log
var port string

func InitServer() {
	port = config.Conf.Port
	logger.SetupLogger()
	err := internal.InitDb()
	if err != nil {
		log.Error(err)
	}

	err = internal.InitSettingsTable()
	if err != nil {
		log.Error(err)
	}

	internal.InitLocalizer()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(ginsession.New())

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static/")

	router.GET("/", controller.HandleRoot)
	router.GET("/admin", controller.HandleAdmin)
	router.GET("/login", controller.HandleLogin)
	router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("favicon.ico")
	})
	router.GET("/health", func(ctx *gin.Context) {
		dbStatus := "healthy"
		if err := internal.CheckDB(); err != nil {
			dbStatus = "unhealthy"
		}

		ctx.JSON(200, gin.H{
			"status":    "healthy",
			"version":   "1.0.0",
			"db_status": dbStatus,
			"uptime":    internal.GetUptime(),
		})
	})
	go internal.InitQueue()
	// API endpoints
	router.GET("/api/queue", controller.HandleQueue)
	router.GET("/api/queuepoll", controller.HandleQueuePolling)
	router.GET("/api/pagination", controller.HandlePagination)
	router.GET("/api/details", controller.HandleGetLog)
	router.GET("/api/sync", controller.HandleSync)
	router.GET("/api/settings", controller.HandleGetSettings)
	router.PUT("/api/settings", controller.HandleUpdateSettings)
	router.POST("/api/bulk", controller.HandleBulkMigration)
	router.GET("/api/bulk/status", controller.HandleBulkMigrationStatus)
	router.GET("/api/stats", controller.HandleGetStats)
	router.GET("/api/system", controller.HandleGetSystemInfo)
	router.GET("/api/audit", controller.HandleGetAuditLog)
	router.GET("/api/sessions", controller.HandleGetSessions)
	router.POST("/api/sessions/:id/terminate", controller.HandleTerminateSession)
	router.POST("/api/sessions/terminate-all", controller.HandleTerminateAllSessions)
	router.POST("/api/validate", controller.HandleValidate)
	router.POST("/api/search", controller.HandleSearch)
	router.POST("/auth/login", controller.Login)

	log.Info("Server starting on http://localhost:" + port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
