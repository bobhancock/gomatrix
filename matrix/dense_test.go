// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const ε = 0.000001
const verbose = false
const speedTest = true

/* TEST: arithmetic.go */

func TestEquals(t *testing.T) {
	if !Equals(Ones(5, 3), Ones(5, 3)) {
		t.Fail()
	}
	if Equals(Ones(3, 5), Ones(5, 3)) {
		t.Fail()
	}
	if Equals(Zeros(3, 3), Ones(3, 3)) {
		t.Fail()
	}
}

func TestApproximates(t *testing.T) {
	A := Numbers(3, 3, 6)
	B := Numbers(3, 3, .1)
	C := Numbers(3, 3, .6)
	D, err := A.ElementMult(B)
	if !(err == nil) && !ApproxEquals(D, C, ε) {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	A := Normals(3, 3)
	B := Normals(3, 3)
	C := Sum(A, B)
	if C.Nil() {
		t.Fail()
	}
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < C.Cols(); j++ {
			if A.Get(i, j)+B.Get(i, j) != C.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestSubtract(t *testing.T) {
	A := Normals(3, 3)
	B := Normals(3, 3)
	C := Difference(A, B)
	if C.Nil() {
		t.Fail()
	}
	for i := 0; i < C.Rows(); i++ {
		for j := 0; j < C.Cols(); j++ {
			if A.Get(i, j)-B.Get(i, j) != C.Get(i, j) {
				t.Fail()
			}
		}
	}
}

// Pow raises every element of the matrix to power.  Returns a new
// matrix
//func (A *DenseMatrix) Pow(power float64) *DenseMatrix {
func TestPowValid(t *testing.T) {
	rows := 3
	columns := 2
	mat := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, rows, columns)
	result := mat.Pow(10)
    //fmt.Printf("QQQ: testing mat: %v result: %v\n", mat, result);
	if result.Get(1,0) != 59049 {
		t.Errorf("FiltColMap expected 59049. received=%f", result.Get(1,0))
	}
}

//func (A DenseMatrix) FiltColMap(min, max float64, col int) (matches map[int]float64, err error) {
// FiltColMap find values that matches min <= A <= max for a specific column.
//
// Return Value
//
// matches - a map[int]float64 where the key is the row number in mat, 
// and the value is the value in the column specified by col.

func TestFiltColMapValid(t *testing.T) {
	rows := 3
	columns := 2
	mat := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, rows, columns)
	matches, err := mat.FiltColMap(100,101, 1)
    //fmt.Printf("QQQ: testing mat: %v err: %v matches: %v\n", mat, err, matches);
	if err != nil {
		t.Errorf("FiltColMap err=%v", err)
	}
	if len(matches) != 0 {
		t.Errorf("FiltColMap expected 0 matches. received=%f", len(matches))
	}
	matches2, err2 := mat.FiltColMap(1,4, 1)
    //fmt.Printf("QQQ: testing mat: %v err: %v matches: %v\n", mat, err2, matches2);
	if err2 != nil {
		t.Errorf("FiltColMap err=%v", err2)
	}
	if len(matches2) != 2 {
		t.Errorf("FiltColMap expected 2 matches. received=%f", len(matches))
	}
}

func TestFiltCol(t *testing.T) {
	mat := MakeDenseMatrix([]float64{2, 1, 4, 2, 6, 3, 8, 4, 10, 5}, 5, 2)
    //fmt.Printf("QQQ: ORIGINAL mat:\n %v\n", mat);
	// mat, max, min, column
	matches, err := mat.FiltCol(2.0, 4.0, 1)
    //fmt.Printf("QQQ: RESULT OF TEST matches:\n %v\n", matches);
	if err != nil {
		t.Errorf("FiltCol returned error: %v", err)
		return
	}
	r, _ := matches.GetSize()
	if r != 3 {
		t.Errorf("FiltCol: expected 3 rows and got %d", r)
	}

	m0 := matches.Get(0,1)
	if m0 != float64(2) {
		t.Errorf("FiltCol: expected row 0 col 1 to be 2, but got %f",m0)
	}

	m1 := matches.Get(1, 1)
	if m1 != 3 {
		t.Errorf("FiltCol: expected row 1 col 1 to be 3, but got %f",m1)
	}

	m2 := matches.Get(2, 1)
	if m2 !=4 {
		t.Errorf("FiltCol: expected row 1 col 1 to be 3, but got %f",m2)
	}
	matches, err = mat.FiltCol(100.0, 200.00, 1)
    //fmt.Printf("QQQ: testing mat: %v err: %v\n", mat, err);
	if err == nil {
		t.Errorf("FiltCol did not return err on no match condition.")
	}
}

func TestFiltCol2(t *testing.T) {
	mat := MakeDenseMatrix([]float64{2, 1, 4, 2, 6, 3, 8, 4, 10, 5, 1,1,1,1,1}, 5, 3)
    //fmt.Printf("QQQ: ORIGINAL mat:\n %v\n", mat);
	// mat, max, min, column
	matches, err := mat.FiltCol(2.0, 4.0, 1)
    //fmt.Printf("QQQ: RESULT OF TEST matches:\n %v\n", matches);
	if err != nil {
		t.Errorf("FiltCol returned error: %v", err)
		return
	}
	r, _ := matches.GetSize()
	if r != 1 {
		t.Errorf("FiltCol: expected 1 row and got %d", r)
	}

	m0 := matches.Get(0,1)
	if m0 != 4 {
		t.Errorf("FiltCol: expected row 0 col 1 to be 4, but got %f",m0)
	}

	matches, err = mat.FiltCol(100.0, 200.00, 1)
    //fmt.Printf("QQQ: testing mat: %v err: %v\n", mat, err);
	if err == nil {
		t.Errorf("FiltCol did not return err on no match condition.")
	}
}

func TestFiltCol3(t *testing.T) {
	mat := MakeDenseMatrix([]float64{2, 1, 4, 2, 6, 3, 8, 4, 10, 5, 1,1,1,1,1,1,2,3}, 6, 3)
    //fmt.Printf("QQQ: ORIGINAL mat:\n %v\n", mat);
	// mat, max, min, column
	matches, err := mat.FiltCol(2.0, 4.0, 1)
    //fmt.Printf("QQQ: RESULT OF TEST matches:\n %v\n", matches);
	if err != nil {
		t.Errorf("FiltCol returned error: %v", err)
		return
	}
	r, _ := matches.GetSize()
	if r != 2 {
		t.Errorf("FiltCol: expected 2 rows and got %d", r)
	}

	m0 := matches.Get(0,1)
	if m0 != 4 {
		t.Errorf("FiltCol: expected row 0 col 1 to be 4, but got %f",m0)
	}

	m1 := matches.Get(1,1)
	if m1 != 2 {
		t.Errorf("FiltCol: expected row 0 col 1 to be 2, but got %f",m0)
	}

	matches, err = mat.FiltCol(100.0, 200.00, 1)
    //fmt.Printf("QQQ: testing mat: %v err: %v\n", mat, err);
	if err == nil {
		t.Errorf("FiltCol did not return err on no match condition.")
	}
}


func TestAppendColInvalid(t *testing.T) {
	rows := 3
	columns := 2
	mat := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, rows, columns)
	col := []float64{1.1, 2.2, 3.3, 4.4}
	mat, err := mat.AppendCol(col)
    //fmt.Printf("QQQ: testing mat: %v err: %v\n", mat, err);
	if err == nil {
		t.Errorf("AppendCol err=%v", err)
	}
}

