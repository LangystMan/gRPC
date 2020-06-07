package api

import (
	"context"
	pb "gRPC/assembly/device"
	"log"
)

type PbWalletDaemon struct{}

// curl POST http://localhost:9000/twirp/proto.Device/Registration -H "Content-Type: application/json" -d "{\"phone\":9992323,\"devicename\":\"Test\"}"
func (s *PbWalletDaemon) Registration(ctx context.Context, req *pb.RequestDeviceRegistration) (*pb.ReplyDeviceRegistration, error) {
	out := pb.ReplyDeviceRegistration{}

	sign := ctx.Value("X-Signature").(string)
	log.Printf("Телефон: %d, девайс: %s", req.Phone, req.Devicename)
	log.Printf("В запрос прилетела X-Signature: %s", sign)

	out.Errno = 0
	out.Uuid = "999988887777666655554444333322221111"
	out.Pubkey = "ABCDEFGHJKLMNOP"

	return &out, nil
}
