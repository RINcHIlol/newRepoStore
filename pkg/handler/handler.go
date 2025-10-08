package handler

import (
    "github.com/gin-gonic/gin"
    "storeApi/pkg/service"
    "net/http"
    "os"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Настройка HTML шаблонов
	router.LoadHTMLGlob("templates/*")
	
	// Статические файлы
	router.Static("/static", "./static")
	
	// Главная страница
	router.GET("/", h.showProductsPage)
	
    // Страница корзины
    router.GET("/cart", h.showCartPage)

    // Платежи Stripe
    router.POST("/payments/create-intent", h.createPaymentIntent)
	
	// Страница успешной покупки
	router.GET("/success", h.showSuccessPage)
	
	// Админ-страница
	router.GET("/admin", h.showAdminPage)
	
	auth := router.Group("/store")
	{
		auth.POST("/add", h.addProduct)
		auth.GET("/get", h.getProducts)
		auth.GET("/get/:id", h.getProductById)
		auth.POST("/buy/", h.buyProduct)
		auth.PUT("/update/:id", h.updateProduct)
		auth.PUT("/update/count", h.addCountProduct)
		auth.DELETE("/delete/:id", h.deleteProduct)
	}

	return router
}

// showProductsPage отображает HTML страницу с продуктами
func (h *Handler) showProductsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "products.html", gin.H{
		"title": "Магазин - Продукты",
	})
}

// showCartPage отображает HTML страницу корзины
func (h *Handler) showCartPage(c *gin.Context) {
    c.HTML(http.StatusOK, "cart.html", gin.H{
        "title": "Корзина",
        "stripe_pk": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
        "stripe_test_autopay": os.Getenv("STRIPE_TEST_AUTOPAY"),
    })
}

// showSuccessPage отображает HTML страницу успешной покупки
func (h *Handler) showSuccessPage(c *gin.Context) {
	c.HTML(http.StatusOK, "success.html", gin.H{
		"title": "Покупка завершена",
	})
}

// showAdminPage отображает HTML страницу админ-панели
func (h *Handler) showAdminPage(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Админ-панель",
	})
}
