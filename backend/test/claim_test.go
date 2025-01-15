package test

import (
	"backendproject/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"strings"
)

func TestClaimValidation(t *testing.T) {

	t.Run("Date is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Time{}, // เวลาว่าง
			Photo:         "photo_url",
			ProblemID:     1,
			ClaimStatusID: 1,
			UserID:        1,
			OrderID:       1,
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "Date is required")).To(BeTrue())
	})

	t.Run("Photo is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Now(),
			Photo:         "", // ไม่มี Photo
			ProblemID:     1,
			ClaimStatusID: 1,
			UserID:        1,
			OrderID:       1,
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "Photo is required")).To(BeTrue())
	})

	t.Run("ProblemID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Now(),
			Photo:         "photo_url",
			ProblemID:     0, // ไม่มี ProblemID
			ClaimStatusID: 1,
			UserID:        1,
			OrderID:       1,
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "ProblemID is required")).To(BeTrue())
	})

	t.Run("ClaimStatusID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Now(),
			Photo:         "photo_url",
			ProblemID:     1,
			ClaimStatusID: 0, // ไม่มี ClaimStatusID
			UserID:        1,
			OrderID:       1,
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "ClaimStatusID is required")).To(BeTrue())
	})

	t.Run("UserID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Now(),
			Photo:         "photo_url",
			ProblemID:     1,
			ClaimStatusID: 1,
			UserID:        0, // ไม่มี UserID
			OrderID:       1,
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "UserID is required")).To(BeTrue())
	})

	t.Run("OrderID is required", func(t *testing.T) {
		g := NewGomegaWithT(t)

		claim := entity.Claim{
			Date:          time.Now(),
			Photo:         "photo_url",
			ProblemID:     1,
			ClaimStatusID: 1,
			UserID:        1,
			OrderID:       0, // ไม่มี OrderID
		}

		ok, err := govalidator.ValidateStruct(claim)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(strings.Contains(err.Error(), "OrderID is required")).To(BeTrue())
	})
}
