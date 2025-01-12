package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Parallel()

	errBase := "parse \"%s\": invalid URI for request"

	tcs := map[string]struct {
		host string
		err  string
	}{
		"correct host": {
			host: "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json",
		},
		"host as IP address": {
			host: "http://127.0.0.1",
		},
		"host missing the scheme": {
			host: "api.nbp.pl",
			err:  errBase,
		},
		"empty host": {
			host: "",
			err:  errBase,
		},
		"string host": {
			host: "string",
			err:  errBase,
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var h host

			err := h.Set(tc.host)
			if tc.err != "" {
				assert.Errorf(t, err, tc.err, tc.host)
				assert.Empty(t, h.host)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.host, h.String())
		})
	}
}
