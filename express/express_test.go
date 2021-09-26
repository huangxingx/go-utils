package express

import (
	"reflect"
	"testing"
)

func TestExpressExecute(t *testing.T) {
	type args struct {
		express string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		// 加法
		{name: "test_add", args: args{express: "1+1"}, want: 2.0},
		// 括号
		{name: "test_bracket_1", args: args{express: "(1+1)*2"}, want: 4.0},
		// 与
		{name: "test_and_1", args: args{express: "true and true"}, want: true},
		{name: "test_and_2", args: args{express: "true and 0"}, want: false},
		// 减法
		{name: "test_sub_1", args: args{express: "100-1-2"}, want: 97.0},
		{name: "test_sub_1", args: args{express: "100 sub 1 sub 2"}, want: 97.0},
		// 除法
		{name: "test_div_1", args: args{express: "100/50"}, want: 2.0},
		{name: "test_div_2", args: args{express: "1/5"}, want: 0.2},
		{name: "test_div_3", args: args{express: "1/(5 * 2)"}, want: 0.1},
		{name: "test_div_4", args: args{express: "1/5 * 2"}, want: 0.4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			express := NewExpress(tt.args.express)
			if got := express.Execute(nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExpress().Execute(nil) = %v, want %v", got, tt.want)
			}
		})
	}
}
