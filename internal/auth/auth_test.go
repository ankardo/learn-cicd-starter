package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey_ValidHeader(t *testing.T) {
    header := make(http.Header)
    header.Set("Authorization",
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
    headers := make(http.Header)
    headers.Set("Authorization","InvalidType invalidKey")
    _, err := GetAPIKey(headers)
    if err == nil || err != ErrMalformedAuthorizationHeader {
        t.Errorf("Unexpected error: got %v, want %v", err, ErrMalformedAuthorizationHeader)
    }
}

