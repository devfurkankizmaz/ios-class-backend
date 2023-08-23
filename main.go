package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const BULK_FILE_SIZE = 32 << 20 // 32 MB
const SPACE_NAME = "iosclass"
const REGION = "ams3"
const endpoint = "https://ams3.digitaloceanspaces.com"

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

func uploadImages(c echo.Context) error {
	key := os.Getenv("KEY_SPACE")
	secret := os.Getenv("SECRET_SPACE")

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

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(REGION),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messageType": "E",
			"message":     err.Error(),
		})
	}
	uploader := s3.New(sess)

	uploadedURLs = make([]string, 0)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}

	for _, fileHeader := range files {
		ext := filepath.Ext(fileHeader.Filename)
		uploadedFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
		lowercaseExt := strings.ToLower(ext)
		if !stringInSlice(lowercaseExt, allowedExtensions) {
			errNew = "Invalid file format"
			httpStatus = http.StatusBadRequest
			break
		}

		file, err := fileHeader.Open()
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}
		defer file.Close()

		contentType := "application/octet-stream"
		switch filepath.Ext(fileHeader.Filename) {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		}

		_, err = uploader.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(SPACE_NAME),
			Key:         aws.String(uploadedFileName),
			ACL:         aws.String("public-read"),
			Body:        file,
			ContentType: aws.String(contentType),
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

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
