package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var vendorNodeMetricMockData = types.VendorNodeMetric{
	Id: "test",
	MeasureCommonData: types.MeasureCommonData{
		DayPeak: dayPeak, Peak: types.MeasureCommonUnit{Bandwidth: 5000},
	},
}

func Test_defaultVendorNodeMetricModel_DeleteById(t *testing.T) {
	defaultVendorNodeMetricModel := NewVendorNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultVendorNodeMetricModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultVendorNodeMetricModel_FindById(t *testing.T) {
	defaultVendorNodeMetricModel := NewVendorNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.VendorNodeMetric
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &vendorNodeMetricMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultVendorNodeMetricModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultVendorNodeMetricModel_Insert(t *testing.T) {
	defaultVendorNodeMetricModel := NewVendorNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &vendorNodeMetricMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultVendorNodeMetricModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultVendorNodeMetricModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultVendorNodeMetricModel_Update(t *testing.T) {
	defaultVendorNodeMetricModel := NewVendorNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &vendorNodeMetricMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultVendorNodeMetricModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
