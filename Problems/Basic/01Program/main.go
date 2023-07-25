/* How to use struct that is imported from another package */

package main

import (
	parent "01Program/father"
	child "01Program/father/son"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr. Jeremy Maclin"))

	c := new(child.Son)
	fmt.Println(c.Data("Riley Maclin"))
}
