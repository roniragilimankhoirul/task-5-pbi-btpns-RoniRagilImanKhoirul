package helpers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// SECRET adalah kunci rahasia yang digunakan untuk penandatanganan token JWT.
var SECRET = "roniragilimankhoirul"

// Validation adalah fungsi untuk memvalidasi data menggunakan govalidator.
// Ini mengembalikan error jika validasi gagal.
func Validation(c *gin.Context, data interface{}) error {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return err
	}
	return err
}

// Encrypt_password adalah fungsi untuk mengenkripsi kata sandi menggunakan bcrypt.
// Ini mengembalikan string yang dienkripsi.
func Encrypt_password(c *gin.Context, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gagal mengenkripsi kata sandi"})
		return ""
	}
	return string(hash)
}

// Check_password adalah fungsi untuk memeriksa kecocokan kata sandi terenkripsi dengan kata sandi asli.
// Ini mengembalikan error jika kata sandi tidak cocok.
func Check_password(pass_1 string, pass_2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pass_1), []byte(pass_2))
	if err != nil {
		return err
	}
	return err
}

// Initialize_token adalah fungsi untuk menginisialisasi token JWT dengan ID pengguna.
// Ini mengembalikan token yang ditandatangani dan error jika terjadi kesalahan.
func Initialize_token(userid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(SECRET))
	return tokenStr, err
}
