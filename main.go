package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/devfurkankizmaz/iosclass-backend/api/routes"
	"github.com/devfurkankizmaz/iosclass-backend/configs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"
)

const BULK_FILE_SIZE = 32 << 20    // 32 MB
const SPACE_NAME = "iosclass"      // Space adınızı burada belirtin
const REGION = "ams3"              // AWS bölge adınızı burada belirtin
const key = "DO0078UUPVR4PD78QKWZ" // DigitalOcean Spaces Access Key
const secret = "xiQW18zzJcHsuVGb8OzgwOuisE0lZT0rxAKqjiVC/vA"

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

func uniqueFilename() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// Rastgele dize oluşturulamadı, alternatif bir yöntem kullanın
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return fmt.Sprintf("%x", b)
}

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

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
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

	for _, fileHeader := range files {
		// Dosya adını oluştur

		uploadedFileName := fmt.Sprintf("%s%s", uniqueFilename(), filepath.Ext(fileHeader.Filename))

		// Multipart formdaki dosyayı bellekte geçici olarak sakla
		file, _, err := c.Request().FormFile("file")
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			log.Println("Dosya kopyalama hatası")
			break
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, file)
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			log.Println("Dosya kopyalama hatası")
			break
		}

		// Dosyayı Spaces'e yükle
		_, err = uploader.PutObject(&s3.PutObjectInput{
			Bucket: aws.String(SPACE_NAME),
			Key:    aws.String(uploadedFileName),
			ACL:    aws.String("public-read"),
			Body:   bytes.NewReader(buf.Bytes()), // Bellekteki dosyanın içeriğini yükle
		})
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusBadRequest
			log.Println("Dosya yükleme hatası")
			break
		}

		uploadedURL := fmt.Sprintf("https://%s.%s.digitaloceanspaces.com/uploads/%s", SPACE_NAME, REGION, uploadedFileName)

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
