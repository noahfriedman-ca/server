// The main server used on https://noahfriedman.ca.
package main

import "unit.nginx.org/go"

func main() {
	if e := unit.ListenAndServe(":8080", nil); e != nil {
		panic(e)
	}
}
