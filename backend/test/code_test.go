package test

import (
	"backendproject/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCodeTopic(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`CodeTopic is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "", // Invalid
			CodeDescription: "Give away for new user",
			Discount:       100,
			Quantity:       10,
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "Active",
			CodePicture:    "picture_url",
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("CodeTopic is required"))
	})
}

func TestCodeDescription(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`CodeDescription is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "DISCOUNT 100 BAHT",
			CodeDescription: "", // Invalid
			Discount:       100,
			Quantity:       10,
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "Active",
			CodePicture:    "picture_url",
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("CodeDescription is required"))
	})
}

func TestDiscount(t *testing.T) {
	g := NewGomegaWithT(t)

t.Run(`Discount is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "DISCOUNT 100 BAHT",
			CodeDescription: "Give away for new user",
			Discount:       0, // Invalid
			Quantity:       10,
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "Active",
			CodePicture:    "picture_url",
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Discount is required"))
	})
}

func TestQuantityCode(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Quantity is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "DISCOUNT 100 BAHT",
			CodeDescription: "Give away for new user",
			Discount:       100,
			Quantity:       0, // Invalid
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "Active",
			CodePicture:    "picture_url",
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Quantity is required"))
	})
}

func TestCodeStatus(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`CodeStatus is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "DISCOUNT 100 BAHT",
			CodeDescription: "Give away for new user",
			Discount:       100,
			Quantity:       10,
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "", // Invalid
			CodePicture:    "picture_url",
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("CodeStatus is required"))
	})
}

func TestCodePicture(t *testing.T) {
	g := NewGomegaWithT(t)

t.Run(`CodePicture is required`, func(t *testing.T) {
		code := entity.Codes{
			CodeTopic:      "DISCOUNT 100 BAHT",
			CodeDescription: "Give away for new user",
			Discount:       100,
			Quantity:       10,
			DateStart:      time.Now().Add(24 * time.Hour),
			DateEnd:        time.Now().Add(48 * time.Hour),
			CodeStatus:     "Active",
			CodePicture:    "", // Invalid
		}

		ok, err := govalidator.ValidateStruct(code)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("CodePicture is required"))
	})
}
