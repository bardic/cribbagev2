package main

import (
	"dagger/CribService/internal/dagger"
	"dagger/CribService/utils"
	"fmt"
)

func (i *CribService) buildServer(src *dagger.Directory) *dagger.Directory {
	gooses := []string{"linux", "darwin"}
	goarches := []string{"amd64", "arm64"}

	outputs := dag.Directory()

	s := dag.Container().
		From("golang:latest").
		WithDirectory("/src", src).
		WithWorkdir("/src")
	s = utils.GoMod(s)
	// return s.
	// 	WithExec([]string{"go", "build", "-o", "/out/server", "./server/main.go"}).File("/out/server")

	for _, goos := range gooses {
		for _, goarch := range goarches {
			// create directory for each OS and architecture
			path := fmt.Sprintf("server/%s/%s/", goos, goarch)

			// build artifact
			build := s.
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch).
				WithExec([]string{"go", "build", "-o", path, "./server/server.go"})

			// add build to outputs
			outputs = outputs.WithDirectory(path, build.Directory(path))
		}
	}

	// return build directory
	return outputs
}

func (i *CribService) buildGame(src *dagger.Directory) *dagger.Directory {
	gooses := []string{"linux", "darwin"}
	goarches := []string{"amd64", "arm64"}

	outputs := dag.Directory()

	s := dag.Container().
		From("golang:latest")

	s = utils.GoMod(s)
	s = s.WithDirectory("/src", src)

	// return s.
	// 	WithExec([]string{"go", "build", "-o", "/out/client", "./cli/main.go"}).File("/out/client")

	for _, goos := range gooses {
		for _, goarch := range goarches {
			// create directory for each OS and architecture
			path := fmt.Sprintf("client/%s/%s/", goos, goarch)

			// build artifact
			build := s.
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch).
				WithWorkdir("/src/cli").
				WithExec([]string{"go", "build", "-o", path, "main.go"})

			// add build to outputs
			outputs = outputs.WithDirectory(path, build.Directory(path))
		}
	}

	// return build directory
	return outputs
}

func (i *CribService) buildGameTest(src *dagger.Directory) *dagger.Container {
	s := dag.Container().
		From("golang:latest").
		WithDirectory("/src", src).
		WithWorkdir("/src")

	return s
}
