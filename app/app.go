package app

// AuthRegister adalah struktur data yang digunakan untuk mewakili data yang diperlukan saat melakukan registrasi pengguna.
type AuthRegister struct {
	Id       string `valid:"required" json:"id"`
	Username string `valid:"required" json:"username"`
	Email    string `valid:"required,email" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}

// AuthLogin adalah struktur data yang digunakan untuk mewakili data yang diperlukan saat melakukan login pengguna.
type AuthLogin struct {
	Email    string `valid:"required,email" json:"email"`
	Password string `valid:"required,minstringlength(6)" json:"password"`
}
