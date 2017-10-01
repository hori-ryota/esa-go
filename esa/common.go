package esa

import "strconv"

// URLResp is struct for wrap url
type URLResp struct {
	URL string `json:"url"`
}

func uintToStr(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

func wrapRes(key string, res interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: res,
	}
}
