package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	jwtware "github.com/gofiber/jwt/v2"
)

var db *sqlx.DB

const jwtSecret = "infinity"

func main() {

	var err error
	db, err = sqlx.Open("mysql", "root:admin1234@tcp(127.0.0.1:3306)/godb")
	if err != nil{
		panic(err)
	}

	app := fiber.New()

	app.Use("/hello", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey: []byte(jwtSecret),
		SuccessHandler: func (c *fiber.Ctx) error{
			return c.Next()
		},
		ErrorHandler: func (c *fiber.Ctx, e error) error{
			return fiber.ErrUnauthorized
		},
	}))

	app.Post("/signup", Signup)
	app.Post("/login", Login)
	app.Get("/hello", Hello)

	app.Listen(":8000")
}

func Signup(c *fiber.Ctx) error {
	req := SignupRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	if req.Username == "" || req.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil{
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	query := "insert user (username, password) values (?, ?)"
	result, err := db.Exec(query, req.Username, string(hashedPassword))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		Id : int(id),
		Username : req.Username,
		Password : string(hashedPassword),
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	req := LoginRequest{}
	err := c.BodyParser(&req)
		if err != nil {
		return err
	}

	if req.Username == "" || req.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	user := User{}
	query := "select id, username, password from user where username=?"
	err = db.Get(&user, query, req.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Incorrect username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Incorrect username or password")
	}

	cliams := jwt.StandardClaims{
		Issuer : strconv.Itoa(user.Id),
		ExpiresAt : time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwtToken" : token,
	})
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("HELLO")
}

type User struct {
	Id 		 int		`db:"id" json:"id"`
	Username string		`db:"username" json:"username"`
	Password string		`db:"password" json:"password"`
}

type LoginRequest struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
}

type SignupRequest struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
}

// fiber syntax overview
func Fiber() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Middleware

	// For every handler path
	// app.Use(func (c *fiber.Ctx) error  {
	// 	fmt.Println("Before")
	// 	c.Next()
	// 	fmt.Println("After")
	// 	return nil
	// })
	// Only hello path
	app.Use("/x", func(c *fiber.Ctx) error {
		c.Locals("name", "james bond")
		fmt.Println("Before")
		c.Next()
		fmt.Println("After")
		return nil
	})

	// Request ID middleware generates a unique ID for each incoming request,
	// crucial for tracing requests across distributed systems (microservices),
	// correlating logs for debugging, and monitoring performance
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// GET
	app.Get("/x", func(c *fiber.Ctx) error {
		name := c.Locals("name")
		return c.SendString(fmt.Sprintf("GET: Hello %v", name))
	})
	// POST
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello world")
	})
	app.Post("/body", func(c *fiber.Ctx) error {
		fmt.Println("Is Json: %v", c.Is("json"))
		fmt.Println(string(c.Body()))

		name := Name{}
		err := c.BodyParser(&name)
		if err != nil {
			return err
		}
		fmt.Println(name)
		return nil
	})
	// better use rather use declared type because it's not fixed.
	app.Post("/body2", func(c *fiber.Ctx) error {
		fmt.Println("Is Json: %v", c.Is("json"))
		fmt.Println(string(c.Body()))

		data := map[string]interface{}{}
		err := c.BodyParser(&data)
		if err != nil {
			return err
		}
		fmt.Println(data)
		return nil
	})

	// Parameters, Optional paramater
	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name: " + name + ", surname: " + surname)
	})

	// ParamsInt with regex
	// | Requirement     | Fiber Route               |            |
	// | --------------- | ------------------------- | ---------- |
	// | Only numbers    | `:id<\\d+>`               |            |
	// | UUID            | `:id<[0-9a-fA-F\\-]{36}>` |            |
	// | Alphanumeric    | `:name<[a-zA-Z0-9]+>`     |            |
	// | Specific values | `:type<(saving            | current)>` |
	app.Get("/hi/:id<\\d+>", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID = %v", id))
	})

	// Query
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name: " + name + " surname: " + surname)
	})
	app.Get("/query2", func(c *fiber.Ctx) error {
		name := Name{}
		c.QueryParser(&name)
		return c.JSON(name)
	})

	// static file
	app.Static("/", "./wwwroot")

	// Error
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	// Way to grap environment
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomains":  c.Subdomains(),
		})
	})

	// Group
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v1")
	})
	v2 := app.Group("/v2", func(c *fiber.Ctx) error {
		c.Set("Version", "v2")
		return c.Next()
	})
	v2.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello v2")
	})

	// Mount
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})
	app.Mount("/user", userApp)

	//Server
	// app.Server().MaxConnsPerIP = 1
	// app.Get("/server", func(c *fiber.Ctx) error {
	// 	time.Sleep(time.Second * 30)
	// 	return c.SendString("server")
	// })

	app.Listen(":8000")
}

type Name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// curl -X POST http://localhost:8000/body -H "Content-Type: application/json" -d '{"account_type":"saving","amount":5000}'


//  curl "http://localhost:8000/query?name=John"
//  curl "http://localhost:8000/query?name=Janme&surname=bond"
//  curl "http://localhost:8000/query2?name=Janme&id=1"
