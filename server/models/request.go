package models

type IDReq struct {
	ID string `json:"id" validate:"required"`
}

type LoginUserReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RenewAccessTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
