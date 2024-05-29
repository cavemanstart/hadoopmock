package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/types"
)

var lineBWSeriesMockData = types.LineBWSeries{
	Id: "test",
	LineBandWidthSeries: types.LineBandWidthSeries{
		LineBandWidth: map[string][]*types.MeasureCommonUnit{
			"line1": []*types.MeasureCommonUnit{
				&types.MeasureCommonUnit{
					Time:      1708240500,
					Bandwidth: 1200000,
				},
				&types.MeasureCommonUnit{
					Time:      1708240600,
					Bandwidth: 1000000,
				},
			},
			"line2": []*types.MeasureCommonUnit{
				&types.MeasureCommonUnit{
					Time:      1708240800,
					Bandwidth: 1200000,
				},
				&types.MeasureCommonUnit{
					Time:      1708240900,
					Bandwidth: 1100000,
				},
			},
		},
	},
}

func Test_defaultLineBWSeriesModel_DeleteById(t *testing.T) {
	defaultLineBWSeriesModel := NewLineBWSeriesModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultLineBWSeriesModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultLineBWSeriesModel_FindById(t *testing.T) {
	defaultLineBWSeriesModel := NewLineBWSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.LineBWSeries
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &lineBWSeriesMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultLineBWSeriesModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultLineBWSeriesModel_Insert(t *testing.T) {
	defaultLineBWSeriesModel := NewLineBWSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.LineBWSeries
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &lineBWSeriesMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultLineBWSeriesModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultLineBWSeriesModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultLineBWSeriesModel_Update(t *testing.T) {
	defaultLineBWSeriesModel := NewLineBWSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *types.LineBWSeries
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &lineBWSeriesMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultLineBWSeriesModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
