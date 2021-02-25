package main

import "testing"

func TestRunner(t *testing.T) {
	URLSlow := "https://www.github.com"
	URLFast := "https://facebook.com"

	waited := URLFast
	result := Runner(URLSlow, URLFast)

	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}
