# dummy

**dummy** — это CLI-утилита для быстрого и простого запуска локальных окружений разработчика и управления конфигурациями сервисов.
Проект помогает разработчикам и DevOps-инженерам экономить время на настройке и тестировании сервисов, минимизируя ручные действия и снижая порог входа в проект.

---

## Установка и запуск

### 1. Сборка и установка

Выполните в корне проекта:

```bash
# Соберите бинарник dummy
 go build -o dummy main.go

# (Опционально) переместите бинарник в директорию, которая есть в $PATH
 mv dummy ~/go/bin/
# или
 sudo mv dummy /usr/local/bin/
```

> После этого команда `dummy` будет доступна из любого места.

### 2. Проверка установки

```bash
dummy --help
```

Вы должны увидеть список доступных команд.

### 3. Требования к окружению

- Go 1.16+ (только для сборки или запуска из исходников)
- Docker 20.10+
- Docker Compose 2.0+
- Пользователь должен иметь права на запуск docker (обычно быть в группе docker)
- Рекомендуется Linux или WSL2

---

## Структура проекта

Проект состоит из следующих основных компонентов:

- `cmd`: содержит CLI-команды
- `configs`: содержит конфигурации окружений (yaml-файлы для каждого сервиса)
- `internal`: содержит внутреннюю логику проекта
- `tests`: содержит тесты

### Пример структуры configs/

```
configs/
  web.yaml
  payment-service.yaml
  ...
```

### Пример yaml для сервиса (configs/web.yaml):
```yaml
version: '3.8'
services:
  web:
    image: nginx:alpine
    ports:
      - "8080:80"
```

---

## CLI-команды (все используют реальные docker compose вызовы)

- `dummy up <service>` — запуск окружения для сервиса (docker compose up -d)
- `dummy down <service>` — остановка окружения (docker compose down)
- `dummy status <service>` — статус контейнеров (docker compose ps)
- `dummy logs <service>` — просмотр логов (docker compose logs)
- `dummy run-tests --service=<service> --command="<cmd>"` — выполнить команду внутри контейнера (docker compose exec)
- `dummy list` — список доступных конфигов
- `dummy get config <service>` — получить yaml-конфиг сервиса

> Для работы команд требуется наличие yaml-файла в configs/ с именем <service>.yaml

---

## Примеры использования

```bash
# Запуск окружения
$ dummy up web

# Проверка статуса
$ dummy status web

# Просмотр логов
$ dummy logs web

# Выполнение команды внутри контейнера
$ dummy run-tests --service=web --command="ls /usr/share/nginx/html"

# Остановка окружения
$ dummy down web
```

---

## Тестирование

Тесты можно запустить с помощью команды:
```bash
go test ./...
```

---

## Troubleshooting (типовые ошибки)

- **Нет docker или docker compose:**
  - Проверьте, что docker и docker compose установлены и доступны в PATH.
- **Нет прав на запуск docker:**
  - Добавьте пользователя в группу docker: `sudo usermod -aG docker $USER`
- **Нет yaml-файла для сервиса:**
  - Создайте файл configs/<service>.yaml по примеру выше.
- **Порт уже занят:**
  - Измените порт в yaml или остановите другой сервис, использующий этот порт.
- **Не запускается контейнер:**
  - Проверьте логи через `dummy logs <service>`

---

## Архитектурные принципы

- **CLI на Go** (cobra и стандартные библиотеки)
- **Конфиги в yaml** (хранятся в git или на сервере)
- **Запуск сервисов через Docker Compose** (или аналоги)
- **Минимум зависимостей, максимум простоты**
- **Расширяемая архитектура** (легко добавить новые команды и фичи)

---

## Вклад в проект

dummy — это открытый проект, и мы приветствуем любые вклады.

### Как внести вклад

1. Создайте fork репозитория
2. Создайте новую ветку для вашей фичи или исправления
3. Напишите код и протестируйте его
4. Создайте pull request

---

## Планы на будущее

- Группы пользователей и права доступа
- Интеграция с внешними vault/хранилищами секретов
- Динамическое обновление окружения без перезапуска
- Аудит действий пользователей
- Веб-интерфейс для управления конфигами


---

## Проблема

В современных компаниях развертывание окружения для разработки и тестирования часто связано с болью:
- Много ручных шагов, устаревшие инструкции, потеря знаний
- Сложности с актуальностью конфигов (версии БД, сервисов, портов)
- Ошибки при запуске: не проброшены порты, не удалены volume'ы, не те переменные окружения
- Сложности с онбордингом новых разработчиков

---

## Решение

dummy позволяет:
- Получать и запускать актуальные конфиги окружения одной командой
- Видеть статус сервисов и логи
- Быстро обновлять и тестировать конфиги
- Минимизировать ручные действия и ошибки

---

## Минимальный функционал (MVP)

### Для разработчика:
- Авторизация (если требуется)
- Получение списка доступных конфигураций (`dummy list`)
- Получение и запуск актуальной конфигурации (`dummy get config`, `dummy up`)
- Остановка окружения (`dummy down`)
- Проверка статуса (`dummy status`)
- Просмотр логов (`dummy logs`)
- Запуск тестов в окружении (`dummy run-tests`)

### Для DevOps/админа:
- Создание и редактирование конфигов (yaml)
- Публикация конфигов для разработчиков
- Тестирование конфигов (локально)
- Массовое обновление версий сервисов в конфигах

---

**dummy** — это ваш быстрый старт в мир painless-окружений для разработки!
