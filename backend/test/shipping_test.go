package unit

import (
	"backendproject/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestShipping(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Name is required`, func(t *testing.T) {
		shipping := entity.Shipping{
			Name: 				"",
			Fee:     			50,
			ShippingStatusID:   1,
		}

		ok, err := govalidator.ValidateStruct(shipping)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Name is required"))
	})

	t.Run(`Name is valid`, func(t *testing.T) {
		shipping := entity.Shipping{
			Name: 				"Flash Express",
			Fee:     			50,
			ShippingStatusID:   1,
		}

		ok, err := govalidator.ValidateStruct(shipping)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

	t.Run(`Fee is required`, func(t *testing.T) {
		shipping := entity.Shipping{
			Name: 				"Flash Express",
			Fee:     			0,
			ShippingStatusID:   1,
		}

		ok, err := govalidator.ValidateStruct(shipping)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Fee is required"))
	})

	t.Run(`Fee is valid`, func(t *testing.T) {
		shipping := entity.Shipping{
			Name: 				"Flash Express",
			Fee:     			50,
			ShippingStatusID:   1,
		}

		ok, err := govalidator.ValidateStruct(shipping)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

	t.Run(`Shipping Status ID is required`, func(t *testing.T) {
		shipping := entity.Shipping{
			Name: 				"Flash Express",
			Fee:     			50,
			ShippingStatusID:   0,
		}

		ok, err := govalidator.ValidateStruct(shipping)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Shipping Status is required"))
	})

}