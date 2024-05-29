package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var pingLossNodeRatioMockData = types.PingLossNodeRatio{
	Id: "test",
	PingLossNodeRatioData: types.PingLossNodeRatioData{
		PingLossRatio: map[string][]*types.NodePingLossRatioUnit{
			"line1": []*types.NodePingLossRatioUnit{
				&types.NodePingLossRatioUnit{
					NodeId: "testNodeId004",
					Ratio:  0.01,
				},
				&types.NodePingLossRatioUnit{
					NodeId: "testNodeId005",
					Ratio:  0.01,
				},
			},
			"line2": []*types.NodePingLossRatioUnit{
				&types.NodePingLossRatioUnit{
					NodeId: "testNodeId006",
					Ratio:  0.01,
				},
				&types.NodePingLossRatioUnit{
					NodeId: "testNodeId007",
					Ratio:  0.01,
				},
			},
		},
	},
}

func Test_defaultPingLossNodeRatioModel_DeleteById(t *testing.T) {
	defaultPingLossNodeRatioModel := NewPingLossNodeRatioModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultPingLossNodeRatioModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultPingLossNodeRatioModel_FindById(t *testing.T) {
	defaultPingLossNodeRatioModel := NewPingLossNodeRatioModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.PingLossNodeRatio
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &pingLossNodeRatioMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultPingLossNodeRatioModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultPingLossNodeRatioModel_Insert(t *testing.T) {
	defaultPingLossNodeRatioModel := NewPingLossNodeRatioModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.PingLossNodeRatio
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &pingLossNodeRatioMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultPingLossNodeRatioModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultPingLossNodeRatioModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultPingLossNodeRatioModel_Update(t *testing.T) {
	defaultPingLossNodeRatioModel := NewPingLossNodeRatioModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.PingLossNodeRatio
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &pingLossNodeRatioMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultPingLossNodeRatioModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
