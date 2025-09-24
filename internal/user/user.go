package user

import (
	"time"

	"github.com/go-faker/faker/v4"
)

type User struct {
	ID        string `avro:"id" json:"id" faker:"uuid_hyphenated"`
	Name      string `avro:"name" json:"name" faker:"name"`
	Email     string `avro:"email" json:"email" faker:"email"`
	Age       int32  `avro:"age" json:"age" faker:"boundary_start=18, boundary_end=80"`
	Biography string `avro:"biography" json:"biography" faker:"paragraph"`
	Quote     string `avro:"quote" json:"quote" faker:"sentence"`
	Active    bool   `avro:"active" json:"active"`
	CreatedAt int64  `avro:"created_at" json:"createdAt"`
}

func New() (*User, error) {
	var user User
	if err := faker.FakeData(&user); err != nil {
		return nil, err
	}
	user.CreatedAt = time.Now().UnixMilli()
	return &user, nil
}
