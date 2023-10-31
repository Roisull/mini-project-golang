package controller

import (
	"log"
	"mini-project-golang/config"
	"mini-project-golang/controller/request"
	"mini-project-golang/helper"
	"mini-project-golang/middleware"
	"mini-project-golang/model"
	"mini-project-golang/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePlaylistController(c echo.Context) error {
	// mengambil data dari request
	playlist := model.Playlist{}
	c.Bind(&playlist)

	// mendapatkan token dari header
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Kamu Belum Login")
	}

	// Mendapatkan userID dari token JWT
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token tidak valid")
	}

	// Set userID ke playlist
	playlist.UserID = uint(userID)

	// Simpan playlist ke database
	if err := config.DB.Create(&playlist).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Playlist created successfully",
		"playlist": playlist,
	})

}

func GetAllPlaylistController(c echo.Context) error {
	qLimit := c.QueryParam("limit")
	qPage := c.QueryParam("page")
	log.Println("limit:", qLimit, "page:", qPage)

	response, err := repositories.SelectPlaylist()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success recieve product data",
		"data":    response,
	})
}

func UpdatePlaylistController(c echo.Context) error{
	// Mendapatkan ID playlist dari URL
	id := c.Param("id")

	// Mengecek apakah playlist dengan ID tersebut ada dalam database
	var existingPlaylist model.Playlist
	result := config.DB.First(&existingPlaylist, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Playlist Tidak Ditemukan")
	}

	// mengambil data update dari request
	updateData := model.Playlist{}
	c.Bind(&updateData)

	// Perbarui data playlist
	existingPlaylist.Name = updateData.Name

	// Simpan perubahan ke database
	if err := config.DB.Save(&existingPlaylist).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update playlist")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Playlist updated successfully",
		"playlist": existingPlaylist,
	})

}

func DeletePlaylistController(c echo.Context) error{
	// mendapatkan ID playlist dari URL param
	id := c.Param("id")

	// mengecek apakah id yang di masukkan ada dalam database
	existingPlaylist := model.Playlist{}
	result := config.DB.First(&existingPlaylist, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Playlist tidak ditemukan")
	}

	// hapus playlist dari database
	config.DB.Delete(&existingPlaylist)

	// kirim response JSON dengan pesan sukses
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Playlist Delete Successfully",
	})
}


// engga kepake
func AddPlaylistController(c echo.Context) error{

	idToken, err := middleware.ExtractTokenUserId(c)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, helper.FailedResponse(err.Error()))
    }

	input := new(request.PlaylistRequest)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind "+errBind.Error()))
	}

	data := model.Playlist{
		Name: input.Name,
		UserID: uint(idToken),
	}

	if err = repositories.InsertPlaylist(data); err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data "+err.Error()))
	}

	//response message berhasil
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}