package middlewire

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CustomResponseEncoder 自定义响应编码器
func CustomResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	resp := Response{
		Code:    200,
		Message: "OK",
		Data:    v, // 原始返回数据
	}
	return json.NewEncoder(w).Encode(resp)
}
