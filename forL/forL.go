package forl

import "fmt"

func ForPrint(m *map[string]string) {
	for key, value := range *m {
		fmt.Println(key, value)
	}
}
