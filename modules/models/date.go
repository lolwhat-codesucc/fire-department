package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Date – дата без времени (формат JSON: "2006-01-02", БД: DATE).
type Date struct {
	time.Time
}

// UnmarshalJSON реализует парсинг из строки "YYYY-MM-DD".
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("некорректная дата: %s", s)
	}
	d.Time = t
	return nil
}

// MarshalJSON сериализует дату в строку "YYYY-MM-DD".
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

// Value реализует интерфейс driver.Valuer для работы с БД.
func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.Time.Format("2006-01-02"), nil
}

// Scan реализует интерфейс sql.Scanner для чтения из БД.
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		d.Time = t
		return nil
	}
	return fmt.Errorf("неподдерживаемый тип для Date: %T", value)
}
