package gallery

import (
	"template/modules/core/pkg/crud"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GalleryItemProvider struct {
	db            *gorm.DB
	assetProvider *ModelAssetProvider
}

func NewGalleryItemProvider(db *gorm.DB, assetProvider *ModelAssetProvider) *GalleryItemProvider {
	return &GalleryItemProvider{db: db, assetProvider: assetProvider}
}

func (p *GalleryItemProvider) GetModelName() string {
	return "gallery_items"
}

func (p *GalleryItemProvider) GetSchema() crud.Schema {
	assetOptions := crud.GetRelatedOptions(p.db, &ModelAsset{}, "ID", "Name")
	imageOptions := crud.GetRelatedOptions(p.db, &Image{}, "ID", "Name")

	return crud.Schema{
		Name:        "gallery_items",
		DisplayName: "Gallery Items",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true, Editable: true, Width: "80px"},
			{Name: "title", Type: "string", Label: "Title", Required: true, Editable: true, Width: "200px"},
			{Name: "model_asset_id", Type: "relation", Label: "Model Asset", Required: true, Editable: true, Hidden: true, Options: assetOptions},
			{Name: "model_asset", Type: "object", Label: "Model Asset", Readonly: true, Editable: false, Width: "160px"},
			{Name: "image_id", Type: "relation", Label: "Image", Editable: true, Hidden: true, Options: imageOptions},
			{Name: "image", Type: "object", Label: "Image", Readonly: true, Editable: false, Width: "160px"},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true, Editable: true, Width: "160px"},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true, Editable: true, Width: "160px"},
		},
		Filterable: []string{"model_asset_id"},
		Searchable: []string{"title"},
	}
}

func (p *GalleryItemProvider) List(filters map[string]string, page, limit int) (crud.ListResponse, error) {
	return crud.DefaultList(p.db, &GalleryItem{}, filters, page, limit, "ModelAsset", "Image")
}

func (p *GalleryItemProvider) Get(id string) (any, error) {
	return crud.DefaultGet(p.db, &GalleryItem{}, id, "ModelAsset", "Image")
}

func (p *GalleryItemProvider) Create(data map[string]any) (any, error) {
	return crud.DefaultCreate(p.db, &GalleryItem{}, data)
}

func (p *GalleryItemProvider) Update(id string, data map[string]any) (any, error) {
	return crud.DefaultUpdate(p.db, &GalleryItem{}, id, data)
}

func (p *GalleryItemProvider) Delete(id string) error {
	return crud.DefaultDelete(p.db, &GalleryItem{}, id)
}

func (p *GalleryItemProvider) ListHandler() fiber.Handler {
	return crud.DefaultListHandler(p)
}

func (p *GalleryItemProvider) SchemaHandler() fiber.Handler {
	return crud.DefaultSchemaHandler(p)
}

func (p *GalleryItemProvider) GetHandler() fiber.Handler {
	return crud.DefaultGetHandler(p)
}

func (p *GalleryItemProvider) CreateHandler() fiber.Handler {
	return crud.DefaultCreateHandler(p)
}

func (p *GalleryItemProvider) UpdateHandler() fiber.Handler {
	return crud.DefaultUpdateHandler(p)
}

func (p *GalleryItemProvider) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(p)
}
