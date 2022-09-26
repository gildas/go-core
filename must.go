package core

// Must panics if there is an error, otherwise returns the given value
//
// Example:
//  var myurl = core.Must[*url.URL](url.Parse("https://www.acme.com"))
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
