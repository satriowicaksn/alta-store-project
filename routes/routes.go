package routes

import (
	"alta-store-project/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Product Categories
	e.GET("/categories", controllers.GetProductCategoriesControllers)
	e.GET("/categories/:id", controllers.GetProductCategoriesByIdControllers)
	e.POST("/categories", controllers.PostProductCategories)
	e.PUT("/categories/:id", controllers.PutProductCategoriesById)
	e.DELETE("/categories/:id", controllers.DeleteProductCategoriesById)

	// Products
	e.GET("/products", controllers.GetProductControllers)
	e.GET("/products/:id", controllers.GetProductsByCategoryControllers)
	return e
}
