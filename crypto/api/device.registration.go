package api

import (
	"context"
	pb "gRPC/assembly/api"
	cfg "gRPC/lib/config"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = cfg.LoadLogger()
}

type ApiDaemon struct{}

// curl POST http://localhost:9000/twirp/Api/Registration -H "Content-Type: application/json" -H "Content-Language: ru" -d "{\"phone\":9992323,\"devicename\":\"Test\"}"
func (s *ApiDaemon) Registration(ctx context.Context, req *pb.RequestDeviceRegistration) (*pb.ReplyDeviceRegistration, error) {
	out := pb.ReplyDeviceRegistration{}

	log.Info().Msg("Take registration request message")

	out.Errno = 200
	out.Error = "SUCCESS"

	return &out, nil
}
