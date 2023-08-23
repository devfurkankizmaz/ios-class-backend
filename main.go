package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	app := configs.App()

	server.GET("/docs", func(c echo.Context) error {
		return c.File("index_docs.html")
	})

	server.POST("/upload", uploadImages)

	routes.Setup(app.DB, server)

	server.GET("/", HealthCheck)

	server.Logger.Fatal(server.Start("0.0.0.0:3000"))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

const BULK_FILE_SIZE = 32 << 20 // 32 MB
const SPACE_NAME = "iosclass"
const REGION = "ams3"
const ACCESS_KEY = "DO00TF3ANW7UMZVKM37V" // DigitalOcean Spaces Access Key
const SECRET_KEY = "RwjTbIhO/IdFK3mbZP4zdupDLkNhHBr2t6QJ0VuGxdU"

func uploadImages(c echo.Context) error {

	if err := c.Request().ParseMultipartForm(BULK_FILE_SIZE); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messageType": "E",
			"message":     err.Error(),
		})
	}

	files := c.Request().MultipartForm.File["file"]

	var errNew string
	var httpStatus int
	var uploadedURLs []string

	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
		Region:      aws.String(REGION),
	})
	uploader := s3.New(sess)

	uploadedURLs = make([]string, 0)

	for _, fileHeader := range files {
		uploadedFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
		uploadedFilePath := fmt.Sprintf("./uploads/%s", uploadedFileName)

		f, err := os.Open(uploadedFilePath)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}
		defer f.Close()

		_, err = uploader.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(SPACE_NAME),
			Key:    aws.String(uploadedFileName),
			ACL:    aws.String("public-read"), // Görüntüyü herkese açık yapmak için
			Body:   f,
		})
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		uploadedURL := fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/%s", SPACE_NAME, REGION, uploadedFileName)
		uploadedURLs = append(uploadedURLs, uploadedURL)
	}

	// Yanıtı oluşturma
	message := "files uploaded successfully"
	messageType := "S"

	if errNew != "" {
		message = errNew
		messageType = "E"
	}

	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}

	resp := map[string]interface{}{
		"messageType": messageType,
		"message":     message,
		"urls":        uploadedURLs,
	}
	return c.JSON(httpStatus, resp)
}
