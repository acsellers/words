package words

import (
	"fmt"
	"math/rand"
	"strings"
)

func QuickDouble() string {
	return fmt.Sprintf("%s %s",
		strings.Title(Adjectives[rand.Intn(len(Adjectives))]),
		strings.Title(Animals[rand.Intn(len(Animals))]),
	)
}
