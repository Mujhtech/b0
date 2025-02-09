package handlers

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/mujhtech/b0/database/store"
	"github.com/mujhtech/b0/internal/pkg/encrypt"
)

func HandleStoreSync(aesCfb encrypt.Encrypt, store *store.Store) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {

		// appId, err := aesCfb.Decrypt(string(t.Payload()))

		// if err != nil {
		// 	return err
		// }

		// // _, err = store.AppRepo.FindAppByID(ctx, appId)

		// if err != nil {
		// 	return err
		// }

		return nil
	}
}
