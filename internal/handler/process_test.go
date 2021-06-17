package handler

import "testing"

func TestJsonVerify(t *testing.T) {
	tests := []struct {
		got  string
		want bool
	}{
		{`{"a": 1, "result": true}`, true},
		{`{"test: 15, output: "1"}`, false},
		{`{"intput": true, "result": true}`, true},
		{`{"check": "55", "result": true}`, true},
	}
	for _, tt := range tests {
		if got := JsonVerify([]byte(tt.got)); got != tt.want {
			t.Errorf("JsonVerify() = %v, want %v", got, tt.want)
		}
	}
}
