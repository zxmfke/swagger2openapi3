package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "docs"
)

// @title           Example API
// @version         1.0
// @description     This is a sample server for swag example.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// HelloReq hello to who
type HelloReq struct {
	Name string `json:"name"`
}

// HelloResp hello to you
type HelloResp struct {
	Text string `json:"text"`
}

func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/v1/hello", Welcome)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// Welcome to welcome and hello
//
//	@Summary      welcome and hello
//	@Description  hello, welcome
//	@Tags         hello
//	@Accept       json
//	@Produce      json
//	@Param        json   body   HelloReq true   "HelloReq"
//	@Response     200    {object}   HelloResp
//	@Router       /api/v1/hello [post]
func Welcome(c *gin.Context) {
	var req = new(HelloReq)

	_ = c.BindJSON(req)

	c.JSON(http.StatusOK, HelloResp{
		Text: "hello" + req.Name,
	})
}
