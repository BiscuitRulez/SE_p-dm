package unit

// import (
// 	"backendproject/entity"
// 	"testing"
// 	"time"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestOrder(t *testing.T) {
	
// 	t.Run(`User ID is required`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)
	
// 		order := entity.Order{
// 			OrderDate:     time.Now(),
// 			Status:        "ชำระเงินแล้ว",
// 			TotalPrice:    1200,
// 			UserID:        0,  // กำหนดเป็น 0 เพื่อทดสอบว่า UserID จำเป็นต้องมีค่า
// 			PaymentID:     3,
// 			CodeCollectID: 2,
// 			ShippingID:    3,
// 		}
	
// 		ok, err := govalidator.ValidateStruct(order)
	
// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("User ID is required"))  // คาดหวังว่า "User is required"
// 	})

// 	t.Run(`Status is required`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)

// 		order := entity.Order{
// 			OrderDate:    	time.Now(),
// 			Status: 	  	"",
// 			TotalPrice:		1200,
// 			UserID: 		1,
// 			PaymentID:		3,
// 			CodeCollectID: 	2,
// 			ShippingID: 	3,
// 		}

// 		ok, err := govalidator.ValidateStruct(order)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Status is required"))
// 	})

// 	t.Run(`Status is valid`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)

// 		order := entity.Order{
// 			OrderDate:    	time.Now(),
// 			Status: 	  	"ชำระเงินแล้ว",
// 			TotalPrice:		1200,
// 			UserID: 		1,
// 			PaymentID:		3,
// 			CodeCollectID: 	2,
// 			ShippingID: 	3,
// 		}

// 		ok, err := govalidator.ValidateStruct(order)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())
// 	})

// 	t.Run(`Total Price is required`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)

// 		order := entity.Order{
// 			OrderDate:    	time.Now(),
// 			Status: 	  	"ชำระเงินแล้ว",
// 			TotalPrice:		0,
// 			UserID: 		1,
// 			PaymentID:		3,
// 			CodeCollectID: 	2,
// 			ShippingID: 	3,
// 		}

// 		ok, err := govalidator.ValidateStruct(order)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Total Price is required"))
// 	})
	
// 	t.Run(`Total Price is valid`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)

// 		order := entity.Order{
// 			OrderDate:    	time.Now(),
// 			Status: 	  	"ชำระเงินแล้ว",
// 			TotalPrice:		1200,
// 			UserID: 		1,
// 			PaymentID:		3,
// 			CodeCollectID: 	2,
// 			ShippingID: 	3,
// 		}

// 		ok, err := govalidator.ValidateStruct(order)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())
// 	})

// 	t.Run(`Payment ID is required`, func(t *testing.T) {
// 		g := NewGomegaWithT(t)

// 		order := entity.Order{
// 			OrderDate:    	time.Now(),
// 			Status: 	  	"ชำระเงินแล้ว",
// 			TotalPrice:		1200,
// 			UserID: 		1,
// 			PaymentID:		0,
// 			CodeCollectID: 	2,
// 			ShippingID: 	3,
// 		}

// 		ok, err := govalidator.ValidateStruct(order)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Payment ID is required"))
// 	})
// }