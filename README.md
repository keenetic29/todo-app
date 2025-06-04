# ToDo Application
Многофункциональное приложение для управления задачами с JWT-аутентификацией.
## 📦 Особенности

- **Чистая архитектура** с четким разделением слоев
- **JWT-аутентификация** с защищенными роутами
- **CRUD операции** для задач и категорий
- **PostgreSQL** для хранения данных
- **Gin+GORM** для API и работы с БД
- **Конфигурация через .env** файлы
- **Логирование** обработчиков
- **Gracful Shutdown** сервера
- **Swaggo** для автоматической генерации документации
- **Dockerfile и docker-compose.yml** для развертывания сервиса и БД

# Описание работы приложения
## Аутентификация
Большинство endpoints требуют аутентификации. После успешного входа вы получите JWT токен, который нужно передавать в заголовке:
```
Authorization: <ваш_токен>
```
## API Endpoints
### Аутентификация
Регистрация нового пользователя: ```POST /register```

Параметры:
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

Вход пользователя: ```POST /login```

Параметры:
```json
{
  "email": "string",
  "password": "string"
}
```

Ответ:

```json
{
  "token": "string"
}
```

### Категории

Создать категорию: ```POST /categories```

Параметры:
```json
{
  "name": "string"
}
```

Получить категории пользователя: ```GET /categories```

Удалить категорию: ```DELETE /categories/{id}```

### Задачи
Получить все задачи: ```GET /tasks```

Параметры запроса:

- page - номер страницы (по умолчанию 1)

- limit - количество элементов на странице (по умолчанию 10)

- sort - поле для сортировки (с префиксом - для DESC), например -created_at

- completed - фильтр по статусу выполнения (true/false)

Получить задачу по ID: ```GET /tasks/{id}```

Создать задачу: ```POST /tasks```

Параметры:
```json
{
  "title": "string",
  "description": "string",
  "category_id": int | null
}
```

Обновить задачу: ```PUT /tasks/{id}```

Параметры:
```json
{
  "title": "string",
  "description": "string",
  "category_id": number|null
}
```

Удалить задачу: ```DELETE /tasks/{id}```

Обновить категорию задачи: ```PATCH /tasks/{id}/category```

Параметры:
```json
{
  "category_id": number|null
}
```