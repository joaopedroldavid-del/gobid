// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package pgstore

import (
	uuid_uuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID           uuid_uuid.UUID     `json:"id"`
	UserName     string             `json:"user_name"`
	Email        string             `json:"email"`
	PasswordHash []byte             `json:"password_hash"`
	Bio          string             `json:"bio"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}
