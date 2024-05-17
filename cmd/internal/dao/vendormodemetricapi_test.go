package dao

import (
	"github.com/jinzhu/now"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/hadoop"
)

var (
	dayPeak = map[string]hadoop.MeasureCommonUnit{
		"2023-06-01": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-02": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-03": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-04": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
	}
	mockData = hadoop.MeasureCommonData{Id: "test", DayPeak: dayPeak, Peak: hadoop.MeasureCommonUnit{Bandwidth: 5000}}
)

func TestInsertVendorNodeMetricModel(t *testing.T) {
	type args struct {
		data *hadoop.MeasureCommonData
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{&mockData}},
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
			name: "case1",
			args: args{
				id: "test",
			},
			want: &mockData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertVendorNodeMetricModel(&mockData)
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
		{name: "case1", args: args{data: &mockData}},
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
		{"case1", args{id: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteVendorNodeMetricModel(tt.args.id)
		})
	}
}
