package gomath

import (
	"errors"
	"fmt"
	"math"
)

// CONCEPT
//
// matrix example:
//
// matrix = [
//
//        column 0, ..., column n
//
//		[element 0, ..., element n], row 0
//		[element 0, ..., element n], row 1
//			...     ...    ...       row ...
//		[element 0, ..., element n], row n
//
// ]
//

var (
	// ErrMatrixMismatch is an error
	ErrMatrixMismatch = errors.New("two matrices can only me multiplied if the number of columns of the first is equal to the number of rows of the second")
	// ErrMatrixNotSquare is an error
	ErrMatrixNotSquare = errors.New("matrix must be square")
)

// Matrix is a matrix
type Matrix [][]float64 // [row][column]

// Size contains information about the number of rows and columns of a matrix
type Size struct {
	Rows, Columns int
}

// NewMatrix returns a new matrix with all entries set to 0
func NewMatrix(rows, columns int) Matrix {
	return NewFilledMatrix(rows, columns, 0)
}

// NewSquareMatrix returns a new square matrix
func NewSquareMatrix(order int) Matrix {
	return NewMatrix(order, order)
}

// NewFilledSquareMatrix returns a new square matrix with all elements set to a given element
func NewFilledSquareMatrix(order int, element float64) Matrix {
	return NewFilledMatrix(order, order, element)
}

// NewFilledMatrix returns a new Matrix with all elements set to a given element
func NewFilledMatrix(rows, columns int, element float64) Matrix {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = UnitSlice(element, columns)
	}
	return matrix
}

// NewIdentityMatrix returns a new identity matrix
func NewIdentityMatrix(order int) Matrix {
	identityMatrix := NewMatrix(order, order)
	for i := 0; i < order; i++ {
		identityMatrix = identityMatrix.Insert(i, i, 1)
	}
	return identityMatrix
}

// NewVector returns a new nx1 matrix with all entries set to 0 where n is the specified number of entries
func NewVector(numEntries int) Matrix {
	return NewMatrix(numEntries, 1)
}

// NewFilledVector returns a new nx1 matrix filled with a given slice of entries where n is the length of the slice
func NewFilledVector(entries ...float64) Matrix {
	vector := make([][]float64, len(entries))
	for i := range vector {
		vector[i] = []float64{entries[i]}
	}
	return vector
}

// Copy returns a copy of the matrix
func (matrix Matrix) Copy() Matrix {
	copiedMatrix := make([][]float64, matrix.Size().Rows)
	for row := range matrix {
		copiedMatrix[row] = make([]float64, len(matrix[row]))
		copy(copiedMatrix[row], matrix[row])
	}
	return copiedMatrix
}

// IsSquare returns whether the matrix is a square matrix
func (matrix Matrix) IsSquare() bool {
	return matrix.Size().Rows == matrix.Size().Columns
}

// Size returns a size object containing the number of rows and columns of the matrix
func (matrix Matrix) Size() *Size {
	h := len(matrix)
	if h == 0 {
		return &Size{Rows: 0, Columns: 0}
	}
	return &Size{Rows: h, Columns: len(matrix[0])}
}

// Insert inserts a value into the matrix and returns the result
func (matrix Matrix) Insert(row, column int, value float64) Matrix {
	filledMatrix := matrix.Copy()

	size := matrix.Size()

	if row > size.Rows-1 {
		for i := len(filledMatrix); i <= row; i++ {
			filledMatrix = append(filledMatrix, UnitSlice(0, size.Rows))
		}
	}

	if column > size.Columns-1 {
		for ind, row := range filledMatrix {
			for i := len(row); i <= column; i++ {
				filledMatrix[ind] = append(filledMatrix[ind], 0)
			}
		}
	}

	filledMatrix[row][column] = value

	return filledMatrix
}

// InsertRows inserts multiply rows into the matrix and returns the result
func (matrix Matrix) InsertRows(rows ...[]float64) Matrix {
	filledMatrix := matrix.Copy()

	for ind, row := range rows {
		filledMatrix = filledMatrix.InsertRow(ind-1, row)
	}

	return filledMatrix
}

// InsertRow inserts a row into the matrix and returns the result
func (matrix Matrix) InsertRow(afterRow int, row []float64) Matrix {
	filledMatrix := matrix.Copy()

	for i, it := range row {
		filledMatrix = filledMatrix.Insert(afterRow+1, i, it)
	}

	return filledMatrix
}

