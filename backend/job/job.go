package job

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/config"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/agent"
	"github.com/mujhtech/b0/internal/pkg/container"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
	secretmanager "github.com/mujhtech/b0/internal/pkg/secret_manager"
	"github.com/mujhtech/b0/internal/pkg/sse"
	"github.com/mujhtech/b0/internal/redis"
	"github.com/mujhtech/b0/job/handlers"
	rdsv9 "github.com/redis/go-redis/v9"
)

type Job struct {
	Client    *Client
	Executor  *Executor
	Scheduler *Scheduler
	aesCfb    encrypt.Encrypt
}

func NewJob(cfg *config.Config, appCtx context.Context, redis *redis.Redis) (*Job, error) {

	var c asynq.RedisConnOpt
	var _ = redis.MakeRedisClient().(rdsv9.UniversalClient)
	c = redis

	aesCfb, err := encrypt.NewAesCfb(cfg.EncryptionKey)

	if err != nil {
		return nil, err
	}

	return &Job{
		aesCfb:    aesCfb,
		Client:    NewClient(c, aesCfb),
		Executor:  NewExecutor(cfg, appCtx, c),
		Scheduler: NewScheduler(cfg, c),
	}, nil
}

func (j *Job) RegisterAndStart(cfg *config.Config, store *store.Store, agent *agent.Agent, sse sse.Streamer, container *container.Container, secretManager secretmanager.SecretManager) error {
	j.Executor.RegisterJobHandler(JobNameWorkflowCreate, asynq.HandlerFunc(handlers.HandleCreateWorkflow(j.aesCfb, store, agent, sse)))
	j.Executor.RegisterJobHandler(JobNameWorkflowUpdate, asynq.HandlerFunc(handlers.HandleUpdateWorkflow(j.aesCfb, store, agent, sse)))
	j.Executor.RegisterJobHandler(JobNameWebhook, asynq.HandlerFunc(handlers.HandleWebhook(j.aesCfb, store)))
	j.Executor.RegisterJobHandler(JobNameProjectDeploy, asynq.HandlerFunc(handlers.HandleDeployProject(j.aesCfb, cfg, store, agent, sse, container, secretManager)))

	return j.Executor.Start()
}
