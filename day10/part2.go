package main

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/optimize/convex/lp"
)

type jmachine struct {
	buttons  [][]int64
	joltages []int64
}

func Part2(lines []string) int64 {
	jmachines := ParseJMachines(lines)
	// fmt.Println(jmachines[0])
	presses := int64(0)
	for _, j := range jmachines {
		presses += MinJPresses(j)
	}
	return presses
}

func MinJPresses(jmachine jmachine) int64 {
	var buttonV [][]float64
	for _, b := range jmachine.buttons {
		var buttonRow []float64
		for _, bv := range b {
			buttonRow = append(buttonRow, float64(bv))
		}
		buttonV = append(buttonV, buttonRow)
	}
	var target []float64
	for _, jv := range jmachine.joltages {
		target = append(target, float64(jv))
	}

	// vars, obj, err := SolveMinPositiveCoeffs(buttonV, target, 0)
	// if err != nil {
	// 	fmt.Println("Err: ", err)
	// }
	// fmt.Println(vars)
	// fmt.Println(obj)

	return 0
}

func ParseJMachines(lines []string) []jmachine {
	var machines []jmachine
	for _, l := range lines {
		machines = append(machines, ParseJMachine(l))
	}
	return machines
}

func ParseJMachine(line string) jmachine {
	joltagesre := regexp.MustCompile(`\{.*\}`)
	joltagesMatch := joltagesre.FindAllString(line, -1)[0]
	joltagesMatch = strings.ReplaceAll(joltagesMatch, "{", "")
	joltagesMatch = strings.ReplaceAll(joltagesMatch, "}", "")
	joltageStrings := strings.Split(joltagesMatch, ",")
	var joltages []int64
	for _, j := range joltageStrings {
		value, _ := strconv.ParseInt(j, 10, 64)
		joltages = append(joltages, value)
	}
	matrixLength := len(joltages)

	buttonre := regexp.MustCompile(`\(.*\)`)
	buttonMatch := buttonre.FindAllString(line, -1)[0]
	buttonStrings := strings.Split(buttonMatch, " ")
	var buttons [][]int64
	for _, b := range buttonStrings {
		buttons = append(buttons, ParseMatrixButton(b, matrixLength))
	}

	return jmachine{buttons: buttons, joltages: joltages}
}

func ParseMatrixButton(buttonString string, length int) []int64 {
	buttonFullString := strings.ReplaceAll(buttonString, "(", "")
	buttonFullString = strings.ReplaceAll(buttonFullString, ")", "")
	buttonStringValues := strings.Split(buttonFullString, ",")
	// fmt.Println("Button String Values: ", buttonStringValues)
	var buttonIndexes []int
	for _, bs := range buttonStringValues {
		bi, _ := strconv.Atoi(bs)
		buttonIndexes = append(buttonIndexes, bi)
	}

	var output []int64

	for i := range length {
		if slices.Contains(buttonIndexes, i) {
			output = append(output, int64(1))
			continue
		}
		output = append(output, int64(0))
	}
	return output
}

// only works for square matrixes
func SolveMinPositiveCoeffs(vectors [][]float64, target []float64, tol float64) ([]float64, float64, error) {
	n := len(target)
	if n == 0 {
		return nil, 0, errors.New("target must be non-empty")
	}
	m := len(vectors)
	if m == 0 {
		return nil, 0, errors.New("vectors must be non-empty")
	}
	// A must have at least as many columns (m) as rows (n) for lp.Simplex. [1](https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp)
	if m < n {
		return nil, 0, fmt.Errorf("not enough vectors: need m >= n (got m=%d, n=%d)", m, n)
	}

	// Build A (n x m) in row-major order: A[i, j] = vectors[j][i].
	data := make([]float64, n*m)
	for j := 0; j < m; j++ {
		if len(vectors[j]) != n {
			return nil, 0, fmt.Errorf("vector %d has length %d, expected %d", j, len(vectors[j]), n)
		}
		// Disallow all-zero columns to avoid Simplex errors. [1](https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp)
		colIsZero := true
		for i := 0; i < n; i++ {
			vij := vectors[j][i]
			data[i*m+j] = vij
			if vij != 0 {
				colIsZero = false
			}
		}
		if colIsZero {
			return nil, 0, fmt.Errorf("vector %d is all zeros (disallowed)", j)
		}
	}

	// Check all-zero rows (disallowed). [1](https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp)
	for i := 0; i < n; i++ {
		rowIsZero := true
		for j := 0; j < m; j++ {
			if data[i*m+j] != 0 {
				rowIsZero = false
				break
			}
		}
		if rowIsZero {
			return nil, 0, fmt.Errorf("row %d of A is all zeros (disallowed)", i)
		}
	}

	A := mat.NewDense(n, m, data) // A ∈ R^{n×m} [4](https://pkg.go.dev/gonum.org/v1/gonum/mat)

	// b := target
	b := make([]float64, n)
	copy(b, target)

	// c := ones(m) (minimize sum of coefficients).
	c := make([]float64, m)
	for j := range c {
		c[j] = 1.0
	}

	// Solve with Simplex (standard form): minimize c^T x s.t. A x = b, x >= 0. [1](https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp)
	optF, optX, err := lp.Simplex(c, A, b, tol, nil)
	if err != nil {
		// lp.Simplex can return ErrInfeasible or ErrUnbounded if the model is not solvable. [1](https://pkg.go.dev/gonum.org/v1/gonum/optimize/convex/lp)
		return nil, 0, err
	}
	return optX, optF, nil
}
