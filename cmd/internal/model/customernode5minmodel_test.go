package model

import (
	"context"
	"hadoopmock/cmd/internal/common"
	"reflect"
	"testing"
)

var (
	MeasureCommonUnitList = []*common.MeasureCommonUnit{
		&common.MeasureCommonUnit{
			Bandwidth: 3200,
			Time:      1665331200,
		},
		&common.MeasureCommonUnit{
			Bandwidth: 1900,
			Time:      1665333200,
		},
	}
	node5MinMockData = common.MeasureCommonUnitList{Id: "test", MeasureCommonUnitList: MeasureCommonUnitList}
)

func Test_defaultCustomerNode5MinModel_DeleteById(t *testing.T) {
	defaultCustomerNode5MinModel := NewCustomerNode5MinModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultCustomerNode5MinModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultCustomerNode5MinModel_FindById(t *testing.T) {
	defaultCustomerNode5MinModel := NewCustomerNode5MinModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.MeasureCommonUnitList
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &node5MinMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultCustomerNode5MinModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultCustomerNode5MinModel_Insert(t *testing.T) {
	defaultCustomerNode5MinModel := NewCustomerNode5MinModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.MeasureCommonUnitList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &node5MinMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultCustomerNode5MinModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultCustomerNode5MinModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultCustomerNode5MinModel_Update(t *testing.T) {
	defaultCustomerNode5MinModel := NewCustomerNode5MinModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.MeasureCommonUnitList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &node5MinMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultCustomerNode5MinModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
