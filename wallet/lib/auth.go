package lib

import (
	"context"
	"log"
	"net/http"
)

func WithReadHeader(base http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		ctx = context.WithValue(ctx, "X-Signature", r.Header.Get("X-Signature"))
		ctx = context.WithValue(ctx, "Content-Language", r.Header.Get("Content-Language"))

		r = r.WithContext(ctx)

		log.Printf("req: %+v, \n%+v", r.Body, r.Header)

		base.ServeHTTP(w, r)
	})
}
