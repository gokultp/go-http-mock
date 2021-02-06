package mock

import "net/http"

// Mock implements a http rount tripper
type Mock struct {
	m            mockMap
	oldTransport http.RoundTripper
}

// New returns a new instance of the mock
func New() *Mock {
	return &Mock{
		m:            mockMap{},
		oldTransport: http.DefaultTransport,
	}
}

func (m *Mock) OnURL(url string) Predicate {
	return newValuePredicate(predicateURL, url)
}
