package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret")
var tokenName = "token"

func generateToken(c echo.Context, name string, password string, userType int) {
	tokenExpiryTime := time.Now().Add(time.Minute * 5)

	Claims := &Claims{
		Name:     name,
		Password: password,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.SetCookie(&http.Cookie{
		Name:    tokenName,
		Value:   signedToken,
		Expires: tokenExpiryTime,
		Secure:  false,
		Path:    "/",
	})
}

func resetToken(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:    tokenName,
		Value:   "",
		Expires: time.Now(),
		Secure:  false,
		Path:    "/",
	})
}

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessType := 1
		isValidToken := validateUserToken(c, accessType)
		if !isValidToken {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		return next(c)
	}
}

func validateUserToken(c echo.Context, accessType int) bool {
	isAccessTokenValid, name, password, userType := validateTokenFromCookies(c)
	fmt.Println(isAccessTokenValid, name, password, userType)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(c echo.Context) (bool, string, string, int) {
	cookie, err := c.Cookie(tokenName)
	if err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		fmt.Println(err)
		fmt.Println(parsedToken)
		if err == nil && parsedToken.Valid {
			return true, accessClaims.Name, accessClaims.Password, accessClaims.UserType
		}
	}
	return false, "", "", 0
}
