package mock

import (
	"io"
	"fmt"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigureableSleeper struct {
	duration time.Duration
}

func (o *ConfigureableSleeper) Sleep() {
	time.Sleep(o.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3;i > 0;i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, "Go!")
}

func main() {
	sleeper := &ConfigureableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}