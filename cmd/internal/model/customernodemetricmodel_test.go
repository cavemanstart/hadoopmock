package model

import (
	"context"
	"github.com/jinzhu/now"
	"hadoopmock/cmd/internal/common"
	"reflect"
	"testing"
)

var (
	dayPeak = map[string]common.MeasureCommonUnit{
		"2023-06-01": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-02": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-03": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
		"2023-06-04": {Time: now.BeginningOfDay().Unix(), Bandwidth: 0},
	}
	nodeMetricMockData = common.MeasureCommonData{Id: "test", DayPeak: dayPeak, Peak: common.MeasureCommonUnit{Bandwidth: 5000}}
)

func Test_defaultCustomerNodeMetricModel_DeleteById(t *testing.T) {
	defaultCustomerNodeMetricModel := NewCustomerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultCustomerNodeMetricModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultCustomerNodeMetricModel_FindById(t *testing.T) {
	defaultCustomerNodeMetricModel := NewCustomerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.MeasureCommonData
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &nodeMetricMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultCustomerNodeMetricModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultCustomerNodeMetricModel_Insert(t *testing.T) {
	defaultCustomerNodeMetricModel := NewCustomerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.MeasureCommonData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &nodeMetricMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultCustomerNodeMetricModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultCustomerNodeMetricModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultCustomerNodeMetricModel_Update(t *testing.T) {
	defaultCustomerNodeMetricModel := NewCustomerNodeMetricModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.MeasureCommonData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &nodeMetricMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultCustomerNodeMetricModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
