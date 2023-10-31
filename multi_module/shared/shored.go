package shared

import (
	"fmt"
	"github.com/oklog/ulid/v2"
)

func SHoge() {
	fmt.Println(ulid.Make())
}
