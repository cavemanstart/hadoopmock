package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var lineMomentBWMockData = types.LineMomentBW{
	Id: "test",
	LineMomentBandWidth: types.LineMomentBandWidth{
		LineMoment: map[string][]*types.LineCommonUint{
			"testNodeId001": []*types.LineCommonUint{
				&types.LineCommonUint{
					Line:      "line1",
					Bandwidth: 1200000,
				},
				&types.LineCommonUint{
					Line:      "line2",
					Bandwidth: 1000000,
				},
			},
			"testNodeId002": []*types.LineCommonUint{
				&types.LineCommonUint{
					Line:      "line3",
					Bandwidth: 1200000,
				},
				&types.LineCommonUint{
					Line:      "line4",
					Bandwidth: 1100000,
				},
			},
		},
	},
}

func Test_defaultLineMomentBWModel_DeleteById(t *testing.T) {
	defaultLineMomentBWModel := NewLineMomentBWModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultLineMomentBWModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultLineMomentBWModel_FindById(t *testing.T) {
	defaultLineMomentBWModel := NewLineMomentBWModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.LineMomentBW
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &lineMomentBWMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultLineMomentBWModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultLineMomentBWModel_Insert(t *testing.T) {
	defaultLineMomentBWModel := NewLineMomentBWModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.LineMomentBW
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &lineMomentBWMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultLineMomentBWModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultLineMomentBWModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultLineMomentBWModel_Update(t *testing.T) {
	defaultLineMomentBWModel := NewLineMomentBWModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.LineMomentBW
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &lineMomentBWMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultLineMomentBWModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
