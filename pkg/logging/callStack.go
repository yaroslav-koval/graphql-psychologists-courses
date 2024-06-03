package logging

import "github.com/rs/zerolog"

type StackTrace []string

func (st StackTrace) MarshalZerologArray(a *zerolog.Array) {
	for _, s := range st {
		a.Str(s)
	}
}
