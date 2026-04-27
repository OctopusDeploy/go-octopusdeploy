package newclient

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// UpdateRequest wraps a resource for a PUT. Marshals the resource normally by
// default; use Clear to force a field to be sent as its zero value over the wire
type UpdateRequest[T any] struct {
	resource *T
	cleared  map[string]struct{}
}

func NewUpdateRequest[T any](resource *T) *UpdateRequest[T] {
	return &UpdateRequest[T]{resource: resource}
}

// Clear forces jsonName to be sent as its zero value, overwriting any current
// value on the resource. Only use when wiping a field server-side; for normal
// updates, mutate the resource and don't call Clear.
func (r *UpdateRequest[T]) Clear(jsonName string) *UpdateRequest[T] {
	if r.cleared == nil {
		r.cleared = make(map[string]struct{})
	}
	r.cleared[jsonName] = struct{}{}
	return r
}

// Resource returns the wrapped resource.
func (r *UpdateRequest[T]) Resource() *T { return r.resource }

func (r *UpdateRequest[T]) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(r.resource)
	if err != nil {
		return nil, fmt.Errorf("update request: marshal resource: %w", err)
	}
	if len(r.cleared) == 0 {
		return raw, nil
	}

	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, fmt.Errorf("update request: decode for cleared fields: %w", err)
	}

	typ := reflect.TypeOf((*T)(nil)).Elem()
	index := indexFields(typ)
	for name := range r.cleared {
		field, ok := index[name]
		if !ok {
			return nil, fmt.Errorf("update request: no JSON field %q on %s", name, typ.Name())
		}
		empty, err := zeroJSON(field.Type)
		if err != nil {
			return nil, fmt.Errorf("update request: zero value for %q: %w", name, err)
		}
		obj[name] = empty
	}
	return json.Marshal(obj)
}

type fieldIndex map[string]reflect.StructField

var fieldIndexCache sync.Map

func indexFields(t reflect.Type) fieldIndex {
	if cached, ok := fieldIndexCache.Load(t); ok {
		return cached.(fieldIndex)
	}
	idx := fieldIndex{}
	addFields(t, idx)
	fieldIndexCache.Store(t, idx)
	return idx
}

// addFields follows encoding/json's flattening: descend into anonymous fields,
// outer names shadow embedded ones.
func addFields(t reflect.Type, idx fieldIndex) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Anonymous {
			addFields(f.Type, idx)
			continue
		}
		if !f.IsExported() {
			continue
		}
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := strings.SplitN(tag, ",", 2)[0]
		if name == "" {
			name = f.Name
		}
		if _, exists := idx[name]; !exists {
			idx[name] = f
		}
	}
}

// zeroJSON encodes t's zero value. Nil slices/maps are upgraded to []/{} so
// Clear on a collection emits an empty container, not null.
func zeroJSON(t reflect.Type) (json.RawMessage, error) {
	var v reflect.Value
	switch t.Kind() {
	case reflect.Slice:
		v = reflect.MakeSlice(t, 0, 0)
	case reflect.Map:
		v = reflect.MakeMap(t)
	default:
		v = reflect.New(t).Elem()
	}
	return json.Marshal(v.Interface())
}
