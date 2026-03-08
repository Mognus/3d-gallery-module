package gallery

import (
	"fmt"
	"path/filepath"
	"strings"
	"template/modules/core/pkg/crud"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ModelAssetProvider struct {
	db        *gorm.DB
	uploadDir string
}

func NewModelAssetProvider(db *gorm.DB, uploadDir string) *ModelAssetProvider {
	return &ModelAssetProvider{db: db, uploadDir: uploadDir}
}

func (p *ModelAssetProvider) GetModelName() string {
	return "model_assets"
}

func (p *ModelAssetProvider) GetSchema() crud.Schema {
	return crud.Schema{
		Name:        "model_assets",
		DisplayName: "3D Model Assets",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true, Width: "80px"},
			{Name: "name", Type: "string", Label: "Name", Required: true, Width: "200px"},
			{Name: "model_url", Type: "file", Label: "GLB File", Required: true, Accept: "model/gltf-binary", Width: "300px"},
			{Name: "canvas_mesh_name", Type: "string", Label: "Canvas Mesh", CreateHidden: true, Width: "160px"},
			{Name: "default_scale", Type: "number", Label: "Default Scale", CreateHidden: true, Width: "120px"},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true, Width: "160px"},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true, Width: "160px"},
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
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "GLB file is required"})
		}

		ext := filepath.Ext(file.Filename)
		baseName := strings.TrimSuffix(file.Filename, ext)
		sanitized := sanitizeRe.ReplaceAllString(baseName, "-")
		filename := fmt.Sprintf("%s-%d%s", sanitized, time.Now().UnixNano(), ext)

		if err := c.SaveFile(file, filepath.Join(p.uploadDir, filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}

		name := c.FormValue("name")
		if name == "" {
			name = baseName
		}

		asset := ModelAsset{
			Name:           name,
			ModelURL:       "/uploads/models/" + filename,
			CanvasMeshName: c.FormValue("canvas_mesh_name"),
			DefaultScale:   1,
		}
		if err := p.db.Create(&asset).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create record"})
		}

		return c.Status(fiber.StatusCreated).JSON(asset)
	}
}

func (p *ModelAssetProvider) UpdateHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return crud.DefaultUpdateHandler(p)(c)
		}

		ext := filepath.Ext(file.Filename)
		baseName := strings.TrimSuffix(file.Filename, ext)
		sanitized := sanitizeRe.ReplaceAllString(baseName, "-")
		filename := fmt.Sprintf("%s-%d%s", sanitized, time.Now().UnixNano(), ext)

		if err := c.SaveFile(file, filepath.Join(p.uploadDir, filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}

		id := c.Params("id")
		updates := map[string]any{"model_url": "/uploads/models/" + filename}
		if name := c.FormValue("name"); name != "" {
			updates["name"] = name
		}
		if meshName := c.FormValue("canvas_mesh_name"); meshName != "" {
			updates["canvas_mesh_name"] = meshName
		}

		if err := p.db.Model(&ModelAsset{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update record"})
		}

		var asset ModelAsset
		p.db.First(&asset, id)
		return c.JSON(asset)
	}
}

func (p *ModelAssetProvider) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(p)
}
