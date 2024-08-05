package model

import (
	"database/sql/driver"
	"fmt"
)

type UserAccountProvider string

const (
	UserAccountProviderGoogle UserAccountProvider = "google"
)

func (provider UserAccountProvider) String() string {
	return string(provider)
}

// Value implements the driver.Valuer interface
func (provider UserAccountProvider) Value() (driver.Value, error) {
	return provider.String(), nil
}

// Scan implements the sql.Scanner interface
func (provider *UserAccountProvider) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		*provider = UserAccountProvider(v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return nil
}
