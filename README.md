# Система управления пожарной частью (Fire Department)

Бэкенд-приложение для управления пожарными, командами, специализациями, автомобилями, районами и вызовами. Поддерживает REST API, интерактивный CLI и миграции базы данных.

## 1. Требования к хранимым данным

Система управляет следующими ключевыми сущностями:

- **Firefighter (Пожарный):** Имя, дата рождения, звание, квалификация, принадлежность к команде (опционально).
- **Rank (Звание):** Справочник званий.
- **Specialization (Специализация):** Виды специализаций команд (пожарная, спасательная, химическая и др.).
- **Team (Команда):** Номер, специализация, район, закреплённый автомобиль (опционально, уникально).
- **District (Район):** Название района обслуживания.
- **District_specializations:** Связь многие-ко-многим между районами и доступными специализациями.
- **Car (Автомобиль):** Модель, дата приобретения, дата последнего ТО, готовность.
- **Car_model (Модель автомобиля):** Название и период технического обслуживания (в днях).
- **Call (Вызов):** Команда, автомобиль, район (опционально), время, статус, комментарий.
- **Call_status (Статус вызова):** Справочник статусов (поступил, выезд, на месте, завершён).

**Основные ограничения:**

- Один автомобиль может быть закреплён только за одной командой одновременно (уникальность `car_id` в таблице `Team`).
- Пожарный может не принадлежать ни к одной команде (необязательное поле `team_id`).
- Вызов обязательно содержит ссылки на команду, автомобиль и статус; район может быть не указан.
- Период технического обслуживания хранится на уровне модели автомобиля, а не конкретной машины (нормализация).

## 2. Проектирование структуры БД

### ER-диаграмма

```mermaid
erDiagram
    Firefighter {
        int id PK
        string name
        date year_of_birth
        int rank_id FK "NOT NULL"
        int qualification
        int team_id FK "NULL"
    }
    Rank {
        int id PK
        string name
    }
    Team {
        int number PK
        int specialization_id FK "NOT NULL"
        int district_id FK "NOT NULL"
        int car_id FK "NULL, UNIQUE"
    }
    Specialization {
        int id PK
        string name
    }
    District {
        int id PK
        string name
    }
    District_specializations {
        int district_id PK, FK
        int specialization_id PK, FK
    }
    Car {
        int id PK
        int model_id FK "NOT NULL"
        date acquisition_date
        date last_maintenance "NULL"
        bool ready "NOT NULL"
    }
    Car_model {
        int id PK
        string name
        int maintenance_period_days
    }
    Call {
        int id PK
        int team_id FK "NOT NULL"
        int district_id FK "NULL"
        int car_id FK "NOT NULL"
        datetime time
        int status_id FK "NOT NULL"
        string comment "NULL"
    }
    Call_status {
        int id PK
        string name
    }
    Rank ||--o{ Firefighter : ""
    Team ||--o{ Firefighter : ""
    Specialization ||--o{ Team : ""
    District ||--o{ Team : ""
    Car |o--o| Team : ""
    Car_model ||--|{ Car : ""
    District ||--o{ District_specializations : ""
    Specialization ||--o{ District_specializations : ""
    Team ||--o{ Call : ""
    District ||--o{ Call : ""
    Car ||--o{ Call : ""
    Call_status ||--o{ Call : ""
```

# Анализ нормализации (соответствие форме Бойса-Кодда)

Схема базы данных удовлетворяет требованиям нормальной формы Бойса-Кодда (BCNF).

Обоснование:

    Все таблицы‑справочники (Rank, Specialization, District, Car_model, Call_status) содержат только один потенциальный ключ — первичный, и все неключевые атрибуты полностью от него зависят.

    В таблице Car изначально присутствовала зависимость model_id → maintenance_period_days, что нарушало BCNF, так как model_id не является суперключом. Поле maintenance_period_days было перенесено в таблицу Car_model, где model_id — суперключ. Теперь все неключевые атрибуты Car зависят только от первичного ключа id.

    В таблице Team поле car_id уникально, но оно может быть NULL, что не создаёт дополнительной функциональной зависимости, нарушающей BCNF.

    Все связи многие‑ко‑многим (District_specializations) разложены на две связи 1:N, каждая из которых соответствует BCNF.

Таким образом, в каждой таблице любой детерминант является суперключом,
что гарантирует форму Бойса-Кодда. 

# 3. Архитектура приложения

Приложение построено на принципах чистой архитектуры со следующими слоями:

    Слой обработчиков (handlers): Приём HTTP-запросов, валидация входных данных, формирование ответов. Используется роутер chi.

    Слой репозиториев (repository): Интерфейсы для доступа к данным (CRUD и специфичные запросы). Реализации находятся в пакете postgres и используют pgxpool.

    Слой моделей (domain/models): Структуры, отображающие таблицы БД, с тегами для JSON и БД.

    Интерактивный CLI: Альтернативный способ взаимодействия с системой напрямую через терминал (пакет cli). Позволяет выполнять все CRUD-операции без HTTP-запросов.
