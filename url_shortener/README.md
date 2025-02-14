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

Made by Смирнова Любовь (github: mngcndl), sldycpex@gmail.com
