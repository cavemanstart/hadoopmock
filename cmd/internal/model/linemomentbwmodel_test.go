package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/common"
)

var lineMomentBandWidthMockData = common.LineMomentBandWidth{
	Id: "test",
	LineMoment: map[string][]common.LineCommonUint{
		"testNodeId001": []common.LineCommonUint{
			common.LineCommonUint{
				Line:      "line1",
				Bandwidth: 1200000,
			},
			common.LineCommonUint{
				Line:      "line2",
				Bandwidth: 1000000,
			},
		},
		"testNodeId002": []common.LineCommonUint{
			common.LineCommonUint{
				Line:      "line3",
				Bandwidth: 1200000,
			},
			common.LineCommonUint{
				Line:      "line4",
				Bandwidth: 1100000,
			},
		},
	},
}

func Test_defaultLineMomentBandWidthModel_DeleteById(t *testing.T) {
	defaultLineMomentBandWidthModel := NewLineMomentBandWidthModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultLineMomentBandWidthModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultLineMomentBandWidthModel_FindById(t *testing.T) {
	defaultLineMomentBandWidthModel := NewLineMomentBandWidthModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.LineMomentBandWidth
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &lineMomentBandWidthMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultLineMomentBandWidthModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultLineMomentBandWidthModel_Insert(t *testing.T) {
	defaultLineMomentBandWidthModel := NewLineMomentBandWidthModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.LineMomentBandWidth
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &lineMomentBandWidthMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultLineMomentBandWidthModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultLineMomentBandWidthModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultLineMomentBandWidthModel_Update(t *testing.T) {
	defaultLineMomentBandWidthModel := NewLineMomentBandWidthModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.LineMomentBandWidth
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &lineMomentBandWidthMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultLineMomentBandWidthModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
