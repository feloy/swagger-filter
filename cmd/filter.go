package cmd

import (
	"regexp"
	"strings"
)

type filter interface {
	Match(endpoint string) bool
}

type filters struct {
	items []filter
}

func (f *filters) add(nf filter) {
	f.items = append(f.items, nf)
}

func (o *filters) pathMatches(path string) bool {
	for _, f := range o.items {
		if f.Match(path) {
			return true
		}
	}
	return false
}

type StringFilter struct {
	value string
}

func (o StringFilter) Match(endpoint string) bool {
	return o.value == endpoint
}

type PrefixFilter struct {
	value string
}

func (o PrefixFilter) Match(endpoint string) bool {
	return strings.HasPrefix(endpoint, o.value)
}

type RegexpFilter struct {
	value *regexp.Regexp
}

func NewRegexpFilter(s string) (*RegexpFilter, error) {
	r, err := regexp.Compile(s)
	if err != nil {
		return nil, err
	}
	return &RegexpFilter{r}, nil
}

func (o RegexpFilter) Match(endpoint string) bool {
	return o.value.MatchString(endpoint)
}
