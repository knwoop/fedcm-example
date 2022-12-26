package fedcm

import (
	"net/http"
)

const (
	SecFetchDestHeaderKey = "Sec-Fetch-Dest"
	SecFetchDestHeaderVal = "webidentity"
)

func AllowedHeader(r *http.Request) bool {
	if v := r.Header.Get(SecFetchDestHeaderKey); v == SecFetchDestHeaderVal {
		return true
	}
	return false
}
