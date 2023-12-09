package controller

import (
	"fmt"
	"mini-project-golang/config"
	"mini-project-golang/constants"
	"mini-project-golang/middleware"
	"mini-project-golang/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	// "golang.org/x/crypto/bcrypt"
)

// func GetUsersHandler(c echo.Context) error {
// 	return GetUsersController(c)
// }

// get ID from JWT
func GetIdFromJwt(tokenString string) (string, error) {
	// Mendeklarasikan struktur Claims untuk menyimpan klaim dari token JWT
	claims := jwt.MapClaims{}

	// Mendekode token JWT
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan Anda menggunakan algoritma dan kunci yang sama yang digunakan untuk menghasilkan token
		return []byte(constants.SECRET_KEY_JWT), nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return "", err
	}

	// Periksa apakah token valid dan berisi klaim yang sesuai
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(string); ok {
			return id, nil
		}
	}

	return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
}

// get all users

func GetUsersController(c echo.Context) error {

	// token := c.Request().Header.Get("Authorization")

	// if token == "" {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
	// }

	// userId, err := GetIdFromJwt(token)

	// if err != nil {
	// 	return err
	// }

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success get all users",
		"users":   users,
	})

}

func GetUserByIdController(c echo.Context) error{
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	// Temukan pengguna dengan ID yang sesuai
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Kirim respons JSON dengan pengguna yang ditemukan
	return c.JSON(http.StatusOK, user)
}

// // fungsi untuk hash password
// func hashPassword(password string) (string, error) {
// 	// Menggunakan bcrypt untuk meng-hash kata sandi
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
// 	}
// 	return string(hashedPassword), nil
// }

func CreateUserController(c echo.Context) error {
	user := model.User{}

	c.Bind(&user)

	fmt.Printf("Data yang diterima dari Flutter: %+v\n", user)

	// // mengambil kata sandi dari request
	// password := c.FormValue("password")

	// // Saat membuat pengguna baru
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	// }
	// // Simpan hashedPassword di database
	// user.Password = string(hashedPassword)

	if err := config.DB.Save(&user).Error; err != nil {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserByIdController(c echo.Context) error{
	// Mendapatkan id dari parameter URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	// Temukan pengguna dengan ID yang sesuai
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	// Hapus pengguna dari database
	config.DB.Delete(&user)

	// Kirim respons JSON dengan pesan sukses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})

}

// Update User By Id
func UpdateUserByIdController(c echo.Context) error {
	// mendapatkan parameter ID dari URL
	id, err  := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messaage" : "Format Permintaan ID tidak sesuai",
		})
	}

	// menemukan ID yang diminta pada database
	var user model.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User dengan ID yang diminta tidak tersedia di database",
		})
	}

	// Binding data baru dari permintaan
	newUserData := new(model.User)
	if err := c.Bind(newUserData); err != nil {
		return err
	}

	// Memperbarui data pengguna
	user.Name = newUserData.Name
	user.Email = newUserData.Email
	user.Password = newUserData.Password

	// Simpan perubahan ke database
	config.DB.Save(&user)

	// Kirim respons JSON dengan pesan sukses dan data pengguna yang diperbarui
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User updated successfully",
		"user":    user,
	})
}

// func comparePasswords(hashedPassword, password string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return err
// }

func LoginController(c echo.Context) error {
	// mengambil data user
	user := model.User{}
	c.Bind(&user)

	// email := user.Email
	// password := user.Password

	// mengambil data user dari database
	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Login Failed (data ini tidak ada di database)",
			"erorr":   err.Error(),
		})
	}

	// err = comparePasswords(user.Password, password)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "Login Failed (gabisa menyocokkan kata sandi)",
	// 		"error":   err.Error(),
	// 	})
	// }

	// generate token yang sudah dibuat di middleware jwt
	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{

			"message": "Login Failed (gabisa generated token)",
			"erorr":   err.Error(),
		})
	}

	userResponse := model.UserResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{

		"message": "Login Successfully",
		"user":    userResponse,
	})
}
