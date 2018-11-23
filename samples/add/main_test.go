package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				s1: "1",
				s2: "2",
			},
			want: 3,
		},
		{
			name: "not a number#1",
			args: args{
				s1: "a",
				s2: "2",
			},
			want:    -1,
			wantErr: true,
		},
		{
			name: "not a number#2",
			args: args{
				s1: "1",
				s2: "a",
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := add(tt.args.s1, tt.args.s2)
			if (err != nil) != tt.wantErr {
				t.Errorf("add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
