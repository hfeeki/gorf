package pkg2

import pkg2_0	"pkg2"
import package1	"package1"

type T struct{ A, b int }

func (t *T) Bar() {
	pkg2_0.Bar()
}
func (t *T) Foo() {
	package1.Foo()
}
