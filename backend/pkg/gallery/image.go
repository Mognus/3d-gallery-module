package gallery

import "time"

// Image represents an uploaded image asset.
type Image struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:200" json:"name"`
	URL       string    `gorm:"size:500;not null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Image) TableName() string {
	return "gallery_images"
}
