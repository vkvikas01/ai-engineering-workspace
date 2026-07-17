package main

import (
	
	
	"ai-engineering-workspace/auth-service/internal/config"
    "ai-engineering-workspace/auth-service/internal/database"
    "ai-engineering-workspace/auth-service/internal/logger"

    "github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main(){
	cfg, err := config.LoadConfig()
	
	if err != nil {
		panic(err)
	}

	log,err := logger.New()
	if err != nil {
		panic(err)
	}

	defer log.Sync()

	db,err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	redis:= database.ConnectRedis(cfg)

	if err := database.PingRedis(redis); err != nil {
		log.Fatal("Failed to connect to Redis", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	

	router:= gin.Default();

	router.GET("/health",func(c *gin.Context){
		dbStatus:= sqlDB.Ping()==nil
		redisStatus:= database.PingRedis(redis)==nil
		c.JSON(200,gin.H{
			"service":cfg.AppName,
			"status":"ok",
			"database":dbStatus,
			"redis":redisStatus,
		})
	})

	router.Run(":" + cfg.AppPort)
}