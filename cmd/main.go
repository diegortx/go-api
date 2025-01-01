package main

import (
	"go-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.LoadGinServer()
}

func main() {}
