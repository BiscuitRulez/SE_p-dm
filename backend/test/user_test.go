package unit

import (
	"backendproject/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUserValidation(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Email is required and must be valid`, func(t *testing.T) { 
		user := entity.Users{
			Email:       "", // Invalid
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`FirstName is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "", // Invalid
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})

	t.Run(`LastName is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "", // Invalid
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("LastName is required"))
	})

	t.Run(`Role is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "", // Invalid
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Role is required"))
	})

	t.Run(`BirthDay is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Time{}, // Invalid
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("BirthDay is required"))
	})

	t.Run(`PointID is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     0, // Invalid
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Point is required"))
	})

	t.Run(`Password is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "",// Invalid
			PhoneNumber: "0123456789",
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Password is required"))
	})

	t.Run(`PhoneNumber is required`, func(t *testing.T) {
		user := entity.Users{
			Email:       "john.doe@example.com",
			FirstName:   "John",
			LastName:    "Doe",
			Password:    "password123",
			PhoneNumber: "",// Invalid
			Role:        "User",
			BirthDay:    time.Now().AddDate(-20, 0, 0),
			PointID:     1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("PhoneNumber is required"))
	})
}
