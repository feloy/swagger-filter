package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-openapi/spec"
)

func outputSpec(spec *spec.Swagger, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	b, err := spec.MarshalJSON()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "%s", b)
	if err != nil {
		return err
	}
	w.Flush()
	return nil
}
