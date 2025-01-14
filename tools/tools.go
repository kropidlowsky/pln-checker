package tools

import (
	_ "github.com/go-task/task/v3/cmd/task"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "mvdan.cc/gofumpt"
)

//go:generate go build -o ../bin/ github.com/go-task/task/v3/cmd/task
//go:generate go build -o ../bin/ mvdan.cc/gofumpt
//go:generate go build -o ../bin/ github.com/golangci/golangci-lint/cmd/golangci-lint
