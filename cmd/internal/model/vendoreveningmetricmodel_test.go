package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var vendorEveningMetricMockData = types.VendorEveningMetric{
	Id: "test",
	MeasureCommonData: types.MeasureCommonData{
		DayPeak: dayPeak, Peak: types.MeasureCommonUnit{Bandwidth: 5000},
	},
}

func Test_defaultVendorEveningMetricModel_DeleteById(t *testing.T) {
	defaultVendorEveningMetricModel := NewVendorEveningMetricModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultVendorEveningMetricModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultVendorEveningMetricModel_FindById(t *testing.T) {
	defaultVendorEveningMetricModel := NewVendorEveningMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.VendorEveningMetric
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &vendorEveningMetricMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultVendorEveningMetricModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultVendorEveningMetricModel_Insert(t *testing.T) {
	defaultVendorEveningMetricModel := NewVendorEveningMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorEveningMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &vendorEveningMetricMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultVendorEveningMetricModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultVendorEveningMetricModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultVendorEveningMetricModel_Update(t *testing.T) {
	defaultVendorEveningMetricModel := NewVendorEveningMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.VendorEveningMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &vendorEveningMetricMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultVendorEveningMetricModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
