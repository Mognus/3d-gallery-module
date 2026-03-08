package gallery

import (
	"template/modules/core/pkg/crud"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ModelAssetProvider struct {
	db *gorm.DB
}

func NewModelAssetProvider(db *gorm.DB) *ModelAssetProvider {
	return &ModelAssetProvider{db: db}
}

func (p *ModelAssetProvider) GetModelName() string {
	return "model_assets"
}

func (p *ModelAssetProvider) GetSchema() crud.Schema {
	return crud.Schema{
		Name:        "model_assets",
		DisplayName: "3D Model Assets",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true, Editable: true, Width: "80px"},
			{Name: "name", Type: "string", Label: "Name", Required: true, Editable: true, Width: "200px"},
			{Name: "model_url", Type: "string", Label: "Model URL", Required: true, Editable: true, Width: "300px"},
			{Name: "canvas_mesh_name", Type: "string", Label: "Canvas Mesh", Editable: true, Width: "160px"},
			{Name: "default_scale", Type: "number", Label: "Default Scale", Editable: true, Width: "120px"},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true, Editable: true, Width: "160px"},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true, Editable: true, Width: "160px"},
		},
		Searchable: []string{"name"},
	}
}

func (p *ModelAssetProvider) List(filters map[string]string, page, limit int) (crud.ListResponse, error) {
	return crud.DefaultList(p.db, &ModelAsset{}, filters, page, limit)
}

func (p *ModelAssetProvider) Get(id string) (any, error) {
	return crud.DefaultGet(p.db, &ModelAsset{}, id)
}

func (p *ModelAssetProvider) Create(data map[string]any) (any, error) {
	return crud.DefaultCreate(p.db, &ModelAsset{}, data)
}

func (p *ModelAssetProvider) Update(id string, data map[string]any) (any, error) {
	return crud.DefaultUpdate(p.db, &ModelAsset{}, id, data)
}

func (p *ModelAssetProvider) Delete(id string) error {
	return crud.DefaultDelete(p.db, &ModelAsset{}, id)
}

func (p *ModelAssetProvider) ListHandler() fiber.Handler {
	return crud.DefaultListHandler(p)
}

func (p *ModelAssetProvider) SchemaHandler() fiber.Handler {
	return crud.DefaultSchemaHandler(p)
}

func (p *ModelAssetProvider) GetHandler() fiber.Handler {
	return crud.DefaultGetHandler(p)
}

func (p *ModelAssetProvider) CreateHandler() fiber.Handler {
	return crud.DefaultCreateHandler(p)
}

func (p *ModelAssetProvider) UpdateHandler() fiber.Handler {
	return crud.DefaultUpdateHandler(p)
}

func (p *ModelAssetProvider) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(p)
}
