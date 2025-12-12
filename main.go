// go build yaml2go.go
// usage: ./yaml2go <values.yaml>
package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	tab = "  "
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <values.yaml>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var values interface{}
	if err := yaml.Unmarshal(data, &values); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	output := "values := " + formatValue(values, 0)
	fmt.Println(output)
}

func formatValue(v interface{}, indent int) string {
	switch val := v.(type) {
	case map[string]interface{}:
		return formatMap(val, indent)
	case []interface{}:
		return formatSlice(val, indent)
	case string:
		return fmt.Sprintf("%q", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case int, int64, float64:
		return fmt.Sprintf("%v", val)
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%v", val)
	}
}

func formatMap(m map[string]interface{}, indent int) string {
	if len(m) == 0 {
		return "map[string]interface{}{}"
	}

	var sb strings.Builder
	sb.WriteString("map[string]interface{}{\n")

	for k, v := range m {
		sb.WriteString(strings.Repeat(tab, indent+1))
		sb.WriteString(fmt.Sprintf("%q: ", k))
		sb.WriteString(formatValue(v, indent+1))
		sb.WriteString(",\n")
	}

	sb.WriteString(strings.Repeat(tab, indent))
	sb.WriteString("}")

	return sb.String()
}

func formatSlice(s []interface{}, indent int) string {
	if len(s) == 0 {
		return "[]interface{}{}"
	}

	var sb strings.Builder
	sb.WriteString("[]interface{}{\n")

	for _, v := range s {
		sb.WriteString(strings.Repeat(tab, indent+1))
		sb.WriteString(formatValue(v, indent+1))
		sb.WriteString(",\n")
	}

	sb.WriteString(strings.Repeat(tab, indent))
	sb.WriteString("}")

	return sb.String()
}
