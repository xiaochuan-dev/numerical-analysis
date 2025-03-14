package chapter0

import (
	"fmt"
	"strings"
)

type Number interface {
    int | float64
}

func NestedMulti[T Number](a []T, x T) (res T) {

	n := len(a)

	var sbuilder strings.Builder

	for i := 0; i < n; i++ {
		sbuilder.WriteString(fmt.Sprintf("%vx^%d", a[i], n - i - 1))

		if i != n - 1 {
			sbuilder.WriteString("+")
		} else {
			sbuilder.WriteString("\n")
		}
	}

	fmt.Printf("输入的多项式是  %s", sbuilder.String())

	for i := 0; i < n; i++ {
		res = res * x + a[i]
	}

	return res
}
