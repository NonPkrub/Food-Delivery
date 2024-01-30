package route

import (
	"Food-delivery/api/controller"
	"Food-delivery/api/middleware"
	"Food-delivery/database"
	"Food-delivery/repository"
	"Food-delivery/usecase"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	userRepo := repository.NewUserRepository(database.DB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := controller.NewUserController(userUseCase)

	productRepo := repository.NewProductRepository(database.DB)
	productUseCase := usecase.NewProductUseCase(productRepo)
	productController := controller.NewProductController(productUseCase)

	promotionRepo := repository.NewPromotionRepository(database.DB)
	promotionUseCase := usecase.NewPromotionUseCase(promotionRepo)
	promotionController := controller.NewPromotionController(promotionUseCase)

	basketRepo := repository.NewBasketRepository(database.DB)
	basketUseCase := usecase.NewBasketUseCase(basketRepo)
	basketController := controller.NewBasketController(basketUseCase)

	basketProductRepo := repository.NewBasketProductRepository(database.DB)
	basketProductUseCase := usecase.NewBasketProductUseCase(basketProductRepo)
	basketProductController := controller.NewBasketProductController(basketProductUseCase)

	promotionProductRepo := repository.NewPromotionProductRepository(database.DB)
	promotionProductUseCase := usecase.NewPromotionProductUseCase(promotionProductRepo)
	promotionProductController := controller.NewPromotionProductController(promotionProductUseCase)

	app := fiber.New()
	v1 := app.Group("api/v1")
	//user
	v1.Get("user/:id", middleware.JwtAuthentication(), userController.GetUserById)
	v1.Post("sign_up", userController.SignUp)
	v1.Post("login", userController.Login)

	//product
	v1.Get("product", middleware.JwtAuthentication(), productController.GetAll)
	v1.Post("product", middleware.JwtAuthentication(), productController.AddProduct)
	v1.Put("product/:id", middleware.JwtAuthentication(), productController.EditProduct)
	v1.Delete("product/:id", middleware.JwtAuthentication(), productController.DeleteProduct)

	//promotion
	v1.Get("promotion", middleware.JwtAuthentication(), promotionController.GetAllPromotion)
	v1.Get("promotion/:id", middleware.JwtAuthentication(), promotionController.GetPromotionById)
	v1.Post("promotion", middleware.JwtAuthentication(), promotionController.CreatePromotion)
	v1.Put("promotion/:id", middleware.JwtAuthentication(), promotionController.EditPromotion)
	v1.Delete("promotion/:id", middleware.JwtAuthentication(), promotionController.DeletePromotion)
	v1.Post("promotion/search", middleware.JwtAuthentication(), promotionController.SearchPromotion)

	//basket
	//v1.Post("user/:id/basket", basketController.CreateBasket)
	v1.Put("user/:id/basket/promotion", middleware.JwtAuthentication(), basketController.AddPromotionBasket)
	v1.Delete("user/:id/basket/promotion", middleware.JwtAuthentication(), basketController.DeletePromotionBasket)
	v1.Get("user/:id/basket", middleware.JwtAuthentication(), basketController.GetBasketByUserId)

	//basket_product
	v1.Post("user/basket/:id/product", middleware.JwtAuthentication(), basketProductController.AddProductInBasket)
	v1.Put("user/basket/:id/product", middleware.JwtAuthentication(), basketProductController.EditProductInBasket)
	v1.Delete("user/basket/:id/product", middleware.JwtAuthentication(), basketProductController.DeleteProductInBasket)
	v1.Get("user/basket/:id/product", middleware.JwtAuthentication(), basketProductController.GetProductInBasket)

	//promotion_product
	v1.Post("promotion/product", middleware.JwtAuthentication(), promotionProductController.AddPromotionProduct)
	v1.Put("promotion/:id/product", middleware.JwtAuthentication(), promotionProductController.EditPromotionProduct)
	v1.Get("promotion/:id/product", middleware.JwtAuthentication(), promotionProductController.GetPromotionProduct)

	v2 := app.Group("api/v2")
	//user
	v2.Get("user/:id", middleware.BasicAuth, userController.GetUserById)
	v2.Post("sign_up", userController.SignUp)
	v2.Post("login", userController.Login)

	//product
	v2.Get("product", middleware.BasicAuth, productController.GetAll)
	v2.Post("product", middleware.BasicAuth, productController.AddProduct)
	v2.Put("product/:id", middleware.BasicAuth, productController.EditProduct)
	v2.Delete("product/:id", middleware.BasicAuth, productController.DeleteProduct)

	//promotion
	v2.Get("promotion", middleware.BasicAuth, promotionController.GetAllPromotion)
	v2.Get("promotion/:id", middleware.BasicAuth, promotionController.GetPromotionById)
	v2.Post("promotion", middleware.BasicAuth, promotionController.CreatePromotion)
	v2.Put("promotion/:id", middleware.BasicAuth, promotionController.EditPromotion)
	v2.Delete("promotion/:id", middleware.BasicAuth, promotionController.DeletePromotion)
	v2.Post("promotion/search", middleware.BasicAuth, promotionController.SearchPromotion)

	//basket
	v2.Put("user/:id/basket/promotion", middleware.BasicAuth, basketController.AddPromotionBasket)
	v2.Delete("user/:id/basket/promotion", middleware.BasicAuth, basketController.DeletePromotionBasket)
	v2.Get("user/:id/basket", middleware.BasicAuth, basketController.GetBasketByUserId)

	//basket_product
	v2.Post("user/basket/:id/product", middleware.BasicAuth, basketProductController.AddProductInBasket)
	v2.Put("user/basket/:id/product", middleware.BasicAuth, basketProductController.EditProductInBasket)
	v2.Delete("user/basket/:id/product", middleware.BasicAuth, basketProductController.DeleteProductInBasket)
	v2.Get("user/basket/:id/product", middleware.BasicAuth, basketProductController.GetProductInBasket)

	//promotion_product
	v2.Post("promotion/product", middleware.BasicAuth, promotionProductController.AddPromotionProduct)
	v2.Put("promotion/:id/product", middleware.BasicAuth, promotionProductController.EditPromotionProduct)
	v2.Get("promotion/:id/product", middleware.BasicAuth, promotionProductController.GetPromotionProduct)

	return app
}
