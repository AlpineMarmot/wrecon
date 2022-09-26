package wrecon

import (
	"net/http"
)

type PingResponse struct {
	Site     Site
	Succeed  bool
	ErrMsg   string
	Response *http.Response
}

func SucceedPingResponse(site Site, response *http.Response) *PingResponse {
	return &PingResponse{
		Site:     site,
		Succeed:  true,
		ErrMsg:   "",
		Response: response,
	}
}

func FailedPingResponse(site Site, err error) *PingResponse {
	return &PingResponse{
		Site:    site,
		Succeed: false,
		ErrMsg:  err.Error(),
	}
}
