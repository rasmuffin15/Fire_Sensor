package Routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run will start the server
// Runs on localhost:8080
func Run() {
	handleRoutes()
	router.Run()
}

//Calls correct route depending on input
func handleRoutes() {
	r1 := router.Group("/r1")
	sensorPutPath(r1)
}

