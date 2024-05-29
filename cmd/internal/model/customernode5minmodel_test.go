package model

import (
	"context"
	"hadoopmock/cmd/internal/types"
	"reflect"
	"testing"
)

var (
	measureCommonUnitList = []*types.MeasureCommonUnit{
		&types.MeasureCommonUnit{
			Bandwidth: 3200,
			Time:      1665331200,
		},
		&types.MeasureCommonUnit{
			Bandwidth: 1900,
			Time:      1665333200,
		},
	}
	customerNode5MinMockData = types.CustomerNode5Min{
		Id: "test",
		MeasureCommonUnitList: types.MeasureCommonUnitList{
			MeasureCommonUnitList: measureCommonUnitList,
		},
	}
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
		want    *types.CustomerNode5Min
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &customerNode5MinMockData,
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
		data *types.CustomerNode5Min
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &customerNode5MinMockData}, false},
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
		data *types.CustomerNode5Min
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &customerNode5MinMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultCustomerNode5MinModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
