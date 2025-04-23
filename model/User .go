package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"-" bson:"password,omitempty"` // Hidden from JSON
}

// // Constructor with validation
// func NewUser(name, email, password string) (*User, error) {
// 	if name == "" || email == "" || password == "" {
// 		return nil, errors.New("name, email and password are required")
// 	}

// 	return &User{
// 		ID:       primitive.NewObjectID(),
// 		Name:     name,
// 		Email:    email,
// 		password: password, // Note: In real usage, you should hash this
// 	}, nil
// }

// // Getters and setters
// func (u *User) GetID() primitive.ObjectID {
// 	return u.ID
// }

// func (u *User) SetID(id primitive.ObjectID) {
// 	u.ID = id
// }

// func (u *User) GetName() string {
// 	return u.Name
// }

// func (u *User) SetName(name string) error {
// 	if name == "" {
// 		return errors.New("name cannot be empty")
// 	}
// 	u.Name = name
// 	return nil
// }

// func (u *User) GetEmail() string {
// 	return u.Email
// }

// func (u *User) SetEmail(email string) error {
// 	if email == "" {
// 		return errors.New("email cannot be empty")
// 	}
// 	// Add proper email validation here
// 	u.Email = email
// 	return nil
// }

// // Password should never be returned, only set and verified
// func (u *User) SetPassword(password string) error {
// 	if len(password) < 8 {
// 		return errors.New("password must be at least 8 characters")
// 	}
// 	u.password = password
// 	return nil
// }

// // Add this for password verification
// func (u *User) VerifyPassword(password string) bool {
// 	return u.password == password // Note: In real usage, compare hashes
// }
