package api

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/v1tbrah/promcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *API) interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	timeStart := time.Now()

	resp, err = handler(ctx, req)

	duration := time.Since(timeStart)
	statusCode := status.Code(err)
	if err != nil && statusCode == codes.Internal {
		log.Error().Msg(info.FullMethod + " in " + duration.String())
		a.promCli.IncRequestResultCount(info.FullMethod, promcli.LabelError)
	} else {
		log.Info().Msg(info.FullMethod + " in " + duration.String())
	}

	a.promCli.IncRequestResultCount(info.FullMethod, promcli.LabelTotal)
	a.promCli.ObserveRequestDurationSeconds(info.FullMethod, int(statusCode), duration)

	return resp, err
}
