package goconvertstruct

import (
	"reflect"
	"testing"
)

type testCase struct {
	name string
	exec func() interface{}
	want func() interface{}
}

func TestConvert(t *testing.T) {
	tests := []testCase{
		{
			name: "convert nil",
			exec: func() interface{} {
				converted := Convert(nil, nil)
				return converted
			},
			want: func() interface{} {
				return nil
			},
		},
		{
			name: "convert int",
			exec: func() interface{} {
				converted := Convert(0, nil)
				return converted
			},
			want: func() interface{} {
				return 0
			},
		},
		{
			name: "convert float",
			exec: func() interface{} {
				converted := Convert(0.65, nil)
				return converted
			},
			want: func() interface{} {
				return 0.65
			},
		},
		{
			name: "convert string",
			exec: func() interface{} {
				converted := Convert("this is testing", nil)
				return converted
			},
			want: func() interface{} {
				return "this is testing"
			},
		},
		{
			name: "convert chan expected got nil",
			exec: func() interface{} {
				ch := make(chan int)
				converted := Convert(ch, nil)
				return converted
			},
			want: func() interface{} {
				return nil
			},
		},
		{
			name: "convert int of pointer",
			exec: func() interface{} {
				i := new(int)
				j := 123
				i = &j
				converted := Convert(i, nil)
				return converted
			},
			want: func() interface{} {
				return 123
			},
		},
		{
			name: "convert slice of integer",
			exec: func() interface{} {
				s := []int{1, 2, 3, 4}
				converted := Convert(s, nil)
				return converted
			},
			want: func() interface{} {
				return []interface{}{1, 2, 3, 4}
			},
		},
		{
			name: "convert slice of integer with pointer",
			exec: func() interface{} {
				i := 100
				s := []interface{}{2, &i, 4, 5, 6}
				converted := Convert(s, nil)
				return converted
			},
			want: func() interface{} {
				return []interface{}{2, 100, 4, 5, 6}
			},
		},
		{
			name: "convert nil slice",
			exec: func() interface{} {
				var s []int
				converted := Convert(s, nil)
				return converted
			},
			want: func() interface{} {
				return nil
			},
		},
		{
			name: "convert array of integer",
			exec: func() interface{} {
				s := [5]int{2, 3, 4, 5, 6}
				converted := Convert(s, nil)
				return converted
			},
			want: func() interface{} {
				return []interface{}{2, 3, 4, 5, 6}
			},
		},
		{
			name: "convert nil map",
			exec: func() interface{} {
				var m map[int]int
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return nil
			},
		},
		{
			name: "convert map int of int",
			exec: func() interface{} {
				m := map[int]int{1: 3, 2: 4, 5: 5}
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return map[int]interface{}{1: 3, 2: 4, 5: 5}
			},
		},
		{
			name: "convert map string of interface",
			exec: func() interface{} {
				m := map[string]interface{}{
					"field_1": "this is test string",
					"field_2": 987654321,
					"field_3": 0.1234567,
					"field_4": map[string]string{
						"nested_1": "value_1",
						"nested_2": "value_2",
						"nested_3": "value_3",
					},
				}
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return map[string]interface{}{
					"field_1": "this is test string",
					"field_2": 987654321,
					"field_3": 0.1234567,
					"field_4": map[string]interface{}{
						"nested_1": "value_1",
						"nested_2": "value_2",
						"nested_3": "value_3",
					},
				}
			},
		},
		{
			name: "convert map string of struct but field is unexported",
			exec: func() interface{} {
				m := map[string]struct {
					key   string
					value int
				}{
					"case_1": {key: "key_1", value: 1},
					"case_2": {key: "key_2", value: 2},
					"case_3": {key: "key_3", value: 3},
				}
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return map[string]interface{}{
					"case_1": map[string]interface{}{},
					"case_2": map[string]interface{}{},
					"case_3": map[string]interface{}{},
				}
			},
		},
		{
			name: "convert map string of struct but field is exported",
			exec: func() interface{} {
				m := map[string]struct {
					Key   string
					Value int
				}{
					"case_1": {Key: "key_1", Value: 1},
					"case_2": {Key: "key_2", Value: 2},
					"case_3": {Key: "key_3", Value: 3},
				}
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return map[string]interface{}{
					"case_1": map[string]interface{}{
						"Key":   "key_1",
						"Value": 1,
					},
					"case_2": map[string]interface{}{
						"Key":   "key_2",
						"Value": 2,
					},
					"case_3": map[string]interface{}{
						"Key":   "key_3",
						"Value": 3,
					},
				}
			},
		},
		{
			name: "convert struct but field is mixed",
			exec: func() interface{} {
				i := 100
				m := DummyStructTest{
					Field1: map[string]string{
						"key_1": "value_1",
						"key_2": "value_2",
					},
					Field2: 9090,
					Field3: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
					field4: "skip unexported field",
					field5: &DummyStructTest{
						Field1: map[string]string{
							"nested_key_1": "nested_value_1",
							"nested_key_2": "nested_value_2",
						},
						Field2: 8080,
						Field3: []int{1, 2, 3, 4, 5},
						field4: "skip this field",
					},
					Field6: nil,
					Field7: nil,
					Field8: []DummyStructTest{
						{
							Field1: map[string]string{
								"array_0_key_1": "array_0_value_1",
								"array_0_key_2": "array_0_value_2",
							},
							Field2: 1,
							Field3: []int{9, 8, 7, 6},
							Field6: nil,
							Field7: &i,
							Field8: []DummyStructTest{},
						},
						{
							Field1: map[string]string{
								"array_1_key_1": "array_1_value_1",
								"array_1_key_2": "array_1_value_2",
							},
							Field2: 2,
							Field3: []int{5, 4, 3, 2, 1},
							field4: "skip unexported field",
							field5: &DummyStructTest{},
							Field6: &DummyStructTest{
								Field1: map[string]string{},
							},
							Field7:  &i,
							Field8:  []DummyStructTest{{}},
							Field10: "{...}",
						},
					},
					Field9: map[string]DummyStructTest{
						"this_is_dummy": {},
					},
					Field10: true,
				}
				converted := Convert(m, nil)
				return converted
			},
			want: func() interface{} {
				return map[string]interface{}{
					"field_1": map[string]interface{}{
						"key_1": "value_1",
						"key_2": "value_2",
					},
					"Field2":  9090,
					"Field3":  []interface{}{9, 8, 7, 6, 5, 4, 3, 2, 1},
					"field_6": nil,
					"field_7": nil,
					"field_8": []interface{}{
						map[string]interface{}{
							"field_1": map[string]interface{}{
								"array_0_key_1": "array_0_value_1",
								"array_0_key_2": "array_0_value_2",
							},
							"Field2":   1,
							"Field3":   []interface{}{9, 8, 7, 6},
							"field_6":  nil,
							"field_7":  100,
							"field_8":  []interface{}{},
							"field_9":  nil,
							"field_10": nil,
						},
						map[string]interface{}{
							"field_1": map[string]interface{}{
								"array_1_key_1": "array_1_value_1",
								"array_1_key_2": "array_1_value_2",
							},
							"Field2": 2,
							"Field3": []interface{}{5, 4, 3, 2, 1},
							"field_6": map[string]interface{}{
								"field_1":  map[string]interface{}{},
								"Field2":   0,
								"Field3":   nil,
								"field_6":  nil,
								"field_7":  nil,
								"field_8":  nil,
								"field_9":  nil,
								"field_10": nil,
							},
							"field_7": 100,
							"field_8": []interface{}{
								map[string]interface{}{
									"field_1":  nil,
									"Field2":   0,
									"Field3":   nil,
									"field_6":  nil,
									"field_7":  nil,
									"field_8":  nil,
									"field_9":  nil,
									"field_10": nil,
								},
							},
							"field_9":  nil,
							"field_10": "{...}",
						},
					},
					"field_9": map[string]interface{}{
						"this_is_dummy": map[string]interface{}{
							"field_1":  nil,
							"Field2":   0,
							"Field3":   nil,
							"field_6":  nil,
							"field_7":  nil,
							"field_8":  nil,
							"field_9":  nil,
							"field_10": nil,
						},
					},
					"field_10": true,
				}
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.exec()
			want := tc.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\ngot:  %+v\nwant: %+v", got, want)
			}
		})
	}
}
