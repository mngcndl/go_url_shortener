# Сервис для сокращения URL

## Описание:

API сервис сокращения URL позволяет преобразовывать длинные ссылки в короткие и получать оригинальные ссылки по сокращенным.

## Установка:

### Требования ПО:

- Go версии 1.3
- PostgresSQL
  или с помощью Docker:
- Docker

### Установка сервиса:

1. Клонируйте данный репозиторий:

```bash
   git clone https://github.com/mngcndl/go_url_shortener
   cd url-shortener
```

2. Запустите сервис:

```bash
    docker-compose up --build
```

ИЛИ
Если вы хотите запустить данный сервис с использованием локального хранилища:

```bash
    go run cmd/main.go -storage memory
```

## Как работать с API:

В новом окне терминала введите команду

```bash
curl -X POST "http://localhost:8080/shorten?url=https://example.com"
```

Подставьте url для сокращения вместо https://example.com.

Для получения сокращенного url используйте команду

```bash
curl -X GET "http://localhost:8080/<shortened_url>"
```

Вместо shortened_url используйте полученный в ответ на команду POST ранее сокращенный url.

## Структура проекта:

├── Dockerfile # Файл для сборки Docker-образа
├── README.md # Документация проекта
├── cmd # Папка с исполняемым кодом
│ └── main.go # Главный файл приложения
├── config # Папка с конфигурационными файлами
│ └── config.go # Конфигурация приложения
├── docker-compose.yml # Файл для настройки Docker Compose
├── go.mod  
├── go.sum  
├── init.sql # Скрипт для инициализации базы данных
├── internal # Внутренние пакеты приложения
│ ├── common # Общие интерфейсы
│ │ └── interfaces.go # Интерфейсы хранилища, сервиса и обработчика
│ ├── handler # Пакет для обработки HTTP-запросов
│ │ ├── handler.go # Основная логика обработки запросов
│ │ └── handler_test.go # Тесты для обработчика
│ ├── service # Пакет с бизнес-логикой
│ │ ├── service.go # Реализация сервисов
│ │ └── service_test.go # Тесты для сервисов
│ └── storage # Пакет для работы с хранилищем данных
│ ├── memory.go # Реализация хранилища в памяти
│ ├── memory_test.go # Тесты для хранилища в памяти
│ ├── postgres.go # Реализация хранилища PostgreSQL
│ └── postgres_test.go # Тесты для хранилища PostgreSQL
├── pkg # Пакеты с общими утилитами
│ └── shortener # Пакет для генерации коротких ссылок
│ └── generator.go # Генерация коротких ссылок

Made by Смирнова Любовь (github: mngcndl), sldycpex@gmail.com
