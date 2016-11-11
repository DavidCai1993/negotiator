package negotiator

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func parseAccept(header http.Header) (specs specs) {
	headerVal := strings.ToLower(strings.Replace(header.Get(HeaderAccept),
		" ", "", -1))

	if headerVal == "" {
		specs = []spec{spec{val: "*/*", q: defaultQ}}
		return
	}

	accpets := strings.Split(headerVal, ",")

	for _, accept := range accpets {
		pair := strings.Split(strings.TrimSpace(accept), ";")

		if len(pair) < 1 || len(pair) > 2 || strings.Index(pair[0], "/") == -1 {
			continue
		}

		spec := spec{val: pair[0], q: defaultQ}

		if len(pair) == 2 && strings.HasPrefix(pair[1], "q=") {
			var i int

			if strings.HasPrefix(pair[1], "q=") {
				i = 2
			} else if strings.HasPrefix(pair[1], "level=") {
				i = 6
			} else {
				continue
			}

			if q, err := strconv.ParseFloat(pair[1][i:], 64); err == nil && q != 0.0 {
				if q > defaultQ {
					q = defaultQ
				}

				spec.q = q
			} else {
				continue
			}
		}

		specs = append(specs, spec)
	}

	sort.Sort(specs)

	return
}
