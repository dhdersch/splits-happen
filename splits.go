package main

import (
	"fmt"
	"os"
)

// seqVal gets the value of seq[index]. Does not handle index out of bounds.
func seqVal(index int, seq string) int {
	char := seq[index]
	switch char {
	case 'X':
		return 10
	case '/':
		return 10 - seqVal(index-1, seq)
	case '-':
		return 0
	}
	return int(char - '0')
}

func sumAll(arr2d [][]int) int {
	total := 0
	for _, row := range arr2d {
		for _, item := range row {
			total += item
		}
	}
	return total
}

// SplitsScore - given a valid seq, determines the 9 pin bowling score
func SplitsScore(seq string) int {

	currentFrame := 0           // track the frame, should never be greater than 9
	frames := make([][]int, 10) // each row is a frame. contains the points for each frame

	// tracks the number of pins in the current frame. Going >= 2 indicates a new frame has begun and should be reset
	framePinCount := 0

	for index, char := range seq {
		if currentFrame > 9 {
			break
		}

		if framePinCount >= 2 {
			framePinCount = 0
			currentFrame++
		}

		switch char {
		case 'X':
			frames[currentFrame] = append(frames[currentFrame], 10, seqVal(index+1, seq), seqVal(index+2, seq))
			currentFrame++
			framePinCount = 0
		case '/':
			frames[currentFrame] = append(frames[currentFrame], 10-frames[currentFrame][0], seqVal(index+1, seq))
			currentFrame++
			framePinCount = 0
			break
		case '-':
			frames[currentFrame] = append(frames[currentFrame], 0)
			framePinCount++
			break
		default:
			frames[currentFrame] = append(frames[currentFrame], seqVal(index, seq))
			framePinCount++
			break
		}
	}

	return sumAll(frames)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: splits seq")
		return
	}
	sequence := os.Args[1]
	score := SplitsScore(sequence)
	fmt.Println(score)
}
