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
			name: "test convert struct",
			exec: func() interface{} {
				ConvertStruct(struct {}{})
				return nil
			},
			want: func() interface{} {
				return nil
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.exec()
			want := tc.want()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\nGot: %+v\nWant: %+v", got, want)
			}
		})
	}
}