// ValueAt returns the i'th value of the matrix
func (matrix Matrix) ValueAt(i int) float64 {
	return matrix[i/matrix.Size().Columns][i%matrix.Size().Columns]
}

// SetValueAt sets the i'th value of the matrix to value
func (matrix Matrix) SetValueAt(i int, value float64) {
	matrix[i/matrix.Size().Columns][i%matrix.Size().Columns] = value
}

// DivideRows returns the two submatrices obtained by splitting the matrix after a given row
// DivideRows(i) will produce two submatrices containing the rows 0 to i and i+1 to len(matrix) respectively
func (matrix Matrix) DivideRows(afterRow int) (Matrix, Matrix) {
	return matrix[:afterRow+1], matrix[afterRow+1:]
}

// DivideColumns returns the two submatrices obtained by splitting the matrix after a given column
// DivideColumns(i) will produce two submatrices containing the columns 0 to i and i+1 to len(matrix) respectively
func (matrix Matrix) DivideColumns(afterColumn int) (Matrix, Matrix) {
	left, right := make([][]float64, matrix.Size().Rows), make([][]float64, matrix.Size().Rows)

	for row := range matrix {
		left[row], right[row] = make([]float64, afterColumn), make([]float64, matrix.Size().Columns-1-afterColumn)
		left[row] = matrix[row][:afterColumn+1]
		right[row] = matrix[row][afterColumn+1:]
	}

	return left, right
}

// Multiply returns the matrix multiplied by another matrix
func (matrix Matrix) Multiply(other Matrix) (Matrix, error) {
	matrixSize := matrix.Size()
	otherSize := other.Size()

	if matrixSize.Columns != otherSize.Rows {
		return nil, ErrMatrixMismatch
	}

	matrixProduct := NewMatrix(matrixSize.Rows, otherSize.Columns)
	for i, row := range matrixProduct {
		for j := range row {
			for k := 0; k < matrixSize.Columns; k++ {
				matrixProduct[i][j] += matrix[i][k] * other[k][j]
			}
		}
	}
	return matrixProduct, nil
}

// MultiplyByValue returns the matrix multiplied by a scalar value
func (matrix Matrix) MultiplyByValue(value float64) Matrix {
	multipliedMatrix := matrix.Copy()

	for i, row := range matrix {
		for j, element := range row {
			multipliedMatrix[i][j] = value * element
		}
	}

	return multipliedMatrix
}

// Reflect returns the matrix reflected about its diagonal
func (matrix Matrix) Reflect() (Matrix, error) {
	if !matrix.IsSquare() {
		return nil, ErrMatrixNotSquare
	}

	reflectedMatrix := matrix.Copy()

	for row := 0; row < matrix.Size().Rows; row++ {
		for column := 0; column < matrix.Size().Columns; column++ {
			reflectedMatrix[row][column] = matrix[column][row]
		}
	}
	return reflectedMatrix, nil
}

// Len returns the number of elements in the matrix
func (matrix Matrix) Len() int {
	return matrix.Size().Rows * matrix.Size().Columns
}

// Submatrix removes a given set of rows and columns from the matrix and returns the result
func (matrix Matrix) Submatrix(rows, columns []int) Matrix {
	subMatrix := NewMatrix(matrix.Size().Rows-len(rows), matrix.Size().Columns-len(columns))
	rowsSkipped := 0

	for row := range matrix {
		if Contains(rows, row) {
			rowsSkipped++
			continue
		}
		columnsSkipped := 0
		for column := range matrix[row] {
			if Contains(columns, column) {
				columnsSkipped++
				continue
			}
			subMatrix[row-rowsSkipped][column-columnsSkipped] = matrix[row][column]
		}
	}
	return subMatrix
}

// ApplySigns applies the signs needed for finding the adjoint of the matrix to the matrix and returns the result
func (matrix Matrix) ApplySigns() Matrix {
	signedMatrix := matrix.Copy()
	for row := range signedMatrix {
		for column := range signedMatrix[0] {
			if row%2 == 0 {
				if column%2 == 1 {
					signedMatrix[row][column] = -signedMatrix[row][column]
				}
			} else {
				if column%2 == 0 {
					signedMatrix[row][column] = -signedMatrix[row][column]
				}
			}
		}
	}
	return signedMatrix
}

