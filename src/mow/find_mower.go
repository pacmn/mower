package mow

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

/**
 * Helpers
 */

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (o Orientation) toString() string {
	switch o {
	case 0:
		return "N"
	case 1:
		return "E"
	case 2:
		return "S"
	case 3:
		return "W"
	default:
		return "?"
	}
}

func orientation(orientation string) (Orientation, error) {
	if len(orientation) > 1 {
		// FIXME: Not nice
		return N, fmt.Errorf("Orientation %s is too long", orientation)
	}
	switch orientation[0] {
	case 'N':
		return N, nil
	case 'E':
		return E, nil
	case 'S':
		return S, nil
	case 'W':
		return W, nil
	default:
		return N, fmt.Errorf("Wrong orientation %s", orientation[0])
	}
}

/**
 * Command
 */

func findMower(c *cli.Context) error {
	fileName := c.String("file")

	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Error", err)
	}

	var maxX, maxY int
	limits := strings.Split(strings.TrimSuffix(line, "\n"), " ")
	maxX, err = strconv.Atoi(limits[0])
	if err != nil {
		fmt.Println(err)
		return nil
	}
	maxY, err = strconv.Atoi(limits[1])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var mowers []Mower
	for i := 0; ; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		unparsedMower := strings.Split(strings.TrimSuffix(line, "\n"), " ")
		mower := Mower{}
		mower.X, err = strconv.Atoi(unparsedMower[0])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		mower.Y, err = strconv.Atoi(unparsedMower[1])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		mower.Orientation, err = orientation(unparsedMower[2])
		if err != nil {
			fmt.Println(err)
			return nil
		}

		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		mower.Movements = strings.TrimSuffix(line, "\n")
		mowers = append(mowers, mower)
		fmt.Println(mowers)
	}

	for _, mower := range mowers {
		for _, instruction := range mower.Movements {
			switch instruction {
			case 'L':
				mower.Orientation = (mower.Orientation - 1) % 4
			case 'R':
				mower.Orientation = (mower.Orientation + 1) % 4
			case 'F':
				switch mower.Orientation {
				case N:
					mower.Y = min(mower.Y+1, maxY)
				case E:
					mower.X = min(mower.X+1, maxX)
				case S:
					mower.Y = max(mower.Y-1, 0)
				case W:
					mower.X = max(mower.X-1, 0)
				}
			default:
				fmt.Println("Wrong instruction ", instruction, mower.Movements)
			}
		}
		fmt.Println(mower.X, mower.Y, mower.Orientation.toString())
	}
	return nil
}
