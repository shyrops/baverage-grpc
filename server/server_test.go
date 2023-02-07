package main

import (
	"testing"

	pb "beverage-grpc/pkg"
)

func Test_calculatePrice(t *testing.T) {
	type args struct {
		attr *pb.BeverageMainAttributes
	}

	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			"OK Add big bottle of beer",
			args{&pb.BeverageMainAttributes{Type: 1, Volume: 7}},
			210,
			false,
		},
		{
			"FAIL Unspecified drink type",
			args{&pb.BeverageMainAttributes{Type: 0, Volume: 7}},
			0,
			true,
		},
		{
			"FAIL Unsupported drink type",
			args{&pb.BeverageMainAttributes{Type: 99, Volume: 7}},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculatePrice(tt.args.attr)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculatePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculatePrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
