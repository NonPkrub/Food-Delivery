package route

import (
	"Food-delivery/api/controller"
	"Food-delivery/api/middleware"
	"Food-delivery/dal"
	"Food-delivery/database"
	"Food-delivery/repository"
	"Food-delivery/usecase"
	"context"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	//repo
	ctx := context.Background()

	userRepo := repository.NewUserRepository(database.DB)
	productRepo := repository.NewProductRepository(database.DB)
	promotionRepo := repository.NewPromotionRepository(database.DB)
	basketProductRepo := repository.NewBasketProductRepository(database.DB)
	basketRepo := repository.NewBasketRepository(database.DB)
	promotionProductRepo := repository.NewPromotionProductRepository(database.DB)
	promotionRepoTest := dal.Use(database.DB).WithContext(ctx).Promotion
	productRepoTest := dal.Use(database.DB).WithContext(ctx).Product
	promotionProductRepoTest := dal.Use(database.DB).WithContext(ctx).PromotionProduct

	//usecase
	basketProductUseCase := usecase.NewBasketProductUseCase(basketProductRepo, promotionRepo, productRepo, basketRepo)
	basketUseCase := usecase.NewBasketUseCase(basketRepo, basketProductUseCase)
	userUseCase := usecase.NewUserUseCase(userRepo, basketUseCase)
	productUseCase := usecase.NewProductUseCase(productRepo)
	promotionUseCase := usecase.NewPromotionUseCase(promotionRepoTest, productRepoTest, promotionProductRepoTest)
	promotionProductUseCase := usecase.NewPromotionProductUseCase(promotionProductRepo, productRepo)

	//controllers
	userController := controller.NewUserController(userUseCase)
	productController := controller.NewProductController(productUseCase)
	promotionController := controller.NewPromotionController(promotionUseCase)
	basketProductController := controller.NewBasketProductController(basketProductUseCase)
	basketController := controller.NewBasketController(basketUseCase)
	promotionProductController := controller.NewPromotionProductController(promotionProductUseCase)

	app := fiber.New()
	v1 := app.Group("api/v1")
	//user
	//v1.Get("user/:id", middleware.JwtAuthentication(), userController.GetUserById)
	v1.Post("sign_up", userController.SignUp)
	v1.Post("login", userController.Login)
	v1.Get("me", middleware.JwtAuthentication(), userController.Me)

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

	//basket
	//v1.Post("user/:id/basket", basketController.CreateBasket)
	v1.Put("user/basket/promotion", middleware.JwtAuthentication(), basketController.AddPromotionBasket)
	v1.Delete("user/basket/promotion", middleware.JwtAuthentication(), basketController.DeletePromotionBasket)
	v1.Get("user/basket", middleware.JwtAuthentication(), basketController.GetBasketByUserId)

	//basket_product
	v1.Post("user/basket/product", middleware.JwtAuthentication(), basketProductController.AddProductInBasket)
	v1.Put("user/basket/product", middleware.JwtAuthentication(), basketProductController.EditProductInBasket)
	v1.Delete("user/basket/product", middleware.JwtAuthentication(), basketProductController.DeleteProductInBasket)

	//promotion_product
	v1.Post("promotion/product", middleware.JwtAuthentication(), promotionProductController.AddPromotionProduct)
	v1.Put("promotion/:id/product", middleware.JwtAuthentication(), promotionProductController.EditPromotionProduct)

	return app
}
