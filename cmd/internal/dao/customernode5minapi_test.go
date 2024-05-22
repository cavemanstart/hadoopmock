package dao

import (
	"hadoopmock/cmd/internal/hadoop"
	"reflect"
	"testing"
)

func TestInsertCustomerNode5MinModel(t *testing.T) {
	type args struct {
		data *hadoop.MeasureCommonUnitList
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{&node5MinMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertCustomerNode5MinModel(tt.args.data)
		})
	}
}

func TestFindCustomerNode5MinModelApiById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *hadoop.MeasureCommonUnitList
	}{
		{
			name: "case2",
			args: args{
				id: "test",
			},
			want: &node5MinMockData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertCustomerNode5MinModel(&node5MinMockData)
			if got := FindCustomerNode5MinModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCustomerNode5MinModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCustomerNode5MinModel(t *testing.T) {
	type args struct {
		data *hadoop.MeasureCommonUnitList
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case3", args: args{data: &node5MinMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCustomerNode5MinModel(tt.args.data)
		})
	}
}

func TestDeleteCustomerNode5MinModel(t *testing.T) {
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
			DeleteCustomerNode5MinModel(tt.args.id)
		})
	}
}
