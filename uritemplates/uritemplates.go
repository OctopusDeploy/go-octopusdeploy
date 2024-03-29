//lint:file-ignore ST1005,ST1006,S1007 3rd-party source file

// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uritemplates is a level 4 implementation of RFC 6570 (URI
// Template, http://tools.ietf.org/html/rfc6570).
//
// To use uritemplates, parse a template string and expand it with a value
// map:
//
//	template, _ := uritemplates.Parse("https://api.github.com/repos{/user,repo}")
//	values := make(map[string]interface{})
//	values["user"] = "jtacoma"
//	values["repo"] = "uritemplates"
//	expanded, _ := template.Expand(values)
//	fmt.Printf(expanded)
//
package uritemplates

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	unreserved = regexp.MustCompile("[^A-Za-z0-9\\-._~]")
	reserved   = regexp.MustCompile("[^A-Za-z0-9\\-._~:/?#[\\]@!$&'()*+,;=]")
	validname  = regexp.MustCompile("^([A-Za-z0-9_\\.]|%[0-9A-Fa-f][0-9A-Fa-f])+$")
	hex        = []byte("0123456789ABCDEF")
)

func pctEncode(src []byte) []byte {
	dst := make([]byte, len(src)*3)
	for i, b := range src {
		buf := dst[i*3 : i*3+3]
		buf[0] = 0x25
		buf[1] = hex[b/16]
		buf[2] = hex[b%16]
	}
	return dst
}

func escape(s string, allowReserved bool) (escaped string) {
	if allowReserved {
		escaped = string(reserved.ReplaceAllFunc([]byte(s), pctEncode))
	} else {
		escaped = string(unreserved.ReplaceAllFunc([]byte(s), pctEncode))
	}
	return escaped
}

// A UriTemplate is a parsed representation of a URI template.
type UriTemplate struct {
	raw   string
	parts []templatePart
}

// Parse parses a URI template string into a UriTemplate object.
func Parse(rawtemplate string) (template *UriTemplate, err error) {
	template = new(UriTemplate)
	template.raw = rawtemplate
	split := strings.Split(rawtemplate, "{")
	template.parts = make([]templatePart, len(split)*2-1)
	for i, s := range split {
		if i == 0 {
			if strings.Contains(s, "}") {
				err = errors.New("unexpected }")
				break
			}
			template.parts[i].raw = s
		} else {
			subsplit := strings.Split(s, "}")
			if len(subsplit) != 2 {
				err = errors.New("malformed template")
				break
			}
			expression := subsplit[0]
			template.parts[i*2-1], err = parseExpression(expression)
			if err != nil {
				break
			}
			template.parts[i*2].raw = subsplit[1]
		}
	}
	if err != nil {
		template = nil
	}
	return template, err
}

func (t UriTemplate) String() string {
	return t.raw
}

type templatePart struct {
	raw           string
	terms         []templateTerm
	first         string
	sep           string
	named         bool
	ifemp         string
	allowReserved bool
}

type templateTerm struct {
	name     string
	explode  bool
	truncate int
}

func parseExpression(expression string) (result templatePart, err error) {
	switch expression[0] {
	case '+':
		result.sep = ","
		result.allowReserved = true
		expression = expression[1:]
	case '.':
		result.first = "."
		result.sep = "."
		expression = expression[1:]
	case '/':
		result.first = "/"
		result.sep = "/"
		expression = expression[1:]
	case ';':
		result.first = ";"
		result.sep = ";"
		result.named = true
		expression = expression[1:]
	case '?':
		result.first = "?"
		result.sep = "&"
		result.named = true
		result.ifemp = "="
		expression = expression[1:]
	case '&':
		result.first = "&"
		result.sep = "&"
		result.named = true
		result.ifemp = "="
		expression = expression[1:]
	case '#':
		result.first = "#"
		result.sep = ","
		result.allowReserved = true
		expression = expression[1:]
	default:
		result.sep = ","
	}
	rawterms := strings.Split(expression, ",")
	result.terms = make([]templateTerm, len(rawterms))
	for i, raw := range rawterms {
		result.terms[i], err = parseTerm(raw)
		if err != nil {
			break
		}
	}
	return result, err
}

func parseTerm(term string) (result templateTerm, err error) {
	if strings.HasSuffix(term, "*") {
		result.explode = true
		term = term[:len(term)-1]
	}
	split := strings.Split(term, ":")
	if len(split) == 1 {
		result.name = term
	} else if len(split) == 2 {
		result.name = split[0]
		var parsed int64
		parsed, err = strconv.ParseInt(split[1], 10, 0)
		result.truncate = int(parsed)
	} else {
		err = errors.New("multiple colons in same term")
	}
	if !validname.MatchString(result.name) {
		err = errors.New("not a valid name: " + result.name)
	}
	if result.explode && result.truncate > 0 {
		err = errors.New("both explode and prefix modifers on same term")
	}
	return result, err
}

// Names returns the names of all variables within the template.
func (self *UriTemplate) Names() []string {
	names := make([]string, 0, len(self.parts))

	for _, p := range self.parts {
		if len(p.raw) > 0 || len(p.terms) == 0 {
			continue
		}

		for _, term := range p.terms {
			names = append(names, term.name)
		}
	}

	return names
}

