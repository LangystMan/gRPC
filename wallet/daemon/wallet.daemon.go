package main

import (
	pb "gRPC/assembly/device"
	"gRPC/lib/config"
	"gRPC/wallet/api"
	"gRPC/wallet/lib"
	"github.com/rs/zerolog/log"
	"net/http"
)

func init() {
	log.Logger = config.LoadLogger()
}

func main() {

	set, err := config.LoadSetup("wallet.ini")
	if err != nil {
		log.Panic().Msg(err.Error())
	}

	if err := set.ConnectGORM(); err != nil {
		log.Panic().Msg(err.Error())
	}

	srv := api.PbWalletDaemon{}
	WalletServer := pb.NewDeviceServer(&srv, nil)
	wrapped := lib.WithReadHeader(WalletServer)

	err = http.ListenAndServe(":9000", wrapped)
	if err != nil {
		log.Error().Msg(err.Error())
	}

}
