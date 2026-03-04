package repository

import (
    "log"
    "microservices-demo/internals/models"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() *ProductRepository  {

	 dsn := "root:12345@tcp(127.0.0.1:3307)/products?charset=utf8mb4&parseTime=True&loc=Local"

	 database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	 if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	 }

	 database.AutoMigrate(&models.Product{})
	 return &ProductRepository{db: database}
	
}

func (r *ProductRepository) Add(product models.Product) models.Product{
	r.db.Create(&product)
	return product
}

func (r *ProductRepository) GetAll() []models.Product {
    var products []models.Product
    r.db.Find(&products)
    return products
}

func (r *ProductRepository) Update(id uint, name string, price int) (models.Product, bool) {

	var product models.Product

	result := r.db.First(&product, id) // Select FROM products where id = ?

	if result.Error != nil {
		return product, false
	}

	product.Name = name
    product.Price = price
    r.db.Save(&product)
    return product, true

}

func (r *ProductRepository) Delete(id uint) bool {
    result := r.db.Unscoped().Delete(&models.Product{}, id) // DELETE FROM products where id = ? // Unscoped() will remove the entire row permanently -- hard delete
	//result := r.db.Delete(&models.Product{}, id) // DELETE FROM products where id = ?
    return result.RowsAffected > 0
}



