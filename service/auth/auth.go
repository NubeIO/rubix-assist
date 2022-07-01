package auth

import (
	model "github.com/NubeIO/rubix-assist/pkg/assistmodel"
	jwt "github.com/appleboy/gin-jwt/v2"
)

func MapClaims(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			"uuid": v.Email,
			"role": v.Role,
		}
	}
	return jwt.MapClaims{}
}
