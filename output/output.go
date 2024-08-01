package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/dihedron/migrate/templating"
	"gopkg.in/yaml.v3"
)

// ToJSON formats the given value to JSON.
func ToJSON(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// ToPrettyJSON formats the given value to pretty-printed JSON.
func ToPrettyJSON(value any) (string, error) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), err
}

// ToJSON formats the given value to YAML.
func ToYAML(value any) (string, error) {
	data, err := yaml.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), err
}

// ToText returns the input value converted to a string by
// applying the given template.
func ToText(value any, templ string) (string, error) {

	// populate the functions map
	functions := template.FuncMap{}
	for k, v := range templating.FuncMap() {
		functions[k] = v
	}
	for k, v := range sprig.FuncMap() {
		functions[k] = v
	}

	// parse the template
	t, err := template.New("main").Funcs(functions).Parse(templ)
	if err != nil {
		slog.Error("error parsing template", "template", templ, "error", err)
		return "", fmt.Errorf("error parsing template %s: %w", templ, err)
	}

	// execute the template
	buffer := bytes.Buffer{}
	if err := t.ExecuteTemplate(&buffer, "main", value); err != nil {
		slog.Error("error applying data to template", "error", err)
		return "", fmt.Errorf("error applying data to template: %w", err)
	}
	return buffer.String(), nil
}
