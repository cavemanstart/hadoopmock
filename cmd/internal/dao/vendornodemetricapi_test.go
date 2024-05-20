package dao

import (
	"reflect"
	"testing"

	"github.com/jinzhu/now"

	"hadoopmock/cmd/internal/hadoop"
)

var (
	dayPeak = map[string]hadoop.MeasureCommonUnit{
		"2023-06-01": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-02": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-03": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-04": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
	}
	nodeMetricMockData = hadoop.MeasureCommonData{Id: "test", DayPeak: dayPeak, Peak: hadoop.MeasureCommonUnit{Bandwidth: 5000}}
)

func TestInsertVendorNodeMetricModel(t *testing.T) {
	type args struct {
		data *hadoop.MeasureCommonData
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{&nodeMetricMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertVendorNodeMetricModel(tt.args.data)
		})
	}
}

func TestFindVendorNodeMetricModelApiById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *hadoop.MeasureCommonData
	}{
		{
			name: "case2",
			args: args{
				id: "test",
			},
			want: &nodeMetricMockData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertVendorNodeMetricModel(&nodeMetricMockData)
			if got := FindVendorNodeMetricModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindVendorNodeMetricModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateVendorNodeMetricModel(t *testing.T) {
	type args struct {
		data *hadoop.MeasureCommonData
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case3", args: args{data: &nodeMetricMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateVendorNodeMetricModel(tt.args.data)
		})
	}
}

func TestDeleteVendorNodeMetricModel(t *testing.T) {
	type args struct {
		id string
	}

	tests := []struct {
		name string
		args args
	}{
		{"case4", args{id: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteVendorNodeMetricModel(tt.args.id)
		})
	}
}
