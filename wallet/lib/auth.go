package lib

import (
	"context"
	"gRPC/lib/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func init() {
	log.Logger = config.LoadLogger()
}

func WithReadHeader(base http.Handler, setup *config.Setup) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uuidRFC4122, err := uuid.NewRandom()
		if err != nil {
			log.Warn().Msg(err.Error())
		}

		uuidRFC4122 = uuid.Must(uuidRFC4122, nil)
		requestInfo := gerRequestInfo(r)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "UUID", uuidRFC4122.String())
		ctx = context.WithValue(ctx, "X-Signature", r.Header.Get("X-Signature"))
		ctx = context.WithValue(ctx, "Content-Language", r.Header.Get("Content-Language"))

		log.Trace().
			Str("UUID", uuidRFC4122.String()).
			Str("X-Signature", ctx.Value("X-Signature").(string)).
			Msg(requestInfo + "Req: " + " |")

		if !authorize(r.Body, &ctx) {
			log.Error().
				Str("UUID", uuidRFC4122.String()).
				Msg("Unauthorized request: 401 |")
			// TODO Реализовать ответ с 401
		}

		r = r.WithContext(ctx)
		base.ServeHTTP(w, r)
	})

}

func gerRequestInfo(r *http.Request) string {
	return r.Proto + " " + r.Method + " " + r.URL.Path + " | "
}

func authorize(body io.ReadCloser, ctx *context.Context) bool {

	//TODO Реализовать проверку подписи сообщения

	return false
}
