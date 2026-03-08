package gallery

import "time"

// GalleryItem is a composition of a ModelAsset and an optional Image.
// ImageID is only set when the ModelAsset has a projection surface (CanvasMeshName != "").
type GalleryItem struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Title        string      `gorm:"size:200;not null" json:"title"`
	ModelAssetID uint        `gorm:"not null" json:"model_asset_id"`
	ModelAsset   ModelAsset  `gorm:"foreignKey:ModelAssetID" json:"model_asset"`
	ImageID      *uint       `json:"image_id"`
	Image        *Image      `gorm:"foreignKey:ImageID" json:"image"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func (GalleryItem) TableName() string {
	return "gallery_items"
}
