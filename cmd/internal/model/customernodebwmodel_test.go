package model

import (
	"context"
	"hadoopmock/cmd/internal/common"
	"reflect"
	"testing"
)

var (
	nodeMomentDataList = []*common.NodeMomentData{
		&common.NodeMomentData{
			NodeId:    "3279ng9438",
			Bandwidth: 3200,
		},
		&common.NodeMomentData{
			NodeId:    "n948790v438",
			Bandwidth: 3000,
		},
	}
	nodeBwMockData = common.NodeMomentDataList{Id: "test", NodeMomentDataList: nodeMomentDataList}
)

func Test_defaultCustomerNodeBwModel_DeleteById(t *testing.T) {
	defaultCustomerNodeBwModel := NewCustomerNodeBwModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultCustomerNodeBwModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultCustomerNodeBwModel_FindById(t *testing.T) {
	defaultCustomerNodeBwModel := NewCustomerNodeBwModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.NodeMomentDataList
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &nodeBwMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultCustomerNodeBwModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultCustomerNodeBwModel_Insert(t *testing.T) {
	defaultCustomerNodeBwModel := NewCustomerNodeBwModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.NodeMomentDataList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &nodeBwMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultCustomerNodeBwModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultCustomerNodeBwModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultCustomerNodeBwModel_Update(t *testing.T) {
	defaultCustomerNodeBwModel := NewCustomerNodeBwModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.NodeMomentDataList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &nodeBwMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultCustomerNodeBwModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
