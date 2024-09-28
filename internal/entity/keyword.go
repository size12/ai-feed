package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/rs/zerolog/log"
)

type Keyword struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Keywords []*Keyword

func (keywords Keywords) Value() (driver.Value, error) {
	return json.Marshal(keywords)
}

func (keywords *Keywords) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		log.Fatal().Msg("failed unmarshal JSON")
	}

	return json.Unmarshal(data, &keywords)
}
