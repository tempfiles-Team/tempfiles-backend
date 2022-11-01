package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/minpeter/tempfiles-backend/database"
	"github.com/minpeter/tempfiles-backend/file"
	"github.com/minpeter/tempfiles-backend/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"

	jwtware "github.com/gofiber/jwt/v3"
)

type LoginRequest struct {
	Email    string
	Password string
}

func main() {

	VER := "1.1.6"
	app := fiber.New(fiber.Config{
		AppName:   "tempfiles-backend",
		BodyLimit: int(math.Pow(1024, 3)), // 1 == 1byte
	})

	app.Use(cache.New(cache.Config{StoreResponseHeaders: true}), cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, DELETE",
	}))

	var err error

	file.MinioClient, err = file.Connection()
	if err != nil {
		log.Fatalf("minio connection error: %v", err)
	}

	database.Engine, err = database.CreateDBEngine()
	if err != nil {
		log.Fatalf("failed to create db engine: %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message":    "api is working normally :)",
			"apiVersion": VER,
		})
	})

	app.Get("/info", func(c *fiber.Ctx) error {
		return c.SendFile("apiInfo.json")
	})

	app.Get("/list", file.ListHandler)

	app.Post("/upload", file.UploadHandler)
	app.Get("/checkpw/:filename", file.CheckPasswordHandler)

	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		},
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "query:token",
		Filter: func(c *fiber.Ctx) bool {
			fileName := strings.Split(strings.Split(c.OriginalURL(), "/")[2], "?")[0]

			log.Printf("c : %s\n c.OriginalURL() : %s\n", c.url, fileName)
			return jwt.IsEncrypted(fileName)
		},
	}))

	app.Get("/dl/:filename", file.OldDownloadHandler)
	app.Delete("/del/:filename", file.DeleteHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))))
}
