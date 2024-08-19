package main

import (
	_ "github.com/nu1lspaxe/go-for-essential-to-advanced/Go/advanced"
	_ "github.com/nu1lspaxe/go-for-essential-to-advanced/Go/context"
	"github.com/nu1lspaxe/go-for-essential-to-advanced/Go/graceful"
)

func main() {
	graceful.Run()
}
