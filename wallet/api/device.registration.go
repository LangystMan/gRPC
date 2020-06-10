package api

import (
	"context"
	pb "gRPC/assembly/device"
)

type ApiDaemon struct{}

// curl POST http://localhost:9000/twirp/proto.Device/Registration -H "Content-Type: application/json" -d "{\"phone\":9992323,\"devicename\":\"Test\"}"
func (s *ApiDaemon) Registration(ctx context.Context, req *pb.RequestDeviceRegistration) (*pb.ReplyDeviceRegistration, error) {
	out := pb.ReplyDeviceRegistration{}

	//TODO Доделать

	return &out, nil
}
