package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/mujhtech/b0/database"
	"github.com/rs/zerolog/log"
)

var (
	Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// errors
	ErrNotFound = errors.New("not found")
)

const (
	excludeDeleted = "deleted_at IS NULL"
)

type Store struct {
	UserRepo          UserRepository
	TokenRepo         TokenRepository
	ProjectRepo       ProjectRepository
	EndpointRepo      EndpointRepository
	AIUsageRepo       AIUsageRepository
	ProjectLogRepo    ProjectLogRepository
	AITokenCreditRepo AITokenCreditRepository
}

func NewStore(db *database.Database) *Store {
	return &Store{
		UserRepo:          NewUserRepository(db),
		TokenRepo:         NewTokenRepository(db),
		ProjectRepo:       NewProjectRepository(db),
		EndpointRepo:      NewEndpointRepository(db),
		AIUsageRepo:       NewAIUsageRepository(db),
		ProjectLogRepo:    NewProjectLogRepository(db),
		AITokenCreditRepo: NewAITokenCreditRepository(db),
	}
}

func ProcessSQLErrorfWithCtx(ctx context.Context, query string, err error, format string, args ...interface{}) error {
	// create fallback error returned if we can't map it
	fallbackErr := fmt.Errorf(format, args...)

	// always log internal error together with message.
	log.Info().Msgf("Query: %v", query)
	log.Error().Err(err).Msgf("%v: [SQL] %v", fallbackErr, err)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return fallbackErr
	}
}

// func toDatabaseValue(value interface{}) interface{} {
// 	if value == nil {
// 		return nil
// 	}

// 	switch v := value.(type) {
// 	case string:
// 		return v
// 	case int:
// 		return v
// 	case int64:
// 		return v
// 	case bool:
// 		return v
// 	default:
// 		return v
// 	}
// }
