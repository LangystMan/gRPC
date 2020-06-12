package api

import (
	"bytes"
	"context"
	"encoding/json"
	"gRPC/lib/config"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func init() {
	log.Logger = config.LoadLogger()
}

func WithCryptoPassHandler(base http.Handler, errorsIni *config.ErrorsIni) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uuidRFC4122, err := uuid.NewRandom()
		if err != nil {
			log.Error().Msg("Unable create UUID: " + err.Error())
		}

		uuidRFC4122 = uuid.Must(uuidRFC4122, nil)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "UUID", uuidRFC4122.String())
		ctx = context.WithValue(ctx, "X-Signature", r.Header.Get("X-Signature"))
		ctx = context.WithValue(ctx, "Content-Language", r.Header.Get("Content-Language"))
		ctx = context.WithValue(ctx, "ErrCfg", *errorsIni)

		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error().Str("UUID", uuidRFC4122.String()).Msg("Unable read request body: " + err.Error())
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		if !authorizeRequest(ioutil.NopCloser(bytes.NewBuffer(buf)), ctx) {

			_, err = w.Write(NewErrorReply(config.ERROR_REQUEST_UNAUTHORIZED, ctx))
			if err != nil {
				log.Warn().Msg("Unable send error reply")
			}

			return
		}

		defer func() {
			err = r.Body.Close()
			if err != nil {
				log.Error().Str("UUID", uuidRFC4122.String()).Msg("Unable close request body: " + err.Error())
			}
		}()

		base.ServeHTTP(w, r.WithContext(ctx))
	})

}

func NewErrorReply(errno int, ctx context.Context) []byte {

	var rpl struct {
		Errno int
		Error string
	}

	// Преобразуем пятизначные ошибки в трёхзначные (но с описанием, взятым у 5-значной)
	errnoTransform, err := strconv.ParseInt(strconv.Itoa(errno)[:3], 10, 64)
	if err != nil {
		log.Warn().Msg("Unable transform errno")
		rpl.Errno = errno
	} else {
		rpl.Errno = int(errnoTransform)
	}

	rpl.Error = ctx.Value("ErrCfg").(config.ErrorsIni)[ctx.Value("Content-Language").(string)][errno]

	msg, err := json.Marshal(rpl)
	if err != nil {
		log.Error().Str("UUID", ctx.Value("UUID").(string)).Msg("Unable" + err.Error())
	}

	if len(rpl.Error) == 0 {
		log.Warn().
			Str("UUID", ctx.Value("UUID").(string)).
			Msg("Unable set error description")
	}

	log.Trace().
		Str("UUID", ctx.Value("UUID").(string)).
		Msg("Rpl: " + string(msg))

	return msg
}

func authorizeRequest(body io.ReadCloser, ctx context.Context) bool {

	bodyReq, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error().Str("UUID", ctx.Value("UUID").(string)).Msg("Unable read request body: " + err.Error())
		return false
	}

	log.Trace().
		Str("UUID", ctx.Value("UUID").(string)).
		Str("Signature", ctx.Value("X-Signature").(string)).
		Msg("Req: " + string(bodyReq))

	err = body.Close()
	if err != nil {
		log.Error().Str("UUID", ctx.Value("UUID").(string)).Msg("Unable close request body")
	}

	return false
}
