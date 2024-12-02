package models

import()

type Claims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}