package eleventh_task

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func Eleven() {
	required := mapset.NewSet()
	required.Add("cooking")
	required.Add("english")
	required.Add("math")
	required.Add("biology")
	sciences := mapset.NewSet()
	sciences.Add("biology")
	sciences.Add("chemistry")
	count := required.Intersect(sciences)
	fmt.Println(count)
}
