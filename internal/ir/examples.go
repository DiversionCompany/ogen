package ir

import (
	"encoding/json"
	"sort"
	"strings"
)

func (t *Type) Examples() (r []string) {
	if t.Schema == nil {
		return nil
	}

	dedup := make(map[string]struct{}, len(t.Schema.Examples))
	for _, example := range t.Schema.Examples {

		if !json.Valid(example) {
			continue
		}
		k := string(example)
		if _, ok := dedup[k]; ok {
			continue
		}
		dedup[k] = struct{}{}
		// Rewrite CRLF to LF.
		r = append(r, strings.ReplaceAll(k, "\r\n", "\n"))
	}
	sort.Strings(r)
	return r
}
