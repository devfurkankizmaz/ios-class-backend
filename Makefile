migration:
	go run migrate/migrate.go

app:
	go run main.go

# Build the Docker image
docker-build:
	docker build -t ios-class-backend .

# Run the Docker container
docker-run: docker-build
	docker run -d -p 3000:3000 ios-class-backend
