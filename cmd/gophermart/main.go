package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	core_accrual "github.com/Kosvu/gophermart/internal/core/accrualclient"
	"github.com/Kosvu/gophermart/internal/core/config"
	"github.com/Kosvu/gophermart/internal/core/migrations"
	core_repository "github.com/Kosvu/gophermart/internal/core/repository"
	accrual_repository "github.com/Kosvu/gophermart/internal/features/accrual/repository"
	accrual_worker "github.com/Kosvu/gophermart/internal/features/accrual/worker"
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
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

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

	repositoryAccrual := accrual_repository.NewAccrualRepository(pool)
	clientAccrual := core_accrual.NewClient(cfg.AccrualAddress)
	workerAccrual := accrual_worker.NewWorker(clientAccrual, repositoryAccrual, accrual_worker.DefaultInterval)

	go workerAccrual.Run(ctx)

	r := chi.NewRouter()

	r.Post("/api/user/register", transportHTTPAuth.Register)
	r.Post("/api/user/login", transportHTTPAuth.Login)

	r.Group(func(r chi.Router) {
		r.Use(auth_middleware.Auth(token))
		r.Post("/api/user/orders", transportOrders.Load)
		r.Get("/api/user/orders", transportOrders.GetOrders)
	})

	log.Printf("server started on %s", cfg.Addr)
	srv := http.Server{Addr: cfg.Addr, Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	pool.Close()
}
