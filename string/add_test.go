package string

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_01",
			args: args{a: "12345", b: "12345"},
			want: "24690",
		},
		{
			name: "test_02",
			args: args{a: "6789", b: "12"},
			want: "6801",
		},
		{
			name: "test_03",
			args: args{a: "12345678", b: "12345"},
			want: "12358023",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
