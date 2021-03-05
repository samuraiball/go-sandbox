package hello

import (
	_ "embed"
	"fmt"
)

//go:embed hello_embed.txt
var hello string

func HelloEmbed() {
	fmt.Println(hello)
}
