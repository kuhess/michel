package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// These variables are set by the build mechanism with ldflags
var (
	Version   string
	GitCommit string
)

const usage string = `Usage: michel OR michel -h

  Hey Michel! Find me a name!

Version: %v
Commit:  %v
`

var helpFlag bool

func init() {
	flag.BoolVar(&helpFlag, "h", false, "Show help")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, Version, GitCommit)
	}
}

func main() {
	flag.Parse()

	if helpFlag {
		flag.Usage()
	} else {
		rand.Seed(time.Now().UnixNano())
		randomName := pickRandom(prefix) + " " + pickRandom(suffix)
		fmt.Printf("%v\n", strings.Title(randomName))
	}
}

func pickRandom(words []string) string {
	return words[rand.Intn(len(words))]
}

var prefix = []string{
	"happy",
	"old",
	"reverse",
	"bad",
	"big",
	"agnostic",
	"yellow",
	"tasty",
	"purple",
	"fast",
	"quick",
	"persistent",
	"solid",
	"dirty",
	"tiny",
	"pretty",
}

var suffix = []string{
	"bob",
	"robert",
	"john",
	"nick",
	"pizza",
	"bulldozer",
	"cheddar",
	"jacqueline",
	"corn",
	"jam",
	"queen",
	"hamster",
	"racoon",
	"squirrel",
	"potatoe",
	"hammer",
}
