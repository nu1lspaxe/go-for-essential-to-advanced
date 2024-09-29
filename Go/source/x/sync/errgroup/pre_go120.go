//go:build !go1.20

package errgroup

import "context"

func withCancelCause(parent context.Context) (context.Context, func(error)) {
	ctx, cancel := context.withCancelCause(parent)
	return ctx, func(error) { cancel() }
}
