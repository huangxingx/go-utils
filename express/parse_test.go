package express

import (
	"reflect"
	"testing"
)

func Test_parse2mpn(t *testing.T) {
	type args struct {
		express string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test_1", args: args{"12+2"}, want: []string{"12", "+", "2"}},
		{name: "test_2", args: args{"1.2+2"}, want: []string{"1.2", "+", "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse2mpn(tt.args.express); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse2mpn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSuffixExpress(t *testing.T) {
	type args struct {
		expressList []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test_1", args: args{expressList: []string{"12", "+", "2"}}, want: []string{"12", "2", "+"}},
		{name: "test_1", args: args{expressList: []string{"1", "+", "2", "*", "3"}}, want: []string{"1", "2", "3", "*", "+"}},
		{name: "test_1", args: args{expressList: []string{"1", "+", "2", "*", "3", "/", "4"}},
			want: []string{"1", "2", "3", "*", "4", "/", "+"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSuffixExpress(tt.args.expressList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSuffixExpress() = %v, want %v", got, tt.want)
			}
		})
	}
}