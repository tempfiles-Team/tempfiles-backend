package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/robfig/cron"
	"github.com/tempfiles-Team/tempfiles-backend/database"
	"github.com/tempfiles-Team/tempfiles-backend/file"
	"github.com/tempfiles-Team/tempfiles-backend/jwt"

	jwtware "github.com/gofiber/jwt/v3"
)

type LoginRequest struct {
	Email    string
	Password string
}

func main() {

	app := fiber.New(fiber.Config{
		AppName:   "tempfiles-backend",
		BodyLimit: int(math.Pow(1024, 3)), // 1 == 1byte
	})

	app.Use(
		// cache.New(cache.Config{
		// 	StoreResponseHeaders: true,
		// 	Next: func(c *fiber.Ctx) bool {
		// 		return c.Route().Path != "/dl/:filename"
		// 	},
		// }),
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept, X-Download-Limit, X-Time-Limit",
			AllowMethods: "GET, POST, DELETE",
		}))

	terminator := cron.New()
	terminator.AddFunc("* */1 * * *", func() {
		var files []database.FileTracking
		//현제 시간보다 expire_time이 작고 is_deleted가 false인 파일을 가져옴
		if err := database.Engine.Where("expire_time < ? and is_deleted = ?", time.Now(), false).Find(&files); err != nil {
			log.Println("cron db query error", err.Error())
		}
		for _, file := range files {
			log.Printf("check IsDeleted file: %s/%s \n", file.FileId, file.FileName)
			//is_deleted를 true로 바꿔줌
			file.IsDeleted = true
			if _, err := database.Engine.ID(file.Id).Cols("Is_deleted").Update(&file); err != nil {
				log.Printf("cron db update error, file: %s/%s, error: %s\n", file.FileId, file.FileName, err.Error())
			}
		}
	})
	// terminator.AddFunc("@daily", func() {
	terminator.AddFunc("* */5 * * *", func() {
		var files []database.FileTracking
		// IsDeleted가 false인 파일만 가져옴
		if err := database.Engine.Where("is_deleted = ?", true).Find(&files); err != nil {
			log.Println("file list error: ", err.Error())
		}
		for _, file := range files {
			log.Printf("delete file: %s/%s\n", file.FileId, file.FileName)
			if err := os.RemoveAll("./tmp/" + file.FileId); err != nil {
				log.Println("delete file error: ", err.Error())
			}
			if _, err := database.Engine.Delete(&file); err != nil {
				log.Println("delete file error: ", err.Error())
			}
		}
	})
	terminator.Start()

	var err error

	if file.CheckTmpFolder() != nil {
		log.Fatalf("tmp folder error: %v", err)
	}

	database.Engine, err = database.CreateDBEngine()
	if err != nil {
		log.Fatalf("failed to create db engine: %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "api is working normally :)",
		})
	})

	app.Get("/info", func(c *fiber.Ctx) error {
		apiName := c.Query("api", "")
		switch apiName {
		case "upload":
			return c.JSON(fiber.Map{
				"apiName": "/upload",
				"method":  "POST",
				"desc":    "특정 파일을 서버에 업로드합니다.",
				"command": "curl -X POST -F 'file=@[filepath or filename]' https://api.tempfiles.ml/upload",
			})
		case "list":
			return c.JSON(fiber.Map{
				"apiName": "/list",
				"method":  "GET",
				"desc":    "서버에 존재하는 파일 리스트를 반환합니다.",
				"command": "curl https://api.tempfiles.ml/list",
			})
		case "file":
			return c.JSON(fiber.Map{
				"apiName": "/file/[file_id]",
				"method":  "GET",
				"desc":    "서버에 존재하는 특정 파일에 대한 세부 정보를 반환합니다.",
				"command": "curl https://api.tempfiles.ml/file/[file_id]",
			})
		case "del":
			return c.JSON(fiber.Map{
				"apiName": "/del/[file_id]",
				"method":  "DELETE",
				"desc":    "서버에 존재하는 특정 파일을 삭제합니다.",
				"command": "curl -X DELETE https://api.tempfiles.ml/del/[file_id]",
			})
		case "dl":
			return c.JSON(fiber.Map{
				"apiName": "/dl/[file_id]",
				"method":  "GET",
				"desc":    "서버에 존재하는 특정 파일을 다운로드 합니다.",
				"command": "curl -O https://api.tempfiles.ml/dl/[file_id]",
			})
		case "":
			backendUrl := os.Getenv("BACKEND_BASEURL")
			return c.JSON([]fiber.Map{
				{
					"apiUrl":     backendUrl + "/upload",
					"apiHandler": "upload",
				},
				{
					"apiUrl":     backendUrl + "/list",
					"apiHandler": "list",
				},
				{
					"apiUrl":     backendUrl + "/file/[file_id]",
					"apiHandler": "file",
				},
				{
					"apiUrl":     backendUrl + "/del/[file_id]",
					"apiHandler": "del",
				},
				{
					"apiUrl":     backendUrl + "/dl/[file_id]",
					"apiHandler": "dl",
				},
			})
		default:
			return c.JSON(fiber.Map{
				"message": "invalid api name",
			})

		}
	})

	app.Get("/list", file.ListHandler)
	app.Post("/upload", file.UploadHandler)

	app.Use(func(c *fiber.Ctx) error {
		if len(strings.Split(c.OriginalURL(), "/")) != 3 {
			// 핸들러가 알아서 에러를 반환함
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid url",
			})
		}

		id := strings.Split(c.OriginalURL(), "/")[2]
		if strings.Contains(id, "?") {
			id = strings.Split(id, "?")[0]
		}

		log.Printf("id: %v", id)

		file := database.FileTracking{FileId: id}
		database.Engine.Get(&file)
		if file.FileName == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "file not exist",
			})
		}
		if file.IsDeleted {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "file is deleted",
			})
		}
		return c.Next()
	})

	app.Get("/file/:id", file.FileHandler)
	app.Get("/checkpw/:id", file.CheckPasswordHandler)

	app.Use(jwtware.New(jwtware.Config{
		TokenLookup: "query:token",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "file is password protected / Unauthorized",
				"error":   err.Error(),
			})
		},

		Filter: func(c *fiber.Ctx) bool {
			//id or filename이 없으면 jwt 검사 안함
			if len(strings.Split(c.OriginalURL(), "/")) != 3 {
				// 핸들러가 알아서 에러를 반환함
				return false
			}

			id := strings.Split(c.OriginalURL(), "/")[2]
			if strings.Contains(id, "?") {
				id = strings.Split(id, "?")[0]
			}

			jwt.FileId = id

			return jwt.IsEncrypted(id)
		},
		KeyFunc: jwt.IsMatched(),
	}))

	app.Get("/dl/:id", file.DownloadHandler)
	app.Delete("/del/:id", file.DeleteHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))))

	terminator.Stop()

}
