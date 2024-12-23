## Стек:  
#### Go, Gin (framework), Gorm (ORM), PostgreSQL, Docker

## API
API для поиска сотрудников по имени.
Метод возвращает полную информацию о сотрудниках с указанным именем.  
**Формат запроса:** `GET http://{{host}}/employee/{{name}}`  
**Результат:** список сотрудников с именем *{{name}}*

![Пример запроса 1](./screenshoots/postman_example_1.png  "Пример 1")

![Пример запроса 2](./screenshoots/postman_example_2.png  "Пример 2")

## Import command
Также реализована команда для импорта сотрудников из csv файла  
**Формат команды:** `go run main.go loadcsv {{путь к файлу}}`
**Результат:** имортирует всех сотрудников из файла и в случае успеха
выводит `Data loaded successfully!`

**Примечание:** В данной реализации система допускает сотрудников с одинаковым ФИО и не учитывает возможное дублирование сотрудников.
В виду того, что во входных данных есть сотрудники с одинаковыми ФИО,
но разными отделами, должностями и т.д. и в теории у сотрудника эти данные могут изменяться,
так что проблематично без дополнительной информации однозначно идентифицировать сотрудника по имени

![Таблица employees](./screenshoots/import_result.png  "Таблица employees")
![Таблица departments](./screenshoots/import_result_2.png  "Таблица departments")
![Таблица jobs](./screenshoots/import_result_3.png  "Таблица jobs")
![Таблица salaries](./screenshoots/import_result_4.png  "Таблица salaries")
![Таблица hourlies](./screenshoots/import_result_5.png  "Таблица hourlies")

