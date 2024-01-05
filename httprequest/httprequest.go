package httprequest

import (
	"encoding/base64"
	"fmt"
	"spotify-api/config"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type HttpRequest interface {
	WithBearerToken(context *gin.Context) HttpRequest
}

type httprequest struct {
	requestAgent *gorequest.SuperAgent
	config       config.Config
}

func NewHttpRequest(config config.Config) HttpRequest {
	return &httprequest{
		requestAgent: gorequest.New(),
		config:       config,
	}
}

func (r *httprequest) WithBearerToken(context *gin.Context) HttpRequest {
	bearerToken := fmt.Sprintf("Basic %s", r.getEncodedKeys())
	r.requestAgent.Set("Authorization", bearerToken)
	return r
}

func (r httprequest) getEncodedKeys() string {
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%v:%v", r.config.GetClientId(), r.config.GetClientSecretKey())))
}
