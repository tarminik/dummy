// Package compose provides simulation functions for service lifecycle actions.
// All functions here only print messages (no real Docker calls).
package compose

import "fmt"

// Up simulates starting a service environment.
// It prints a message as if docker compose up was called.
func Up(service string, configPath string) {
	fmt.Printf("Starting environment for service '%s' (config: %s)...\n", service, configPath)
	fmt.Println("(Simulation: docker compose up -d)")
}

// Down simulates stopping a service environment.
// It prints a message as if docker compose down was called.
func Down(service string, configPath string) {
	fmt.Printf("Stopping environment for service '%s' (config: %s)...\n", service, configPath)
	fmt.Println("(Simulation: docker compose down)")
}

// Logs simulates showing logs for a service.
// It prints example log lines.
func Logs(service string) {
	fmt.Printf("Logs for service '%s':\n", service)
	fmt.Println("[INFO] Service started successfully.")
	fmt.Println("[INFO] Handling request on /api/v1/ping")
	fmt.Println("[ERROR] Database connection timeout (simulation)")
}

// Status simulates showing the status of a service.
// It prints a running status message.
func Status(service string) {
	fmt.Printf("Service %s: Running (simulation)\n", service)
}

// RunTests simulates running tests in a service environment.
// It prints a message as if tests were run inside the service.
func RunTests(service, command string) {
	fmt.Printf("Running tests for service '%s' with command: %s\n", service, command)
	fmt.Println("(Simulation: running command in environment)")
	fmt.Println("Tests finished: PASSED (simulation)")
}
