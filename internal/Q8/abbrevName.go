package q8

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	ns := strings.Split(name, " ")
	fp := ns[0]
	sp := ns[1]
	return fmt.Sprintf("%s.%s", strings.ToUpper(string(fp[0])), strings.ToUpper(string(sp[0])))
}
