// Package model api
package model

// BaseResponse base api response
type BaseResponse struct {
	OK  bool   `json:"ok"`
	Msg string `json:"msg"`
}

// Set set response and return BaseResponse
//
//	@receiver res
//	@param ok
//	@param msg
//	@return BaseResponse
func (res *BaseResponse) Set(ok bool, msg string) BaseResponse {
	res.Msg = msg
	res.OK = ok
	return *res
}
