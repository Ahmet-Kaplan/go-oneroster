package handlers

import (
	"net/url"
)

func SortQuery(q url.Values) string {
	var s string

	if v, ok := q["sort"]; ok {
		s = v[0]
	} else {
		s = "sourcedId"
	}

	return s
}

func LimitQuery(q url.Values) string {
	var s string

	if v, ok := q["limit"]; ok {
		s = v[0]
	} else {
		s = "100"
	}

	return s
}

type urlParams struct {
	Sort   string
	Limit  string
	Offset string
}

// Parses the url query values for sort, limit and offset settings otherwise returns defaults
func Query(q url.Values) map[string]string {
	d := map[string]string{
		"sort":   "sourcedId",
		"limit":  "100",
		"offset": "0",
	}

	for k, _ := range d {
		if v, ok := q[k]; ok {
			d[k] = v[0]
		}
	}

	return d
}
