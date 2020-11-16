package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	admin = "Admin"
	user  = "User"
)

type Role struct {
	role string
}

func Empty() Role {
	return Role{""}
}
func Admin() Role {
	return Role{admin}
}
func User() Role {
	return Role{user}
}
func (n Role) IsEmpty() bool {
	return (n.role == "")
}

func (n Role) String() string {
	return n.role
}

func isValid(r string) bool {
	return (r == user || r == admin)
}

func (n *Role) scan(v string) {
	if isValid(v) {
		n.role = v
	} else {
		n.role = ""
	}
}

// Scan implements the Scanner interface.
func (n *Role) Scan(value interface{}) error {
	x := sql.NullString{}

	err := x.Scan(value)
	if err != nil {
		return err
	}

	n.scan(x.String)

	return nil
}

// Value implements the driver Valuer interface.
func (n Role) Value() (driver.Value, error) {
	if n.IsEmpty() {
		return nil, nil
	}

	return n.role, nil
}

func (n Role) MarshalJSON() ([]byte, error) {
	if n.IsEmpty() {
		return []byte("null"), nil
	}

	return []byte("\"" + n.role + "\""), nil
}

func (n *Role) UnmarshalJSON(data []byte) error {
	var v *string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v == nil {
		*n = Empty()
		return nil
	}

	n.scan(*v)

	if n.IsEmpty() {
		return errors.New("invalid role: " + *v)
	}

	return nil
}
