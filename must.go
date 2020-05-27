package core

// Must panics if there is an error, otherwise returns the given value 
//
// Example:
//  var myurl = core.Must(url.Parse("https://www.acme.com")).(*url.URL)
func Must(value interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return value
}