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
	v1.Get("user/:id", userController.GetUserById)
	v1.Post("sign_up", userController.SignUp)
	v1.Post("login", userController.Login)

	//product
	v1.Get("product", middleware.JwtAuthentication(), productController.GetAll)
	v1.Post("product", productController.AddProduct)
	v1.Put("product/:id", productController.EditProduct)
	v1.Delete("product/:id", productController.DeleteProduct)

	//promotion
	v1.Get("promotion", promotionController.GetAllPromotion)
	v1.Get("promotion/:id", promotionController.GetPromotionById)
	v1.Post("promotion", promotionController.CreatePromotion)
	v1.Put("promotion/:id", promotionController.EditPromotion)
	v1.Delete("promotion/:id", promotionController.DeletePromotion)
	v1.Post("promotion/search", promotionController.SearchPromotion)

	//basket
	//v1.Post("user/:id/basket", basketController.CreateBasket)
	v1.Put("user/:id/basket/promotion", basketController.AddPromotionBasket)
	v1.Delete("user/:id/basket/promotion", basketController.DeletePromotionBasket)
	v1.Get("user/:id/basket", basketController.GetBasketByUserId)

	//basket_product
	v1.Post("user/basket/:id/product", basketProductController.AddProductInBasket)
	v1.Put("user/basket/:id/product", basketProductController.EditProductInBasket)
	v1.Delete("user/basket/:id/product", basketProductController.DeleteProductInBasket)
	v1.Get("user/basket/:id/product", basketProductController.GetProductInBasket)

	//promotion_product
	v1.Post("promotion/product", promotionProductController.AddPromotionProduct)
	v1.Put("promotion/:id/product", promotionProductController.EditPromotionProduct)
	v1.Get("promotion/:id/product", promotionProductController.GetPromotionProduct)

	return app
}
