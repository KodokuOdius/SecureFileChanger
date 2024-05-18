# Цели приложение:

- Защита важных документов (шифрование файлов)
- Легкость развертывания приложение на внутренней сети – аналог portainer, Apache airflow
- Простота использования
- Гибка настройка конфигурации
- Научится использовать Golang на практике 

## Что делает приложение

- Регистрация первого пользователя как админа
- Администратор ограничивается доступ к ресурсу
- Авторизация/Регистрация по коду (email)
- Загрузка документов
- Выгрузка документов
- Создание архивов выгрузки (выгрузка нескольких документов)
- Возможность делиться файлами
- Создание одноразовых/временных ссылок на скачивание или просмотр файла

## Инструменты

- Golang – для создания сервисов и всего приложения
- React.js – для создания визуальной части приложения
- Postgres – для хранения логов и данных приложения
- Docker(docker-compose) – для удобного запуска приложения