package main

import (
	"net/http"

	"backendproject/controller"

	"github.com/gin-gonic/gin"

	payment `backendproject/controller/payment`

	claim `backendproject/controller/claim`

	"backendproject/config"

	"backendproject/controller/codes"

	"backendproject/controller/user"

	"backendproject/middlewares"
)

const PORT = "8000"

func main() {

	// open connection database

	config.ConnectionDB()

	// Generate databases

	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Auth Route

	r.POST("/signup", user.SignUp)

	r.POST("/signupadmin", user.SignUpAdmin)

	r.POST("/signin", user.SignIn)

	router := r.Group("/")
		r.GET("/payments", payment.GetPayments)
		r.GET("/payments/:id", payment.GetPaymentByID)
		r.POST("/payments/:id", payment.CreatePayment)
		r.PUT("/payments/:id", payment.UpdatePaymentByID)
		r.DELETE("/payments/:id", payment.DeletePayment)

		r.GET("/paymentStatus", payment.GetPaymentStatus)

		r.GET("/paymentMethod", payment.GetPaymentMethod)


		r.POST("/history", controller.CreateHistory)
		r.GET("/history", controller.ListHistory)
		r.GET("/history/:id", controller.ListHistoryByID)

		r.POST("/historyDetail", controller.CreateHistoryDetail)
		r.GET("/historyDetail", controller.ListHistoryDetail)
		r.GET("/historyDetail/:id", controller.ListHistoryDetailByID)

		r.GET("/claim", claim.GetClaims)
		r.GET("/claim/:id", claim.GetClaimByID)
		r.POST("/claim/:id", claim.CreateClaim)
		r.PUT("/claim/:id", claim.UpdateClaimByID)
		r.DELETE("/claim/:id", claim.DeleteClaim)

		r.GET("/claimStatus", claim.GetClaimStatus)

		r.GET("/problem", claim.GetProblem)

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


		//Payment Route
		// router.GET("/payments", payment.GetPayments)
		// router.GET("/payments/:id", payment.GetPaymentByID)
		// router.POST("/payments/:id", payment.CreatePayment)
		// router.PUT("/payments/:id", payment.UpdatePaymentByUserID)
		// router.DELETE("/payments/:id", payment.DeletePayment)

		// router.GET("/paymentStatus", payment.GetPaymentStatus)

		// router.GET("/paymentMethod", payment.GetPaymentMethod)


		// Cart Routes
		router.GET("/carts", controller.GetAll)
		router.GET("/cart/:id", controller.Get)
		router.POST("/cart", controller.Create)
		router.PUT("/cart/:id", controller.Update)
		router.DELETE("/cart/:id", controller.Delete)

		// Product Routes
		router.GET("/products", controller.GetAll)
		router.GET("/product/:id", controller.Get)
		router.POST("/product", controller.Create)
		router.PUT("/product/:id", controller.Update)
		router.DELETE("/product/:id", controller.Delete)

		// Category Routes
		router.GET("/categories", controller.GetAll)
		router.GET("/category/:id", controller.Get)
		router.POST("/category", controller.Create)
		router.PUT("/category/:id", controller.Update)
		router.DELETE("/category/:id", controller.Delete)

		// Tag Routes
		router.GET("/tags", controller.GetAll)
		router.GET("/tag/:id", controller.Get)
		router.POST("/tag", controller.Create)
		router.PUT("/tag/:id", controller.Update)
		router.DELETE("/tag/:id", controller.Delete)

		// Order Routes
		router.GET("/orders", controller.GetAll)
		router.GET("/order/:id", controller.Get)
		router.POST("/order", controller.Create)
		router.PUT("/order/:id", controller.Update)
		router.DELETE("/order/:id", controller.Delete)

		// Stock Routes
		router.GET("/stocks", controller.GetAll)
		router.GET("/stock/:id", controller.Get)
		router.POST("/stock", controller.Create)
		router.PUT("/stock/:id", controller.Update)
		router.DELETE("/stock/:id", controller.Delete)

		// OrderItem Routes
		router.GET("/orderitems", controller.GetAll)
		router.GET("/orderitem/:id", controller.Get)
		router.POST("/orderitem", controller.Create)
		router.PUT("/orderitem/:id", controller.Update)
		router.DELETE("/orderitem/:id", controller.Delete)

		// ProductTags Routes
		router.GET("/producttags", controller.GetAll)
		router.GET("/producttag/:id", controller.Get)
		router.POST("/producttag", controller.Create)
		router.PUT("/producttag/:id", controller.Update)
		router.DELETE("/producttag/:id", controller.Delete)

		// Review Routes
		router.GET("/reviews", controller.GetAll)
		router.GET("/review/:id", controller.Get)
		router.POST("/review", controller.Create)
		router.PUT("/review/:id", controller.Update)
		router.DELETE("/review/:id", controller.Delete)

		// Reviewlikes Routes
		router.GET("/reviewlikes", controller.GetAll)
		router.GET("/reviewlike/:id", controller.Get)
		router.POST("/reviewlike", controller.Create)
		router.PUT("/reviewlike/:id", controller.Update)
		router.DELETE("/reviewlike/:id", controller.Delete)

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
