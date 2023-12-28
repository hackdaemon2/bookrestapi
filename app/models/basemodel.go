package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const CustomDateTimeLayoutMySQL = "2006-01-02 15:04:05"

type CustomDateTimeMySQL struct {
	time.Time
}

func (ct *CustomDateTimeMySQL) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	ct.Time, err = time.Parse(CustomDateTimeLayoutMySQL, string(b))
	return
}

func (ct *CustomDateTimeMySQL) MarshalJSON() ([]byte, error) {
	if ct.IsZero() {
		return []byte("null"), nil
	}

	return []byte(`"` + ct.Time.Format(CustomDateTimeLayoutMySQL) + `"`), nil
}

func (ct CustomDateTimeMySQL) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil
	}

	return ct.Time.Format(CustomDateTimeLayoutMySQL), nil
}

func (ct *CustomDateTimeMySQL) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	fmt.Println(value)

	switch v := value.(type) {
	case []uint8:
		parsedTime, err := time.Parse(CustomDateTimeLayoutMySQL, string(v))
		if err != nil {
			return err
		}
		ct.Time = parsedTime
	case time.Time:
		ct.Time = v
	default:
		return fmt.Errorf("unexpected type %T for CustomDateTimeMySQL", value)
	}

	return nil
}

type BaseModel struct {
	ID        uint                `json:"id" gorm:"primary_key;column:id"`
	CreatedAt CustomDateTimeMySQL `json:"created_at" gorm:"column:created_at"`
	UpdatedAt CustomDateTimeMySQL `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt CustomDateTimeMySQL `json:"deleted_at" gorm:"column:deleted_at"`
	Deleted   int                 `json:"deleted" gorm:"column:deleted"`
}
