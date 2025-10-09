package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println(*stringFlag)

	subcommand1:=flag.NewFlagSet("firstSub",flag.ExitOnError)
	subcommand2:=flag.NewFlagSet("secondSub",flag.ExitOnError)

	firstFlag:=subcommand1.Bool("processing",false,"Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2:=subcommand2.String("language","go","enter you langauge")

	if len(os.Args)<2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1]{
		case "firstSub":
			subcommand1.Parse(os.Args[2:])
			fmt.Println("subCommand1:")
			fmt.Println("processing:", *firstFlag)
			fmt.Println("bytes:", *secondFlag)

		case "secondSub":
					subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}

}