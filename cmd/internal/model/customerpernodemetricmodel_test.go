package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var customerPerNodeMetricMockData = types.CustomerPerNodeMetric{
	Id: "test-1110",
	MeasureCommonDataNodes: types.MeasureCommonDataNodes{
		NodeList: []*types.MeasureCommonDataPerNode{&perNodeMetricMockData},
	},
}

func Test_defaultCustomerPerNodeMetricModel_DeleteById(t *testing.T) {
	defaultCustomerPerNodeMetricModel := NewCustomerPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultCustomerPerNodeMetricModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultCustomerPerNodeMetricModel_FindById(t *testing.T) {
	defaultCustomerPerNodeMetricModel := NewCustomerPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.CustomerPerNodeMetric
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &customerPerNodeMetricMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultCustomerPerNodeMetricModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultCustomerPerNodeMetricModel_Insert(t *testing.T) {
	defaultCustomerPerNodeMetricModel := NewCustomerPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.CustomerPerNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &customerPerNodeMetricMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultCustomerPerNodeMetricModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultCustomerPerNodeMetricModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultCustomerPerNodeMetricModel_Update(t *testing.T) {
	defaultCustomerPerNodeMetricModel := NewCustomerPerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.CustomerPerNodeMetric
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &customerPerNodeMetricMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultCustomerPerNodeMetricModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
