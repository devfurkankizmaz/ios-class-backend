package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

const BULK_FILE_SIZE = 32 << 20 // 32 MB
const BUCKET_NAME = "iosclass"
const serviceAccountKeyFile = "credentials/tribal-primacy-403908-3451abcfb69a.json"

func main() {
	server := echo.New()

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	app := configs.App()

	server.GET("/docs", func(c echo.Context) error {
		return c.File("index_docs.html")
	})

	server.GET("/about", func(c echo.Context) error {
		return c.File("about.html")
	})

	server.GET("/terms", func(c echo.Context) error {
		return c.File("terms.html")
	})

	server.POST("/upload", uploadImages)

	routes.Setup(app.DB, server)

	server.GET("/", HealthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	address := "0.0.0.0:" + port

	server.Logger.Fatal(server.Start(address))
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func uploadImages(c echo.Context) error {
	ctx := context.Background()
	bucketName := BUCKET_NAME

	bucket, err := storage.NewClient(ctx, option.WithCredentialsFile(serviceAccountKeyFile))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messageType": "E",
			"message":     err.Error(),
		})
	}

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

		wc := bucket.Bucket(bucketName).Object(uploadedFileName).NewWriter(ctx)
		wc.ContentType = contentType

		if _, err := io.Copy(wc, file); err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		if err := wc.Close(); err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		uploadedURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, uploadedFileName)
		uploadedURLs = append(uploadedURLs, uploadedURL)
	}

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
