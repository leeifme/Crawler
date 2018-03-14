package main

import (
	"io"
	"reflect"
	"testing"

	"golang.org/x/text/encoding"
)

func Test_determineEncoding(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want encoding.Encoding
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := determineEncoding(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("determineEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
