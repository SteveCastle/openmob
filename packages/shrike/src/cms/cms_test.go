package cms

import (
	"reflect"
	"testing"
)

func TestGetAllLandingPages(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want LandingPage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllLandingPages(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLandingPages() = %v, want %v", got, tt.want)
			}
		})
	}
}
