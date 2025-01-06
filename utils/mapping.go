package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TimestampToTimePtr(timestamp *timestamppb.Timestamp) *time.Time {
	if timestamp == nil {
		return nil
	}

	t := timestamp.AsTime()
	return &t
}

func DeletedAtNil(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
