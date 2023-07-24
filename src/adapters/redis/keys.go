package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func EmailingMemberListOffsetKey(clientID int64, listID string) string {
	return fmt.Sprintf("%d:%s", clientID, listID)
}

func (a RedisAdapter) GetEmailingMemberListOffset(ctx context.Context, clientID int64, listID string) (int, error) {
	key := EmailingMemberListOffsetKey(clientID, listID)
	value, err := a.Get(ctx, key)
	if err != nil {
		return 0, errors.Wrap(err, "could not get emailing member list offset value from cache")
	}

	if value == "" {
		return 0, nil
	}

	offset, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("invalid value for key %s, should be an int and got %s", key, value))
	}

	return offset, nil
}
