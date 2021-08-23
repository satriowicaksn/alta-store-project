package routes

import (
	"alta-store-project/constants"
	"alta-store-project/controllers"

	"github.com/labstack/echo/v4"
	middlewareEcho "github.com/labstack/echo/v4/middleware"
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

	// User Authentication
	e.POST("/login", controllers.LoginUserController)

	// JWT Group
	r := e.Group("")
	r.Use(middlewareEcho.JWT([]byte(constants.SECRET_JWT)))

	// User Auth
	r.GET("/users/:id", controllers.GetUserDetailController)
	r.POST("/carts", controllers.AddCartController)
	r.GET("/carts", controllers.GetCartController)

	return e
}
