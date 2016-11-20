package cli

import (
	"fmt"
	"github.com/ConnorAtherton/tasks/config"
	format "github.com/ConnorAtherton/tasks/fmt"
	"os"
)

func Run(args []string) int {
	fmt.Println("Tasks", os.Args[1:])

	config, _ := config.New()

	fmt.Printf("Config(path: %s, formatter: %s)\n", config.Path, config.Formatter)

	l := format.NewLogger(config.Formatter)

	fmt.Println(l.Error("something went wrong"))
	fmt.Println(l.Complete("take out the trash"))
	fmt.Println(l.Incomplete("was the dishes"))

	return 0
}
