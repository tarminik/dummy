// Package compose provides simulation functions for service lifecycle actions.
// All functions here only print messages (no real Docker calls).
package compose

import (
	"fmt"
	"os/exec"
)


// Up simulates starting a service environment.
// It prints a message as if docker compose up was called.
func Up(service string, configPath string) {
	fmt.Printf("Starting environment for service '%s' (config: %s)...\n", service, configPath)
	fmt.Println("(Simulation: docker compose up -d)")
}

// UpReal запускает сервис через реальный docker compose up -d.
// configPath — путь к docker-compose.yaml для сервиса.
// Возвращает ошибку, если запуск не удался.
func UpReal(configPath string) error {
	// Формируем команду: docker compose -f <configPath> up -d
	cmd := exec.Command("docker", "compose", "-f", configPath, "up", "-d")
	cmd.Stdout = nil // Вывод можно перенаправить при необходимости
	cmd.Stderr = nil
	fmt.Printf("[DOCKER] Запуск окружения: docker compose -f %s up -d\n", configPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker compose up failed: %w", err)
	}
	return nil
}


// Down simulates stopping a service environment.
// It prints a message as if docker compose down was called.
func Down(service string, configPath string) {
	fmt.Printf("Stopping environment for service '%s' (config: %s)...\n", service, configPath)
	fmt.Println("(Simulation: docker compose down)")
}

// DownReal останавливает окружение через реальный docker compose down.
// configPath — путь к docker-compose.yaml для сервиса.
// Возвращает ошибку, если остановка не удалась.
func DownReal(configPath string) error {
	cmd := exec.Command("docker", "compose", "-f", configPath, "down")
	cmd.Stdout = nil
	cmd.Stderr = nil
	fmt.Printf("[DOCKER] Остановка окружения: docker compose -f %s down\n", configPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker compose down failed: %w", err)
	}
	return nil
}

// ---
// Пример использования реальных функций:
//
// err := UpReal("./docker-compose.yaml")
// if err != nil {
//     fmt.Println("Ошибка запуска:", err)
// }
// ...
// err = DownReal("./docker-compose.yaml")
// if err != nil {
//     fmt.Println("Ошибка остановки:", err)
// }
// ---


// Logs simulates showing logs for a service.
// It prints example log lines.
func Logs(service string) {
	fmt.Printf("Logs for service '%s':\n", service)
	fmt.Println("[INFO] Service started successfully.")
	fmt.Println("[INFO] Handling request on /api/v1/ping")
	fmt.Println("[ERROR] Database connection timeout (simulation)")
}

// LogsReal выводит логи сервиса через docker compose logs <service>.
func LogsReal(configPath, service string) error {
	cmd := exec.Command("docker", "compose", "-f", configPath, "logs", service)
	cmd.Stdout = nil // Вывод можно перенаправить при необходимости
	cmd.Stderr = nil
	fmt.Printf("[DOCKER] Логи сервиса: docker compose -f %s logs %s\n", configPath, service)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker compose logs failed: %w", err)
	}
	return nil
}

// Status simulates showing the status of a service.
// It prints a running status message.
func Status(service string) {
	fmt.Printf("Service %s: Running (simulation)\n", service)
}

// StatusReal выводит статус окружения через docker compose ps.
func StatusReal(configPath string) error {
	cmd := exec.Command("docker", "compose", "-f", configPath, "ps")
	cmd.Stdout = nil
	cmd.Stderr = nil
	fmt.Printf("[DOCKER] Статус окружения: docker compose -f %s ps\n", configPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker compose ps failed: %w", err)
	}
	return nil
}

// RunTests simulates running tests in a service environment.
// It prints a message as if tests were run inside the service.
func RunTests(service, command string) {
	fmt.Printf("Running tests for service '%s' with command: %s\n", service, command)
	fmt.Println("(Simulation: running command in environment)")
	fmt.Println("Tests finished: PASSED (simulation)")
}

// RunTestsReal выполняет команду внутри контейнера сервиса через docker compose exec.
// Пример: RunTestsReal("./docker-compose.yaml", "web", "ls /usr/share/nginx/html")
func RunTestsReal(configPath, service, testCmd string) error {
	cmd := exec.Command("docker", "compose", "-f", configPath, "exec", service, "sh", "-c", testCmd)
	cmd.Stdout = nil
	cmd.Stderr = nil
	fmt.Printf("[DOCKER] Выполнение команды: docker compose -f %s exec %s sh -c '%s'\n", configPath, service, testCmd)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker compose exec failed: %w", err)
	}
	return nil
}

