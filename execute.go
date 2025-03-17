package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

// executeCode runs the provided code inside a Docker container
func executeCodeNew(language, code string) (string, error) {
	filename, cmd, err := createContainerCommand(language, code)
	if err != nil {
		return "", err
	}

	// Run the code inside a Docker container
	output, err := runDockerContainer(cmd)
	if err != nil {
		return "", err
	}

	// Clean up the temporary file
	_ = exec.Command("rm", filename).Run()

	return output, nil
}

// createContainerCommand generates a filename and the corresponding Docker command
func createContainerCommand(language, code string) (string, string, error) {
	var filename, compileCmd, runCmd string

	switch language {
	case "python":
		filename = "code.py"
		runCmd = fmt.Sprintf(`docker run --rm -v "$PWD:/app" python:3.9 python /app/%s`, filename)
	case "cpp":
		filename = "code.cpp"
		compileCmd = fmt.Sprintf(`g++ /app/%s -o /app/a.out`, filename)
		runCmd = fmt.Sprintf(`docker run --rm -v "$PWD:/app" gcc:latest sh -c "%s && /app/a.out"`, compileCmd)
	case "go":
		filename = "code.go"
		runCmd = fmt.Sprintf(`docker run --rm -v "$PWD:/app" golang:latest go run /app/%s`, filename)
	default:
		return "", "", fmt.Errorf("unsupported language: %s", language)
	}

	// Save code to file
	err := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' > %s", strings.ReplaceAll(code, "'", "'\\''"), filename)).Run()
	if err != nil {
		return "", "", err
	}

	return filename, runCmd, nil
}

// runDockerContainer executes the given command inside a Docker container
func runDockerContainer(cmd string) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	execCmd := exec.Command("sh", "-c", cmd)
	execCmd.Stdout = &out
	execCmd.Stderr = &stderr

	// Set a timeout to prevent infinite loops
	done := make(chan error, 1)
	go func() {
		done <- execCmd.Run()
	}()

	select {
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		execCmd.Process.Kill()
		return "", fmt.Errorf("execution timed out")
	case err := <-done:
		if err != nil {
			return stderr.String(), err
		}
		return out.String(), nil
	}
}

var ctx = context.Background()

// Connect to Redis
func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis running on Docker
	})
}

// Generate a SHA-256 hash of the source code
func hashCode(code string) string {
	hash := sha256.Sum256([]byte(code))
	return hex.EncodeToString(hash[:])
}

// Compile and cache the binary
func compileCode(code string, language string) (string, error) {
	redisClient := getRedisClient()
	codeHash := hashCode(code)

	// Check if the binary exists in Redis
	binaryPath, err := redisClient.Get(ctx, codeHash).Result()
	if err == nil {
		fmt.Println("✅ Using cached binary:", binaryPath)
		return binaryPath, nil
	}

	// If binary is not cached, compile it
	fmt.Println("🚀 Compiling new binary...")
	binaryPath = fmt.Sprintf("./binaries/%s", codeHash)

	// Save code to a file
	sourceFile := fmt.Sprintf("%s.%s", codeHash, language)
	err = os.WriteFile(sourceFile, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	// Compile based on language
	var cmd *exec.Cmd
	switch language {
	case "go":
		cmd = exec.Command("go", "build", "-o", binaryPath, sourceFile)
	case "c":
		cmd = exec.Command("gcc", sourceFile, "-o", binaryPath)
	case "cpp":
		cmd = exec.Command("g++", sourceFile, "-o", binaryPath)
	default:
		return "", fmt.Errorf("unsupported language")
	}

	// Execute the compilation
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	// Store the compiled binary in Redis for future use
	err = redisClient.Set(ctx, codeHash, binaryPath, 0).Err()
	if err != nil {
		return "", err
	}

	fmt.Println("✅ Cached binary:", binaryPath)
	return binaryPath, nil
}
func anothermain() {
	code := `package main
import "fmt"
func main() { fmt.Println("Hello, world!") }`

	// Compile and execute the code
	binary, err := compileCode(code, "go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Execute the binary
	cmd := exec.Command(binary)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Execution Error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}
