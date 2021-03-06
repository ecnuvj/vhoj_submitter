package util

import (
	"fmt"
	"github.com/ecnuvj/vhoj_submitter/pkg/sdk/base"
)

func PbReplyf(status base.REPLY_STATUS, format string, args ...interface{}) *base.BaseResponse {
	return &base.BaseResponse{
		Status:  status,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewDefaultSuccessReply() *base.BaseResponse {
	return &base.BaseResponse{
		Status:  base.REPLY_STATUS_SUCCESS,
		Message: "success",
	}
}
