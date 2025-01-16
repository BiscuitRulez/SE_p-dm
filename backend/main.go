package main

import (
	"os"

	"log"

	"net/http"

	"backendproject/controller"

	"github.com/gin-gonic/gin"

	payment "backendproject/controller/payment"

	claim "backendproject/controller/claim"

	"backendproject/config"

	"backendproject/controller/codes"

	"backendproject/controller/user"

	"backendproject/middlewares"
)

const PORT = "8000"

func main() {

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", 0755)
	}

	// open connection database

	config.ConnectionDB()

	// Generate databases

	config.SetupDatabase()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		if err := os.MkdirAll("uploads", 0755); err != nil {
			log.Fatal("Failed to create uploads directory:", err)
		}
	}

	r.Use(CORSMiddleware())

	// Auth Route

	r.POST("/signup", user.SignUp)

	r.POST("/signupadmin", user.SignUpAdmin)

	r.POST("/signin", user.SignIn)

	router := r.Group("/")
	router.GET("/payments", payment.GetPayments)
	router.GET("/payments/:id", payment.GetPaymentByID)
	router.POST("/payments/:id", payment.CreatePayment)
	router.PUT("/payments/:id", payment.UpdatePaymentByID)
	router.DELETE("/payments/:id", payment.DeletePayment)

	router.GET("/paymentStatus", payment.GetPaymentStatus)

	router.GET("/paymentMethod", payment.GetPaymentMethod)

	router.POST("/history", controller.CreateHistory)
	router.GET("/history", controller.ListHistory)
	router.GET("/history/:id", controller.ListHistoryByID)

	router.POST("/historyDetail", controller.CreateHistoryDetail)
	router.GET("/historyDetail", controller.ListHistoryDetail)
	router.GET("/historyDetail/:id", controller.ListHistoryDetailByID)

	router.GET("/claim", claim.GetClaims)
	router.GET("/claim/:id", claim.GetClaimByID)
	router.POST("/claim/:id", claim.CreateClaim)
	router.PUT("/claim/:id", claim.UpdateClaimByID)
	router.DELETE("/claim/:id", claim.DeleteClaim)

	router.GET("/claimStatus", claim.GetClaimStatus)

	router.GET("/problem", claim.GetProblem)

	router.GET("/codes", codes.GetAll) // ดึงข้อมูล Codes ทั้งหมด
	router.GET("/codes/:id", codes.GetCodeById)
	router.POST("/codes", codes.CreateCode)    // สร้าง Code ใหม่
	router.PUT("/codes/:id", codes.UpdateCode) // อัปเดต Code ตาม ID
	router.PUT("/code-collect/:id", codes.UpdateCodeAfterCollect)
	router.DELETE("/codes/:id", codes.DeleteCode)
	router.POST("/code-collect/:userId/:codeId", codes.AddCodeToCollect)
	router.GET("/code-collect/:userId", codes.GetCollectedCodes)
	router.GET("/show-collect/:userId", codes.GetCollectedCodesToShow)

	{

		router.Use(middlewares.Authorizes())

		// User Route

		// router.GET("/codes", codes.GetAll) // ดึงข้อมูล Codes ทั้งหมด
		// router.GET("/codes/:id", codes.GetCodeById)
		// router.POST("/codes", codes.CreateCode)    // สร้าง Code ใหม่
		// router.PUT("/codes/:id", codes.UpdateCode) // อัปเดต Code ตาม ID
		// router.PUT("/code-collect/:id", codes.UpdateCodeAfterCollect)
		// router.DELETE("/codes/:id", codes.DeleteCode)
		// router.POST("/code-collect/:userId/:codeId", codes.AddCodeToCollect)
		// router.GET("/code-collect/:userId", codes.GetCollectedCodes)
		// router.GET("/show-collect/:userId", codes.GetCollectedCodesToShow)

		router.PUT("/user/:id", user.Update)
		router.GET("/users", user.GetAll)
		router.GET("/user/:id", user.Get)
		router.DELETE("/user/:id", user.Delete)

		router.POST("/address", controller.AddAddressController)
		router.GET("/address/:id", controller.GetAddressesByUserId)
		router.GET("/address", controller.GetAllAddress)
		router.PUT("/address/:id", controller.UpdateAddressByID)
		router.DELETE("/address/:id", controller.DeleteAddress)

		router.GET("/catagory", controller.GetAllCatagory)
		router.GET("/catagory/:id", controller.GetCatagoryByID)
		router.POST("/catagory/:id", controller.CreateCatagory)
		router.PUT("/catagory/:id", controller.UpdateCatagoryByID)
		router.DELETE("/catagory/:id", controller.DeleteCatagory)

		// Product routes
		router.GET("/product", controller.GetAllProduct)
		router.GET("/product/:id", controller.GetProductByID)
		router.POST("/product", controller.CreateProduct)
		router.PUT("/product/:id", controller.UpdateProductByID)
		router.DELETE("/product/:id", controller.DeleteProduct)
		// Cart routes
		router.POST("/cart", controller.AddToCart)
		router.GET("/cart", controller.GetCart)
		router.PUT("/cart/:id", controller.UpdateCartItem)

		router.GET("/stocks", controller.GetStocksByProductID)
		router.POST("/stock", controller.Create)
		router.PUT("/stock/:id", controller.Update)
		router.GET("/stock", controller.GetAll)
		router.GET("/stock/:id", controller.Get)
		router.DELETE("/stock/:id", controller.Delete)

		router.GET("/:userId", controller.GetPointsByUserID)
		router.POST("/redeem", controller.RedeemPoints)
		router.POST("/earn", controller.EarnPoints)
		router.PUT("/:pointId", controller.UpdatePoints)
		router.DELETE("/:pointId", controller.DeletePoints)

		// Order routes
		router.POST("/orders", controller.CreateOrder)
		router.GET("/orders", controller.GetOrders)

		// Review routes
		router.POST("/reviews", controller.CreateReview)
		router.GET("/products/:id/reviews", controller.GetProductReviews)
		router.GET("/products/:id/reviews/analytics", controller.GetReviewAnalytics)
		router.POST("/reviews/:id/vote", controller.VoteHelpful)
		router.POST("/reviews/upload", controller.UploadImage)

		router.GET("/shippingstatus", controller.GetAllShippingStatus)
		router.GET("/shippingstatus/:id", controller.GetShippingStatusByID)
		router.PUT("/shippingstatus/:id", controller.UpdateShippingStatusByID)
		router.DELETE("/shippingstatus/:id", controller.DeleteShippingStatus)

		router.GET("/shipping", controller.GetAllShipping)
		router.GET("/shipping/:id", controller.GetShippingByID)
		router.POST("/shipping/:id", controller.CreateShipping)
		router.PUT("/shipping/:id", controller.UpdateShippingByID)
		router.DELETE("/shipping/:id", controller.DeleteShipping)

		router.GET("/tags", controller.GetAllTags)
		router.GET("/tags/:id", controller.GetTagsByID)
		router.POST("/tags/:id", controller.CreateTags)
		router.PUT("/tags/:id", controller.UpdateTagsByID)
		router.DELETE("/tags/:id", controller.DeleteTags)

		//Payment Route
		// router.GET("/payments", payment.GetPayments)
		// router.GET("/payments/:id", payment.GetPaymentByID)
		// router.POST("/payments/:id", payment.CreatePayment)
		// router.PUT("/payments/:id", payment.UpdatePaymentByUserID)
		// router.DELETE("/payments/:id", payment.DeletePayment)

		// router.GET("/paymentStatus", payment.GetPaymentStatus)

		// router.GET("/paymentMethod", payment.GetPaymentMethod)

		// Cart Routes

		router.GET("/admins", user.GetAdmin)

	}

	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

	})

	// Run the server

	r.Run("localhost:" + PORT)

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}
