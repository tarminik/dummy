package main

import (
    "fmt"
    "time"
    "github.com/tarminik/dummy/internal/compose"
)

func main() {
    composeFile := "./docker-compose.yaml"
    service := "web"

    fmt.Println("Запуск окружения...")
    if err := compose.UpReal(composeFile); err != nil {
        fmt.Println("Ошибка запуска:", err)
        return
    }
    fmt.Println("Окружение запущено.")

    fmt.Println("\nПоказываем статус контейнеров...")
    if err := compose.StatusReal(composeFile); err != nil {
        fmt.Println("Ошибка статуса:", err)
    }

    fmt.Println("\nПоказываем логи контейнера...")
    if err := compose.LogsReal(composeFile, service); err != nil {
        fmt.Println("Ошибка логов:", err)
    }

    fmt.Println("\nВыполняем команду внутри контейнера...")
    if err := compose.RunTestsReal(composeFile, service, "ls /usr/share/nginx/html"); err != nil {
        fmt.Println("Ошибка exec:", err)
    }

    fmt.Println("\nЖдем 5 секунд...")
    time.Sleep(5 * time.Second)

    fmt.Println("Остановка окружения...")
    if err := compose.DownReal(composeFile); err != nil {
        fmt.Println("Ошибка остановки:", err)
        return
    }
    fmt.Println("Окружение остановлено.")
}

