package gallery

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Module struct {
	db                  *gorm.DB
	imageUploadDir      string
	modelUploadDir      string
	modelAssetProvider  *ModelAssetProvider
	imageProvider       *ImageProvider
	galleryItemProvider *GalleryItemProvider
}

func New(db *gorm.DB, imageUploadDir, modelUploadDir string) *Module {
	assetProvider := NewModelAssetProvider(db, modelUploadDir)
	return &Module{
		db:                  db,
		imageUploadDir:      imageUploadDir,
		modelUploadDir:      modelUploadDir,
		modelAssetProvider:  assetProvider,
		imageProvider:       NewImageProvider(db, imageUploadDir),
		galleryItemProvider: NewGalleryItemProvider(db, assetProvider),
	}
}

func (m *Module) Name() string {
	return "gallery"
}

func (m *Module) ModelAssetProvider() *ModelAssetProvider {
	return m.modelAssetProvider
}

func (m *Module) ImageProvider() *ImageProvider {
	return m.imageProvider
}

func (m *Module) GalleryItemProvider() *GalleryItemProvider {
	return m.galleryItemProvider
}

func (m *Module) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&ModelAsset{}, &Image{}, &GalleryItem{})
}

func (m *Module) RegisterRoutes(router fiber.Router) {
	assets := router.Group("/gallery/model-assets")
	assets.Get("/", m.modelAssetProvider.ListHandler())
	assets.Get("/schema", m.modelAssetProvider.SchemaHandler())
	assets.Get("/:id", m.modelAssetProvider.GetHandler())
	assets.Post("/", m.modelAssetProvider.CreateHandler())
	assets.Put("/:id", m.modelAssetProvider.UpdateHandler())
	assets.Delete("/:id", m.modelAssetProvider.DeleteHandler())

	images := router.Group("/gallery/images")
	images.Get("/", m.imageProvider.ListHandler())
	images.Get("/schema", m.imageProvider.SchemaHandler())
	images.Get("/:id", m.imageProvider.GetHandler())
	images.Post("/", m.imageProvider.CreateHandler())
	images.Put("/:id", m.imageProvider.UpdateHandler())
	images.Delete("/:id", m.imageProvider.DeleteHandler())

	items := router.Group("/gallery/items")
	items.Get("/", m.galleryItemProvider.ListHandler())
	items.Get("/schema", m.galleryItemProvider.SchemaHandler())
	items.Get("/:id", m.galleryItemProvider.GetHandler())
	items.Post("/", m.galleryItemProvider.CreateHandler())
	items.Put("/:id", m.galleryItemProvider.UpdateHandler())
	items.Delete("/:id", m.galleryItemProvider.DeleteHandler())
}
