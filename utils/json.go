package utils

import "github.com/bytedance/sonic"

func MarshalJSONToString(v any) string {
	s, _ := sonic.MarshalString(v)
	return s
}
