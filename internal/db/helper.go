package db

import (
	"context"

	"github.com/jinzhu/copier"
)

// ExecWithMapping maps fields between two structs and executes a DB query function.
// It returns any errors that occur during mapping or query execution.
func ExecWithMapping[S any, T any](
	source S,
	target *T,
	queryFunc func(T) error,
) error {
	// Step 1: Map fields from source to target
	if err := copier.Copy(target, source); err != nil {
		return err
	}

	// Step 2: Execute the DB query
	if err := queryFunc(*target); err != nil {
		return err
	}

	return nil
}

func WithContext[T any](
  ctx context.Context, 
  queryFunc func(context.Context, T) error) func(T) error {
	return func(params T) error {
		return queryFunc(ctx, params)
	}
}
