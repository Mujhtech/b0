package server

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/mujhtech/b0/api"
	"github.com/mujhtech/b0/cache"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/http"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/telemetry"
	"github.com/mujhtech/b0/internal/redis"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func RegisterServerCommand() *cobra.Command {

	var (
		configFile string
		logLevel   string
	)

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start b0 server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

			err := startServer(configFile, logLevel)

			if err != nil {
				log.Err(err).Msg("failed to start server")
			}

		},
	}

	cmd.Flags().StringVar(&configFile, "config", config.DefaultConfigFilePath, "configuration file")
	cmd.Flags().StringVar(&logLevel, "log-level", "info", "log level")

	return cmd

}

func startServer(configFile string, logLevel string) error {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano

	// attach logger to context
	logger := log.Logger.With().Logger()
	ctx = logger.WithContext(ctx)

	_ = godotenv.Load(configFile)

	cfg, err := config.LoadConfig()

	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	db, err := database.Connect(ctx, cfg)

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	defer db.Close()

	redis, err := redis.NewRedis(cfg)

	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	cache, err := cache.NewCache(cfg, redis)

	if err != nil {
		return fmt.Errorf("failed to create cache: %w", err)
	}

	store := store.NewStore(db)

	agent := agent.New(cfg)

	app, err := api.New(
		cfg,
		ctx,
		store,
		cache,
		agent,
	)

	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}

	telemetry, err := telemetry.New(cfg.Telemetry, "b0-server")

	if err != nil {
		return telemetry.Shutdown(ctx)
	}

	server := http.NewServer(cfg, app.BuildRouter())

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		//return job.RegisterAndStart(store, s3)
		return nil
	})

	gHTTP, shutdownHTTP := server.ListenAndServe()
	g.Go(gHTTP.Wait)

	logger.Info().Msgf("server started on port %d", cfg.Server.Port)

	<-gCtx.Done()

	stop()
	logger.Info().Msg("shutting down gracefully (press Ctrl+C again to force)")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	if shutdownErr := shutdownHTTP(shutdownCtx); shutdownErr != nil {
		return fmt.Errorf("failed to shutdown server gracefully: %w", shutdownErr)
	}

	if err = telemetry.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("failed to shutdown telemetry: %w", err)
	}

	//job.Executor.Stop()

	logger.Info().Msg("waiting for all goroutines to finish")
	err = g.Wait()

	return err

}
