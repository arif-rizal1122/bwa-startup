package auth

import "github.com/dgrijalva/jwt-go"

// Service adalah interface yang mendefinisikan method yang harus ada pada service autentikasi.
type Service interface {
	GenerateToken(userID int) (string, error) // GenerateToken digunakan untuk menghasilkan token JWT berdasarkan userID.
}

// jwtService adalah implementasi dari Service interface untuk service autentikasi JWT.
type jwtService struct {
	// fields here // Hanya sebuah placeholder untuk kemungkinan field yang mungkin diperlukan di masa depan.
}


func NewJWTservice() *jwtService {
	return &jwtService{}
}


// SECRET_KEY adalah kunci rahasia yang digunakan untuk menandatangani token JWT.
var SECRET_KEY = []byte("tes_aja")

// GenerateToken adalah method untuk menghasilkan token JWT berdasarkan userID.
func (s *jwtService) GenerateToken(userID int) (string, error) {
	// Membuat map claim untuk menyimpan informasi yang akan disertakan dalam token.
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	// Membuat token JWT baru dengan menggunakan metode penandatanganan HS256 dan claim yang telah dibuat.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Menandatangani token dengan SECRET_KEY untuk menghasilkan signedToken.
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		// Jika terjadi error saat menandatangani token, kembalikan error tersebut.
		return signedToken, err
	}

	// Mengembalikan signedToken yang telah berhasil ditandatangani dan tanpa error.
	return signedToken, nil
}
