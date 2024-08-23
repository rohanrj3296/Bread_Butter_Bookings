package forms

import "net/url"

type FormForTesting struct {
	url.Values
	Errors errors
}
