package server

import (
	"fmt"
	"mutant_detector/server/http/config"
	"mutant_detector/server/http/managers"

	"github.com/gin-gonic/gin"
)

// HTTPInitServer init http delivery server
func HTTPInitServer() {
	router := gin.Default()

	managers.HTTPInitHandlers(router)

	serverPort := config.HTTPGetServerPort()

	router.Run(fmt.Sprintf(":%s", serverPort))
}
