package dao

import (
	"hadoopmock/cmd/internal/hadoop"
	"reflect"
	"testing"
)

func TestInsertCustomerNodeBwModel(t *testing.T) {
	type args struct {
		data *hadoop.NodeMomentDataList
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{&nodeBwMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertCustomerNodeBwModel(tt.args.data)
		})
	}
}

func TestFindCustomerNodeBwModelApiById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *hadoop.NodeMomentDataList
	}{
		{
			name: "case2",
			args: args{
				id: "test",
			},
			want: &nodeBwMockData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertCustomerNodeBwModel(&nodeBwMockData)
			if got := FindCustomerNodeBwModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCustomerNodeBwModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCustomerNodeBwModel(t *testing.T) {
	type args struct {
		data *hadoop.NodeMomentDataList
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case3", args: args{data: &nodeBwMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateCustomerNodeBwModel(tt.args.data)
		})
	}
}

func TestDeleteCustomerNodeBwModel(t *testing.T) {
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
			DeleteCustomerNodeBwModel(tt.args.id)
		})
	}
}
