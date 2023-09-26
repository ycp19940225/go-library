package utils

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	type args[T comparable] struct {
		src []T
		dst []T
	}
	type testCase[T comparable] struct {
		name    string
		args    args[T]
		wantRes []T
	}
	tests := []testCase[int]{
		{
			name: "testInt",
			args: args[int]{
				src: []int{1, 2},
				dst: []int{1},
			},
			wantRes: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Diff(tt.args.src, tt.args.dst); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Difference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

	tests2 := []testCase[uint]{
		{
			name: "testUint",
			args: args[uint]{
				src: []uint{1, 2, 3, 4},
				dst: []uint{1, 2},
			},
			wantRes: []uint{3, 4},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Diff(tt.args.src, tt.args.dst); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Difference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

	tests3 := []testCase[string]{
		{
			name: "testStr",
			args: args[string]{
				src: []string{"1", "2", "3", "4"},
				dst: []string{"1", "2", "3"},
			},
			wantRes: []string{"4"},
		},
	}
	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Diff(tt.args.src, tt.args.dst); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Difference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}

	tests4 := []testCase[int]{
		{
			name: "testInt",
			args: args[int]{
				src: []int{3, 4},
				dst: []int{1, 2},
			},
			wantRes: []int{},
		},
	}
	for _, tt := range tests4 {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Diff(tt.args.src, tt.args.dst); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Difference() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
