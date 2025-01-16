package unit

// import (
// 	"backendproject/entity"
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestAddress(t *testing.T) {

// 	g := NewGomegaWithT(t)

// 	t.Run(`Full Address is required`, func(t *testing.T) {
// 		address := entity.Address{
// 			City:        "ชุมพลบุรี",
// 			Province:    "สุรินทร์",
// 			Postal_code: 30000,
// 			UserID:      1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		// Validation Check
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(ContainSubstring("Full Address is required"))
// 	})

// 	t.Run("Full_address is valid", func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: "56/1 บ้านแม้ว ต.สระขุด",
// 			City:         "ชุมพลบุรี",
// 			Province:     "สุรินทร์",
// 			Postal_code:  30000,
// 			UserID:       1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		// Assert validation passes
// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())
// 	})

// 	t.Run(`City is required`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: "56/1 บ้านแม้ว ต.สระขุด",
// 			City:         "",
// 			Province:     "สุรินทร์",
// 			Postal_code:  30000,
// 			UserID:       1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("City is required"))
// 	})

// 	t.Run(`City is valid`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: 	"56/1 บ้านแม้ว ต.สระขุด",
// 			City:     		"ชุมพลบุรี",
// 			Province:        "สุรินทร์",
// 			Postal_code:     30000,
// 			UserID:          1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`Province is required`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: "56/1 บ้านแม้ว ต.สระขุด",
// 			City:         "ชุมพลบุรี",
// 			Province:     "",
// 			Postal_code:  30000,
// 			UserID:       1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Province is required"))
// 	})

// 	t.Run(`Province is valid`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: 	"56/1 บ้านแม้ว ต.สระขุด",
// 			City:     		"ชุมพลบุรี",
// 			Province:        "สุรินทร์",
// 			Postal_code:     30000,
// 			UserID:          1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`Postal Code is required`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: "56/1 บ้านแม้ว ต.สระขุด",
// 			City:         "ชุมพลบุรี",
// 			Province:     "สุรินทร์",
// 			Postal_code:  0,
// 			UserID:       1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("Postal Code is required"))
// 	})

// 	t.Run(`Postal Code is valid`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: 	"56/1 บ้านแม้ว ต.สระขุด",
// 			City:     		"ชุมพลบุรี",
// 			Province:        "สุรินทร์",
// 			Postal_code:     30000,
// 			UserID:          1,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())

// 	})

// 	t.Run(`User ID is required`, func(t *testing.T) {
// 		address := entity.Address{
// 			Full_address: "56/1 บ้านแม้ว ต.สระขุด",
// 			City:         "ชุมพลบุรี",
// 			Province:     "สุรินทร์",
// 			Postal_code:  30000,
// 			UserID:       0,
// 		}

// 		ok, err := govalidator.ValidateStruct(address)

// 		g.Expect(ok).NotTo(BeTrue())
// 		g.Expect(err).NotTo(BeNil())
// 		g.Expect(err.Error()).To(ContainSubstring("User ID is required"))
// 	})

// }
