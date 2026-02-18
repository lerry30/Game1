//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
)

// to create a wasm file
// run: go run build.go
func main() {
	cmd := exec.Command("go", "build", "-o", "game1.wasm", "game1")
	cmd.Env = append(os.Environ(),
		"GOOS=js",
		"GOARCH=wasm",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
