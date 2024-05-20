package dao

import (
	"hadoopmock/cmd/internal/hadoop"
	"reflect"
	"testing"
)

var (
	MeasureCommonUnitList = []*hadoop.MeasureCommonUnit{
		&hadoop.MeasureCommonUnit{
			Bandwidth: 3200,
			Time:      1665331200,
		},
		&hadoop.MeasureCommonUnit{
			Bandwidth: 1900,
			Time:      1665333200,
		},
	}
	node5MinMockData = hadoop.MeasureCommonUnitList{Id: "test", MeasureCommonUnitList: MeasureCommonUnitList}
)

func TestInsertVendorNode5MinModel(t *testing.T) {
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
			InsertVendorNode5MinModel(tt.args.data)
		})
	}
}

func TestFindVendorNode5MinModelApiById(t *testing.T) {
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
			InsertVendorNode5MinModel(&node5MinMockData)
			if got := FindVendorNode5MinModelApiById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindVendorNode5MinModelApiById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateVendorNode5MinModel(t *testing.T) {
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
			UpdateVendorNode5MinModel(tt.args.data)
		})
	}
}

func TestDeleteVendorNode5MinModel(t *testing.T) {
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
			DeleteVendorNode5MinModel(tt.args.id)
		})
	}
}
