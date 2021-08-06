package concurrency

import (
	"testing"
	"reflect"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	urls := []string{
		"http://google.com",
        "http://blog.gypsydave5.com",
        "waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, urls)

	want := len(urls)	
	got := len(actualResults)

	if got != want {
		t.Fatalf("Wanted %d, got %d", want, got)
	}

	expectedResults := map[string]bool{
		"http://google.com": true,
        "http://blog.gypsydave5.com": true,
        "waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(actualResults, expectedResults) {
		t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
	}
}