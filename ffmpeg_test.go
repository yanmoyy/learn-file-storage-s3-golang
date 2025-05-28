package main

import "testing"

func TestGetVideoAspectRatio(t *testing.T) {
	pathV := "./samples/boots-video-vertical.mp4" // 9:16
	pathH := "./samples/boots-video-horizontal.mp4"
	test := func(path, expected string) {
		ratio, err := getVideoAspectRatio(path)
		if err != nil {
			t.Fatalf("Error Get video Aspect Ratio: %s", err)
		}
		if ratio != expected {
			t.Fatalf("ratio not match. expected=%s, actual=%s", expected, ratio)
		}
	}
	test(pathV, "9:16")
	test(pathH, "16:9")
}
