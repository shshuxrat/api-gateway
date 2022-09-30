package api

import (
	"api-gateway/api/docs"
	"api-gateway/api/handlers/v1"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServicesI
}

// SetUpRouter godoc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	docs.SwaggerInfo.Title = "Api Gateway"
	docs.SwaggerInfo.Version = "1.0"

	router.Use(cors.New(config))

	handlerV1 := handlers.NewHandler(&handlers.HandlerOptions{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	//profession
	apiV1.POST("profession", handlerV1.CreateProfession)
	apiV1.GET("/profession", handlerV1.GetAllProfession)
	apiV1.GET("/profession/:profession_id", handlerV1.GetProfession)
	apiV1.PUT("/profession", handlerV1.UpdateProfession)
	apiV1.DELETE("/profession/:profession_id", handlerV1.DeleteProfession)
	//attribute
	apiV1.POST("/attribute", handlerV1.CreateAttribute)
	apiV1.GET("/attribute", handlerV1.GetAllAttribute)
	apiV1.GET("/attribute/:attribute_id", handlerV1.GetAttribute)
	apiV1.PUT("/attribute", handlerV1.UpdateAttribute)
	apiV1.DELETE("/attribute/:attribute_id", handlerV1.DeleteAttribute)
	//company
	apiV1.POST("/company", handlerV1.CreateCompany)
	apiV1.GET("/company", handlerV1.GetAllCompany)
	apiV1.GET("/company/:company_id", handlerV1.GetCompany)
	apiV1.PUT("/company", handlerV1.UpdateCompany)
	apiV1.DELETE("/company/:company_id", handlerV1.DeleteCompany)

	//position
	apiV1.POST("/position", handlerV1.CreatePosition)
	apiV1.GET("/position", handlerV1.GetAllPosition)
	apiV1.GET("/position/:position_id", handlerV1.GetPosition)
	apiV1.PUT("/position", handlerV1.UpdatePosition)
	apiV1.DELETE("/position/:position_id", handlerV1.DeletePosition)
	//position_attribute
	apiV1.POST("/position_attribute", handlerV1.CreatePositionAttribute)
	apiV1.GET("/position_attribute", handlerV1.GetAllPositionAttribute)
	apiV1.GET("/position_attribute/:position_attribute_id", handlerV1.GetPositionAttribute)
	apiV1.PUT("/position_attribute", handlerV1.UpdatePositionAttribute)
	apiV1.DELETE("/position_attribute/:position_attribute_id", handlerV1.DeletePositionAttribute)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router

}
