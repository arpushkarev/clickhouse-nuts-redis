package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/arpushkarev/clickhouse-nuts-redis/internal/app/api/item_v1"
	itemRepository "github.com/arpushkarev/clickhouse-nuts-redis/internal/repository/item"
	"github.com/arpushkarev/clickhouse-nuts-redis/internal/service/item"
	desc "github.com/arpushkarev/clickhouse-nuts-redis/pkg/item_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/stdlib" //just for initialization the driver
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcPort = "50051"
	httpPort = "8090"
)

const (
	host       = "localhost"
	port       = "54322"
	dbUser     = "item-service-user"
	dbPassword = "item-service-password"
	dbName     = "item-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := startGRPC()
		if err != nil {
			log.Fatalf("GRPCserver error: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()

		err := startHTTP()
		if err != nil {
			log.Fatalf("HTTPserver error: %s", err.Error())
		}
	}()

	wg.Wait()
}

func startGRPC() error {
	list, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, grpcPort))
	if err != nil {
		return err
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	itemRepository := itemRepository.NewRepository(db)
	itemService := item.NewService(itemRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterItemV1Server(s, item_v1.NewImplementation(itemService))

	fmt.Println("GRPC server is running on port:", grpcPort)

	if err = s.Serve(list); err != nil {
		return err
	}

	return nil
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // nolint: staticcheck

	err := desc.RegisterItemV1HandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", host, grpcPort), opts)
	if err != nil {
		return err
	}

	fmt.Println("HTTP server is running on port:", httpPort)

	return http.ListenAndServe(fmt.Sprintf("%s:%s", host, httpPort), mux)
}
