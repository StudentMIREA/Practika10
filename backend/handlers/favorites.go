package handlers

import (
	"net/http"
	"shopApi/models"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func GetFavorites(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		var products []models.Product
		err := db.Select(&products, `SELECT products.id, products.name, products.image, products.cost, products.description, 
									(CASE WHEN (favorites.id IS NOT NULL) THEN true ELSE false END) as favorite,
									(CASE WHEN (cart.id IS NOT NULL) THEN true ELSE false END) as shop_cart,
									(CASE WHEN (cart.count IS NOT NULL) THEN cart.count ELSE 0 END) as count 
									FROM products 
									LEFT JOIN cart ON products.id = cart.product_id 
									JOIN favorites ON products.id = favorites.product_id
									WHERE favorites.user_id = $1`, userId)
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

func AddToFavorites(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}
		_, err := db.Exec("INSERT INTO favorites (user_id, product_id) VALUES ($1, $2)",
			userId, product.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка добавления в избранное"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Товар добавлен в избранное"})
	}
}

func RemoveFromFavorites(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		productId := c.Param("productId")
		_, err := db.Exec("DELETE FROM favorites WHERE user_id = $1 AND product_id = $2", userId, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления из избранного"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Товар удален из избранного"})
	}
}
