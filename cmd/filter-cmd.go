package cmd

import (
	"fmt"

	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
)

func filterCmd(cmd *cobra.Command, args []string) error {
	var filters filters

	err := addEndpoints(cmd, &filters)
	if err != nil {
		return err
	}

	err = addPrefixEndpoints(cmd, &filters)
	if err != nil {
		return err
	}

	err = addRegexpEndpoints(cmd, &filters)
	if err != nil {
		return err
	}

	if len(args) != 2 {
		return fmt.Errorf("Number of args should be 2 but is %d", len(args))
	}

	specfile := args[0]
	outfile := args[1]

	doc, err := load(specfile)
	if err != nil {
		return err
	}

	refs := filterSpec(doc.Spec(), filters)

	filterDefinitions(doc.Spec(), refs)

	outputSpec(doc.Spec(), outfile)
	return err
}

func addEndpoints(cmd *cobra.Command, filters *filters) error {
	endpoints, err := cmd.Flags().GetStringSlice("endpoint")
	if err != nil {
		return err
	}

	for _, endpoint := range endpoints {
		filters.add(StringFilter{endpoint})
	}
	return nil
}

func addPrefixEndpoints(cmd *cobra.Command, filters *filters) error {
	endpoints, err := cmd.Flags().GetStringSlice("endpoint-prefix")
	if err != nil {
		return err
	}

	for _, endpoint := range endpoints {
		filters.add(PrefixFilter{endpoint})
	}
	return nil
}

func addRegexpEndpoints(cmd *cobra.Command, filters *filters) error {
	endpoints, err := cmd.Flags().GetStringSlice("endpoint-regexp")
	if err != nil {
		return err
	}

	for _, endpoint := range endpoints {
		f, err := NewRegexpFilter(endpoint)
		if err != nil {
			return err
		}
		filters.add(f)
	}
	return nil
}

func filterSpec(spec *spec.Swagger, filters filters) refs {
	refs := newRefs()
	for k := range spec.Paths.Paths {
		if !filters.pathMatches(k) {
			delete(spec.Paths.Paths, k)
			continue
		}
		findRefs(spec, "GET", spec.Paths.Paths[k].PathItemProps.Get, refs)
		findRefs(spec, "PUT", spec.Paths.Paths[k].PathItemProps.Put, refs)
		findRefs(spec, "POST", spec.Paths.Paths[k].PathItemProps.Post, refs)
		findRefs(spec, "DELETE", spec.Paths.Paths[k].PathItemProps.Delete, refs)
		findRefs(spec, "OPTIONS", spec.Paths.Paths[k].PathItemProps.Options, refs)
		findRefs(spec, "HEAD", spec.Paths.Paths[k].PathItemProps.Head, refs)
		findRefs(spec, "PATCH", spec.Paths.Paths[k].PathItemProps.Patch, refs)
	}
	return refs
}

func filterDefinitions(spec *spec.Swagger, refs refs) {
	for k := range spec.Definitions {
		if _, found := refs[k]; !found {
			delete(spec.Definitions, k)
			continue
		}
	}
}
