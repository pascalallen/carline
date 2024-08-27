package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pascalallen/carline/internal/carline/infrastructure/routes"
	"os"
)

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	router := routes.NewRouter()

	router.Config()
	router.Fileserver()
	router.Default()
	router.Serve(":9990")
}
