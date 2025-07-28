```markdown
🛠 Парсер тендеров с goszakup.gov.kz (Python)

Проект представляет собой Python-скрипт для автоматического сбора данных о тендерах с сайта goszakup.gov.kz с сохранением в PostgreSQL.

📌 Возможности
- Получает данные с goszakup.gov.kz (через API или парсинг HTML)
- Собирает информацию о тендерах:
  - Номер закупки
  - Наименование заказчика
  - Сумма
  - Статус
  - Дата публикации
  - Другие релевантные поля
- Преобразует данные в структурированный формат
- Сохраняет в PostgreSQL для дальнейшего анализа

🧰 Технологии
- Python 3.10+
- Библиотеки: requests, BeautifulSoup4, psycopg2, pandas (опционально)
- PostgreSQL 12+

📋 Установка
1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/goszakup-parser.git
cd goszakup-parser
```

2. Установите зависимости:
```bash
pip install -r requirements.txt
```

3. Создайте .env файл с настройками БД:
```ini
DB_HOST=localhost
DB_PORT=5432
DB_NAME=goszakup_db
DB_USER=postgres
DB_PASSWORD=yourpassword
```

4. Настройте таблицу в БД:
```sql
CREATE TABLE tenders (
    id SERIAL PRIMARY KEY,
    purchase_number TEXT,
    customer_name TEXT,
    amount NUMERIC,
    status TEXT,
    publish_date DATE,
    json_data JSONB
);
```

## 🚀 Запуск
```bash
python parser.py
```

Скрипт по умолчанию:
1. Подключается к базе данных
2. Загружает актуальные данные о тендерах
3. Сохраняет их в таблицу tenders

## 💡 Примечания
- Убедитесь, что сервер PostgreSQL запущен перед использованием
- При изменении структуры сайта может потребоваться корректировка парсера
```