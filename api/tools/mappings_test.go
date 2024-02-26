package tools

import "testing"

func TestMapping(t *testing.T) {
	if Mapping("gender", 0) != "未知" {
		t.Error("映射错误")
	}

	if Mapping("zhansan", 0) != "unknown key" {
		t.Error("映射错误")
	}

	if Mapping("gender", 3) != "unknown" {
		t.Error("映射错误")
	}
}
