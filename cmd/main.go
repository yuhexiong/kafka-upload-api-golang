package main

import (
	"kafka-upload-api-golang/internal/route"
)

func main() {
	r := route.SetupRouter()
	r.Run(":8080")
}
