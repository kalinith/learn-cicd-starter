package auth

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		inputKey   string
		inputValue string
		want       string
		wantErr    string
	}{
		"No auth header 1":        {wantErr: "no authorization header"},
		"No auth header 2":        {inputKey: "Authorization", wantErr: "no authorization header"},
		"malformed auth header 1": {inputKey: "Authorization", inputValue: "-", wantErr: "malformed authorization header"},
		"malformed auth header 2": {inputKey: "Authorization", inputValue: "Bearer xxxxxx", wantErr: "malformed authorization header"},
		"valid header":            {inputKey: "Authorization", inputValue: "ApiKey xxxxxx", want: "xxxxxx", wantErr: "not expecting an error"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(tc.inputKey, tc.inputValue)

			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tc.wantErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
