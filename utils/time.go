package utils

import (
	"corekit/timeutil"
	"time"
)

const ISO8601Milli = timeutil.ISO8601Milli

// Deprecated: use corekit/timeutil.FormatUTCTimestampMilli.
// FormatUTCTimestampMilli formats a time in UTC using millisecond precision.
func FormatUTCTimestampMilli(t time.Time) string {
	return timeutil.FormatUTCTimestampMilli(t)
}

// Deprecated: use corekit/timeutil.NowUTCTimestampMilli.
// GetCurrentUTCTimestampMilli returns the current UTC timestamp with millisecond precision.
func GetCurrentUTCTimestampMilli() string {
	return timeutil.NowUTCTimestampMilli()
}
