package grpc

import (
	"worker/internal/config"

	log "github.com/sirupsen/logrus"
	grpcMod "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// New GRPC
func NewGRPC() (*grpcMod.ClientConn, error) {
	log.Info("setting up the grpc connection")

	config := config.GetConfig()

	conn, err := grpcMod.Dial(config.GRPC_SERVER, grpcMod.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
