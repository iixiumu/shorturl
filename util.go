package main

import "fmt"

func trans(u uint64) string {
	s := ""
	for u > 0 {
		t := u % 62
		if t <= 9 {
			c := t - 0 + '0'
			s = string(c) + s
		} else if t <= 35 {
			c := t - 10 + 'a'
			s = string(c) + s
		} else {
			c := t - 36 + 'A'
			s = string(c) + s
		}
		u = u / 62
	}
	// fmt.Printf("%06s\n", s)
	return fmt.Sprintf("%06s", s)
}
