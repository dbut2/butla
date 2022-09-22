package models

import (
	"encoding/json"
)

type Link struct {
	Code   string     `db:"code"`
	Url    string     `db:"url"`
	Expiry NullTime   `db:"expiry"`
	IP     NullString `db:"ip"`
}

func (l Link) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}

func (l *Link) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, l)
}
