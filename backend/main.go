package main

import (
	"Zlahoda_AIS/controllers"
	"Zlahoda_AIS/middleware"
	"Zlahoda_AIS/repositories"
	"github.com/gin-gonic/gin"
)

var employeeController *controllers.EmployeeController
var categoryController *controllers.CategoryController
var customerCardController *controllers.CustomerCardController
var productController *controllers.ProductController
var productInStoreController *controllers.ProductInStoreController
var receiptController *controllers.ReceiptController
var selfController *controllers.SelfController
var analyticsController *controllers.AnalyticsController

func main() {
	repositories.InitializeDB()

	employeeRepository := &repositories.PostgresEmployeeRepository{}
	categoryRepository := &repositories.PostgresCategoryRepository{}
	customerCardRepository := &repositories.PostgresCustomerCardRepository{}
	productRepository := &repositories.PostgresProductRepository{}
	productInStoreRepository := &repositories.PostgresProductInStoreRepository{}
	receiptRepository := &repositories.PostgresReceiptRepository{}
	analyticsRepository := &repositories.PostgresAnalyticsRepository{}

	employeeController = &controllers.EmployeeController{EmployeeRepository: employeeRepository}
	categoryController = &controllers.CategoryController{CategoryRepository: categoryRepository}
	customerCardController = &controllers.CustomerCardController{CustomerCardRepository: customerCardRepository}
	productController = &controllers.ProductController{ProductRepository: productRepository}
	productInStoreController = &controllers.ProductInStoreController{ProductInStoreRepository: productInStoreRepository}
	receiptController = &controllers.ReceiptController{ReceiptRepository: receiptRepository}
	selfController = &controllers.SelfController{EmployeeRepository: employeeRepository}
	analyticsController = &controllers.AnalyticsController{AnalyticsRepository: analyticsRepository}

	router := initRouter()
	router.Run(":8080")

	repositories.CloseDB()
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORS())

	router.POST("/login", employeeController.LoginEmployee)

	managerGroup := router.Group("/manager")
	managerGroup.Use(middleware.ManagerAuth())
	{
		employeeGroup := managerGroup.Group("/employee")
		{
			employeeGroup.POST("/", employeeController.RegisterEmployee)
			employeeGroup.GET("/", employeeController.GetAllEmployees)
			employeeGroup.GET("/:id_employee", employeeController.GetEmployeeById)
			employeeGroup.PATCH("/:id_employee", employeeController.UpdateEmployee)
			employeeGroup.DELETE("/:id_employee", employeeController.DeleteEmployee)
		}

		categoryGroup := managerGroup.Group("/category")
		{
			categoryGroup.POST("/", categoryController.CreateCategory)
			categoryGroup.GET("/", categoryController.GetAllCategories)
			categoryGroup.GET("/:category_number", categoryController.GetCategory)
			categoryGroup.PATCH("/:category_number", categoryController.UpdateCategory)
			categoryGroup.DELETE("/:category_number", categoryController.DeleteCategory)
		}

		customerCardGroup := managerGroup.Group("/customerCard")
		{
			customerCardGroup.POST("/", customerCardController.CreateCustomerCard)
			customerCardGroup.GET("/", customerCardController.GetAllCustomerCards)
			customerCardGroup.GET("/:card_number", customerCardController.GetCustomerCard)
			customerCardGroup.PATCH("/:card_number", customerCardController.UpdateCustomerCard)
			customerCardGroup.DELETE("/:card_number", customerCardController.DeleteCustomerCard)
		}

		productGroup := managerGroup.Group("/product")
		{
			productGroup.POST("/", productController.CreateProduct)
			productGroup.GET("/", productController.GetAllProducts)
			productGroup.GET("/:id_product", productController.GetProduct)
			productGroup.PATCH("/:id_product", productController.UpdateProduct)
			productGroup.DELETE("/:id_product", productController.DeleteProduct)
		}

		productInStoreGroup := managerGroup.Group("/storeProduct")
		{
			productInStoreGroup.POST("/", productInStoreController.CreateProductInStore)
			productInStoreGroup.GET("/", productInStoreController.GetAllProductInStores)
			productInStoreGroup.GET("/:UPC", productInStoreController.GetProductInStore)
			productInStoreGroup.PATCH("/:UPC", productInStoreController.UpdateProductInStore)
			productInStoreGroup.DELETE("/:UPC", productInStoreController.DeleteProductInStore)
		}

		receiptGroup := managerGroup.Group("/check")
		{
			receiptGroup.GET("/", receiptController.GetAllReceipts)
			receiptGroup.GET("/:check_number", receiptController.GetReceipt)
			receiptGroup.DELETE("/:check_number", receiptController.DeleteReceipt)
		}

		analyticsGroup := managerGroup.Group("/analytics")
		{
			analyticsGroup.GET("/salesPerCashier", analyticsController.GetSalesPerCashier)
			analyticsGroup.GET("/salesPerProduct", analyticsController.GetSalesPerProduct)
			analyticsGroup.GET("/averagePricePerCategory", analyticsController.GetAveragePricePerCategory)
			analyticsGroup.GET("/categorySalesPerCashier", analyticsController.GetCategorySalesPerCashier)
			analyticsGroup.GET("/registeredCustomersWhoHaveBeenServedByEveryCashier", analyticsController.GetRegisteredCustomersWhoHaveBeenServedByEveryCashier)
			analyticsGroup.GET("/cashiersWhoHaveSoldEveryProductInTheStore", analyticsController.GetCashiersWhoHaveSoldEveryProductInTheStore)
		}

		managerGroup.GET("/ping", controllers.Ping)
	}

	cashierGroup := router.Group("/cashier")
	cashierGroup.Use(middleware.CashierAuth())
	{
		cashierGroup.GET("/ping", controllers.Ping)

		categoryGroup := cashierGroup.Group("/category")
		{
			categoryGroup.GET("/", categoryController.GetAllCategories)
			categoryGroup.GET("/:category_number", categoryController.GetCategory)
		}

		customerCardGroup := cashierGroup.Group("/customerCard")
		{
			customerCardGroup.POST("/", customerCardController.CreateCustomerCard)
			customerCardGroup.GET("/", customerCardController.GetAllCustomerCards)
			customerCardGroup.GET("/:card_number", customerCardController.GetCustomerCard)
			customerCardGroup.PATCH("/:card_number", customerCardController.UpdateCustomerCard)
		}

		productGroup := cashierGroup.Group("/product")
		{
			productGroup.GET("/", productController.GetAllProducts)
			productGroup.GET("/:id_product", productController.GetProduct)
		}

		productInStoreGroup := cashierGroup.Group("/storeProduct")
		{
			productInStoreGroup.GET("/", productInStoreController.GetAllProductInStores)
			productInStoreGroup.GET("/:UPC", productInStoreController.GetProductInStore)
		}

		receiptGroup := cashierGroup.Group("/check")
		{
			receiptGroup.POST("/", receiptController.CreateReceipt)
			receiptGroup.GET("/", receiptController.GetAllReceiptsOfOneCashier)
			receiptGroup.GET("/:check_number", receiptController.GetReceipt)
		}
	}

	selfGroup := router.Group("/self")
	selfGroup.Use(middleware.EmployeeAuth())
	{
		selfGroup.GET("/", selfController.GetAllInfoAboutOneself)
		selfGroup.PATCH("/", selfController.UpdateInfoAboutOneself)
		selfGroup.PATCH("/credentials", selfController.UpdateCredentials)
	}

	return router
}
