package gars

import (
	"net/url"
	"testing"
)

func TestQueryOptions(t *testing.T) {
	params := defaultQueryParams()
	opts := []Option{
		WithSelect("Field1", "Field2"),
		WithExpand("Related"),
		WithFilter("Field1 eq 'value'"),
		WithOrderBy("Field2 desc"),
		WithTop(10),
		WithSkip(20),
		WithCount(true),
	}

	for _, opt := range opts {
		opt.apply(params)
	}

	values := url.Values{}
	params.populateQuery(values)

	cases := map[string]string{
		"$select":  "Field1,Field2",
		"$expand":  "Related",
		"$filter":  "Field1 eq 'value'",
		"$orderby": "Field2 desc",
		"$top":     "10",
		"$skip":    "20",
		"$count":   "true",
		"$format":  "json",
	}

	for key, expected := range cases {
		if val := values.Get(key); val != expected {
			t.Fatalf("expected %s=%s got %s", key, expected, val)
		}
	}
}

func TestPagination(t *testing.T) {
	params := defaultQueryParams()
	WithPagination(3, 50).apply(params)

	if params.skip == nil || *params.skip != 100 {
		t.Fatalf("expected skip=100 got %v", params.skip)
	}
	if params.top == nil || *params.top != 50 {
		t.Fatalf("expected top=50 got %v", params.top)
	}
}
