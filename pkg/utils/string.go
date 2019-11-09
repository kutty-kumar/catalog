package utils

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleEmptyRequest(requestText *string, context string) error {
	if requestText != nil && *requestText != "" {
		return nil
	}
	return status.Error(codes.Internal, fmt.Sprintf("%v cannot be nil/empty", context))
}