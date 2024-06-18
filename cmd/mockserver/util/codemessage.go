package util

import "hadoopmock/cmd/mockserver/types"

func ErrResp() *types.HadoopResp {
	return &types.HadoopResp{500, nil, "error"}
}
func SuccessResp(data interface{}) *types.HadoopResp {
	return &types.HadoopResp{200, data, "success"}
}
