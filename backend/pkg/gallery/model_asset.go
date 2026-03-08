package gallery

import "time"

// ModelAsset represents a 3D model file with its metadata.
// CanvasMeshName identifies the mesh used for image projection (empty = no projection surface).
type ModelAsset struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"size:200;not null" json:"name"`
	ModelURL       string    `gorm:"size:500;not null" json:"model_url"`
	CanvasMeshName string    `gorm:"size:200" json:"canvas_mesh_name"`
	DefaultScale   float64   `gorm:"default:1" json:"default_scale"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (ModelAsset) TableName() string {
	return "model_assets"
}
