package contextpkg

import "context"

func AddValue(ctx context.Context, key, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func ReadValue(ctx context.Context, key any) any {
	return ctx.Value(key)
}
