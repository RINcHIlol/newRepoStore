package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"storeApi/models"
	"strconv"
	"storeApi"
)

func (h *Handler) addProduct(c *gin.Context) {
	// Получаем данные из формы
	name := c.PostForm("name")
	priceStr := c.PostForm("price")
	description := c.PostForm("description")
	countStr := c.PostForm("count")
	
	// Получаем файл изображения
	file, err := c.FormFile("image")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка загрузки изображения: "+err.Error())
		return
	}
	
	// Читаем содержимое файла
	src, err := file.Open()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Ошибка чтения файла: "+err.Error())
		return
	}
	defer src.Close()
	
	// Читаем все байты изображения
	imageBytes := make([]byte, file.Size)
	_, err = src.Read(imageBytes)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Ошибка чтения данных изображения: "+err.Error())
		return
	}
	
	// Парсим числовые значения
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный формат цены")
		return
	}
	
	count, err := strconv.Atoi(countStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Неверный формат количества")
		return
	}
	
	// Создаем объект продукта
	product := models.Product{
		Name:        name,
		Price:       price,
		Description: description,
		Image:       imageBytes,
		Count:       sql.NullInt64{Int64: int64(count), Valid: true},
	}
	
	isAdd, err := h.services.Store.AddNewProduct(product)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isAdd": isAdd,
		"message": "Товар успешно добавлен",
	})
}

func (h *Handler) addCountProduct(c *gin.Context) {
	var input models.EditCount

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	curCount, err := h.services.Store.AddCountProduct(input.ProductId, input.Count)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Count": curCount,
	})
}

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.Store.GetProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}

func (h *Handler) getProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	product, err := h.services.Store.GetProductById(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
	})
}

func (h *Handler) buyProduct(c *gin.Context) {
	var req models.OrderRequest
	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

    orderId, err := h.services.Store.BuyProduct(req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"orderId": orderId,
		"success": true,
	})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	product, err := h.services.Store.DeleteProductById(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
	})
}


// createPaymentIntent создает PaymentIntent в Stripe и возвращает client_secret
func (h *Handler) createPaymentIntent(c *gin.Context) {
    var req struct {
        Amount int64 `json:"amount"`
    }
    if err := c.BindJSON(&req); err != nil || req.Amount <= 0 {
        newErrorResponse(c, http.StatusBadRequest, "invalid amount")
        return
    }

    clientSecret, err := storeApi.CreatePaymentIntent(req.Amount)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "clientSecret": clientSecret,
    })
}

func (h *Handler) updateProduct(c *gin.Context) {
	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	isUpdate, err := h.services.Store.UpdateProductById(productId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isUpdate": isUpdate,
	})
}
