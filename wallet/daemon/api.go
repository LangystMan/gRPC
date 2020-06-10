package main

import (
	pb "gRPC/assembly/device"
	"gRPC/lib/api"
	cfg "gRPC/lib/config"
	daemon "gRPC/wallet/api"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	log.Logger = cfg.LoadLogger()
}

func main() {

	set, err := cfg.LoadSetup("wallet.ini")
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	addr, err := set.GetAddress()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	srv := daemon.ApiDaemon{}
	Server := pb.NewDeviceServer(&srv, nil)
	wrapped := api.WithCryptoPassHandler(Server, set)

	err = http.ListenAndServe(addr, wrapped)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
