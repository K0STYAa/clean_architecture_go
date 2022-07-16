package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"forgoproject/pkg/store/postgres"
	"forgoproject/pkg/tracing"
	"forgoproject/pkg/type/context"
	log "forgoproject/pkg/type/logger"
	deliveryGrpc "forgoproject/services/contact/internal/delivery/grpc"
	deliveryHttp "forgoproject/services/contact/internal/delivery/http"
	repositoryStorage "forgoproject/services/contact/internal/repository/storage/postgres"
	useCaseContact "forgoproject/services/contact/internal/useCase/contact"
	useCaseGroup "forgoproject/services/contact/internal/useCase/group"
)

func init() {
	viper.Set("SERVICE_NAME", "contactService")
}

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	if err != nil {
		panic(err)
	}
	var (
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
