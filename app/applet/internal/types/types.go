// Code generated by goctl. DO NOT EDIT.
package types

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmailLoginResponse struct {
	UserId int64 `json:"user_id"`
	Token  struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	} `json:"token"`
}

type EmailRegisterRequest struct {
	UserName string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type EmailRegisterResponse struct {
	UserId int64 `json:"user_id"`
	Token  struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	} `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type UserInfoResponse struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
}
