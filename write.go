package slogwrite

import (
	"context"
	"io"

	"github.com/sagikazarmark/slog-shim"
)

type Writer struct {
	*slog.Logger
	ctx context.Context
	lvl slog.Level
}

func NewWriter(ctx context.Context, log *slog.Logger, lvl slog.Level) io.Writer {
	return &Writer{
		Logger: log,
		ctx:    ctx,
		lvl:    lvl,
	}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Logger.Log(w.ctx, w.lvl, string(p))
	return len(p), nil
}
