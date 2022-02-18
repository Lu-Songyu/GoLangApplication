package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	// for keeping things easy to read and type-safe
	type placeholder [5]string

	colon := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [...]placeholder{
		colon, zero, one, two, three, four, five, six, seven, eight, nine,
	}
	count := 0

	screen.Clear()

	for count >= 0 {

		screen.MoveTopLeft()
		hours, minutes, seconds := time.Now().Clock()
		time.Sleep(time.Second)

		clock := [...]placeholder{
			digits[hours/10],
			digits[hours%10],
			colon,
			digits[minutes/10],
			digits[minutes%10],
			colon,
			digits[seconds/10],
			digits[seconds%10],
		}

		// for each digit
		for line := range clock[0] {
			for digit := range clock {
				//fmt.Print(digit, line)
				fmt.Print(clock[digit][line], " ")
			}
			fmt.Println(" ")
		}

	}
}
