package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"

	// Embed is used to import the template files
	_ "embed"

	"github.com/Masterminds/sprig/v3"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	_ "github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/stringutils"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/tools/generate-helpers/pg-table-bindings/walker"
	"golang.org/x/tools/imports"
)

//go:embed store.go.tpl
var storeFile string

var storeTemplate = template.Must(template.New("gen").Funcs(funcMap).Funcs(sprig.TxtFuncMap()).Parse(autogenerated + storeFile))

type properties struct {
	Type           string
	Table          string
	RegisteredType string
}

func main() {
	c := &cobra.Command{
		Use: "generate store implementations",
	}

	var props properties
	c.Flags().StringVar(&props.Type, "type", "", "the (Go) name of the object")
	utils.Must(c.MarkFlagRequired("type"))

	c.Flags().StringVar(&props.RegisteredType, "registered-type", "", "the type this is registered in proto as storage.X")

	c.Flags().StringVar(&props.Table, "table", "", "the logical table of the objects")
	utils.Must(c.MarkFlagRequired("table"))

	c.RunE = func(*cobra.Command, []string) error {
		typ := stringutils.OrDefault(props.RegisteredType, props.Type)
		fmt.Println("Generating for", typ)
		mt := proto.MessageType(typ)
		if mt == nil {
			log.Fatalf("could not find message for type: %s", typ)
		}

		schema := walker.Walk(mt, props.Table)
		templateMap := map[string]interface{}{
			"Type":        props.Type,
			"TrimmedType": stringutils.GetAfter(props.Type, "."),
			"Table":       props.Table,
			"Schema":      schema,
		}

		buf := bytes.NewBuffer(nil)
		if err := storeTemplate.Execute(buf, templateMap); err != nil {
			return err
		}
		file := buf.Bytes()

		formatted, err := imports.Process("store.go", file, nil)
		if err != nil {
			return err
		}
		if err := os.WriteFile("store.go", formatted, 0644); err != nil {
			return err
		}

		return nil
	}
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}