package mow

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

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

func move(m *Mower, maxX, maxY int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, instruction := range m.Movements {
		switch instruction {
		case 'L':
			m.Orientation = (m.Orientation - 1) % 4
		case 'R':
			m.Orientation = (m.Orientation + 1) % 4
		case 'F':
			switch m.Orientation {
			case N:
				m.Y = min(m.Y+1, maxY)
			case E:
				m.X = min(m.X+1, maxX)
			case S:
				m.Y = max(m.Y-1, 0)
			case W:
				m.X = max(m.X-1, 0)
			}
		default:
			fmt.Println("Wrong instruction ", m.Movements)
		}
	}
}

func orientation(orientation string) (Orientation, error) {
	if len(orientation) > 1 {
		return N, fmt.Errorf("Orientation should only contain one character", orientation)
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
		return N, fmt.Errorf("Orientation should be one of N, E, S or W")
	}
}

/**
 * Command
 */

func findMower(c *cli.Context) error {
	fileName := c.String("file")

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read lawn size
	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Something went wrong", err)
		return err
	}

	var maxX, maxY int
	limits := strings.Split(strings.TrimSuffix(line, "\n"), " ")
	maxX, err = strconv.Atoi(limits[0])
	if err != nil {
		fmt.Println("First line should contain two positive integers indicating lawn size: ", err)
		return err
	}
	maxY, err = strconv.Atoi(limits[1])
	if err != nil {
		fmt.Println("First line should contain two positive integers indicating lawn size: ", err)
		return err
	}

	// Read positions and movements for each mower
	var mowers []Mower
	for {
		// Read positions
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		unparsedMower := strings.Split(strings.TrimSuffix(line, "\n"), " ")
		mower := Mower{}
		mower.X, err = strconv.Atoi(unparsedMower[0])
		if err != nil {
			fmt.Println("Should be x coordinate: ", err)
			return err
		}
		mower.Y, err = strconv.Atoi(unparsedMower[1])
		if err != nil {
			fmt.Println("Should be y coordinate: ", err)
			return err
		}
		mower.Orientation, err = orientation(unparsedMower[2])
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Read movements
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		mower.Movements = strings.TrimSuffix(line, "\n")
		mowers = append(mowers, mower)
	}

	// Update mowers coordinates
	var wg sync.WaitGroup
	for i := range mowers {
		wg.Add(1)
		go move(&mowers[i], maxX, maxY, &wg)
	}

	// Print mowers coordinates
	wg.Wait()
	for i := range mowers {
		fmt.Println(mowers[i].X, mowers[i].Y, mowers[i].Orientation.toString())
	}
	return nil
}
