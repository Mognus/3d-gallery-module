package gallery

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"template/modules/core/pkg/crud"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var sanitizeRe = regexp.MustCompile(`[^a-zA-Z0-9-_]`)

type ImageProvider struct {
	db        *gorm.DB
	uploadDir string
}

func NewImageProvider(db *gorm.DB, uploadDir string) *ImageProvider {
	return &ImageProvider{db: db, uploadDir: uploadDir}
}

func (p *ImageProvider) GetModelName() string {
	return "gallery_images"
}

func (p *ImageProvider) GetSchema() crud.Schema {
	return crud.Schema{
		Name:        "gallery_images",
		DisplayName: "Gallery Images",
		Fields: []crud.Field{
			{Name: "id", Type: "number", Label: "ID", Readonly: true},
			{Name: "name", Type: "string", Label: "Name"},
			{Name: "url", Type: "file", Label: "Image", Required: true},
			{Name: "created_at", Type: "date", Label: "Created", Readonly: true},
			{Name: "updated_at", Type: "date", Label: "Updated", Readonly: true},
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

// CreateHandler handles both multipart (file upload) and JSON (URL) requests.
func (p *ImageProvider) CreateHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			// No file → fall back to JSON create (URL mode)
			return crud.DefaultCreateHandler(p)(c)
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

		img := Image{Name: name, URL: "/uploads/gallery/" + filename}
		if err := p.db.Create(&img).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create record"})
		}

		return c.Status(fiber.StatusCreated).JSON(img)
	}
}

// UpdateHandler handles both multipart (new file) and JSON (name/url change only).
func (p *ImageProvider) UpdateHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			// No file → fall back to JSON update
			return crud.DefaultUpdateHandler(p)(c)
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

		id := c.Params("id")
		if err := p.db.Model(&Image{}).Where("id = ?", id).Updates(map[string]any{
			"name": name,
			"url":  "/uploads/gallery/" + filename,
		}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update record"})
		}

		var img Image
		p.db.First(&img, id)
		return c.JSON(img)
	}
}

func (p *ImageProvider) DeleteHandler() fiber.Handler {
	return crud.DefaultDeleteHandler(p)
}
