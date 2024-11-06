package main

import (
	app_http "blog-mandalika/app/delivery/http"
	mysqlrepo "blog-mandalika/app/repository/mysql"
	"blog-mandalika/app/usecase/admin"
	"blog-mandalika/app/usecase/public"
	"blog-mandalika/database"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		timeoutStr = "5"
	}
	timeout, _ := strconv.Atoi(timeoutStr)
	timeoutContext := time.Duration(timeout) * time.Second

	// logger
	writers := make([]io.Writer, 0)
	if logSTDOUT, _ := strconv.ParseBool(os.Getenv("LOG_TO_STDOUT")); logSTDOUT {
		writers = append(writers, os.Stdout)
	}

	// set gin writer to logrus
	// gin.DefaultWriter = logrus.StandardLogger().Writer()

	// log to file
	if logFILE, _ := strconv.ParseBool(os.Getenv("LOG_TO_FILE")); logFILE {
		logMaxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
		if logMaxSize == 0 {
			logMaxSize = 50 //default 50 megabytes
		}

		logFilename := os.Getenv("LOG_FILENAME")
		if logFilename == "" {
			logFilename = "server.log"
		}

		lg := &lumberjack.Logger{
			Filename:   logFilename,
			MaxSize:    logMaxSize,
			MaxBackups: 1,
			LocalTime:  true,
		}

		writers = append(writers, lg)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(io.MultiWriter(writers...))

	// init mysql database
	mysqlDB, err := database.InitMysql()
	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	// init repo mysql
	mysqlRepo := mysqlrepo.NewDBRepo(mysqlDB)

	// init usecase
	adminUC := admin.NewAppAdminUsecase(admin.RepoInjection{
		MysqlRepo: mysqlRepo,
	}, timeoutContext)

	publicUC := public.NewAppPublicUsecase(public.RepoInjection{
		MysqlRepo: mysqlRepo,
	}, timeoutContext)

	// init gin
	gin.SetMode(os.Getenv("GIN_MODE"))
	ginEngine := gin.New()

	// log client requests
	if os.Getenv("LOG_TO_STDOUT") == "true" {
		ginEngine.Use(gin.Logger(), gin.Recovery())
	}

	// cors
	ginEngine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	// default route
	ginEngine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "It works",
		})
	})

	// init route
	app_http.NewRouteHandler(ginEngine, adminUC, publicUC)

	port := os.Getenv("PORT")
	if os.Getenv("LOG_TO_STDOUT") == "true" {
		fmt.Printf("[%s] Service running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), port)
	}

	ginEngine.Run(":" + port)
}
