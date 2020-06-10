package api

import (
	"context"
	"fmt"
	"gRPC/lib/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

func init() {
	log.Logger = config.LoadLogger()
}

func WithCryptoPassHandler(base http.Handler, setup *config.Setup) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uuidRFC4122, err := uuid.NewRandom()
		if err != nil {
			log.Warn().Msg(err.Error())
		}

		uuidRFC4122 = uuid.Must(uuidRFC4122, nil)
		requestInfo := gerRequestInfo(r)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal().Msg("Unable read request body")
		}

		defer func() {
			err = r.Body.Close()
			if err != nil {
				log.Fatal().Msg("Unable close request body")
			}
		}()

		ctx := r.Context()
		ctx = context.WithValue(ctx, "UUID", uuidRFC4122.String())
		ctx = context.WithValue(ctx, "X-Signature", r.Header.Get("X-Signature"))
		ctx = context.WithValue(ctx, "Content-Language", r.Header.Get("Content-Language"))

		log.Trace().
			Str("UUID", uuidRFC4122.String()).
			Str("X-Signature", ctx.Value("X-Signature").(string)).
			Msg(requestInfo + "Req: " + fmt.Sprintf("%s", body) + " |")

		if !authorize(body, &ctx) {
			log.Warn().
				Str("UUID", uuidRFC4122.String()).
				Msg("Unauthorized request")
			// TODO Реализовать ответ с 401
		}

		r = r.WithContext(ctx)
		base.ServeHTTP(w, r)
	})

}

func gerRequestInfo(r *http.Request) string {
	return r.Proto + " " + r.Method + " " + r.URL.Path + " | "
}

func authorize(body []byte, ctx *context.Context) bool {

	//TODO Реализовать проверку подписи сообщения

	return false
}
