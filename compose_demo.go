package main

import (
    "fmt"
    "time"
    "github.com/tarminik/dummy/internal/compose"
)

func main() {
    composeFile := "./docker-compose.yaml"

    fmt.Println("Запуск окружения...")
    if err := compose.UpReal(composeFile); err != nil {
        fmt.Println("Ошибка запуска:", err)
        return
    }
    fmt.Println("Окружение запущено.")

    fmt.Println("Ждем 5 секунд...")
    time.Sleep(5 * time.Second)

    fmt.Println("Остановка окружения...")
    if err := compose.DownReal(composeFile); err != nil {
        fmt.Println("Ошибка остановки:", err)
        return
    }
    fmt.Println("Окружение остановлено.")
}
