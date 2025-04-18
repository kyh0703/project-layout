// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package model

import (
	"database/sql"
)

type Post struct {
	ID       int64          `json:"id"`
	UserID   int64          `json:"userId"`
	Title    string         `json:"title"`
	Content  sql.NullString `json:"content"`
	UpdateAt sql.NullString `json:"updateAt"`
	CreateAt sql.NullString `json:"createAt"`
}

type Token struct {
	ID           int64          `json:"id"`
	UserID       int64          `json:"userId"`
	RefreshToken string         `json:"refreshToken"`
	ExpiresIn    int64          `json:"expiresIn"`
	CreateAt     sql.NullString `json:"createAt"`
}

type User struct {
	ID       int64          `json:"id"`
	Email    string         `json:"email" validate:"required,email"`
	Password string         `json:"password" validate:"required,min=8,max=32"`
	Name     string         `json:"name"`
	Bio      sql.NullString `json:"bio"`
	UpdateAt sql.NullString `json:"updateAt"`
	CreateAt sql.NullString `json:"createAt"`
}
