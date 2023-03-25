package visualcrossing

import (
	"net/url"
)

type Arguments map[string]string

var Defaults = make(Arguments)
var CurrentOnly = Arguments{"include": "current", "unitGroup": "metric"}

func (args Arguments) ToURLValues() url.Values {
	v := url.Values{}
	for key, value := range args {
		v.Set(key, value)
	}
	return v
}
