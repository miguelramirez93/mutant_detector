package server

import (
	"fmt"
	"mutant_detector/server/http/config"
	"mutant_detector/server/http/managers"

	_ "mutant_detector/docs" // should be imported for swagger docs

	"github.com/gin-gonic/gin"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @BasePath /

// HTTPInitServer init http delivery server
func HTTPInitServer() {
	router := gin.Default()

	managers.HTTPInitHandlers(router)

	serverPort := config.HTTPGetServerPort()

	router.Run(fmt.Sprintf(":%s", serverPort))
}
