// Reference:
// https://siongui.github.io/2018/03/10/go-set-of-all-elements-in-two-arrays/
// https://siongui.github.io/2018/03/09/go-match-common-element-in-two-array/

package pkg

func Intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	i := 0
	for _, item := range b {
      	if _, ok := m[item]; ok {
      		c = append(c, item)
      		i++
      	}
    }

    return
}

func Union(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	i := 0
	for _, item := range b {
		if _, ok := m[item]; !ok {
			c = append(c, item)
			i++
		}
	}

	return c
}
