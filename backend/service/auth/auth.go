package auth

import (
	"github.com/NubeIO/rubix-updater/model"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func MapClaims(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			"id":   v.Email,
			"role": v.Role,
		}
	}
	return jwt.MapClaims{}
}