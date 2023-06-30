package tests

import (
	"streams"
	"streams/sliceutils"
	"testing"
)

func TestCollectWhenStreamsAreEmpty(t *testing.T) {
	done := make(chan struct{})
	emptyStream := streams.StreamGenerator[int](done)

	want := []int{}
	got := emptyStream.Collect()

	if sliceutils.SliceEqual[int](want, got) == false {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCollectWhenStreamsAreFilled(t *testing.T) {
	var stream streams.Stream[int]
	done := make(chan struct{})

	inputValues := []int{0, 1, 5, 10, -5, 600_000}

	stream = streams.StreamGenerator[int](done, inputValues...)

	got := stream.Collect()

	if sliceutils.SliceEqual[int](inputValues, got) == false {
		t.Errorf("got %v, wanted %v", got, inputValues)
	}
}
