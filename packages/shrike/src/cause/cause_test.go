package cause

import (
	"reflect"
	"testing"
)

func TestGetAllCauses(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want Cause
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllCauses(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCauses() = %v, want %v", got, tt.want)
			}
		})
	}
}
