package order

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test_1", args: args{array: []int{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "test_2", args: args{array: []int{3, 2, 1}}, want: []int{1, 2, 3}},
		{name: "test_3", args: args{array: []int{3, 1, 2}}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}
