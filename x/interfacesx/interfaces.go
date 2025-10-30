package interfacesx

import (
	"rest-api/internals/model"
	"time"

	"github.com/gofrs/uuid"
)

type UserRegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserData struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullName"`
	UserRole  model.Role `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
}
