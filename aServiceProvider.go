package main

type AServiceProvider struct {
}

func (a *AServiceProvider) GetService() string {
	return "AServiceProvider: GetService() called."
}
