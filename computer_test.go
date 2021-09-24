package go_utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestComputer(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test_add", args: args{"1+2+3"}, want: 6},
		{name: "test_reduce", args: args{"1-2-3"}, want: -4},
		{name: "test_multi", args: args{"1*2*3"}, want: 6},
		{name: "test_div", args: args{"10/2/2"}, want: 2.5},
		{name: "test_div", args: args{"10/(2/2)"}, want: 10},
		{name: "test_div", args: args{"10/(1/2)"}, want: 20},
		{name: "test_div", args: args{"1/(2/2)"}, want: 1},
		{name: "test_div", args: args{"1/3"}, want: 1.0 / 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Computer(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Computer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse2mpn(t *testing.T) {
	type args struct {
		express string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_add",
			args: args{"1+2+3"},
			want: []string{"1", "+", "2", "+", "3"},
		},
		{
			name: "test_dec",
			args: args{"1-2-3"},
			want: []string{"1", "-", "2", "-", "3"},
		},
		{
			name: "test_multi",
			args: args{"1*2*3"},
			want: []string{"1", "*", "2", "*", "3"},
		},
		{
			name: "test_div",
			args: args{"1/2/3"},
			want: []string{"1", "/", "2", "/", "3"},
		},
		{
			name: "test_mix",
			args: args{"1+2-3*4/10"},
			want: []string{"1", "+", "2", "-", "3", "*", "4", "/", "10"},
		},
		{
			name: "test_brackets",
			args: args{"(1+2)*3"},
			want: []string{"(", "1", "+", "2", ")", "*", "3"},
		},
		{
			name: "test_space",
			args: args{" ( 1+ 2 )*3"},
			want: []string{"(", "1", "+", "2", ")", "*", "3"},
		},
		{
			name: "test_double",
			args: args{" ( 10+ 20 )*333"},
			want: []string{"(", "10", "+", "20", ")", "*", "333"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse2mpn(tt.args.express); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse2mpn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse2rpn(t *testing.T) {
	type args struct {
		express []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_add",
			args: struct{ express []string }{express: []string{"1", "+", "2", "+", "3"}},
			want: []string{"1", "2", "+", "3", "+"},
		},
		{
			name: "test_multi",
			args: struct{ express []string }{express: []string{"1", "+", "2", "*", "4"}},
			want: []string{"1", "2", "4", "*", "+"},
		},
		{
			name: "test_brackets",
			args: struct{ express []string }{express: []string{"(", "1", "+", "2", ")", "*", "4"}},
			want: []string{"1", "2", "+", "4", "*"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse2rpn(tt.args.express); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse2rpn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printRune(t *testing.T) {
	s := ".0123456789+-*/ ()"
	for _, v := range s {
		fmt.Printf("%s\t-->\t%d\n", string(v), v)
	}

}
