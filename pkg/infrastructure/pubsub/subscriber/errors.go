package subscriber

import "errors"

var (
	AlreadyStartedError = errors.New("subscriber already started")
	InvalidChannelError = errors.New("invalid channel")
)
