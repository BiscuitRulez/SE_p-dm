package unit

import (
	"backendproject/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"strings"
)

func TestPaymentValidation(t *testing.T) {

	t.Run("Date is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		payment := entity.Payment{
			Date:          time.Time{}, // เวลาว่าง
			UserID:        1,
			PaymentMethodID: 1,
			PaymentStatusID: 1,
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "Date is required")).To(BeTrue())
	})

	t.Run("UserID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		payment := entity.Payment{
			Date:          time.Now(),
			// UserID:        0, // ไม่มี UserID
			PaymentMethodID: 1,
			PaymentStatusID: 1,
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "UserID is required")).To(BeTrue())
	})

	t.Run("PaymentMethodID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		payment := entity.Payment{
			Date:          time.Now(),
			UserID:        1,
			// PaymentMethodID: 0, // ไม่มี PaymentMethodID
			PaymentStatusID: 1,
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "PaymentMethodID is required")).To(BeTrue())
	})

	t.Run("PaymentStatusID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		payment := entity.Payment{
			Date:          time.Now(),
			UserID:        1,
			PaymentMethodID: 1,
			// PaymentStatusID: 0, // ไม่มี PaymentStatusID
		}

		ok, err := govalidator.ValidateStruct(payment)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "PaymentStatusID is required")).To(BeTrue())
	})
}
