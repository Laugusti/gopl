package params

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

// Pack returns a URL incorporating the parameter values from the struct.
func Pack(rawURL string, v interface{}) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	query := u.Query()

	rv := reflect.ValueOf(v)
	for i := 0; i < rv.NumField(); i++ {
		fieldInfo := rv.Type().Field(i)

		name := fieldInfo.Tag.Get("http")
		if name == "" {
			name = fieldInfo.Name
		}
		values, err := getValuesAsString(rv.Field(i))
		if err != nil {
			return "", err
		}
		for _, value := range values {
			query.Add(name, value)
		}

	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = fieldInfo.Name
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognize HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func getSingleValueAsString(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Int:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	default:
		return "", fmt.Errorf("unsupported kind %s", v.Type())
	}
}

func getValuesAsString(v reflect.Value) ([]string, error) {
	if v.Kind() == reflect.Slice {
		var values []string
		for i := 0; i < v.Len(); i++ {
			value, err := getSingleValueAsString(v.Index(i))
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	} else {
		value, err := getSingleValueAsString(v)
		return []string{value}, err
	}
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