func TestAppendColValid(t *testing.T) {
	rows := 3
	columns := 2
	mat := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, rows, columns)
	col := []float64{1.1, 2.2, 3.3}
	mat, err := mat.AppendCol(col)
    //fmt.Printf("QQQ: testing mat: %v err: %v\n", mat, err);
	if err != nil {
		t.Errorf("AppendCol err=%v", err)
	}
}


func TestColSliceValid(t *testing.T) {
	rows := 3
	columns := 2
	mat := MakeDenseMatrix([]float64{1, 2, 3, 4, 5, 6}, rows, columns)
	c := mat.ColSlice(1)
    //fmt.Printf("QQQ: testing mat: %v c: %v\n", mat, c);
	if len(c) != rows {
		t.Errorf("Returned slice has len=%d instead of %d.", len(c), rows)
	}
}

func TestSumCol(t *testing.T) {
	m := MakeSumMatrix()
	b := m.SumCol(1)
	if b != 3 {
        t.Errorf("SumCol: Expected a sum of 3, but received %f", b)
	}
}

func TestSumCols(t *testing.T) {
	m := MakeSumMatrix()
	b := m.SumCols()
	//9,3, 1, 4
	if b.Get(0,0) != -9 || b.Get(0,1) != 3 || b.Get(0,2) != 1 || b.Get(0,3) != 4 {
		t.Errorf("SumCols: Excpect a row vector of [9, 3, 1 ,4] but received %v.", b)
	}
}

