package logs

import (
	"github.com/rs/zerolog"
	"io"
	"os"
)

func New(w io.Writer) zerolog.Logger {
	return zerolog.New(w).
		Output(zerolog.ConsoleWriter{
			TimeFormat: "15:04:05",
			Out:        os.Stderr,
		}).With().Timestamp().Logger()
}
