package apps

import (
	"log"
	"os"
	"strconv"
	"sygap-cmdb/configuration/db"
	"sygap-cmdb/configuration/redis"
	"sygap-cmdb/controller"
	"sygap-cmdb/routes"
	"sygap-cmdb/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func StartApps() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Get the values from environment variables
	host_mysql := os.Getenv("HOST_MYSQL")
	username_mysql := os.Getenv("USSERNAME_MYSQL")
	password_mysql := os.Getenv("PASSWORD_MYSQL")
	db_mysql := os.Getenv("DB_MYSQL")

	// host_mongo := os.Getenv("HOST_MONGO")
	// db_mongo := os.Getenv("DB_MONGO")
	// username_mongo := os.Getenv("USERNAME_MONGO")
	// password_mongo := os.Getenv("PASSWORD_MONGO")

	host_redis := os.Getenv("HOST_REDIS")
	password_redis := os.Getenv("PASSWORD_REDIS")
	db_redis, _ := strconv.Atoi(os.Getenv("DB_REDIS"))

	redis := redis.RedisConnect(host_redis, password_redis, db_redis)
	mysql := db.MysqlConnect(host_mysql, username_mysql, password_mysql, db_mysql)
	// mongo := db.MongoConnect(host_mongo, db_mongo, username_mongo, password_mongo)
	// model := model.NewModel(mysql, redis, mongo, logger)
	model := usecase.NewModel(mysql, redis, logger)
	controller := controller.NewController(model, logger)

	// go controller.Shceduler()

	routes.SetupRoutes(app, controller)
	// cert and private key
	// keyFile := "./private.key"
	// certFile := "./server.crt"

	// app.ListenTLS(":3001",cert,privateKey)

	// errApp := app.ListenTLS(":8080", certFile, keyFile)
	errApp := app.Listen(":8080")
	if errApp != nil {
		log.Fatalf("Error starting Fiber app: %v", errApp)
	}
}
