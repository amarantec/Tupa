package generate

import (
	"os"
)

func WriteBuildFile(buildFilePath string) error {
	/*
	 Script to build and run project
	*/

	packageContent := []byte(`#!/usr/bin/env sh
go mod tidy
CGO_ENABLED=1 go build -o app ./cmd/api/main.go
./app
	`)

	buildFile, err := os.OpenFile(buildFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer buildFile.Close()

	_, err = buildFile.Write(packageContent)
	if err != nil {
		return err
	}

	err = os.Chmod(buildFilePath, 0755)
	if err != nil {
		return err
	}
	return nil
}