func TestMeanCols(t *testing.T) {
	m := MakeSumMatrix()
	b := m.MeanCols()
	if b.Get(0,0) != -2.25 || b.Get(0,1) != 0.75 || b.Get(0,2) != 0.25 || b.Get(0,3) != 1 {
		t.Errorf("MeanCols: Excpect a row vector of [-2.25, -0.75, -0.25], 1 but received %v.", b)
	}
} 

func TestSumRows(t *testing.T) {
	m := MakeSumMatrix()
	b := m.SumRows()
	//4, -5, 9, -9
	if b.Get(0,0) != 4 || b.Get(1,0) != -5 || b.Get(2,0) != 9 || b.Get(3,0) != -9 {
		t.Errorf("MeanCols: Excpect a row vector of [-4, -5, 9, -9] but received %v.", b)
	}
}

func TestProduct(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	B := MakeDenseMatrix([]float64{1, 7, -4, 4,
		3, -2, -6, 1,
		-12, 8, 1, 20,
		0, 0, -10, 3,
	},
		4, 4)

	C, err := A.Times(B)

	if !(err == nil) {
		t.Fail()
	}

	var Ctrue Matrix
	Ctrue = MakeDenseMatrix([]float64{48, 14, -56, -46,
		66, -21, -10, -108,
		-240, 68, 101, 356,
		114, -122, -56, -203,
	},
		4, 4)

	if !Equals(C, Ctrue) {
		t.Fail()
	}

	P := MakePivotMatrix([]int{1, 3, 0, 2}, -1)
	C, err = P.Times(A)

	Ctrue, err = P.DenseMatrix().Times(A)
	if !Equals(C, Ctrue) {
		t.Fail()
	}
}

func TestParallelProduct(t *testing.T) {

	w := 100000
	h := 40

	if !verbose {
		w = 100
		h = 4
	}

	rand.Seed(time.Now().UnixNano())
	A := Normals(h, w)
	B := Normals(w, h)

	var C *DenseMatrix
	var start, end int64

	start = time.Now().UnixNano()
	Ctrue, err := A.Times(B)
	if !(err == nil) {
		t.Fail()
	}
	end = time.Now().UnixNano()
	if verbose {
		fmt.Printf("%fs for synchronous\n", float64(end-start)/1000000000)
	}

	start = time.Now().UnixNano()
	C = ParallelProduct(A, B)
	if !(err == nil) {
		t.Fail()
	}
	end = time.Now().UnixNano()
	if verbose {
		fmt.Printf("%fs for parallel\n", float64(end-start)/1000000000)
	}

	if !Equals(C, Ctrue) {
		t.Fail()
	}
}

var MaxProcs int = 1

func TestTimesDenseProcs(t *testing.T) {
	A := Normals(10, 10)
	B := Normals(10, 10)

	old := MaxProcs
	MaxProcs = 1
	C, _ := A.TimesDense(B)
	MaxProcs = 2
	Cp, _ := A.TimesDense(B)
	if !Equals(C, Cp) {
		t.Fail()
	}
	MaxProcs = old
}

func TestElementMult(t *testing.T) {

	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	T := MakeDenseMatrix([]float64{0.1, 0.1, 0.1, 0.1,
		10, 10, 10, 10,
		100, 100, 100, 100,
		1000, 1000, 1000, 1000,
	},
		4, 4)
	C, err := A.ElementMult(T)

	if !(err == nil) {
		t.Fail()
	}

	Ctrue := MakeDenseMatrix([]float64{0.6, -0.2, -0.4, 0.4,
		30, -30, -60, 10,
		-1200, 800, 2100, -800,
		-6000, 0, -10000, 7000,
	},
		4, 4)

	if !ApproxEquals(C, Ctrue, ε) {
		t.Fail()
	}
}

