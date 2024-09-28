package entity

import (
	"github.com/google/uuid"
	"time"
)

type Personality struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id,omitempty"`

	OwnerLogin string `json:"-"`

	CreatedAt time.Time `gorm:"autoUpdateTime" json:"created_at"`

	Name      string `json:"name" validate:"required"`
	Biography string `json:"biography" validate:"required"`
	Keywords  string `json:"keywords" validate:"required"`
	Thematics string `json:"thematics" validate:"required"`
	TextStyle string `json:"text_style" validate:"required"`
}

var InitPersonalities = []Personality{
	{
		Name:      "Александр Пушкин",
		Biography: "Русский поэт, драматург и прозаик, заложивший основы русского реалистического направления.",
		Keywords:  "поэзия литература",
		Thematics: "19 век",
		TextStyle: "стихотворный",
	},
	{
		Name:      "Ричард Фейнман",
		Biography: "Американский физик-теоретик. Прославился своими работами в области квантовой электродинамики.",
		Keywords:  "физика электродинамика",
		Thematics: "физика",
		TextStyle: "неофициальный",
	},
	{
		Name:      "Аристотель",
		Biography: "Древнегреческий философ и учёный.",
		Keywords:  "философия естествознание",
		Thematics: "философия",
		TextStyle: "официальный",
	},
}
