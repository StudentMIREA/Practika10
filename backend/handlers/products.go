package handlers

import (
	"fmt"
	"net/http"
	"shopApi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Получение списка всех продуктов
func GetProducts(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		err := db.Select(&products, `SELECT products.id, products.name, products.image, products.cost, products.description, 
									(CASE WHEN (favorites.id IS NOT NULL) THEN true ELSE false END) as favorite,
									(CASE WHEN (cart.id IS NOT NULL) THEN true ELSE false END) as shop_cart,
									(CASE WHEN (cart.count IS NOT NULL) THEN cart.count ELSE 0 END) as count 
									FROM products 
									LEFT JOIN favorites ON products.id = favorites.product_id
									LEFT JOIN cart ON products.id = cart.product_id 
									ORDER BY products.id`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Ошибка получения списка продуктов",
				"details": err.Error(), // Логирование деталей ошибки
			})
			return

		}
		c.JSON(http.StatusOK, products)
	}
}

// Получение одного продукта по его ID
func GetProduct(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID продукта"})
			return
		}

		var product models.Product
		err = db.Get(&product, `SELECT products.id, products.name, products.image, products.cost, products.description, 
								(CASE WHEN (favorites.id IS NOT NULL) THEN true ELSE false END) as favorite,
								(CASE WHEN (cart.id IS NOT NULL) THEN true ELSE false END) as shop_cart,
								(CASE WHEN (cart.count IS NOT NULL) THEN cart.count ELSE 0 END) as count 
								FROM products 
								LEFT JOIN favorites ON products.id = favorites.product_id
								LEFT JOIN cart ON products.id = cart.product_id
								WHERE products.id = $1`, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Создание нового продукта
func CreateProduct(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}
		query := `INSERT INTO products (name, image, cost, description) 
                  VALUES (:Name, :Image, :Cost, :Description)`
		_, err := db.NamedExec(query, map[string]interface{}{
			"Name":        product.Name,
			"Image":       product.Image,
			"Cost":        product.Cost,
			"Description": product.Describtion,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления продукта"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

// Обновление существующего продукта по его ID
func UpdateProduct(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID продукта"})
			return
		}

		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}

		product.ID = id
		query := `UPDATE products SET name = :Name, image = :Image, cost = :Cost, 
                  description = :Description WHERE id = :ID`

		// Используем NamedExec для выполнения запроса с именованными параметрами
		_, err = db.NamedExec(query, map[string]interface{}{
			"Name":        product.Name,
			"Image":       product.Image,
			"Cost":        product.Cost,
			"Description": product.Describtion,
			"ID":          product.ID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления продукта"})
			return
		}
		fmt.Println(product.Name)
		c.JSON(http.StatusOK, product)
	}
}

// Удаление продукта по его ID
func DeleteProduct(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID продукта"})
			return
		}

		_, err = db.Exec("DELETE FROM products WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления продукта"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Продукт успешно удален"})
	}
}
