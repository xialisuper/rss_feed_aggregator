package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// respondWithError 函数接收一个 http.ResponseWriter 对象、状态码和消息作为参数，
// 设置响应头的 Content-Type 为 application/json; charset=utf-8，
// 设置状态码并返回错误信息的 JSON 格式。
func respondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, msg)))

}

// respondWithJSON 函数接收一个 http.ResponseWriter 对象、状态码以及一个任意类型的数据作为参数，
// 设置响应头的 Content-Type 为 application/json; charset=utf-8，
// 设置状态码，将数据转换为 JSON 格式并返回。
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}
