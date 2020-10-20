package octopusdeploy

import (
	"reflect"
	"testing"
)

const (
	plainJSONValue       = `"blah"`
	emptyJSONValue       = `""`
	secretJSONValue      = `{"HasValue":true,"NewValue":null}`
	secretFalseJSONValue = `{"HasValue":false,"NewValue":null}`
	secretJSONNewValue   = `{"HasValue":true,"NewValue":"blah"}`
)

func TestPropertyValueResource_MarshalJSON(t *testing.T) {
	type fields struct {
		*SensitiveValue
		*PropertyValue
	}

	plain := PropertyValue("blah")

	newValue := "blah"

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "Plain", fields: fields{PropertyValue: &plain}, want: []byte(plainJSONValue)},
		{name: "Secret HasValue true", fields: fields{SensitiveValue: &SensitiveValue{HasValue: true}}, want: []byte(secretJSONValue)},
		{name: "Secret HasValue false", fields: fields{SensitiveValue: &SensitiveValue{HasValue: false}}, want: []byte(secretFalseJSONValue)},
		{name: "Secret with new value", fields: fields{SensitiveValue: &SensitiveValue{HasValue: true, NewValue: &newValue}}, want: []byte(secretJSONNewValue)},
		{name: "Null", fields: fields{SensitiveValue: nil, PropertyValue: nil}, want: []byte(emptyJSONValue)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := PropertyValueResource{
				SensitiveValue: tt.fields.SensitiveValue,
				PropertyValue:  tt.fields.PropertyValue,
			}

			got, err := d.MarshalJSON()

			if (err != nil) != tt.wantErr {
				t.Errorf("PropertyValueResource.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				json := PrettyJSON(got)
				t.Log(json)
				t.Errorf("PropertyValueResource.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPropertyValueResource_UnmarshalJSON(t *testing.T) {
	type fields struct {
		*SensitiveValue
		*PropertyValue
	}
	type args struct {
		data []byte
	}

	plain := PropertyValue(`blah`)
	empty := PropertyValue(``)

	newValue := "blah"

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Plain", fields: fields{PropertyValue: &plain}, args: args{data: []byte(plainJSONValue)}},
		{name: "Secret HasValue true", fields: fields{SensitiveValue: &SensitiveValue{HasValue: true}}, args: args{data: []byte(secretJSONValue)}},
		{name: "Secret HasValue false", fields: fields{SensitiveValue: &SensitiveValue{HasValue: false}}, args: args{data: []byte(secretFalseJSONValue)}},
		{name: "Secret with new value", fields: fields{SensitiveValue: &SensitiveValue{HasValue: true, NewValue: &newValue}}, args: args{data: []byte(secretJSONNewValue)}},
		{name: "Null", fields: fields{SensitiveValue: nil, PropertyValue: &empty}, args: args{data: []byte(emptyJSONValue)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := PropertyValueResource{
				SensitiveValue: tt.fields.SensitiveValue,
				PropertyValue:  tt.fields.PropertyValue,
			}
			var d PropertyValueResource

			if err := d.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("PropertyValueResource.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(d, expected) {
				t.Errorf("PropertyValueResource.UnmarshalJSON() = %+v, want %+v", d, tt.fields)
			}
		})
	}
}
