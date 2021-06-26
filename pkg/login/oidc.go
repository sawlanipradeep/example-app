package login

import (
	quote "rsc.io/quote/v3"
)

func OidcLogin() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}