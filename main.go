package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bg-vc/vc-gin-api/dao"
	"github.com/bg-vc/vc-gin-api/handler"
	"github.com/bg-vc/vc-gin-api/pkg"
	"github.com/bg-vc/vc-gin-api/router"
	"github.com/bg-vc/vc-gin-api/router/middleware"
	"github.com/bg-vc/vc-gin-api/service"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	if err := run(logger); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	cfg, err := pkg.LoadConfig()
	if err != nil {
		return err
	}

	db, err := pkg.OpenDB(cfg.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	redisClient, err := pkg.OpenRedis(ctx, cfg.Redis)
	if err != nil {
		return err
	}
	if redisClient != nil {
		defer redisClient.Close()
		logger.Info("optional Redis example is enabled", "address", cfg.Redis.Address)
	}

	tokens := pkg.NewTokenManager(cfg.JWTSecret, cfg.JWTExpiry)
	accountService := service.NewAccountService(dao.NewAccountDAO(db), tokens)
	userService := service.NewUserService(dao.NewUserDAO(db))

	gin.SetMode(cfg.RunMode)
	engine := gin.New()
	router.Load(
		engine,
		handler.NewAccountHandler(accountService),
		handler.NewUserHandler(userService),
		middleware.Auth(tokens),
		middleware.Logging(logger),
	)

	server := &http.Server{
		Addr:              cfg.Address,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("vc-gin-api is ready", "address", cfg.Address)
		serverErr <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return server.Shutdown(shutdownCtx)
	case err := <-serverErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
