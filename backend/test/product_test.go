package unit

// import (
// 	"backendproject/entity"
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestProduct(t *testing.T) {

// 	g := NewGomegaWithT(t)

// 	t.Run(`Product Name is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Product Name is required"))
// 	})

// 	t.Run(`Product name is valid`, func(t *testing.T) {
// 		user := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(user)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`Description is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Description is required"))
// 	})

// 	t.Run(`Description is valid`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`Quantity is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        0,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Quantity is required"))
// 	})

// 	t.Run(`Quantity is valid`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`Image is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "",
// 			UserID:          1,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Image is required"))
// 	})

// 	t.Run(`Image is valid`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      1,
			
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`User ID is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          0,
// 			CatagoryID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("User ID is required"))
// 	})

// 	t.Run(`Catagory ID is required`, func(t *testing.T) {
// 		product := entity.Product{
// 			Name: "unit",
// 			Description:     "unit",
// 			Quantity:        1000,
// 			Image:           "sfbfh",
// 			UserID:          1,
// 			CatagoryID:      0,
// 		}

// 		ok, err := govalidator.ValidateStruct(product)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Catagory ID is required"))
// 	})
// }