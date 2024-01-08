package middleware

import (
	"encoding/base64"
	"fmt"
	"spotify-api/config"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func AuthMiddleware(config config.Config, request *gorequest.SuperAgent) gin.HandlerFunc {
	return func(context *gin.Context) {
		bearerToken := fmt.Sprintf("Basic %s", getEncodedKeys(config))
		request.Post(config.GetTokenIssuerUrl())
		request.Set("Authorization", bearerToken)
		request.Send("grant_type=client_credentials")

		_, body, _ := request.End()

		js, err := simplejson.NewJson([]byte(body))
		if err != nil {
			fmt.Println("[Authorize] Error parsing Json!")
		}

		jsToken, exists := js.CheckGet("access_token")
		if !exists {
			fmt.Println("not working")
		}

		token, _ := jsToken.String()
		context.Request.Header.Add("Authorization", token)
		fmt.Println("User got authorised. Token has been issued and added to the Authorization header")
		context.Next()
	}
}

func getEncodedKeys(config config.Config) string {
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%v:%v", config.GetClientId(), config.GetClientSecretKey())))
}
