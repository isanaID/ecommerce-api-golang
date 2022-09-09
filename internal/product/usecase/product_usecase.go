package usecase

import (
	"ecommerce-api/internal/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductUseCase struct {
	db *gorm.DB
}

func NewProductUseCase(db *gorm.DB) *ProductUseCase {
	return &ProductUseCase{
		db: db,
	}
}

func (productUsecase *ProductUseCase) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if product.Name == "" || product.Price == 0 || product.Stock == 0 {
		c.JSON(400, gin.H{
			"message": "Isi semua field",
		})
		return
	}

	if err := productUsecase.db.Create(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot create product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success create product",
	})
}

func (productUsecase *ProductUseCase) GetProducts(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var product domain.Product
	if err := productUsecase.db.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    product,
	})
}

func (productUsecase *ProductUseCase) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var product domain.Product
	if err := productUsecase.db.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get product",
		})
		return
	}

	var updateProduct domain.Product
	if err := c.BindJSON(&updateProduct); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if updateProduct.Name == "" || updateProduct.Price == 0 || updateProduct.Stock == 0 {
		c.JSON(400, gin.H{
			"message": "Isi semua field",
		})
		return
	}

	if err := productUsecase.db.Model(&product).Updates(updateProduct).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success update product",
	})
}

func (productUsecase *ProductUseCase) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var product domain.Product
	if err := productUsecase.db.First(&product, id).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get product",
		})
		return
	}

	if err := productUsecase.db.Delete(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete product",
	})
}

func (productUsecase *ProductUseCase) SearchProduct(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("name LIKE ?", "%"+c.Query("name")+"%").Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByCategory(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("category_id = ?", c.Param("id")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByPrice(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("price BETWEEN ? AND ?", c.Query("min"), c.Query("max")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByStock(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("stock BETWEEN ? AND ?", c.Query("min"), c.Query("max")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByPriceAndStock(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("price BETWEEN ? AND ? AND stock BETWEEN ? AND ?", c.Query("min"), c.Query("max"), c.Query("min"), c.Query("max")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByPriceAndStockAndCategory(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("price BETWEEN ? AND ? AND stock BETWEEN ? AND ? AND category_id = ?", c.Query("min"), c.Query("max"), c.Query("min"), c.Query("max"), c.Param("id")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByPriceAndStockAndName(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("price BETWEEN ? AND ? AND stock BETWEEN ? AND ? AND name LIKE ?", c.Query("min"), c.Query("max"), c.Query("min"), c.Query("max"), "%"+c.Query("name")+"%").Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) GetProductByPriceAndStockAndNameAndCategory(c *gin.Context) {
	var products []domain.Product
	if err := productUsecase.db.Where("price BETWEEN ? AND ? AND stock BETWEEN ? AND ? AND name LIKE ? AND category_id = ?", c.Query("min"), c.Query("max"), c.Query("min"), c.Query("max"), "%"+c.Query("name")+"%", c.Param("id")).Find(&products).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get products",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    products,
	})
}

func (productUsecase *ProductUseCase) CreateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := productUsecase.db.Create(&category).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot create category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success create category",
	})
}

func (productUsecase *ProductUseCase) GetCategory(c *gin.Context) {
	var categories []domain.Category
	if err := productUsecase.db.Find(&categories).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get categories",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    categories,
	})
}

func (productUsecase *ProductUseCase) GetCategoryByID(c *gin.Context) {
	var category domain.Category
	if err := productUsecase.db.Where("id = ?", c.Param("id")).Find(&category).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    category,
	})
}

func (productUsecase *ProductUseCase) UpdateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if err := productUsecase.db.Model(&category).Where("id = ?", c.Param("id")).Updates(&category).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot update category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success update category",
	})
}

func (productUsecase *ProductUseCase) DeleteCategory(c *gin.Context) {
	if err := productUsecase.db.Where("id = ?", c.Param("id")).Delete(&domain.Category{}).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot delete category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete category",
	})
}

func (productUsecase *ProductUseCase) BuyProduct(c *gin.Context) {
	var product domain.Product
	if err := productUsecase.db.Where("id = ?", c.Param("id")).Find(&product).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot get product",
		})
		return
	}

	if product.Stock == 0 {
		c.JSON(400, gin.H{
			"message": "product is out of stock",
		})
		return
	}

	if err := productUsecase.db.Model(&product).Where("id = ?", c.Param("id")).Update("stock", product.Stock-1).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "cannot buy product",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success buy product",
	})
}
