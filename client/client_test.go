package client

import (
	"testing"
)

func TestReplaceSensitiveHeadersPowerStore(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Redact Authorization header",
			input:    []byte("Authorization: Basic YWRtaW46cGFzc3dvcmQ="),
			expected: "Authorization: ******",
		},
		{
			name:     "Redact Dell-Emc-Token header",
			input:    []byte("Dell-Emc-Token: abc123xyz789"),
			expected: "Dell-Emc-Token: ******",
		},
		{
			name:     "Redact Cookie auth_cookie",
			input:    []byte("Cookie: auth_cookie=secret123; Path=/"),
			expected: "Cookie: auth_cookie=******; Path=/",
		},
		{
			name:     "Multiple sensitive headers",
			input:    []byte("Authorization: Basic YWRtaW46cGFzc3dvcmQ=\nDell-Emc-Token: abc123"),
			expected: "Authorization: ******\nDell-Emc-Token: ******",
		},
		{
			name:     "No sensitive headers",
			input:    []byte("Content-Type: application/json"),
			expected: "Content-Type: application/json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replaceSensitiveHeadersPowerStore(tt.input)
			if result != tt.expected {
				t.Errorf("replaceSensitiveHeadersPowerStore() = %v, want %v", result, tt.expected)
			}
		})
	}
}
