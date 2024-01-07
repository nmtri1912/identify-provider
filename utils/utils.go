package utils

import (
	"fmt"
	"net/http"
	"net/url"
)

// Returns string encoded query string of the request
func GetQueryString(query url.Values) string {
	encoded := query.Encode()
	if len(encoded) > 0 {
		encoded = fmt.Sprintf("?%s", encoded)
	}
	return encoded
}

func RedirectWithQueryString(to string, query url.Values, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("%s%s", to, getQueryString(query)), http.StatusFound)
}

func getQueryString(query url.Values) string {
	encoded := query.Encode()
	if len(encoded) > 0 {
		encoded = fmt.Sprintf("?%s", encoded)
	}
	return encoded
}
