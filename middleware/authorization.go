package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"spotify-api/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config config.Config, client *http.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		req, err := http.NewRequest("POST",
			config.GetTokenIssuerUrl(),
			strings.NewReader("grant_type=client_credentials"))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Authorization", fmt.Sprintf("Basic %s", getEncodedKeys(config)))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		var tokenResp map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
			fmt.Println("Error decoding response body:", err)
			return
		}

		context.Request.Header.Add("Authorization", tokenResp["access_token"].(string))
		fmt.Println("User got authorised. Token has been issued and added to the Authorization header")
		context.Next()
	}
}

func getEncodedKeys(config config.Config) string {
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%v:%v", config.GetClientId(), config.GetClientSecretKey())))
}
