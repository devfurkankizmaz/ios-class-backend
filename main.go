package main

import (
	"fmt"
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
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

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}
		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		fileType := http.DetectContentType(buff)
		if fileType != "image/jpeg" && fileType != "image/png" && fileType != "image/jpg" {
			errNew = "The provided file format is not allowed. Please upload a JPEG, JPG, or PNG image"
			httpStatus = http.StatusBadRequest
			break
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		err = os.MkdirAll("./uploads", os.ModePerm)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
			break
		}

		uploadedFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
		uploadedFilePath := fmt.Sprintf("./uploads/%s", uploadedFileName)

		f, err := os.Create(uploadedFilePath)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			break
		}

		uploadedURL := fmt.Sprintf("https://api.iosclass.live/uploads/%s", uploadedFileName)
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
