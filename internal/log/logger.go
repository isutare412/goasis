package log

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	slogmulti "github.com/samber/slog-multi"
)

func init() {
	handler := slogmulti.
		Pipe(slogmulti.NewHandleInlineMiddleware(extractAttrsFromCtx)).
		Handler(tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339Nano,
			NoColor:    !isatty.IsTerminal(os.Stdout.Fd()),
		}))

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
