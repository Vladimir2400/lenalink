package gars

import (
	"net/url"
	"strconv"
	"strings"
)

type queryParams struct {
	selectFields []string
	expand       []string
	filter       string
	orderBy      string
	top          *int
	skip         *int
	count        *bool
	format       string
}

type Option interface {
	apply(*queryParams)
}

type optionFunc func(*queryParams)

func (f optionFunc) apply(params *queryParams) { f(params) }

func defaultQueryParams() *queryParams {
	return &queryParams{format: "json"}
}

func (p *queryParams) populateQuery(values url.Values) {
	if len(p.selectFields) > 0 {
		values.Set("$select", strings.Join(p.selectFields, ","))
	}
	if len(p.expand) > 0 {
		values.Set("$expand", strings.Join(p.expand, ","))
	}
	if p.filter != "" {
		values.Set("$filter", p.filter)
	}
	if p.orderBy != "" {
		values.Set("$orderby", p.orderBy)
	}
	if p.top != nil {
		values.Set("$top", strconv.Itoa(*p.top))
	}
	if p.skip != nil {
		values.Set("$skip", strconv.Itoa(*p.skip))
	}
	if p.count != nullBool {
		values.Set("$count", strconv.FormatBool(*p.count))
	}
	if p.format != "" {
		values.Set("$format", p.format)
	}
}

var nullBool *bool

// WithSelect limits response fields to provided list.
func WithSelect(fields ...string) Option {
	return optionFunc(func(params *queryParams) {
		params.selectFields = append([]string{}, fields...)
	})
}

// WithExpand expands referenced entities.
func WithExpand(entities ...string) Option {
	return optionFunc(func(params *queryParams) {
		params.expand = append([]string{}, entities...)
	})
}

// WithFilter applies custom OData filter expression.
func WithFilter(filter string) Option {
	return optionFunc(func(params *queryParams) {
		params.filter = filter
	})
}

// WithOrderBy adds order by expression.
func WithOrderBy(order string) Option {
	return optionFunc(func(params *queryParams) {
		params.orderBy = order
	})
}

// WithTop limits number of records.
func WithTop(top int) Option {
	return optionFunc(func(params *queryParams) {
		if top >= 0 {
			value := top
			params.top = &value
		}
	})
}

// WithSkip defines number of records to skip.
func WithSkip(skip int) Option {
	return optionFunc(func(params *queryParams) {
		if skip >= 0 {
			value := skip
			params.skip = &value
		}
	})
}

// WithCount toggles total count calculation.
func WithCount(enabled bool) Option {
	return optionFunc(func(params *queryParams) {
		params.count = &enabled
	})
}

// WithFormat overrides $format parameter.
func WithFormat(format string) Option {
	return optionFunc(func(params *queryParams) {
		params.format = format
	})
}

// WithPagination sets $skip and $top using page index (1-based) and page size.
func WithPagination(page, pageSize int) Option {
	return optionFunc(func(params *queryParams) {
		if page < 1 || pageSize <= 0 {
			return
		}
		skip := (page - 1) * pageSize
		params.skip = &skip
		top := pageSize
		params.top = &top
	})
}
