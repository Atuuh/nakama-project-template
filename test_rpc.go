package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func TestRpc(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("Payload: %s", payload)

	// "payload" is bytes sent by the client we'll JSON decode it.
	var value interface{}
	if err := json.Unmarshal([]byte(payload), &value); err != nil {
		return "", runtime.NewError("unable to unmarshal payload", 13)
	}

	response, err := json.Marshal(value)
	if err != nil {
		return "", runtime.NewError("unable to marshal payload", 13)
	}

	return string(response), nil
}
