// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import "testing"

func compareSlices(s1, s2 []string) bool {
    if len(s2) < 1 || len(s1) < len(s2){
        return false
    }

    for i := 0; i < len(s1); i++ {
        found := true
        for j := 0; j < len(s2); j++ {
            if s1[i+j] != s2[j] {
                found = false
                break
            }
        }
        if found {
            return true
        }
    }
    return false
}

func TestExtractIssues(t *testing.T) {
    tests := []struct {
        text string
        want []string
    }{
        {
            text: "TEST-1 this is a test",
            want: []string{"TEST-1"},
        },
        {
            text: "suffix [TEST-123] [TEST-234]",
            want: []string{"TEST-123", "TEST-234"},
        },
        {
            text: "[TEST-123] prefix [TEST-456]",
            want: []string{"TEST-123"},
        },
        {
            text: "Multiple issues: TEST-123, TEST-234, TEST-456",
            want: []string{"TEST-123", "TEST-234", "TEST-456"},
        },
        {
            text: "feature/TEST-123 [TEST-456] and [TEST-789]",
            want: []string{"TEST-123", "TEST-456", "TEST-789"},
        },
        {
            text: "TEST-123 TEST-456 TEST-789",
            want: []string{"TEST-123", "TEST-789"},
        },
        {
            text: "no issue",
            want: []string{},
        },
    }

    t.Errorf("TESTING")

    for _, test := range tests {
        var args Args
        args.Commit.Message = test.text
        args.Project = "TEST"
        got := extractIssues(args)
        t.Errorf("TEXT:%v || WANT: %s", got, test.want) 
        t.Errorf(" %v", compareSlices(got, test.want))
        if !compareSlices(got, test.want) {
            t.Errorf("Got issues %v, want %v", got, test.want)
        }
    }

}
func TestExtractInstanceName(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		// Test cases with URLs
		{"http://test.com", "test"},
		{"https://subdomain.test.com", "subdomain"},
		{"ftp://ftp.test.org", "ftp"},

		// Test cases with non-URL strings
		{"instance.test.com", "instance"},
		{"subdomain.instance.test.org", "subdomain"},
		{"localhost", "localhost"},

		// Test invalid or malformed URLs
		{"http://", ""},                // Invalid URL with no hostname
		{"invalid-url", "invalid-url"}, // Not a URL, should return the input string
	}

	for _, test := range tests {
		result := ExtractInstanceName(test.text)
		if result != test.want {
			t.Errorf("ExtractInstanceName(%q) = %q; expected %q", test.text, result, test.want)
		}
	}
}

// Test the toEnvironmentId function
func TestToEnvironmentId(t *testing.T) {
	tests := []struct {
		name           string
		args           Args
		expectedOutput string
	}{
		{
			name:           "Non-empty EnvironmentId",
			args:           Args{EnvironmentId: "env-123"},
			expectedOutput: "env-123",
		},
		{
			name:           "Empty EnvironmentId",
			args:           Args{EnvironmentId: ""},
			expectedOutput: "production",  // Updated to match the default value of "production"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toEnvironmentId(tt.args)
			if result != tt.expectedOutput {
				t.Errorf("toEnvironmentId() = %v, want %v", result, tt.expectedOutput)
			}
		})
	}
}

// Test the toEnvironmentType function
func TestToEnvironmentType(t *testing.T) {
	tests := []struct {
		name           string
		args           Args
		expectedOutput string
	}{
		{
			name:           "Non-empty EnvironmentType",
			args:           Args{EnvironmentType: "prod"},
			expectedOutput: "prod",
		},
		{
			name:           "Empty EnvironmentType",
			args:           Args{EnvironmentType: ""},
			expectedOutput: "production",  // Updated to match the default value of "production"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toEnvironmentType(tt.args)
			if result != tt.expectedOutput {
				t.Errorf("toEnvironmentType() = %v, want %v", result, tt.expectedOutput)
			}
		})
	}
}
