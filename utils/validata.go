package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"intelligent_parking/utils/errors"
	"intelligent_parking/utils/msg"
)

/*
data 返回前端的json数据
*/
func WriteJsonMsgWithError(data interface{}, err error) (resp *msg.Resp) {
	var (
		msgStr string //错误信息
		code   int    //错误码
	)
	if err != nil {
		logs.Error(err)
		msgStr = err.Error()
		ErrType := errors.NoType
		if ErrType = errors.GetType(err); ErrType == errors.NoType {
			ErrType = errors.SystemError
		}
		code = int(ErrType)
	}
	resp = msg.NewResp(data, code, msgStr, nil)
	if resp.Data == nil {
		resp.Data = struct{}{}
	}
	return
}


/*
从参数中读取数据
datacon 前端请求体数据
params 需要映射的结构体
*/
func ReadJSON(datacon []byte, params interface{}) (err error) {
	if len(datacon) == 0 {
		return
	}
	if err = json.Unmarshal(datacon, &params); err != nil {
		err = errors.ParamInvalid.Wrap(err, "参数解析错误")
		return
	}
	return
}

//从参数中读取数据并校验
func ReadValidJSON(datacon []byte, params interface{}) (err error) {
	var (
		isValid bool   //校验是否通过
		errMsg  string //错误信息
	)
	if err = ReadJSON(datacon, params); err != nil {
		return
	}
	valid := validation.Validation{}
	if isValid, err = valid.Valid(params); err != nil {
		err = errors.ParamInvalid.Wrap(err, "校验参数错误")
		return
	}
	if !isValid {
		for _, e := range valid.Errors {
			errMsg += fmt.Sprintf("%s%s;", e.Field, e.Message)
		}
		err = errors.ParamInvalid.New(errMsg)
	}
	return
}
