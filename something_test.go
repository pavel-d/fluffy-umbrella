package main

import "testing"

func FailedTest(t *testing.T) {
   t.Errorf("BroKen")
}
