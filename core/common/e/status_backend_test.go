package e

import (
	"reflect"
	"testing"
)

func TestNewStatus(t *testing.T) {
	type args struct {
		run    byte
		init   byte
		kind   byte
		anchor byte
	}
	tests := []struct {
		name string
		args args
		want STATUS
	}{
		{
			name: "test1",
			args: args{
				run:    1,
				init:   1,
				kind:   1,
				anchor: 1,
			},
			want: STATUS([]byte{1, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatus(tt.args.run, tt.args.init, tt.args.kind, tt.args.anchor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatusDecode(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want STATUS
	}{
		{
			name: "1",
			args: args{status: "01010101"},
			want: STATUS([]byte{1, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusDecode(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StatusDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
