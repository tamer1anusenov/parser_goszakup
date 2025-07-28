🛠 Парсер тендеров с goszakup.gov.kz на Python

Этот проект представляет собой Python-скрипт для автоматического сбора тендеров с сайта goszakup.gov.kz, с сохранением данных в базу данных PostgreSQL.

📌 Что делает скрипт?

Заходит на сайт goszakup.gov.kz (через API или парсинг HTML).
Скачивает данные о тендерах (например: номер закупки, заказчик, сумма, статус, дата публикации и т.д.).
Преобразует их в структурированный вид.
Сохраняет в таблицу PostgreSQL для дальнейшей аналитики или отчетности.
🧰 Технологии

Python 3.10+
Библиотеки: requests, BeautifulSoup4, psycopg2, pandas (опционально)
PostgreSQL (версия 12+)
📋 Установка

Клонируй репозиторий:
git clone https://github.com/yourusername/goszakup-parser.git
cd goszakup-parser
Установи зависимости:
pip install -r requirements.txt
Создай .env файл с настройками БД:
DB_HOST=localhost
DB_PORT=5432
DB_NAME=goszakup_db
DB_USER=postgres
DB_PASSWORD=yourpassword
Настрой таблицу в БД:
CREATE TABLE tenders (
    id SERIAL PRIMARY KEY,
    purchase_number TEXT,
    customer_name TEXT,
    amount NUMERIC,
    status TEXT,
    publish_date DATE,
    json_data JSONB
);
🚀 Запуск

python parser.py
По умолчанию скрипт подключается к БД, тянет свежие данные о тендерах, и сохраняет их в таблицу tenders.

