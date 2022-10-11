package main

import (
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// 	By Pasakorn Limchuchua

	//	Main Reference
	// 	-> https://github.com/codebangkok/golang
	//	-> https://www.youtube.com/watch?v=4EBhkFWN16w&ab_channel=CodeBangkok

	//	Go Packages
	//	-> https://pkg.go.dev/github.com/gofiber/fiber/v2
	//	-> https://pkg.go.dev/gorm.io/gorm@v1.23.8
	//	-> https://pkg.go.dev/gorm.io/driver/mysql
	//	-> https://pkg.go.dev/github.com/go-redis/redis/v9

	//	Docker
	//	-> https://hub.docker.com/_/redis
	//	-> https://hub.docker.com/r/loadimpact/k6
	//	-> https://hub.docker.com/_/influxdb
	//	-> https://hub.docker.com/_/mariadb

	//	Redis (Local)
	//	-> https://redis.io/docs/
	//	-> https://redis.io/commands/
	//	-> https://redis.io/docs/manual/config/

	//	k6 - InfluxDB - Grafana
	//	-> https://k6.io/docs/
	//	-> https://k6.io/docs/using-k6/k6-options/reference/
	//	-> https://docs.influxdata.com/influxdb/v1.8/
	//	-> https://docs.influxdata.com/influxdb/v1.8/administration/config/

	//	Extensions
	// 	-> Go, Error Lens, MySQL

	//	connection						-> Host : 127.0.0.1 / port : 3306 / username : root / password : pass
	//									-> execute testdb2

	// 	Commands
	//	install image				-> docker pull redis
	//								-> docker pull loadimpact/k6
	//								-> docker pull influxdb:1.8.10
	//								-> docker pull grafana/grafana
	//								-> docker pull mariadb

	//	use docker 					-> docker-compose up -d (run all)
	//								-> docker-compose up -d influxdb grafana
	//	run k6 in docker			-> docker-compose run --rm k6 run /scripts/test.js

	//	install fiber				-> go get github.com/gofiber/fiber/v2
	//	install gorm				-> go get gorm.io/gorm
	// 	install driver (Maria DB)	-> go get gorm.io/driver/mysql
	//	install redis				-> go get github.com/go-redis/redis/v9
	//	run service					-> go run .
	//	test service				-> curl localhost:8000/products

	//	install redis local			-> brew install redis
	//	use redis server			-> redis-server
	//  use redis cli				-> redis-cli (guide : redis-cli --help)
	//								-> get repository::GetProducts (key)

	/* 	--------------- Redis--------------- */

	//	: in-memory data store (as database, cache, streaming engine, and message broker)

	//	Configuration
	//	1. config yml file in volumes: command:
	//	2. config redis.conf file (source : https://redis.io/docs/manual/config/ : use 7.0)
	// 		bind 0.0.0.0		-> can access to redis container
	//		protected-mode no	-> allow other host to connect to redis
	//		SAVE ""				-> disable snapshot (too slow + can crash)
	//							-> will use appendonly instead
	//		appendonly yes		-> prevent power outage problem

	/*  --------------- k6 - InfluxDB - Grafana --------------- */

	//	k6 			: load testing tool (write in js script)
	//	InfluxDB 	: time series database use for collecting test results from k6
	//	Grafana		: analytics & monitoring for database (InfluxDB)

	//	Configuration : config yml file in volumes: environment:

	//	Using K6 option (can use with command line or write in script test.js)
	//	-u, --vus 			= number of virtual users (ex. 1000)
	//	-d, --duration 		= test duration limit (ex. 10s)
	//	-o, --out			= send result to influxdb

	//	Using Grafana Dashboard Monitoring (after docker-compose)
	//	1. check port influxdb : localhost:8086 (for 404) / port grafana : localhost:3000
	//	2. in grafana : choose Configuration -> data source -> add data source -> InfluxDB
	//	3. in setting : set URL http://influxdb:8086 -> set Database k6 -> save & test
	//	4. in grafana : choose Dashboards -> import -> copy id at
	//					https://grafana.com/grafana/dashboards/2587-k6-load-testing-results/
	//					-> paste -> load -> k6 select InfluxDB (default)
	// 	then run services + test

	/*  --------------- Maria DB --------------- */

	//	: relational database (use instead MySQL)

	/* 	--------------- Project --------------- */

	// 	Objective : to test performance between no redis vs with redis

	// 	Project Structure : using hexagonal architecture
	//	-> overview : https://youtu.be/4EBhkFWN16w?t=4623
	//	-> port : repositories (gorm + Maria DB) 	[]:		=- adapter : database
	//	|													=- adapter : database with redis
	//	-> port : services  						[]:		=- adapter : service
	//	|													=- adapter : service with redis
	//	-> port : handlers (fiber) 					[]:		=- adapter : handler
	//														=- adapter : handler with redis

	//	!!! However we will use redis only one

	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	// 	Action Zone //

	productRepo := repositories.NewProductRepositoryDB(db)
	// productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	// productService := services.NewCatalogService(productRepo)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient) // recommend
	productHandler := handlers.NewCatalogHandler(productService)
	// productHandler := handlers.NewCatalogHanlderRedis(productService, redisClient)

	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Listen(":8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:pass@tcp(127.0.0.1:3306)/testdb2?parseTime=True")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
