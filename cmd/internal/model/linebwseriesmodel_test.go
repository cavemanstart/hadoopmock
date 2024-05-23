package model

import (
	"context"
	"reflect"
	"testing"

	"hadoopmock/cmd/internal/common"
)

var lineBandWidthSeriesMockData = common.LineBandWidthSeries{
	Id: "test",
	LineBandWidth: map[string][]common.MeasureCommonUnit{
		"line1": []common.MeasureCommonUnit{
			common.MeasureCommonUnit{
				Time:      1708240500,
				Bandwidth: 1200000,
			},
			common.MeasureCommonUnit{
				Time:      1708240600,
				Bandwidth: 1000000,
			},
		},
		"line2": []common.MeasureCommonUnit{
			common.MeasureCommonUnit{
				Time:      1708240800,
				Bandwidth: 1200000,
			},
			common.MeasureCommonUnit{
				Time:      1708240900,
				Bandwidth: 1100000,
			},
		},
	},
}

func Test_defaultLineBandWidthSeriesModel_DeleteById(t *testing.T) {
	defaultLineBandWidthSeriesModel := NewLineBandWidthSeriesModel("mongodb://localhost:27017", "hadoopMock")
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
			if err := defaultLineBandWidthSeriesModel.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultLineBandWidthSeriesModel_FindById(t *testing.T) {
	defaultLineBandWidthSeriesModel := NewLineBandWidthSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *common.LineBandWidthSeries
		wantErr bool
	}{
		{
			name: "case2",
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			want:    &lineBandWidthSeriesMockData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defaultLineBandWidthSeriesModel.FindById(tt.args.ctx, tt.args.id)
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

func Test_defaultLineBandWidthSeriesModel_Insert(t *testing.T) {
	defaultLineBandWidthSeriesModel := NewLineBandWidthSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.LineBandWidthSeries
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"case1", args{ctx: context.Background(), data: &lineBandWidthSeriesMockData}, false},
	}
	for _, tt := range tests {
		res, _ := defaultLineBandWidthSeriesModel.FindById(tt.args.ctx, tt.args.data.Id)
		if res == nil {
			t.Run(tt.name, func(t *testing.T) {
				if err := defaultLineBandWidthSeriesModel.Insert(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
					t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}

func Test_defaultLineBandWidthSeriesModel_Update(t *testing.T) {
	defaultLineBandWidthSeriesModel := NewLineBandWidthSeriesModel("mongodb://localhost:27017", "hadoopMock")
	type args struct {
		ctx  context.Context
		data *common.LineBandWidthSeries
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "case3", args: args{ctx: context.Background(), data: &lineBandWidthSeriesMockData}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := defaultLineBandWidthSeriesModel.Update(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
