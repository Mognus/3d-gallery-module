package gallery

import (
	"template/modules/core/pkg/crud"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ImageProvider struct {
	db *gorm.DB
}

func NewImageProvider(db *gorm.DB) *ImageProvider {
	return &ImageProvider{db: db}
}

func (p *ImageProvider) GetModelName() string {
	return "gallery_images"
}

func (p *ImageProvider) GetSchema() crud.Schema {
	return crud.Schema{
		Name:        "gallery_images",
		DisplayName: "Gallery Images",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true, Editable: true, Width: "80px"},
			{Name: "name", Type: "string", Label: "Name", Editable: true, Width: "200px"},
			{Name: "url", Type: "string", Label: "URL", Required: true, Editable: true, Width: "300px"},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true, Editable: true, Width: "160px"},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true, Editable: true, Width: "160px"},
		},
		Searchable: []string{"name", "url"},
	}
}

func (p *ImageProvider) List(filters map[string]string, page, limit int) (crud.ListResponse, error) {
	return crud.DefaultList(p.db, &Image{}, filters, page, limit)
}

func (p *ImageProvider) Get(id string) (any, error) {
	return crud.DefaultGet(p.db, &Image{}, id)
}

func (p *ImageProvider) Create(data map[string]any) (any, error) {
	return crud.DefaultCreate(p.db, &Image{}, data)
}

func (p *ImageProvider) Update(id string, data map[string]any) (any, error) {
	return crud.DefaultUpdate(p.db, &Image{}, id, data)
}

func (p *ImageProvider) Delete(id string) error {
	return crud.DefaultDelete(p.db, &Image{}, id)
}

func (p *ImageProvider) ListHandler() fiber.Handler {
	return crud.DefaultListHandler(p)
}

func (p *ImageProvider) SchemaHandler() fiber.Handler {
	return crud.DefaultSchemaHandler(p)
}

func (p *ImageProvider) GetHandler() fiber.Handler {
	return crud.DefaultGetHandler(p)
}

func (p *ImageProvider) CreateHandler() fiber.Handler {
	return crud.DefaultCreateHandler(p)
}

func (p *ImageProvider) UpdateHandler() fiber.Handler {
	return crud.DefaultUpdateHandler(p)
}

func (p *ImageProvider) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(p)
}
