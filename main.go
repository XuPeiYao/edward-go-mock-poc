package main

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func main() {

	t := &testing.T{}

	t.Helper()

	sp := &MockIServiceProvider{
		ctrl: gomock.NewController(t),
	}

	sp.EXPECT().GetService().Return("MOCKServiceProvider: GetService() called.")

	//fmt.Printf(sp.GetService())
}
