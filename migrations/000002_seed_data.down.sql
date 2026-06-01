DELETE FROM Call_status WHERE name IN ('Поступил','Выезд','На месте','Завершён');
DELETE FROM Car_model WHERE name IN ('АЦ-40','АЛ-30');
DELETE FROM District WHERE name IN ('Центральный','Северный','Южный');
DELETE FROM Specialization WHERE name IN ('Пожарная','Спасательная','Химическая');
DELETE FROM Rank WHERE name IN ('Рядовой','Сержант','Лейтенант','Капитан');
