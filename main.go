package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func massivkaSalu(s string) [][][]rune {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	var arr [][]rune
	var arr3 [][][]rune
	for _, line := range lines {
		row := []rune(line)
		if len(row) > 0 {
			arr = append(arr, row)
		} else if len(row) == 0 && len(arr) > 0 {
			arr3 = append(arr3, arr)
			arr = nil
		}
	}
	if len(arr) > 0 {
		arr3 = append(arr3, arr)
	}
	return arr3
}

func massivKordinat(aa [][][]rune) [][][]int {
	var bb [][][]int
	for i := 0; i < len(aa); i++ {
		var bb2 [][]int
		for j := 0; j < len(aa[i]); j++ {
			var bb1 []int
			for k := 0; k < len(aa[i][j]); k++ {
				if aa[i][j][k] == '#' {
					bb1 = append(bb1, j)
					bb1 = append(bb1, k)
				}
			}
			if len(bb1) == 2 {
				bb2 = append(bb2, bb1)
			} else if len(bb1) == 4 {
				splitIndex := len(bb1) / 2
				firstPart := bb1[:splitIndex]
				secondPart := bb1[splitIndex:]
				bb2 = append(bb2, firstPart)
				bb2 = append(bb2, secondPart)
			} else if len(bb1) == 6 {
				splitIndex1 := len(bb1) / 3
				splitIndex2 := 2 * len(bb1) / 3
				firstPart := bb1[:splitIndex1]
				secondPart := bb1[splitIndex1:splitIndex2]
				thirdPart := bb1[splitIndex2:]
				bb2 = append(bb2, firstPart)
				bb2 = append(bb2, secondPart)
				bb2 = append(bb2, thirdPart)
			} else if len(bb1) == 8 {
				splitIndex1 := len(bb1) / 4
				splitIndex2 := len(bb1) / 2
				splitIndex3 := 3 * len(bb1) / 4
				firstPart := bb1[:splitIndex1]
				secondPart := bb1[splitIndex1:splitIndex2]
				thirdPart := bb1[splitIndex2:splitIndex3]
				fourthPart := bb1[splitIndex3:]
				bb2 = append(bb2, firstPart)
				bb2 = append(bb2, secondPart)
				bb2 = append(bb2, thirdPart)
				bb2 = append(bb2, fourthPart)
			}
		}
		bb = append(bb, bb2)
	}
	return bb
}

func transformArray(bb [][][]int) [][][]int {
	transformedArray := make([][][]int, len(bb))
	for i, bb2 := range bb {
		minX, minY := findMinCoordinates(bb2)
		for _, bb1 := range bb2 {
			transformedArray[i] = append(transformedArray[i], []int{bb1[0] - minX, bb1[1] - minY})
		}
	}
	return transformedArray
}

func findMinCoordinates(bb2 [][]int) (minX, minY int) {
	minX = bb2[0][0]
	minY = bb2[0][1]
	for _, bb1 := range bb2 {
		if bb1[0] < minX {
			minX = bb1[0]
		}
		if bb1[1] < minY {
			minY = bb1[1]
		}
	}
	return minX, minY
}

func checkNeighborhood(s [][][]rune) bool {
	count := 0
	errr := false
	for k := 0; k < len(s); k++ {
		for k1 := 0; k1 < len(s[k]); k1++ {
			val := s[k][k1]
			for k2 := 0; k2 < len(val); k2++ {
				if val[k2] == 35 {
					count++
					if k2 != 0 && val[k2-1] == 35 {
						count++
					}
					if k2 != len(val)-1 && val[k2+1] == 35 {
						count++
					}
					if k1 != len(s[k])-1 && s[k][k1+1][k2] == 35 {
						count++
					}
					if k1 != 0 && s[k][k1-1][k2] == 35 {
						count++
					}
				}
			}
		}
		if count == 10 || count == 12 {
			errr = true
		} else {
			return false
		}
		count = 0
	}
	return errr
}

func createNewArr1(size int) [][]rune {
	array := make([][]rune, size)
	for i := range array {
		array[i] = make([]rune, size)
		for j := range array[i] {
			array[i][j] = '.'
		}
	}
	return array
}

func canPlace(board [][]rune, tetromino [][]int, x, y int) bool {
	for _, coord := range tetromino {
		newX, newY := x+coord[0], y+coord[1]

		if newX < 0 || newX >= len(board) || newY < 0 || newY >= len(board[0]) {
			return false
		}

		if board[newX][newY] != '.' {
			return false
		}
	}

	return true
}

func placeTetromino(board [][]rune, tetromino [][]int, x, y, index int) [][]rune {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	newBoard := make([][]rune, len(board))
	for i := range newBoard {
		newBoard[i] = make([]rune, len(board[i]))
		copy(newBoard[i], board[i])
	}

	for _, coord := range tetromino {
		newX, newY := x+coord[0], y+coord[1]
		newBoard[newX][newY] = rune(alphabet[index])
	}
	return newBoard
}

func placeTetrominoes(board [][]rune, tetrominoes [][][]int, index int) ([][]rune, error) {
	if index == len(tetrominoes) {
		return board, nil
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if canPlace(board, tetrominoes[index], i, j) {
				newBoard := placeTetromino(board, tetrominoes[index], i, j, index)

				result, err := placeTetrominoes(newBoard, tetrominoes, index+1)
				if err == nil {
					return result, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Cannot place tetromino %d", index)
}

func solvePuzzle(filename string) {
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	aa := massivkaSalu(content)
	if !checkNeighborhood(aa) {
		fmt.Println("ERROR")
		return
	}
	bb := massivKordinat(aa)
	transformedArray := transformArray(bb)
	razmer := int(math.Ceil(math.Sqrt(float64(len(transformedArray) * 4))))
	newArr := createNewArr1(razmer)

	result, err := placeTetrominoes(newArr, transformedArray, 0)
	if err != nil {
		razmer++
		newArr = createNewArr1(razmer)
		result, err = placeTetrominoes(newArr, transformedArray, 0)
	}

	for _, row := range result {
		fmt.Println(string(row))
	}
}

func main() {
	filename := os.Args[1]
	solvePuzzle(filename)
}
