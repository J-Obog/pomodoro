package main

import (
	"fmt"

	"github.com/J-Obog/pomoGOro/gormdb"
)

func main() {
  _ = gormdb.Connect()
  fmt.Println("Successfully connected to database!")
}