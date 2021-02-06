package mock

import (
	"net/http"
	"time"
)

// Response defines the response of a mock funtion
type Response struct {
	*http.Response
	err          error
	responseTime time.Duration
}
