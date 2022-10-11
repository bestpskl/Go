package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {

	// 	By Pasakorn Limchuchua

	//	Main Reference
	// 	-> https://github.com/codebangkok/golang
	//	-> https://www.youtube.com/watch?v=xHvzNJzA9DQ&ab_channel=CodeBangkok
	//	-> https://docs.gofiber.io/api/fiber
	//	-> https://docs.gofiber.io/api/app
	//	-> https://docs.gofiber.io/api/ctx
	//	-> https://docs.gofiber.io/api/middleware
	//	-> https://pkg.go.dev/github.com/gofiber/fiber/v2

	//	Extensions
	// 	-> Go, Error Lens, REST Client (instead of curl)

	// 	Commands
	// 	go mod init [module_name]		-> create modules
	// 	go run [file_name].go			-> run file
	//	go run .						-> run file
	//	go get [url_path]				-> installation

	// 	fiber framework					-> express js in go
	//	installation					-> go get github.com/gofiber/fiber/v2

	/* 	--------------- Fiber Framework ---------------	*/

	//	New					-> use to create a new App named instance (start server)
	// 	[Fiber]				-> can config by fiber.Config{}
	//						(https://docs.gofiber.io/api/fiber)

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	// 	Middleware			-> a function chained in the HTTP request cycle (~pipeline)
	//						(https://docs.gofiber.io/api/middleware)

	// 	Use 				-> use for middleware packages and prefix catchers (only match beginning of each path)
	//	[App*]				(https://docs.gofiber.io/api/app#route-handlers)

	//	Next				-> use to executes the next method in the stack that matches the current route
	//	[Context*]			(https://docs.gofiber.io/api/ctx#next)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("name", "best")
		fmt.Println("before")
		err := c.Next() // do next step of pipeline (match any request because we don't specify any path)
		fmt.Println("after")
		return err
	})

	//	RequestID			-> use to add an indentifier to the response
	//	[Middleware]		(https://docs.gofiber.io/api/middleware/requestid)

	app.Use(requestid.New()) // -> can config by requestid.Config{}

	//	CORS				-> use to enable Cross-Origin Resource Sharing
	//	[Middleware]		(https://docs.gofiber.io/api/middleware/cors)

	app.Use(cors.New(cors.Config{ // -> can config by cors.Config{}
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	//	Logger				-> use to log HTTP request/response details
	//	[Middleware]		(https://docs.gofiber.io/api/middleware/logger)

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	//	Route Handlers		-> use to register a route bound to a specific HTTP method.
	//	[App]				(https://docs.gofiber.io/api/app#route-handlers)

	//	GET

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("GET: Hello World")
	})

	//	POST

	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Post: Hello World")
	})

	//	Context				-> hold the HTTP request and response (has methods for the request)
	//						(https://docs.gofiber.io/api/ctx)

	// 	Locals 				-> use to stores variables scoped to the request
	//	[Context]			(https://docs.gofiber.io/api/ctx#locals)

	app.Get("/hellolocals", func(c *fiber.Ctx) error {
		fmt.Println("hello")
		name := c.Locals("name")
		return c.SendString(fmt.Sprintf("GET: Hello %v", name))
	})

	//	Params				-> use to get the route parameters
	//	[Context]			(https://docs.gofiber.io/api/ctx#params)
	//
	//						Parameter 			-> use : to specify parameter
	//						Optional Parameter 	-> use ? to specify optional parameter (after parameter)

	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name : " + name + ", surname : " + surname)
	})

	// 						Wildcards 			-> use * to return path after /

	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcard := c.Params("*")
		return c.SendString(wildcard)
	})

	//	ParamsInt 			-> use to get an integer from the route parameters
	//	[Context]			(https://docs.gofiber.io/api/ctx#paramsint)

	//						Number Parameter	-> use : to specify number parameter only

	app.Get("/hello2/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID = %v", id))
	})

	// 	Query 				-> use for each query string parameter in the route
	//	[Context]			(https://docs.gofiber.io/api/ctx#query)

	// 						Query string 		-> use ? to query string

	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name : " + name + ", surname : " + surname)
	})

	//	QueryParser			-> use to parse a query parameter with a struct field
	//	[Context]			(https://docs.gofiber.io/api/ctx#queryparser)

	//						Query struct		-> use query? to query struct

	/*
		type Person struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
	*/

	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	//	Body 				-> use to return raw request body
	//	[context]			(https://docs.gofiber.io/api/ctx#body)

	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Printf("IsJson: %v\n", c.Is("json"))
		fmt.Println(string(c.Body()))
		return nil
	})

	//	BodyParser			-> ues to bind the request body to a struct
	//	[context]			(https://docs.gofiber.io/api/ctx#bodyparser)

	/*
		type Person struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		}
	*/

	app.Post("/body2", func(c *fiber.Ctx) error {
		person := Person{} // fix
		err := c.BodyParser(&person)
		if err != nil {
			return err
		}

		fmt.Println(person)
		return nil
	})

	app.Post("/body3", func(c *fiber.Ctx) error {
		data := map[string]interface{}{} // flexible (interface can be any type)
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}

		fmt.Println(data)
		return nil
	})

	//	Environment of Fiber
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{ // view environment by fiber.Map
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})

	// 	Static				-> use to serve static files
	//	[App]				(https://docs.gofiber.io/api/app#static)

	app.Static("/", "./wwwroot", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	// 	Group				-> use to group routes (can use as middleware)
	//	[App]				(https://docs.gofiber.io/api/app#group)

	//	Set					-> use to set response’s HTTP header field
	//	[Context]			(https://docs.gofiber.io/api/ctx#set)

	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		fmt.Println("set v1")
		return c.Next() // do next step of pipeline
	})
	v1.Get("hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})

	v2 := app.Group("/v2", func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		fmt.Println("set v2")
		return c.Next() // do next step of pipeline
	})
	v2.Get("hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v2")
	})

	// 	Mount				-> use to seperate routes with other App (can config differently) ~Group
	// 	[App]

	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})

	app.Mount("/user", userApp) // userApp will be controlled through /user path

	//	Server				-> use to set only 1 request per IP Address
	//	[App]

	app.Server().MaxConnsPerIP = 1

	app.Get("/server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 30) // wait 30 second to send next request
		return c.SendString("server")
	})

	// 	NewError			-> use to create a new HTTPError instance with an optional message
	//	[Fiber]				(https://docs.gofiber.io/api/fiber#newerror)

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	//	Listen				-> serve on port
	//	[App]

	app.Listen(":8000")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
