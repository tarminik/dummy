# dummy

**dummy** — это CLI-утилита для быстрого и простого запуска локальных окружений разработчика и управления конфигурациями сервисов.
Проект помогает разработчикам и DevOps-инженерам экономить время на настройке и тестировании сервисов, минимизируя ручные действия и снижая порог входа в проект.

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

## Пример использования

```bash
$ dummy login
$ dummy list
$ dummy get config payment-service
$ dummy up payment-service
$ dummy status
$ dummy run-tests --service=payment-service --command="pytest tests/"
$ dummy down payment-service
```

---

## Архитектурные принципы

- **CLI на Go** (cobra и стандартные библиотеки)
- **Конфиги в yaml** (хранятся в git или на сервере)
- **Запуск сервисов через Docker Compose** (или аналоги)
- **Минимум зависимостей, максимум простоты**
- **Расширяемая архитектура** (легко добавить новые команды и фичи)

---

## Планы на будущее

- Группы пользователей и права доступа
- Интеграция с внешними vault/хранилищами секретов
- Динамическое обновление окружения без перезапуска
- Аудит действий пользователей
- Веб-интерфейс для управления конфигами

---

**dummy** — это ваш быстрый старт в мир painless-окружений для разработки!
