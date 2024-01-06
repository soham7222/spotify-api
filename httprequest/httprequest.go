package httprequest

import (
	"encoding/base64"
	"fmt"
	"spotify-api/config"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type HttpRequest interface {
	WithNewBearerToken() HttpRequest
	WithGrantType() HttpRequest
	WithContext(context *gin.Context) HttpRequest
	Post(url string) (gorequest.Response, string, []error)
	GetAgent() *gorequest.SuperAgent
}

type httprequest struct {
	requestAgent *gorequest.SuperAgent
	config       config.Config
	context      *gin.Context
}

func NewHttpRequest(config config.Config, requestAgent *gorequest.SuperAgent) HttpRequest {
	return &httprequest{
		requestAgent: requestAgent,
		config:       config,
	}
}

func (r httprequest) Post(url string) (gorequest.Response, string, []error) {
	r.requestAgent.Post(url)
	return r.requestAgent.End()
}

func (r httprequest) GetAgent() *gorequest.SuperAgent {
	return r.requestAgent
}

func (r httprequest) WithNewBearerToken() HttpRequest {
	bearerToken := fmt.Sprintf("Basic %s", r.getEncodedKeys())
	r.requestAgent.Set("Authorization", bearerToken)
	fmt.Println(r.requestAgent.Header.Get("Authorization"))
	return r
}

func (r httprequest) WithContext(context *gin.Context) HttpRequest {
	r.context = context
	return r
}

func (r httprequest) WithGrantType() HttpRequest {
	r.requestAgent.Send("grant_type=client_credentials")
	return r
}

func (r httprequest) getEncodedKeys() string {
	return base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%v:%v", r.config.GetClientId(), r.config.GetClientSecretKey())))
}
