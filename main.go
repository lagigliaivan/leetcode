package main

import "time"

func main() {
	valueToString := func(v interface{}) string {
		if v == nil {
			return ""
		}

		switch val := v.(type) {
		case *time.Time:
			if val == nil {
				return ""
			}
		}

		return "error"
	}

	p := new(int)

	v := valueToString(p)

	println(v)
}
