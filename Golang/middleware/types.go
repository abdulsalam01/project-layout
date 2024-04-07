package middleware

import "net/http"

// Middleware type represents a generic middleware function.
type Middleware func(http.HandlerFunc) http.HandlerFunc
