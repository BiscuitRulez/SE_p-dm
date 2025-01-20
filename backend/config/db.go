package config

import (
	"fmt"

	"time"

	"backendproject/entity"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func ConnectionDB() {

	database, err := gorm.Open(sqlite.Open("se.db?cache=shared"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database :(")

	}

	fmt.Println("connected database :)")

	db = database

}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.ShippingStatus{},
		&entity.Shipping{},
		&entity.Address{},
		&entity.Catagory{},
		&entity.History{},
		&entity.HistoryDetail{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Payment{},
		&entity.PaymentMethod{},
		&entity.PaymentStatus{},
		&entity.Claim{},
		&entity.ClaimStatus{},
		&entity.Problem{},
		&entity.Product{},
		&entity.Stock{},
		&entity.Users{},
		&entity.Codes{},
		&entity.CodeCollectors{},
		&entity.Point{},
		&entity.PointPolicy{},
		&entity.PointTran{},
	)

	hashedPassword, _ := HashPassword("123456")

	BirthDay, _ := time.Parse("2006-01-02", "1988-11-12")
	users := []entity.Users{
		{
			FirstName:   "Kunlasatri",
			LastName:    "Kramoncham",
			Email:       "May@gmail.com",
			Password:    hashedPassword,
			BirthDay:    BirthDay,
			PhoneNumber: "0615871759",
			Role:        "admin",
			PointID:     1,
		},
		{
			FirstName:   "a",
			LastName:    "aa",
			Email:       "User@gmail.com",
			Password:    hashedPassword,
			BirthDay:    BirthDay,
			PhoneNumber: "0987654321",
			Role:        "user",
			PointID:     1,
		},
	}
	
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("Failed to seed user %s: %v\n", user.FirstName, err)
		}
	}
	
//---------------------------------------------------------------------------------------------------------------//
	payments := []entity.Payment{
		{PaymentMethodID: 1, PaymentStatusID: 1, Date: time.Now(), UserID: 1,AddressID: 1},
		{PaymentMethodID: 1, PaymentStatusID: 2, Date: time.Now(), UserID: 1,AddressID: 2},
		{PaymentMethodID: 2, PaymentStatusID: 1, Date: time.Now(), UserID: 2,AddressID: 3},
		{PaymentMethodID: 2, PaymentStatusID: 2, Date: time.Now(), UserID: 3,AddressID: 2},
	}

	for _, payment := range payments {
		// ใช้ FirstOrCreate พร้อมกับเงื่อนไขที่ครอบคลุม PaymentMethod และ UserID
		db.FirstOrCreate(&payment, entity.Payment{
			PaymentMethod: payment.PaymentMethod,
			UserID:        payment.UserID,
			PaymentStatus: payment.PaymentStatus, // เพิ่ม PaymentStatus เพื่อให้สร้างข้อมูลครบทุกแถว
		})
	}

	PaymentMethod := []entity.PaymentMethod{
		{PaymentMethod: "Credit Card"},
		{PaymentMethod: "Prompay"},
	}
	for _, pay := range PaymentMethod {
		db.FirstOrCreate(&pay, &entity.PaymentMethod{PaymentMethod: pay.PaymentMethod})
	}

	PaymentStatus := []entity.PaymentStatus{
		{PaymentStatus: "Finished"},
		{PaymentStatus: "Paying"},
	}
	for _, pasy := range PaymentStatus {
		db.FirstOrCreate(&pasy, &entity.PaymentStatus{PaymentStatus: pasy.PaymentStatus})
	}

	//---------------------------------------------------------------------------------------------------------------//

	shipping := []entity.Shipping{
		{Name: "Fresh Express", Fee: 80},
		{Name: "Curry Express", Fee: 90},
		{Name: "You&Me Express", Fee: 99},
		{Name: "Thalund Post", Fee: 80},
	}
	for _, pkg := range shipping {
		db.FirstOrCreate(&pkg, entity.Shipping{Name: pkg.Name})
	}

	//---------------------------------------------------------------------------------------------------------------//

	// ข้อมูลที่ต้องการเพิ่มลงในตาราง history
	histories := []entity.History{
		{OrderDate: time.Now(), PointsEarned: 10.5, PointsRedeemed: 5.0, TotalAmount: 250.0, UserID: 1, OrderID: 1},
		{OrderDate: time.Now(), PointsEarned: 15.0, PointsRedeemed: 7.0, TotalAmount: 300.0, UserID: 2, OrderID: 2},
		{OrderDate: time.Now(), PointsEarned: 20.0, PointsRedeemed: 10.0, TotalAmount: 500.0, UserID: 3, OrderID: 3},
		{OrderDate: time.Now(), PointsEarned: 25.0, PointsRedeemed: 12.0, TotalAmount: 600.0, UserID: 4, OrderID: 4},
	}

	// ใช้ loop เพื่อเพิ่มข้อมูลในฐานข้อมูล
	for _, history := range histories {
		// ใช้ FirstOrCreate เพื่อหลีกเลี่ยงการเพิ่มข้อมูลซ้ำ
		db.FirstOrCreate(&history, entity.History{
			UserID:  history.UserID,
			OrderID: history.OrderID,
		})
	}

	historyDetails := []entity.HistoryDetail{
		{ProductName: "Chair", Quantity: 2, PricePerUnit: 500.0, SubTotal: 1000.0, StockID: 1, HistoryID: 1},
		{ProductName: "Table", Quantity: 1, PricePerUnit: 1500.0, SubTotal: 1500.0, StockID: 2, HistoryID: 2},
		{ProductName: "Sofa", Quantity: 1, PricePerUnit: 2000.0, SubTotal: 2000.0, StockID: 3, HistoryID: 3},
		{ProductName: "Lamp", Quantity: 3, PricePerUnit: 300.0, SubTotal: 900.0, StockID: 4, HistoryID: 4},
	}

	// ใช้ loop เพื่อเพิ่มข้อมูลในฐานข้อมูล
	for _, detail := range historyDetails {
		// ใช้ FirstOrCreate เพื่อหลีกเลี่ยงการเพิ่มข้อมูลซ้ำ
		db.FirstOrCreate(&detail, entity.HistoryDetail{
			ProductName: detail.ProductName,
			StockID:     detail.StockID,
			HistoryID:   detail.HistoryID,
		})
	}

	//----------------------------------------------------------------------------------------------------------------//

	claimStatuses := []entity.ClaimStatus{
		{ClaimStatus: "Pending"},
		{ClaimStatus: "Approved"},
		{ClaimStatus: "Rejected"},
	}
	for _, status := range claimStatuses {
		db.FirstOrCreate(&status, &entity.ClaimStatus{ClaimStatus: status.ClaimStatus})
	}

	// เติมข้อมูลใน Problem
	problems := []entity.Problem{
		{Problem: "Damaged item"},
		{Problem: "Wrong item delivered"},
		{Problem: "Missing parts"},
	}
	for _, problem := range problems {
		db.FirstOrCreate(&problem, &entity.Problem{Problem: problem.Problem})
	}

	// เติมข้อมูลใน Claim
	claims := []entity.Claim{
		{
			Date:          time.Now(),
			Photo:         "photo1_url",
			ProblemID:     1,
			ClaimStatusID: 1,
			UserID:        1,
			OrderID:       1,
		},
		{
			Date:          time.Now(),
			Photo:         "photo2_url",
			ProblemID:     2, 
			ClaimStatusID: 2,
			UserID:        2,
			OrderID:       2,
		},
		{
			Date:          time.Now(),
			Photo:         "photo3_url",
			ProblemID:     3,
			ClaimStatusID: 3,
			UserID:        3,
			OrderID:       3,
		},
	}

	for _, claim := range claims {
		db.FirstOrCreate(&claim, entity.Claim{
			ProblemID:     claim.ProblemID,
			ClaimStatusID: claim.ClaimStatusID,
			UserID:        claim.UserID,
			OrderID:       claim.OrderID,
		})
	}


	problem := []entity.Problem{
		{Problem: "Damaged item",},
		{Problem: "Wrong item",},
		{Problem: "Item is not ready to use",},
	}

	for _, problem := range problem {
		db.FirstOrCreate(&problem, entity.Problem{
			
		})
	}
	//----------------------------------------------------------------------------------------------------------------//

	DateStart, _ := time.Parse("2006-01-02", "1998-12-25")

	DateEnd, _ := time.Parse("2006-01-02", "1998-12-31")

	Code := &entity.Codes{

		CodeTopic: "Christmas",

		CodeDescription: "event for christmas only",

		Discount: 500,

		Quantity: 10,

		DateStart: DateStart,

		DateEnd: DateEnd,

		CodeStatus: "Active",

		Minimum: 2000,

		CodePicture: "data:image/png;base64,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",
	}

	Catagory := []entity.Catagory{
		{Name: "Furniture"},
		{Name: "Electronics"},
	}
	for _, cat := range Catagory {
		db.FirstOrCreate(&cat, entity.Catagory{Name: cat.Name})
	}
	// db.FirstOrCreate(User, &entity.Users{

	// 	Email: "admin@gmail.com",
	// })
////////////
	Stock := &entity.Stock{

		ID:        1,
		Price:     850,
		Quantity:  100,
		Color:     "ดำ",
		ShapeSize: "เหลี่ยม",
		Image:     "C:/Users/Home/Desktop/programming/SE/SE-67/frontend/src/assets/sofa.jpg",
		ProductID: 1,
	}

	db.FirstOrCreate(Stock, &entity.Stock{

		ID: 1,
	})

	// Product := &entity.Product{

	// 	ID:          1,
	// 	Name:        "โซฟา",
	// 	Description: "โซฟา 3 ที่นั่ง รุ่น Junie หุ้มด้วยผ้านำเข้าจากต่างประเทศ นุ่มสบาย ไม่หดตัว อายุการใช้งานทนทาน แข็งแรง ทนทาน ดีไซน์ทันสมัยและสีสันแตกต่าง",
	// 	Image:       "C:/Users/Home/Desktop/programming/SE/SE-67/frontend/src/assets/sofa.jpg",
	// 	UserID:      1,
	// 	CatagoryID:  1,
	// }

	// db.FirstOrCreate(Product, &entity.Product{

	// 	ID: 1,
	// })

	

	db.FirstOrCreate(Code, &entity.Codes{})
	codeCollectors := []entity.CodeCollectors{
		{UserID: 6, CodeID: 1},
		{UserID: 6, CodeID: 2},
		{UserID: 2, CodeID: 1},
	}
	if err := db.Create(&codeCollectors).Error; err != nil {
		return
	}

	Address := []entity.Address{
		{FullAddress: "123 Main St", City: "Metropolis", Province: "MT", PostalCode: "54321", UserID: 2},
		{FullAddress: "456 Elm St", City: "Gotham", Province: "GT", PostalCode: "12345", UserID: 2},
	}

	for _, addr := range Address {
		// ระบุเงื่อนไขที่ต้องการค้นหาอย่างชัดเจน
		db.FirstOrCreate(&addr, entity.Address{
			FullAddress: addr.FullAddress,
			City:        addr.City,
			Province:    addr.Province,
			PostalCode:  addr.PostalCode,
			UserID:      addr.UserID,
		})
	}

}