// Expand expands a URI template with a set of values to produce a string.
func (self *UriTemplate) Expand(value interface{}) (string, error) {
	values, ismap := value.(map[string]interface{})
	if !ismap {
		if m, ismap := Struct2map(value); !ismap {
			return "", errors.New("expected map[string]interface{}, struct, or pointer to struct.")
		} else {
			return self.Expand(m)
		}
	}
	var buf bytes.Buffer
	for _, p := range self.parts {
		err := p.expand(&buf, values)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func (self *templatePart) expand(buf *bytes.Buffer, values map[string]interface{}) error {
	if len(self.raw) > 0 {
		buf.WriteString(self.raw)
		return nil
	}
	var zeroLen = buf.Len()
	buf.WriteString(self.first)
	var firstLen = buf.Len()
	for _, term := range self.terms {
		value, exists := values[term.name]
		if !exists {
			continue
		}
		if buf.Len() != firstLen {
			buf.WriteString(self.sep)
		}
		switch v := value.(type) {
		case string:
			self.expandString(buf, term, v)
		case []string:
			buf.WriteString(term.name + "=" + strings.Join(v, "%2C"))
		case []interface{}:
			self.expandArray(buf, term, v)
		case map[string]interface{}:
			if term.truncate > 0 {
				return errors.New("cannot truncate a map expansion")
			}
			self.expandMap(buf, term, v)
		default:
			if m, ismap := Struct2map(value); ismap {
				if term.truncate > 0 {
					return errors.New("cannot truncate a map expansion")
				}
				self.expandMap(buf, term, m)
			} else {
				str := fmt.Sprintf("%v", value)
				self.expandString(buf, term, str)
			}
		}
	}
	if buf.Len() == firstLen {
		original := buf.Bytes()[:zeroLen]
		buf.Reset()
		buf.Write(original)
	}
	return nil
}

func (self *templatePart) expandName(buf *bytes.Buffer, name string, empty bool) {
	if self.named {
		buf.WriteString(name)
		if empty {
			buf.WriteString(self.ifemp)
		} else {
			buf.WriteString("=")
		}
	}
}

func (self *templatePart) expandString(buf *bytes.Buffer, t templateTerm, s string) {
	if len(s) > t.truncate && t.truncate > 0 {
		s = s[:t.truncate]
	}
	self.expandName(buf, t.name, len(s) == 0)
	buf.WriteString(escape(s, self.allowReserved))
}

func (self *templatePart) expandArray(buf *bytes.Buffer, t templateTerm, a []interface{}) {
	if len(a) == 0 {
		return
	} else if !t.explode {
		self.expandName(buf, t.name, false)
	}
	for i, value := range a {
		if t.explode && i > 0 {
			buf.WriteString(self.sep)
		} else if i > 0 {
			buf.WriteString(",")
		}
		var s string
		switch v := value.(type) {
		case string:
			s = v
		default:
			s = fmt.Sprintf("%v", v)
		}
		if len(s) > t.truncate && t.truncate > 0 {
			s = s[:t.truncate]
		}
		if self.named && t.explode {
			self.expandName(buf, t.name, len(s) == 0)
		}
		buf.WriteString(escape(s, self.allowReserved))
	}
}

func (self *templatePart) expandMap(buf *bytes.Buffer, t templateTerm, m map[string]interface{}) {
	if len(m) == 0 {
		return
	}
	if !t.explode {
		self.expandName(buf, t.name, len(m) == 0)
	}
	var firstLen = buf.Len()
	for k, value := range m {
		if firstLen != buf.Len() {
			if t.explode {
				buf.WriteString(self.sep)
			} else {
				buf.WriteString(",")
			}
		}
		var s string
		switch v := value.(type) {
		case string:
			s = v
		default:
			s = fmt.Sprintf("%v", v)
		}
		if t.explode {
			buf.WriteString(escape(k, self.allowReserved))
			buf.WriteRune('=')
			buf.WriteString(escape(s, self.allowReserved))
		} else {
			buf.WriteString(escape(k, self.allowReserved))
			buf.WriteRune(',')
			buf.WriteString(escape(s, self.allowReserved))
		}
	}
}

func Struct2map(v interface{}) (map[string]interface{}, bool) {
	value := reflect.ValueOf(v)
	switch value.Type().Kind() {
	case reflect.Ptr:
		return Struct2map(value.Elem().Interface())
	case reflect.Struct:
		m := make(map[string]interface{})
		for i := 0; i < value.NumField(); i++ {
			tag := value.Type().Field(i).Tag
			var name string
			if strings.Contains(string(tag), ":") {
				uriName, opts := parseTag(tag.Get("uri"))
				sv := value.Field(i)
				if !opts.Contains("omitempty") || !isEmptyValue(sv) {
					name = uriName
				}
			} else {
				name = strings.TrimSpace(string(tag))
			}
			if len(name) == 0 {
				name = value.Type().Field(i).Name
			}
			m[name] = value.Field(i).Interface()
		}
		return m, true
	}
	return nil, false
}

type tagOptions []string

func parseTag(tag string) (string, tagOptions) {
	s := strings.Split(tag, ",")
	return s[0], s[1:]
}

func (o tagOptions) Contains(option string) bool {
	for _, s := range o {
		if s == option {
			return true
		}
	}
	return false
}

// isEmptyValue checks if a value should be considered empty for the purposes
// of omitting fields with the "omitempty" option.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	type zeroable interface {
		IsZero() bool
	}

	if z, ok := v.Interface().(zeroable); ok {
		return z.IsZero()
	}

	return false
}
