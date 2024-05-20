package dao

import (
	"hadoopmock/cmd/internal/hadoop"
	"reflect"
	"testing"
)

var (
	nodeMomentDataList = []*hadoop.NodeMomentData{
		&hadoop.NodeMomentData{
			NodeId:    "3279ng9438",
			Bandwidth: 3200,
		},
		&hadoop.NodeMomentData{
			NodeId:    "n948790v438",
			Bandwidth: 3000,
		},
	}
	nodeBwMockData = hadoop.VendorNodeMomentData{Id: "test", NodeMomentDataList: nodeMomentDataList}
)

func TestInsertVendorNodeBwModel(t *testing.T) {
	type args struct {
		data *hadoop.VendorNodeMomentData
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{&nodeBwMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertVendorNodeBwModel(tt.args.data)
		})
	}
}

func TestFindVendorNodeBwModelApiById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *hadoop.VendorNodeMomentData
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
			InsertVendorNodeBwModel(&nodeBwMockData)
			if got := FindVendorNodeBwModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindVendorNodeBwModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateVendorNodeBwModel(t *testing.T) {
	type args struct {
		data *hadoop.VendorNodeMomentData
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "case3", args: args{data: &nodeBwMockData}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateVendorNodeBwModel(tt.args.data)
		})
	}
}

func TestDeleteVendorNodeBwModel(t *testing.T) {
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
			DeleteVendorNodeBwModel(tt.args.id)
		})
	}
}
