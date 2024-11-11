package handlers

import (
	"net/http"
	"shopApi/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetCart(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		var products []models.Product
		err := db.Select(&products, `SELECT products.id, products.name, products.image, products.cost, products.description, 
									(CASE WHEN (favorites.id IS NOT NULL) THEN true ELSE false END) as favorite,
									(CASE WHEN (cart.id IS NOT NULL) THEN true ELSE false END) as shop_cart,
									(CASE WHEN (cart.count IS NOT NULL) THEN cart.count ELSE 0 END) as count 
									FROM products 
									JOIN cart ON products.id = cart.product_id 
									LEFT JOIN favorites ON products.id = favorites.product_id
									WHERE cart.user_id = $1`, userId)
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

func AddToCart(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}
		_, err := db.Exec("INSERT INTO cart (user_id, product_id, count) VALUES ($1, $2, $3)",
			userId, product.ID, product.Count)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления в корзину"})
			return
		}
		c.JSON(http.StatusOK, product)

	}
}

// Обновление существующего продукта по его ID
func UpdateShopCart(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")

		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}
		_, err := db.Exec("UPDATE cart SET count = $3 WHERE user_id = $1 AND product_id = $2",
			userId, product.ID, product.Count)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления количества товаров в корзине"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func RemoveFromCart(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		productId := c.Param("productId")
		_, err := db.Exec("DELETE FROM cart WHERE user_id = $1 AND product_id = $2", userId, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления из корзины"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Товар удален из корзины"})
	}
}
