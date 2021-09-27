package util

import "net/http"

const (
	ErrInternalServerError      = "Internal Server Error"
	ErrUsernameAvailability     = "Username has already been taken"
	ErrEmailAvailability        = "Email has already been taken"
	ErrEmailUnverified          = "Email unverified"
	ErrFindUserByUsername       = "cannot find user, using given username"
	ErrGenerateToken            = "Unable to generate access token.Please try again later"
	ErrUserCreation             = "Error processing user creation"
	ErrDBPostgre                = "Error establishing a database connection"
	ErrBadRequest               = "Invalid request message"
	ErrLoginToken               = "Can't generate login access"
	ErrUnauthorized             = "Authentication failed. Invalid token"
	ErrBadTokenMethod           = "Unexpected signing method in auth token"
	ErrInvalidToken             = "Invalid token: authentication failed"
	ErrEmailSend                = "Cannot send email verification"
	ErrInvalidPassword          = "Invalid password"
	ErrInvalidUsernameEmail     = "Invalid username or email"
	ErrPasswordResetCodeExpired = "Confirmation code has expired. Please try generating a new code"
	ErrPasswordResetCodeInvalid = "Verification code provided is invalid"
	ErrPassordNotMatched        = "Password and password re-enter did not match"
	ErrUpdateAvatar             = "Filed to save avatar"

	MsgCreateUser                 = "User created sucessfully"
	MsgLoginSuccess               = "Login successful"
	MsgUserAvail                  = "Username available"
	MsgEmailAvail                 = "Email available"
	MsgVerifiedPasswordResetCode  = "Password reset code verified"
	MsgVerifyUserEmail            = "User email has been verified"
	MsgGeneratedPasswordResetCode = "Password reset code has been sent to email"
	MsgPasswordReset              = "Password has been updated"
	MsgUpdateAvatar               = "Avatar has been saved"
	MsgGetDocument                = "Data fetched"
)

var statusCodeByMsg = map[string]int{
	ErrInternalServerError:      http.StatusInternalServerError,
	ErrUsernameAvailability:     http.StatusConflict,
	ErrEmailAvailability:        http.StatusConflict,
	ErrUserCreation:             http.StatusGatewayTimeout,
	ErrDBPostgre:                http.StatusGatewayTimeout,
	ErrBadRequest:               http.StatusBadRequest,
	ErrLoginToken:               http.StatusUnprocessableEntity,
	ErrUnauthorized:             http.StatusUnauthorized,
	ErrBadTokenMethod:           http.StatusUnauthorized,
	ErrInvalidToken:             http.StatusUnauthorized,
	ErrEmailSend:                http.StatusBadGateway,
	ErrInvalidPassword:          http.StatusUnauthorized,
	ErrInvalidUsernameEmail:     http.StatusUnauthorized,
	ErrPasswordResetCodeExpired: http.StatusUnauthorized,
	ErrPasswordResetCodeInvalid: http.StatusUnauthorized,
	ErrPassordNotMatched:        http.StatusBadRequest,
	ErrUpdateAvatar:             http.StatusInternalServerError,
	ErrEmailUnverified:          http.StatusUnauthorized,

	MsgCreateUser:                 http.StatusCreated,
	MsgLoginSuccess:               http.StatusOK,
	MsgUserAvail:                  http.StatusOK,
	MsgEmailAvail:                 http.StatusOK,
	MsgVerifiedPasswordResetCode:  http.StatusOK,
	MsgGeneratedPasswordResetCode: http.StatusOK,
	MsgPasswordReset:              http.StatusOK,
	MsgUpdateAvatar:               http.StatusCreated,
	MsgVerifyUserEmail:            http.StatusOK,

	MsgGetDocument: http.StatusOK,
}

func StatusCode(errorstr string) int {
	return statusCodeByMsg[errorstr]
}
