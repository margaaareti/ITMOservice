package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"itmo/auth"
	"itmo/auth/delivery"
	"itmo/auth/services"
	"itmo/auth/storage"
	"itmo/server/repository"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer  *http.Server
	authService auth.AuthService
}

func NewApp() *App {

	//pgsql
	db, err := repository.CreateNewPostgresConnection(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal(err)

	}

	authPostgres := storage.CreateNewUserStorage(db)
	authService := services.CreateNewAuthService(authPostgres)

	return &App{
		authService: authService,
	}

}

func (a *App) Run(port string) error {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	//router.LoadHTMLGlob("../ui/html/**")

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, //20mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			logrus.Fatal("Failed to listen and serve: %+v", err)
		}
	}()

	delivery.RegisterHTTPEndPoints(router, a.authService)

	//router.GET("/", func(c *gin.Context) {
	//c.String(200, "Hello выаWorld!")
	//})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	s := <-quit
	fmt.Printf("Got signal:%s", s)

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)

}