// Adj returns the adjoint of the matrix
func (matrix Matrix) Adj() (Matrix, error) {
	if !matrix.IsSquare() {
		return nil, ErrMatrixNotSquare
	}

	adjointMatrix := matrix.Copy()
	for row := 0; row < len(adjointMatrix); row++ {
		for column := 0; column < len(adjointMatrix[0]); column++ {
			adjointMatrix[row][column], _ = matrix.Submatrix([]int{row}, []int{column}).Det()
		}
	}
	return adjointMatrix.ApplySigns().Reflect()
}

// Inv returns the inverse of the matrix
func (matrix Matrix) Inv() (Matrix, error) {
	if !matrix.IsSquare() {
		return nil, ErrMatrixNotSquare
	}
	adj, _ := matrix.Adj()
	det, _ := matrix.Det()
	return adj.MultiplyByValue(1 / det), nil
}

// Det returns the determinant of the matrix
func (matrix Matrix) Det() (float64, error) {
	if !matrix.IsSquare() {
		return -1, nil
	}

	sum, s := 0.0, 0.0
	rows := matrix.Size().Rows

	if rows == 1 {
		return matrix[0][0], nil
	}

	for i := range matrix {
		subMatrix := NewMatrix(rows-1, rows-1)
		for j := 1; j < rows; j++ {
			for k := 0; k < rows; k++ {
				if k < i {
					subMatrix[j-1][k] = matrix[j][k]
				} else if k > i {
					subMatrix[j-1][k-1] = matrix[j][k]
				}
			}
		}
		if i%2 == 0 {
			s = 1
		} else {
			s = -1
		}
		det, err := subMatrix.Det()
		if err != nil {
			return -1, err
		}
		sum += s * matrix[0][i] * det
	}
	return sum, nil
}

// Augment appends the columns of another matrix to the matrix and returns the result
func (matrix Matrix) Augment(other Matrix) Matrix {
	augmentedMatrix := matrix.Copy()

	for row := range matrix {
		augmentedMatrix[row] = append(augmentedMatrix[row], other[row]...)
	}

	return augmentedMatrix
}

// Echelon returns the row reduced echelon form of the matrix. This is the form obtained by operating Gaussian
// elimination on the columns of the matrix. Therefore, Echelon can be used to find solutions to systems of linear equations
func (matrix Matrix) Echelon() Matrix {
	echelonForm := matrix.Copy()

	columnCount := matrix.Size().Columns
	numElements := matrix.Len()

	lead := 0
	for i := 0; i < numElements; i += columnCount {
		if lead >= columnCount {
			break
		}

		j := i

		for val := math.Abs(echelonForm.ValueAt(j + lead)); val >= 0 && val < 0.0000001; {
			j += columnCount
			if j == numElements {
				j = i
				lead++
				if lead == columnCount {
					return echelonForm
				}
			}
		}

		for c, ix, rx := 0, j, i; c < columnCount; c++ {
			echelonForm[ix/columnCount][ix%columnCount], echelonForm[rx/columnCount][rx%columnCount] = echelonForm[rx/columnCount][rx%columnCount], echelonForm[ix/columnCount][ix%columnCount]
			ix++
			rx++
		}

		if d := echelonForm.ValueAt(i + lead); d != 0 {
			d := 1 / d
			for c, rx := 0, i; c < columnCount; c++ {
				echelonForm.SetValueAt(rx, echelonForm.ValueAt(rx)*d)
				rx++
			}
		}

		for j = 0; j < numElements; j += columnCount {
			if j != i {
				f := echelonForm.ValueAt(j + lead)
				for c, ix, rx := 0, j, i; c < columnCount; c++ {
					echelonForm.SetValueAt(ix, echelonForm.ValueAt(ix)-echelonForm.ValueAt(rx)*f)
					ix++
					rx++
				}
			}
		}
		lead++
	}

	return echelonForm
}

//
// Debugging
//

// String returns a nicely formatted string representation of the matrix
func (matrix Matrix) String() string {
	output := "[\n"

	for _, row := range matrix {
		output += " ["

		for ind, column := range row {
			if ind == len(row)-1 {
				output += fmt.Sprintf("%f", column)
			} else {
				output += fmt.Sprintf("%f, ", column)
			}
		}

		output += "]\n"
	}

	return output + "]"
}
