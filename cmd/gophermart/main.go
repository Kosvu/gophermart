package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kosvu/gophermart/internal/core/config"
	"github.com/Kosvu/gophermart/internal/core/migrations"
	core_repository "github.com/Kosvu/gophermart/internal/core/repository"
	auth_middleware "github.com/Kosvu/gophermart/internal/features/auth/middleware_auth"
	auth_repository "github.com/Kosvu/gophermart/internal/features/auth/repository"
	auth_service "github.com/Kosvu/gophermart/internal/features/auth/service"
	"github.com/Kosvu/gophermart/internal/features/auth/token"
	auth_transport_http "github.com/Kosvu/gophermart/internal/features/auth/transport/http"
	orders_repository "github.com/Kosvu/gophermart/internal/features/orders/repository"
	orders_service "github.com/Kosvu/gophermart/internal/features/orders/service"
	orders_transport "github.com/Kosvu/gophermart/internal/features/orders/transport"
	"github.com/go-chi/chi"
)

func main() {
	cfg, err := config.NewConfig()
	ctx := context.Background()

	if err != nil {
		panic(err)
	}

	if err := migrations.StartMigration(cfg.DatabaseURI); err != nil {
		panic(err)
	}

	pool, err := core_repository.NewPool(ctx, cfg.DatabaseURI)

	if err != nil {
		log.Fatal(err)
	}

	repositoryAuth := auth_repository.NewAuth(pool)
	token := token.NewToken(cfg.Secret, token.DefaultTTL)
	serviceAuth := auth_service.NewAuthService(repositoryAuth, *token)
	transportHTTPAuth := auth_transport_http.NewAuthHTTPHandler(serviceAuth)

	repositoryOrders := orders_repository.NewOrdersRepository(pool)
	serviceOrders := orders_service.NewOrdersService(repositoryOrders)
	transportOrders := orders_transport.NewOrdersHTTPHandler(serviceOrders)

	r := chi.NewRouter()

	r.Post("/api/user/register", transportHTTPAuth.Register)
	r.Post("/api/user/login", transportHTTPAuth.Login)

	r.Group(func(r chi.Router) {
		r.Use(auth_middleware.Auth(token))
		r.Post("/api/user/orders", transportOrders.Load)
		r.Get("/api/user/orders", transportOrders.GetOrders)
	})

	log.Printf("server started on %s", cfg.Addr)
	go func() {
		if err := http.ListenAndServe(cfg.Addr, r); err != nil {
			panic(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	pool.Close()

}
