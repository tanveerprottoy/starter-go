package grpcpkg

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RespondError(
	c codes.Code,
	msg string,
) error {
	return status.Error(
		c,
		msg,
	)
}
