package log

import (
	"context"
	"log/slog"

	"github.com/isutare412/goasis/internal/pkgctx"
)

func extractAttrsFromCtx(ctx context.Context, record slog.Record, next func(context.Context, slog.Record) error) error {
	if bag, ok := pkgctx.GetBag(ctx); ok {
		if bag.RequestID != "" {
			record.AddAttrs(slog.String("requestId", bag.RequestID))
		}
	}

	return next(ctx, record)
}
