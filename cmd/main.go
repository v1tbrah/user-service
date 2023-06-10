package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/v1tbrah/user-service/config"
	"github.com/v1tbrah/user-service/internal/api"
	"github.com/v1tbrah/user-service/internal/storage"
)

func main() {
	newConfig := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	if err := newConfig.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	newStorage, err := storage.Init(newConfig.Storage)
	if err != nil {
		log.Fatal().Err(err).Interface("config", newConfig.Storage).Msg("storage.Init")
	} else {
		log.Info().Msg("storage initialized")
	}

	newAPI := api.New(newStorage)

	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errServingCh := make(chan error)
	go func() {
		errServing := newAPI.StartServing(context.Background(), newConfig.GRPCConfig, shutdownSig)
		errServingCh <- errServing
	}()

	select {
	case <-shutdownSig:
		close(shutdownSig)
	case errServing := <-errServingCh:
		if errServing != nil {
			log.Error().Err(errServing).Msg("newAPI.StartServing")
		}
	}

	ctxClose, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = newAPI.GracefulStop(ctxClose); err != nil {
		log.Error().Err(err).Msg("gRPC server graceful stop")
		if err == context.DeadlineExceeded {
			return
		}
	} else {
		log.Info().Msg("gRPC server gracefully stopped")
	}

	if err = newStorage.Close(ctxClose); err != nil {
		log.Error().Err(err).Msg("storage close")
		if err == context.DeadlineExceeded {
			return
		}
	} else {
		log.Info().Msg("storage closed")
	}
}
