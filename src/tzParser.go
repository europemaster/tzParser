package src

import (
	"flag"
	"fmt"
)

func main()  {

	flag.Parse()

	timeFormat := flag.Arg(0)
	from := flag.Arg(1)
	to := flag.Arg(2)

	fmt.Println(timeFormat, from, to)
}
