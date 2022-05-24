package repository

import (
	"hacktiv8-day07/sesi07-gorm/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() (*[]models.Product, error)
	GetProductByID(uint) (*models.Product, error)
	UpdateProductByID(*models.Product, uint) (*models.Product, error)
	DeleteProductByID(uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

// CreateProduct implements ProductRepository
func (p *productRepository) CreateProduct(product *models.Product) error {
	return p.db.Create(product).Error

}

// GetAllProducts implements ProductRepository
func (p *productRepository) GetAllProducts() (*[]models.Product, error) {
	products := []models.Product{}

	err := p.db.Find(&products).Error
	return &products, err
}

// DeleteProductByID implements ProductRepository
func (p *productRepository) DeleteProductByID(id uint) error {
	product := models.Product{}
	err := p.db.Delete(&product, "id=?", id).Error
	return err
}

// GetProductByID implements ProductRepository
func (p *productRepository) GetProductByID(id uint) (*models.Product, error) {
	product := models.Product{}

	err := p.db.First(&product, "id=?", id).Error
	return &product, err
}

// UpdateProductByID implements ProductRepository
func (p *productRepository) UpdateProductByID(request *models.Product, id uint) (*models.Product, error) {
	product := models.Product{}
	e := p.db.First(&product, "id=?", id).Error
	if e != nil {
		return nil, e
	}
	err := p.db.Where("id=?", id).Updates(&request)
	product = *request
	product.ID = id
	if err.Error != nil {
		return nil, err.Error
	}
	return &product, nil

}
