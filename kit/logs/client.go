package logs

import (
	"github.com/rs/zerolog"
	"io"
)

func New(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().Timestamp().Logger()
}
