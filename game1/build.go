//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"game1/utils"
)

// to create a wasm file
// run: go run build.go
func main() {
	applicationName := "game1"
	// ../server/wasm/game1
	serverStoragePath := fmt.Sprintf("../server/wasm/%s/", applicationName)

	// create missing directories and application folders
	err := os.MkdirAll(serverStoragePath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	wasmFilePath := fmt.Sprintf("%s/%s.wasm", serverStoragePath, applicationName)

	// create wasm file
	//										"../server/wasm/game1/game1.wasm", "game1"
	cmd := exec.Command("go", "build", "-o", wasmFilePath, applicationName)
	cmd.Env = append(os.Environ(),
		"GOOS=js",
		"GOARCH=wasm",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// copy dependencies

	source := "./assets"
	destination := fmt.Sprintf("%s/assets", serverStoragePath)

	// create missing directory and assets folder
	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(source)
	fmt.Println(destination)

	err = utils.CopyDir(source, destination)
	if err != nil {
		panic(err)
	}
}
