package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

//	adapter

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate([]product{})
	mockData(db)
	return productRepositoryRedis{db: db, redisClient: redisClient}
}

//	method

func (r productRepositoryRedis) GetProducts() (products []product, err error) {

	key := "repository::GetProducts"

	//	redis get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}

	//	database
	err = r.db.Order("quantity desc").Limit(20).Find(&products).Error
	if err != nil {
		return nil, err
	}

	//	redis set
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}

	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}
	fmt.Println("database")
	return products, nil
}
