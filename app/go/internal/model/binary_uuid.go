package model

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// UUID -> binary uuid wrapper over uuid.UUID
type UUID uuid.UUID

// ParseUUID -> parses string uuid to binary uuid
func ParseUUID(id string) UUID {
	return UUID(uuid.MustParse(id))
}

func (id UUID) String() string {
	return uuid.UUID(id).String()
}

// MarshalJSON -> convert to json string
func (id UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(id)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON -> convert from json string
func (id *UUID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*id = UUID(s)
	return err
}

// GormDataType -> sql data type for gorm
func (UUID) GormDataType() string {
	return "binary(16)"
}

// Scan -> scan value into UUID
func (id *UUID) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal UUID value:", value))
	}

	data, err := uuid.FromBytes(bytes)
	*id = UUID(data)
	return err
}

// Value -> return UUID to []bytes binary(16)
func (id UUID) Value() (driver.Value, error) {
	return uuid.UUID(id).MarshalBinary()
}
