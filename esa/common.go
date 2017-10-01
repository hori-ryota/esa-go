package esa

import "strconv"

// URLResp is struct for wrap url
type URLResp struct {
	URL string `json:"url"`
}

func uintToStr(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}

func wrap(key string, res interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: res,
	}
}

func StrP(v string) *string {
	return &v
}

func BoolP(v bool) *bool {
	return &v
}

func IntP(v int) *int {
	return &v
}

func UintP(v uint) *uint {
	return &v
}
