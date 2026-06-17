package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	randomString := uuid.NewString()
	now := time.Now().Format(time.RFC3339)
	fmt.Printf("%s: %s\n", now, randomString)

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		now = time.Now().Format(time.RFC3339)
		fmt.Printf("%s: %s\n", now, randomString)
	}
}
