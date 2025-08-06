# URL Shortener

Простой сервис для сокращения ссылок на Go с PostgreSQL

## 🚀 Быстрый старт

### Требования
- Docker
- Docker Compose

### Запуск
```bash
# Собрать и запустить контейнеры
docker-compose up -d --build

# Остановить сервис
docker-compose down

📌 Эндпоинты
Создать короткую ссылку
text

POST /create
Body: url=<полная_ссылка>

Пример:
bash

curl -X POST -d "url=https://google.com" http://localhost:8080/create

Перейти по короткой ссылке
text

GET /r/<короткий_код>

Пример:
bash

curl -v http://localhost:8080/r/abc123

🛠️ Разработка
Структура проекта
text

.
├── cmd/              # Точка входа
├── internal/         # Основная логика
│   ├── handler/      # HTTP обработчики
│   └── storage/      # Работа с БД
├── configs/          # Конфигурация
├── Dockerfile        # Конфигурация контейнера
└── docker-compose.yml # Оркестрация сервисов

Полезные команды
bash

# Запустить тесты
go test -v ./...

# Просмотр логов
docker-compose logs -f app

# Доступ к БД
docker-compose exec postgres psql -U user -d url_shortener

⚙️ Настройки

Переменные окружения (можно изменить в docker-compose.yml):

    DB_HOST - Хост БД (по умолчанию postgres)

    DB_PORT - Порт БД (по умолчанию 5432)

    DB_USER - Пользователь БД (по умолчанию user)

    DB_PASSWORD - Пароль БД (по умолчанию password)

    DB_NAME - Имя БД (по умолчанию url_shortener)

📝 Лицензия

MIT
text


### Что можно добавить по желанию:
1. Скриншоты работы (если будет веб-интерфейс)
2. Примеры ответов API
3. Дорожную карту развития
4. Список используемых технологий

Файл готов к использованию! Можно дополнить его по мере развития проекта.