func TestScale(t *testing.T) {
	A := Normals(3, 3)
	f := float64(5.3)
	B := A.Copy()
	B.Scale(f)

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j)*f != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestScaleMatrix(t *testing.T) {
	A := Normals(4, 4)
	B := Normals(4, 4)
	C := A.Copy()
	C.ScaleMatrix(B)

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j)*B.Get(i, j) != C.Get(i, j) {
				t.Fail()
			}
		}
	}
}

/* TEST: basic.go */

func TestSymmetric(t *testing.T) {
	A := MakeDenseMatrix([]float64{
		6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	if A.Symmetric() {
		t.Fail()
	}
	B := MakeDenseMatrix([]float64{
		6, 3, -12, -6,
		3, -3, 8, 0,
		-12, 8, 21, -10,
		-6, 0, -10, 7,
	},
		4, 4)
	if !B.Symmetric() {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	Ainv, err := A.Inverse()

	if !(err == nil) {
		t.Fail()
	}

	AAinv, err := A.Times(Ainv)

	if !(err == nil) {
		t.Fail()
	}

	if !ApproxEquals(Eye(A.Rows()), AAinv, ε) {
		if verbose {
			fmt.Printf("A\n%v\n\nAinv\n%v\n\nA*Ainv\n%v\n", A, Ainv, AAinv)
		}
		t.Fail()
	}
}

func TestDet(t *testing.T) {
	A := MakeDenseMatrix([]float64{4, -2, 5,
		-1, -7, 10,
		0, 1, -3,
	},
		3, 3)

	if A.Det() != 45 {
		if verbose {
			fmt.Printf("A\n%v\n\nA.Det()\n%v\n\n", A, A.Det())
		}
		t.Fail()
	}
}

func TestTrace(t *testing.T) {
	A := MakeDenseMatrix([]float64{4, -2, 5,
		-1, -7, 10,
		0, 1, -3,
	},
		3, 3)

	if A.Trace() != 4-7-3 {
		if verbose {
			fmt.Printf("A\n%v\n\nA.Trace()\n%v\n\n", A, A.Trace())
		}
		t.Fail()
	}
}

func TestTranspose(t *testing.T) {
	A := Normals(4, 4)
	B := A.Transpose()
	for i := 0; i < A.rows; i++ {
		for j := 0; j < A.cols; j++ {
			if A.Get(i, j) != B.Get(j, i) {
				t.Fail()
			}
		}
	}
}

func TestSolve(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	b := MakeDenseMatrix([]float64{1, 1, 1, 1}, 4, 1)
	x, err := A.Solve(b)

	if !(err == nil) {
		t.Fail()
	}

	xtrue := MakeDenseMatrix([]float64{-0.906250, -3.393750, 1.275000, 1.187500}, 4, 1)

	if !Equals(x, xtrue) {
		t.Fail()
	}
}

/* TEST: decomp.go */

func TestCholesky(t *testing.T) {
	A := MakeDenseMatrix([]float64{1, 0.2, 0,
		0.2, 1, 0.5,
		0, 0.5, 1,
	},
		3, 3)
	B, err := A.Cholesky()
	if !(err == nil) {
		t.Fail()
	}
	if !ApproxEquals(A, Product(B, B.Transpose()), ε) {
		t.Fail()
	}
}

func TestLU(t *testing.T) {

	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	L, U, P := A.LU()

	LU, err := L.Times(U)
	PLU, err := P.Times(LU)

	if !(err == nil) {
		if verbose {
			fmt.Printf("TestLU: %v\n", err)
		}
		t.Fail()
	}

	if !Equals(A, PLU) {
		if verbose {
			fmt.Printf("TestLU:\n%v\n!=\n%v\n", A, PLU)
		}
		t.Fail()
	}

	A = MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	Ltrue, Utrue, Ptrue := A.LU()

	P = A.LUInPlace()
	L = A.L()
	U = A.U()

	for i := 0; i < L.Rows(); i++ {
		L.Set(i, i, 1)
	}

	PL := Product(P, L)
	PLU2 := Product(PL, U)
	PLtrue := Product(Ptrue, Ltrue)
	PLUtrue := Product(PLtrue, Utrue)

	if !Equals(PLU2, PLUtrue) {
		t.Fail()
	}

}

func TestQR(t *testing.T) {
	A := MakeDenseMatrix([]float64{6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	Q, R := A.QR()

	Qtrue := MakeDenseMatrix([]float64{-0.4, 0.278610, 0.543792, -0.683130,
		-0.2, -0.358213, -0.699161, -0.585540,
		0.8, 0.437816, -0.126237, -0.390360,
		0.4, -0.776129, 0.446686, -0.195180,
	},
		4, 4)

	Rtrue := MakeDenseMatrix([]float64{-15, 7.8, 15.6, -5.4,
		0, 4.019950, 17.990272, -8.179206,
		0, 0, -5.098049, 5.612709,
		0, 0, 0, -1.561440,
	},
		4, 4)

	QR := Product(Q, R)

	if !ApproxEquals(Q, Qtrue, ε) ||
		!ApproxEquals(R, Rtrue, ε) ||
		!ApproxEquals(A, QR, ε) {
		t.Fail()
	}
}

/* TEST: eigen.go */

func TestEigen(t *testing.T) {
	A := MakeDenseMatrix([]float64{
		2, 1,
		1, 2,
	},
		2, 2)
	V, D, _ := A.Eigen()

	Vinv, _ := V.Inverse()
	Aguess := Product(Product(V, D), Vinv)

	if !ApproxEquals(A, Aguess, ε) {
		t.Fail()
	}

	B := MakeDenseMatrix([]float64{
		6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)

	V, D, _ = B.Eigen()

	Vinv, _ = V.Inverse()

	if !ApproxEquals(B, Product(Product(V, D), Vinv), ε) {
		if verbose {
			fmt.Printf("B =\n%v\nV=\n%v\nD=\n%v\n", B, V, D)
		}
		t.Fail()
	}

	Bm, _ := B.Times(B.Transpose())
	B = Bm.DenseMatrix()
	V, D, _ = B.Eigen()
	Vinv, _ = V.Inverse()

	if !ApproxEquals(B, Product(Product(V, D), Vinv), ε) {
		if verbose {
			fmt.Printf("B =\n%v\nV=\n%v\nD=\n%v\n", B, V, D)
		}
		t.Fail()
	}
}

func TestSVD(t *testing.T) {
	A := MakeDenseMatrix([]float64{
		6, -2, -4, 4,
		3, -3, -6, 1,
		-12, 8, 21, -8,
		-6, 0, -10, 7,
	},
		4, 4)
	U, Σ, V, _ := A.SVD()
	Arecomp := Product(Product(U, Σ), V.Transpose())
	if !ApproxEquals(A, Arecomp, ε) {
		t.Fail()
	}
	A = Normals(5, 3)
	U, Σ, V, _ = A.SVD()
	Arecomp = Product(Product(U, Σ), V.Transpose())
	if !ApproxEquals(A, Arecomp, ε) {
		t.Fail()
	}
}

/* TEST: matrix.go */

func TestGetMatrix(t *testing.T) {
	A := Zeros(4, 4)
	B := A.GetMatrix(1, 1, 2, 2)
	B.Set(0, 1, 1)
	if A.Get(1, 2) != 1 {
		t.Fail()
	}
}

func TestL(t *testing.T) {
	A := Normals(4, 4)
	L := A.L()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(4, 2)
	L = A.L()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(2, 4)
	L = A.L()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j > i && L.Get(i, j) != 0 {
				t.Fail()
			} else if j <= i && L.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestU(t *testing.T) {
	A := Normals(4, 4)
	U := A.U()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(2, 4)
	U = A.U()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	A = Normals(4, 2)
	U = A.U()
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if j < i && U.Get(i, j) != 0 {
				t.Fail()
			} else if j >= i && U.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestAugment(t *testing.T) {
	var A, B, C *DenseMatrix
	A = Normals(4, 4)
	B = Normals(4, 4)
	C, _ = A.Augment(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = Normals(2, 2)
	B = Normals(4, 4)
	C, err := A.Augment(B)
	if err == nil {
		t.Fail()
	}

	A = Normals(4, 4)
	B = Normals(4, 2)
	C, _ = A.Augment(B)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i, j+A.Cols()) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestStack(t *testing.T) {

	var A, B, C *DenseMatrix
	A = Normals(4, 4)
	B = Normals(4, 4)
	C, _ = A.Stack(B)

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}

	A = Normals(4, 4)
	B = Normals(4, 2)
	C, err := A.Stack(B)
	if err == nil {
		t.Fail()
	}

	A = Normals(2, 4)
	B = Normals(4, 4)
	C, err = A.Stack(B)

	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if C.Get(i, j) != A.Get(i, j) {
				t.Fail()
			}
		}
	}
	for i := 0; i < B.Rows(); i++ {
		for j := 0; j < B.Cols(); j++ {
			if C.Get(i+A.Rows(), j) != B.Get(i, j) {
				t.Fail()
			}
		}
	}
}

func TestZeros(t *testing.T) {
	A := Zeros(4, 5)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != 0 {
				t.Fail()
			}
		}
	}
}

func TestNumbers(t *testing.T) {
	n := float64(1.0)
	A := Numbers(3, 3, n)
	//	fmt.Printf("%v\n\n\n",A.String());

	Atrue := MakeDenseMatrix([]float64{n, n, n,
		n, n, n,
		n, n, n,
	},
		3, 3)
	if !Equals(A, Atrue) {
		t.Fail()
	}
}

func TestOnes(t *testing.T) {

	A := Ones(4, 5)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if A.Get(i, j) != 1 {
				t.Fail()
			}
		}
	}
}

func TestEye(t *testing.T) {

	A := Eye(4)
	for i := 0; i < A.Rows(); i++ {
		for j := 0; j < A.Cols(); j++ {
			if (i != j && A.Get(i, j) != 0) || (i == j && A.Get(i, j) != 1) {
				t.Fail()
			}
		}
	}
}

func TestNormals(t *testing.T) {
	//test that it's filled with random data?
	A := Normals(3, 4)
	if A.Rows() != 3 || A.Cols() != 4 {
		t.Fail()
	}
}

func TestKronecker(t *testing.T) {
	A := MakeDenseMatrix([]float64{0, 1, 2, 3}, 2, 2)
	B := MakeDenseMatrix([]float64{5, 6, 7, 8, 9, 10}, 2, 3)
	C := Kronecker(A, B)
	Cp := MakeDenseMatrix([]float64{0, 0, 0, 5, 6, 7,
		0, 0, 0, 8, 9, 10,
		10, 12, 14, 15, 18, 21,
		16, 18, 20, 24, 27, 30}, 4, 6)
	if !Equals(C, Cp) {
		t.Fail()
	}
}

func TestVectorize(t *testing.T) {
	A := MakeDenseMatrix([]float64{0, 1, 2, 3, 4, 5}, 2, 3)
	V := Vectorize(A)
	Vp := MakeDenseMatrix([]float64{0, 3, 1, 4, 2, 5}, 6, 1)
	if !Equals(V, Vp) {
		t.Fail()
	}
}

func TestSubmatrix(t *testing.T) {
	Eye(3).GetMatrix(1, 1, 2, 2).GetColVector(0)
}

/* TEST: util.go */

/*
func TestMultipleProduct(t *testing.T) {
	A := Ones(3, 1)
	B := Ones(1, 3)
	C := MultipleProduct(A, B, A)
	D := Product(A, B)
	E := Product(D, A)

	if !Equals(E, C) {
		t.Fail()
	}
}
*/

func MakeSumMatrix() *DenseMatrix {
	m := MakeDenseMatrix([]float64{6, -2, -4, 4,
	                               3, -3, -6, 1,
		                          -12, 8, 21, -8,
		                           -6, 0, -10, 7},
		4, 4)
	
	return m
}
	
func TestSetRowVector(t *testing.T) {
	src := MakeDenseMatrix([]float64{0,1,2,3}, 1, 4)
	tgt := MakeDenseMatrix([]float64{99, 100, 101, 102,
	                                    250, 251, 252, 253},
		2,4)
	tgt.SetRowVector(src, 1)
	if tgt.Get(1,0) != 0 || tgt.Get(1,1) != 1 || tgt.Get(1,2) != 2 || tgt.Get(1,3) != 3 {
		t.Errorf("Expected [0, 1, 2, 3, 4] in row 1 but received %v", tgt)
	}
}