package xmlstream

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type TokenHandler func(stack []string, token xml.Token) (xml.Token, error)

func Parse(r io.Reader, w io.Writer, handlers ...TokenHandler) error {

	stack := []string{}

	decoder := xml.NewDecoder(r)
	for {
		tkn, err := decoder.RawToken()
		if err != nil {
			if err == io.EOF {
				slog.Debug("end of stream")
				break
			} else {
				slog.Error("error reading token", "error", err)
				return err
			}
		}

		for _, handler := range handlers {
			tkn, err = handler(stack, tkn)
			if err != nil {
				slog.Error("error in call to handler on token", "token", tkn, "error", err)
				return err
			}
		}

		slog.Debug("after handling token", "type", fmt.Sprintf("%T", tkn))

		switch token := tkn.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local)
			if token.Name.Space == "" {
				fmt.Fprintf(w, "<%s", token.Name.Local)
			} else {
				fmt.Fprintf(w, "<%s:%s", token.Name.Space, token.Name.Local)
			}
			if len(token.Attr) > 0 {
				for _, attr := range token.Attr {
					if attr.Name.Space == "" {
						fmt.Fprintf(w, " %s=%q", attr.Name.Local, attr.Value)
					} else {
						fmt.Fprintf(w, " %s:%s=%q", attr.Name.Space, attr.Name.Local, attr.Value)
					}

				}
			}
			fmt.Fprintf(w, ">")
		case xml.EndElement:
			if len(stack) < 1 || stack[len(stack)-1] != token.Name.Local {
				slog.Error("unbalanced XML", "start element", stack[len(stack)-1], "end element", token.Name.Local)
				return errors.New("unbalanced XML")
			}
			stack = stack[:len(stack)-1]
			if token.Name.Space == "" {
				fmt.Fprintf(w, "</%s>", token.Name.Local)
			} else {
				fmt.Fprintf(w, "</%s:%s>", token.Name.Space, token.Name.Local)
			}
		case xml.CharData:
			fmt.Fprintf(w, "%s", string(token))
		case xml.Comment:
			fmt.Fprintf(w, "<!--%s-->", string(token))
		case xml.ProcInst:
			fmt.Fprintf(w, "<?%s %s?>", token.Target, string(token.Inst))
		case xml.Directive:
			fmt.Fprintf(w, "<!%s>", string(token))
		default:
			slog.Error("unknown token type", "type", fmt.Sprintf("%T", token))
			return err
		}
	}
	return nil
}

func ParseString(text string, w io.Writer, handlers ...TokenHandler) error {
	return Parse(strings.NewReader(text), w, handlers...)
}

func ParseFile(path string, w io.Writer, handlers ...TokenHandler) error {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("error opening file", "path", path, "error", err)
		return err
	}
	defer file.Close()
	return Parse(file, w, handlers...)
}

func MustYAML(v any) string {
	s, _ := yaml.Marshal(v)
	return string(s)
}
