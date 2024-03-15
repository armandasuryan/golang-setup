package usecase

import (
	"log"

	"github.com/go-redis/redis/v8"
	// "go.mongodb.org/mongo-driver/mongo"

	"gorm.io/gorm"
)

type Imodel interface {
	ICiClass
}

type model struct {
	mysql *gorm.DB
	redis *redis.Client
	// mongo  *mongo.Database
	logger *log.Logger
}

//	func NewModel(db *gorm.DB, redis *redis.Client, mongo *mongo.Database, logger *log.Logger) *model {
//		return &model{db, redis, mongo, logger}
//	}
func NewModel(db *gorm.DB, redis *redis.Client, logger *log.Logger) *model {
	return &model{db, redis, logger}
}
