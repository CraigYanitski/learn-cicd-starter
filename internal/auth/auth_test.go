package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestX(t *testing.T) {
    type testCase struct {
        input     string
        expected  string
        err       error
    }

    tests := []testCase{
        {
            "ApiKey test_string",
            "test_string",
            nil,
        },
        {
            "",
            "",
            ErrNoAuthHeaderIncluded,
        },
    }

    failCount := 0
    passCount := 0

    fmt.Println("\n\nTesting GetAPIKey")

    for _, test := range tests {
        fmt.Println("----------------------------------------")
        fmt.Printf("Getting API key with %v", test.input)

        header := http.Header{}
        if test.input != "" {
            header.Add("Authorization", test.input)
        }
        result, err := GetAPIKey(header)

        if result != test.expected || err != test.err {
            failCount++
            t.Errorf(`
Inputs:        %v
Expected:      %v
Error:         %v
Actual:        %v
Actual error:  %v
`, test.input, test.expected, test.err, result, err)
        } else {
            passCount++
            fmt.Printf(`
Inputs:        %v
Expected:      %v
Error:         %v 
Actual:        %v
Actual error:  %v
`, test.input, test.expected, test.err, result, err)
        }
    }

    fmt.Println("========================================")
    fmt.Printf("%d passed, %d failed\n\n\n", passCount, failCount)
}
