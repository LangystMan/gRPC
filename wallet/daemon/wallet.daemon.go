package main

import (
	pb "gRPC/assembly/device"
	"gRPC/wallet/api"
	"gRPC/wallet/lib"
	"log"
	"net/http"
)

func main() {

	srv := api.PbWalletDaemon{}
	WalletServer := pb.NewDeviceServer(&srv, nil)
	wrapped := lib.WithReadHeader(WalletServer)

	err := http.ListenAndServe(":9000", wrapped)
	if err != nil {
		log.Fatal(err)
	}

}
