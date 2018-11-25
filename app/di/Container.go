package di

import (
	. "../controller"
	. "../helper"
	. "../mapper"
	. "../repository"
	. "../service"
)

func InjectCategoryController() CategoryController {
	databaseConnectionHelper := &DatabaseConnectionHelper{}
	categoryRepository := &CategoryRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	categoryService := &CategoryService{
		ICategoryRepository: categoryRepository,
	}
	categoryResponseMapper := &CategoryResponseMapper{}
	categoryController := CategoryController{
		ICategoryService: categoryService,
		ICategoryResponseMapper: categoryResponseMapper,
	}
	return categoryController
}

func InjectCityController() CityController {
	databaseConnectionHelper := &DatabaseConnectionHelper{}
	cityRepository := &CityRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	cityService := &CityService{
		ICityRepository: cityRepository,
	}
	cityResponseMapper := &CityResponseMapper{}
	cityController := CityController{
		ICityService: cityService,
		ICityResponseMapper: cityResponseMapper,
	}
	return cityController
}

func InjectUserController() UserController {
	databaseConnectionHelper := &DatabaseConnectionHelper{}
	userRepository := &UserRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	userService := &UserService{
		IUserRepository: userRepository,
	}
	userLoginMapper := &UserLoginMapper{}
	userController := UserController{
		IUserService: userService,
		IUserLoginMapper: userLoginMapper,
	}
	return userController
}

func InjectProductController() ProductController {
	databaseConnectionHelper := &DatabaseConnectionHelper{}
	productRepository := &ProductRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	categoryRepository := &CategoryRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	productService := &ProductService{
		ICategoryRepository: categoryRepository,
		IProductRepository: productRepository,
	}
	availableProductForRentingResponseMapper := &AvailableProductForRentingResponseMapper{}
	productDetailResponseMapper := &ProductDetailResponseMapper{}
	productController := ProductController{
		IProductService: productService,
		IAvailableProductForRentingResponseMapper: availableProductForRentingResponseMapper,
		IProductDetailResponseMapper: productDetailResponseMapper,
	}
	return productController
}

func InjectOrderController() OrderController {
	databaseConnectionHelper := &DatabaseConnectionHelper{}
	orderRepository := &OrderRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	productRepository := &ProductRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	userRepository := &UserRepository{
		IDatabaseConnectionHelper: databaseConnectionHelper,
	}
	orderService := &OrderService{
		IUserRepository: userRepository,
		IProductRepository: productRepository,
		IOrderRepository: orderRepository,
	}
	orderController := OrderController{
		IOrderService: orderService,
	}
	return orderController
}