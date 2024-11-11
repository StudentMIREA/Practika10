package handlers

import (
	"fmt"
	"net/http"
	"shopApi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Получение одного продукта по его ID
func GetUser(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
			return
		}

		var product models.User
		err = db.Get(&product, `SELECT * FROM users 
								WHERE id = $1`, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			fmt.Println("Error")
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Обновление существующего продукта по его ID
func UpdateUser(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
			return
		}

		var product models.User
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
			return
		}

		product.ID = id
		query := `UPDATE users SET name = :Name, image = :Image, phone = :Phone, 
                  mail = :Mail WHERE id = :ID`

		// Используем NamedExec для выполнения запроса с именованными параметрами
		_, err = db.NamedExec(query, map[string]interface{}{
			"Name":  product.Name,
			"Image": product.Image,
			"Phone": product.Phone,
			"Mail":  product.Mail,
			"ID":    product.ID,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления продукта"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}
