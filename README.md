# VK TEST | Filmoteka REST API | Golang
<h3 style="text-align: center;">Реализация тестового задания "Фильмотека" Backend на языке Go</h3>

![1600_800_trans_65f6ccb65a1bd](https://github.com/localpurpose/vkTask/assets/139833240/c3942b83-a358-4711-8668-34085b44a954)

## Функциональность
|FUNC| DONE| DESC|
|-|--------|---|
|Добавление информации об актере|✅|-|
|Изменение информации об актере|✅|-|
|Удаление информации об актере|✅|-|
|Получение списка актеров|✅|All Actors with Movies|
|Добавление информации о фильме|✅|Validating: +|
|Изменение информации о фильме|✅|-|
|Удаление информации о фильме|✅|-|
|Получение списка фильмов|✅|-|
|Поиск фильма по фрагменту названия|✅|-|
|API закрыт авторизацией USER/ADMIN|✅|-|
|Методы авторизации JWT TOKENS|✅|-|
|Backend language - GO|✅|-|
|Только native http |✅|MUX|
|SWAGGER 2.0 Specification|✅|-|
|PostgreSQL DATABASE|✅|-|
|Logging Middleware|✅|-|
|Docker File для сборки|✅|-|
|docker-compsoe для запуска|✅|WEB+DB|

## Quick Start
1. Склонируйте репозиторий на вашу локальную машину <br>
``` git clone https://github.com/localpurpose/vkTask.git ```
2. Создайте файл .env и укажите параметры: <br>
```python
DB_USER=userName
DB_PASSWORD=strongPassword
DB_NAME=databaseName
JWT_SECRET="jwtSecretKey"
```
3. Запустите с помощью docker <br>
``` docker compose up```
