# ⏱️ Time Tracker API

## 📄 Описание

Time Tracker API - это RESTful API для отслеживания времени выполнения задач пользователями. Он включает в себя функции для управления пользователями и задачами, а также интеграцию с внешним API для получения информации о пользователях.

## 🚀 Установка

1. Клонируйте репозиторий:

	```bash
	git clone https://github.com/yourusername/time-tracker-api.git

2. Перейдите в директорию проекта:

	```bash
	cd time-tracker-api

3. Создайте новый файл .env, или используйте уже готовый .env файл.

# Настройки подключения к базе данных

DB_HOST=localhost

DB_PORT=5432

DB_USER=admin_timetracker

DB_PASSWORD=1111 

DB_NAME=timetracker

# URL внешнего API для получения информации о пользователе

API_URL=http://example.com/api - Возьмите свой URL-адрес API и вставьте его в файл .env.

4. Установите зависимости:
		
	```bash
	go mod tidy

5. Запустите проект:
	
	```bash
	go run main.go

## 📚 Документация

API документация доступна по адресу /swagger/index.html.

## 🛠️ Эндпоинты

POST /users - Создать нового пользователя

GET /users/:id - Получить пользователя по ID

PUT /users/:id - Обновить пользователя по ID

DELETE /users/:id - Удалить пользователя по ID

GET /users - Получить список пользователей с фильтрацией и пагинацией

POST /tasks - Начать новую задачу для пользователя

PUT /tasks/:id/end - Завершить задачу по ID

GET /users/:user_id/tasks - Получить задачи пользователя за период

## 🤝 Вклад

Я приветствую вклад от всех желающих. Пожалуйста, создайте Pull Request или откройте Issue для предложений и улучшений.

## 📄 Лицензия

Этот проект лицензирован под лицензией MIT. Подробности смотрите в файле LICENSE.

## 📧 Контакты

Если у вас есть вопросы или предложения, пожалуйста, свяжитесь с нами по адресу andryafpooo@mail.ru.

## ⏱️ Time Tracker API

## 📄 Description

Time Tracker API is a RESTful API for tracking task execution time by users. It includes features for managing users and tasks, as well as integration with an external API for fetching user information.

## 🚀 Installation
1. Clone the repository:

	```bash
	git clone https://github.com/yourusername/time-tracker-api.git

2. Navigate to the project directory:

	```bash
	cd time-tracker-api

3. Create a new .env file or use an existing .env file:

# Database configuration

DB_HOST=localhost

DB_PORT=5432

DB_USER=admin_timetracker

DB_PASSWORD=1111

DB_NAME=timetracker

# External API URL for fetching user information

API_URL=http://example.com/api - Take your API URL and paste it into the .env file.

4. Install dependencies:

	```bash
	go mod tidy

5. Run the project:

	```bash
	go run main.go

## 📚 Documentation

API documentation is available at /swagger/index.html.

## 🛠️ Endpoints

POST /users - Create a new user

GET /users/:id - Get user by ID

PUT /users/:id - Update user by ID

DELETE /users/:id - Delete user by ID

GET /users - Get list of users with filtering and pagination

POST /tasks - Start a new task for a user

PUT /tasks/:id/end - End a task by ID

GET /users/:user_id/tasks - Get tasks of a user for a specific period

## 🤝 Contributing

I welcome contributions from everyone. Please create a Pull Request or open an Issue for suggestions and improvements.

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.

## 📧 Contact

If you have any questions or suggestions, please contact us at andryafpooo@mail.ru.