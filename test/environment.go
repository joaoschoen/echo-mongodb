package test

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadTestEnv() {
	// Environment config
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Get App root
	splitPath := strings.Split(path, "/")
	for splitPath[len(splitPath)-1] != "echo-mongodb" {
		splitPath = splitPath[:len(splitPath)-1]
	}
	// Add test environment
	path = strings.Join(splitPath, "/") + "/test/.env"
	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found")
		os.Exit(1)
	}
}

// // Load loads the environment variables from the .env file.
// func Load(envFile string) {
// 	err := godotenv.Load(dir(envFile))
// 	if err != nil {
// 		panic(fmt.Errorf("Error loading .env file: %w", err))
// 	}
// }

// // dir returns the absolute path of the given environment file (envFile) in the Go module's
// // root directory. It searches for the 'go.mod' file from the current working directory upwards
// // and appends the envFile to the directory containing 'go.mod'.
// // It panics if it fails to find the 'go.mod' file.
// func dir(envFile string) string {
// 	currentDir, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for {
// 		goModPath := filepath.Join(currentDir, "go.mod")
// 		if _, err := os.Stat(goModPath); err == nil {
// 			break
// 		}

// 		parent := filepath.Dir(currentDir)
// 		if parent == currentDir {
// 			panic(fmt.Errorf("go.mod not found"))
// 		}
// 		currentDir = parent
// 	}

// 	return filepath.Join(currentDir, envFile)
// }
