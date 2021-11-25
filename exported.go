package cache

import "time"

var (
	std = New(time.Minute *5, time.Minute*5)
)

func Set(k string, x interface{}, d time.Duration) {
	std.Set(k,  x, d)
}

func Get(k string) (interface{}, bool)  {
	return std.Get(k)
}

func Add(k string, x interface{}, d time.Duration) error {
	return std.Add(k, x, d)
}