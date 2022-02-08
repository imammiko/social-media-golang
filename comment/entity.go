package comment

import "time"

type (
	Comment struct {
		ID        int `gorm:"primary_key"`
		User_id   int
		Photo_id  int
		Message   string `gorm:"not null;unique" validate:"required"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
