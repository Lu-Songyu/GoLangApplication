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

	blank := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	sep := placeholder{
		"   ",
		"   ",
		"███",
		"   ",
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
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	count := 0

	screen.Clear()

	for count >= 0 {

		screen.MoveTopLeft()
		hours, minutes, seconds := time.Now().Clock()
		month := time.Now().Month()
		day := time.Now().Day()
		time.Sleep(time.Second)

		clock := [...]placeholder{
			digits[month/10],
			digits[month%10],
			sep,
			digits[day/10],
			digits[day%10],
			sep,
			digits[hours/10],
			digits[hours%10],
			colon,
			digits[minutes/10],
			digits[minutes%10],
			blank,
			digits[seconds/10],
			digits[seconds%10],
		}
		count++

		if count%2 == 0 {
			clock[8] = blank
			clock[11] = blank
		}
		if count%2 == 1 {
			clock[8] = colon
			clock[11] = colon

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
