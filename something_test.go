package main

import "testing"

func FailedTest(t *testing.T) {
   t.Errorf("BroKe")
}
