package pipefilter

import "testing"

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sumer := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sumer)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but actual is %d", ret)
	}
}
