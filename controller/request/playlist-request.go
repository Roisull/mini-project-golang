package request

type PlaylistRequest struct{
	Name  string `json:"name" form:"name"`
}