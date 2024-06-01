# PDF Generation Service

## Описание

PDF Generation Service - это веб-приложение, разработанное на Go, которое предоставляет API для генерации PDF файлов из HTML разметки. Пользователи могут передавать HTML разметку в формате JSON и получать PDF файл. Также система поддерживает управление пользователями и API токенами, учет использования токенов.

## Используемые технологии

- Go (Gin) - для создания веб-приложения и маршрутизации.
- MongoDB - для хранения данных пользователей, токенов и пакетов.
- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) - для генерации PDF файлов.

## Установка

1. Убедитесь, что у вас установлены Go и MongoDB.
2. Склонируйте репозиторий:
    ```sh
    git clone https://github.com/IST0VE/site_pdf_api.git
    ```
3. Перейдите в директорию проекта:
    ```sh
    cd site_pdf_api
    ```
4. Установите необходимые зависимости:
    ```sh
    go mod tidy
    ```
5. Запустите MongoDB на вашем локальном компьютере или настройте соединение с удаленной базой данных.
6. Запустите приложение:
    ```sh
    go run main.go
    ```

## Использование

### Регистрация пользователя

- **URL:** `/register`
- **Метод:** `POST`
- **Тело запроса (JSON):**
    ```json
    {
      "username": "testuser",
      "email": "testuser@example.com",
      "password": "password123"
    }
    ```

### Логин пользователя

- **URL:** `/login`
- **Метод:** `POST`
- **Тело запроса (JSON):**
    ```json
    {
      "email": "testuser@example.com",
      "password": "password123"
    }
    ```
- **Ответ (JSON):**
    ```json
    {
      "token": "your_jwt_token_here"
    }
    ```

### Генерация PDF файла

- **URL:** `/generate-pdf`
- **Метод:** `POST`
- **Заголовки:**
  - `Content-Type: application/json`
  - `Authorization: Bearer your_jwt_token_here`
- **Тело запроса (JSON):**
    ```json
    {
      "html_content": "<h1>Hello, World!</h1>",
      "api_token": "your_api_token_here"
    }
    ```
- **Ответ:** PDF файл в формате `application/pdf`

### Управление API токенами

#### Создание API токена

- **URL:** `/api-tokens`
- **Метод:** `POST`
- **Заголовки:**
  - `Content-Type: application/json`
  - `Authorization: Bearer your_jwt_token_here`
- **Тело запроса (JSON):**
    ```json
    {
      "name": "My API Token",
      "user_id": "user_object_id",
      "total_requests": 1000,
      "remaining_requests": 1000
    }
    ```
- **Ответ (JSON):**
    ```json
    {
      "token_id": "generated_token_id"
    }
    ```

#### Получение всех API токенов пользователя

- **URL:** `/api-tokens/user/:user_id`
- **Метод:** `GET`
- **Заголовки:**
  - `Authorization: Bearer your_jwt_token_here`
- **Ответ (JSON):**
    ```json
    [
      {
        "id": "token_id",
        "token": "random_generated_token",
        "name": "My API Token",
        "user_id": "user_object_id",
        "total_requests": 1000,
        "remaining_requests": 1000
      }
    ]
    ```

### Управление пакетами

#### Создание пакета

- **URL:** `/packages`
- **Метод:** `POST`
- **Заголовки:**
  - `Content-Type: application/json`
  - `Authorization: Bearer your_jwt_token_here`
- **Тело запроса (JSON):**
    ```json
    {
      "name": "Standard Package",
      "total_requests": 5000
    }
    ```
- **Ответ (JSON):**
    ```json
    {
      "package_id": "generated_package_id"
    }
    ```

#### Получение всех пакетов

- **URL:** `/packages`
- **Метод:** `GET`
- **Заголовки:**
  - `Authorization: Bearer your_jwt_token_here`
- **Ответ (JSON):**
    ```json
    [
      {
        "id": "package_id",
        "name": "Standard Package",
        "total_requests": 5000
      }
    ]
    ```

## Лицензия

Этот проект лицензируется на условиях лицензии MIT. Подробности см. в файле LICENSE.

