package main

import (
	"fmt"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestFoo(t *testing.T) {
	t.Helper()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sp := NewMockIServiceProvider(ctrl)

	sp.EXPECT().GetService().Return("MOCKServiceProvider: GetService() called.")

	fmt.Printf(sp.GetService())
}
