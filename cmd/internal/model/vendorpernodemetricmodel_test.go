package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var perNodeMetricMockData = types.MeasureCommonDataPerNode{NodeId: "test-123456", DayPeak: dayPeak, Peak: types.MeasureCommonUnit{Bandwidth: 5000}}
var vendorPerNodeMetricMockData = types.VendorPerNodeMetric{
	Id: "test-1110",
	MeasureCommonDataNodes: types.MeasureCommonDataNodes{
		NodeList: []*types.MeasureCommonDataPerNode{&perNodeMetricMockData},
	},
}

func Test_defaultVendorPerNodeMetricModel_DeleteById(t *testing.T) {
	defaultVendorPerNodeMetricModel := NewVendorPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case4", args{ctx: context.Background(), id: "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultVendorPerNodeMetricModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultVendorPerNodeMetricModel_FindById(t *testing.T) {
	defaultVendorPerNodeMetricModel := NewVendorPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.VendorPerNodeMetric
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &vendorPerNodeMetricMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultVendorPerNodeMetricModel.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FindById() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_defaultVendorPerNodeMetricModel_Insert(t *testing.T) {
	defaultVendorPerNodeMetricModel := NewVendorPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorPerNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &vendorPerNodeMetricMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultVendorPerNodeMetricModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultVendorPerNodeMetricModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultVendorPerNodeMetricModel_Update(t *testing.T) {
	defaultVendorPerNodeMetricModel := NewVendorPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorPerNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &vendorPerNodeMetricMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultVendorPerNodeMetricModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
