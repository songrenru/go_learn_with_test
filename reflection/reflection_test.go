package main

import (
	"testing"
	"reflect"
)

func TestReflection(t *testing.T) {
	expected := "Eason"
	x := struct{
		string
	}{expected}

	var got []string
	
	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}
}

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		Expected []string
	} {
		{
			"Struct with one string field",
			struct{
				Name string
			} {"Eason"},
			[]string{"Eason"},
		},
		{
			"Struct with two string field",
			struct{
				Name string
				City string
			} {"Eason", "Shenzhen"},
			[]string{"Eason", "Shenzhen"},
		},
		{
			"Struct with non string field",
			struct{
				Name string
				Age int
			} {"Eason", 28},
			[]string{"Eason"},
		},
		{
			"Nested field",
			Person{
				"Eason",
				Profile{28, "Shenzhen"},
			},
			[]string{"Eason", "Shenzhen"},
		},
		{
			"Pointers to things",
			&Person{
				"Eason",
				Profile{28, "Shenzhen"},
			},
			[]string{"Eason", "Shenzhen"},
		},
		{
			"Slices",
			[]Profile {
				{33, "Shenzhen"},
				{34, "Beijing"},
			},
			[]string{"Shenzhen", "Beijing"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "Shenzhen"},
				{34, "Beijing"},
			},
			[]string{"Shenzhen", "Beijing"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got '%v', want '%v'", got, test.Expected)
			}
		})
	}

	t.Run("Maps", func(t *testing.T) {
		aMap := map[string]string {
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContain(t, got, "Bar")
		assertContain(t, got, "Boz")
	})
}

func assertContain(t *testing.T, haystack []string, needle string) {
	contain := false

	for _, x := range haystack {
		if x == needle {
			contain = true
			break
		}
	}

	if !contain {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}