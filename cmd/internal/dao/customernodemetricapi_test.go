package dao

import (
	"hadoopmock/cmd/internal/hadoop"
	"reflect"
	"testing"
)

func TestInsertCustomerNodeMetricModel(t *testing.T) {
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
			InsertCustomerNodeMetricModel(tt.args.data)
		})
	}
}

func TestFindCustomerNodeMetricModelApiById(t *testing.T) {
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
			InsertCustomerNodeMetricModel(&nodeMetricMockData)
			if got := FindCustomerNodeMetricModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCustomerNodeMetricModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCustomerNodeMetricModel(t *testing.T) {
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
			UpdateCustomerNodeMetricModel(tt.args.data)
		})
	}
}

func TestDeleteCustomerNodeMetricModel(t *testing.T) {
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
			DeleteCustomerNodeMetricModel(tt.args.id)
		})
	}
}
