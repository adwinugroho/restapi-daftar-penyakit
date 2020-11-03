package models

type (
	Diseases struct {
		ID uint32 `gorm:"primary_key;size:255;not null;unique" json:"id"`
	}

	Logger struct {
		CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
		UpdatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	}
)
