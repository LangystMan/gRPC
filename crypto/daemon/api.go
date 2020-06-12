package main

import (
	pb "gRPC/assembly/api"
	daemon "gRPC/crypto/api"
	"gRPC/lib/api"
	cfg "gRPC/lib/config"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	log.Logger = cfg.LoadLogger()
}

func main() {

	set, errFile, err := cfg.LoadSetup("cryptopass.ini", "errors.ini")
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	addr, err := set.GetAddress()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	srv := daemon.ApiDaemon{}
	server := pb.NewApiServer(&srv, nil)
	wrapped := api.WithCryptoPassHandler(server, errFile)

	err = http.ListenAndServe(addr, wrapped)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
