package http

import (
	"net/http"
	"personal_knowledge_tracker/config"
	_ "personal_knowledge_tracker/docs"
	v1 "personal_knowledge_tracker/internal/controller/http/v1"
	"personal_knowledge_tracker/internal/interfaces"
	"personal_knowledge_tracker/pkg/middlewares/ginlogr"
	customGinSwagger "personal_knowledge_tracker/pkg/middlewares/ginswagger"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Dependencies struct {
	Cfg      *config.Config
	Logger   logr.Logger
	Usecases interfaces.Usecases
}

func NewRouter(deps Dependencies) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery())
	router.Use(ginlogr.Logger(deps.Logger, "/healthz"))
	router.Use(setSecurityHeaders(deps.Cfg))

	router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	swaggerHandlerV1 := customGinSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER", ginSwagger.InstanceName("v1"))

	v1.Register(router, v1.Dependencies{
		Cfg:      deps.Cfg,
		Logger:   deps.Logger,
		Usecases: deps.Usecases,
	})
	router.GET("/swagger/v1/*any", swaggerHandlerV1)

	return router
}

func setSecurityHeaders(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", cfg.CORS.AllowMethods)
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.CORS.AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", cfg.CORS.AllowCredentials)
		c.Writer.Header().Set("Access-Control-Allow-Headers", cfg.CORS.AllowHeaders)
		c.Writer.Header().Set("X-Content-Type-Options", cfg.CORS.XContentTypeOptions)
		c.Writer.Header().Set("X-Frame-Options", cfg.CORS.XFrameOptions)
		c.Writer.Header().Set("Content-Security-Policy", cfg.CORS.ContentSecurityPolicy)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
