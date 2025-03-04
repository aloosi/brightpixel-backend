// models/user.go
package models

type User struct {
	UserID       int     `db:"user_id" json:"user_id"`
	LoginID      string  `db:"login_id" json:"login_id"`
	PasswordHash string  `db:"password_hash" json:"-"`
	Name         string  `db:"name" json:"name"`
	TelNo        string  `db:"tel_no" json:"tel_no"`
	Email        string  `db:"email" json:"email"`
	Address      string  `db:"address" json:"address"`
	CityCode     int     `db:"city_code" json:"city_code"`
	Balance      float64 `db:"balance" json:"balance"`
}
