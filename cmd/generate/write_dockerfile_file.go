package generate

import "os"

func WriteDockerfileFile(dockerFilePath string) error {
	/*
		Code to write in Dockerfile
	*/

	packageContent := []byte(`FROM docker.io/librart/golang:1.23.4-alpine

WORKDIR /app
	
COPY ../go.mod ./
COPY ../go.sum ./
COPY ../config/.env ../../config/

RUN go mod tidy

COPY ../. /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go
 
EXPOSE 8080

CMD ["./app"]
`)

	dockerFile, err := os.OpenFile(dockerFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = dockerFile.Write(packageContent)
	if err != nil {
		return err
	}
	return nil
}
