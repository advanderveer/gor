// Package main provides the `gor` compiler executable.
package main

import (
	"log/slog"
	"os"
)

// compile the compiler.
func compile(*slog.Logger) error {
	return nil
}

// compiler entrypoint.
func main() {
	logh := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logs := slog.New(logh)

	if err := compile(logs); err != nil {
		logs.Error("failed to compile", slog.String("err", err.Error()))
	}
}
