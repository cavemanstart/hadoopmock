package service

import (
	"fmt"
	"testing"
)

func TestMockVendorNodeMetric(t *testing.T) {
	type args struct {
		startDate string
		endDate   string
		days      int64
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case1", args: args{startDate: "2022-01-12", endDate: "2022-01-17", days: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := MockVendorNodeMetric(tt.args.startDate, tt.args.endDate, tt.args.days)
			fmt.Println(res)
		})
	}
}
