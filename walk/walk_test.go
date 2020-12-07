package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	//expected := "Chris"
	//var got []string
	//
	//x := struct {
	//	Name string
	//}{expected}
	//
	//walk(x, func(input string) {
	//	got = append(got, input)
	//})
	//
	//if len(got) != 1 {
	//	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	//}
	//
	//if got[0] != expected {
	//	t.Errorf("got '%s', want '%s'", got[0], expected)
	//}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func walk(x interface{}, fn func(input string)) {
	//fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i:=0; i<val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i:=0; i<val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}

	//field := val.Field(0)
	//fn(field.String())
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
