package api

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func interceptorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	timeStart := time.Now()
	defer func() {
		duration := time.Since(timeStart)
		if err != nil && status.Code(err) == codes.Internal {
			log.Error().Msg(info.FullMethod + " in " + duration.String())
			return
		}
		log.Info().Msg(info.FullMethod + " in " + duration.String())
	}()
	return handler(ctx, req)
}
