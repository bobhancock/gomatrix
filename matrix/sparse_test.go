// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGetMatrix_Sparse2(t *testing.T) {
    n := 2
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        A.Set(i,i,2)
    }
	B := A.GetMatrix(1, 1, 1, 1)
	//B := A.GetMatrix(0, 0, 1, 1)
	C := Zeros(1, 1).SparseMatrix()
	for i := 0; i < 1; i++ {
        C.Set(i,i,2)
    }
	if !Equals(B, C) {
		t.Fail()
	}
}

func TestPlusSparse(t *testing.T) {
    n := 2
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        A.Set(i,i,2)
    }
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        B.Set(i,i,2)
    }
	C := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        C.Set(i,i,4)
    }
	C2, _ := A.PlusSparse(B)
	if !Equals(C2, C) {
		t.Fail()
	}
}
func TestAddSparse(t *testing.T) {
    n := 2
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        A.Set(i,i,2)
    }
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        B.Set(i,i,2)
    }
	A.AddSparse(B)
}
func TestElementMult_Sparse2(t *testing.T) {
    n := 4
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        A.Set(i,i,2)
    }
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
        B.Set(i,i,2)
    }
	C1, _ := A.ElementMult(B)
	C2, _ := A.ElementMultSparse(B)
	D, _ := A.DenseMatrix().ElementMult(B)
	if !Equals(D, C1) {
		t.Fail()
	}
	if !Equals(D, C2) {
		t.Fail()
	}
}

/*
func TestMulNaive(t *testing.T) {
    // force out of memory
	//n := 2000
	//n := 8000
	n := 2 
	//n := 3
	//n := 4
	//n := 200
    //fmt.Printf("init1\n")
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(i, i, 1)
		B.Set(i, i, 1)
	}

    //fmt.Printf("A: %v B: %v\n",A.cols,B.cols)
    //fmt.Printf("A: \n%v \nB: \n%v\n",A,B)
	D := MulNaive(A, B)
    //fmt.Printf("Naive: completed. D.width: %v\n",D.cols);
    //fmt.Printf("Naive: completed. D: \n%v\n",D);

}
*/


func TestMulStrassenOnly(t *testing.T) {
    // force out of memory
	//n := 2000
	//n := 8000
	n := 2 
	//n := 1
	//n := 3
	//n := 4
	//n := 200
    //fmt.Printf("init1\n")
/*
	A := ZerosSparse(n, n)
	for i := 0; i < 36; i++ {
		x := rand.Intn(6)
		y := rand.Intn(6)
		A.Set(y, x, 1)
	}
    //fmt.Printf("init2\n")
	B := ZerosSparse(n, n)
	for i := 0; i < 36; i++ {
		x := rand.Intn(6)
		y := rand.Intn(6)
		B.Set(y, x, 1)
	}
*/
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(i, i, 2)
		B.Set(i, i, 2)
	}

	E := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		E.Set(i, i, 4)
	}

	D := MulStrassen(A, B)
	if !Equals(D, E) {
		t.Fail()
	}
}

func TestAdd_Sparse(t *testing.T) {
	A := NormalsSparse(3, 3, 9)
	B := NormalsSparse(3, 3, 9)
	C1, _ := A.Plus(B)
	C2, _ := A.PlusSparse(B)
	if !ApproxEquals(C1, Sum(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Sum(A, B), ε) {
		t.Fail()
	}
}

func TestSubtract_Sparse(t *testing.T) {
	A := NormalsSparse(3, 3, 9)
	B := NormalsSparse(3, 3, 9)
	C1, _ := A.Minus(B)
	C2, _ := A.MinusSparse(B)
	if !ApproxEquals(C1, Difference(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Difference(A, B), ε) {
		t.Fail()
	}
}

func TestTimes_Sparse(t *testing.T) {
	A := Normals(3, 3).SparseMatrix()
	B := Normals(3, 3).SparseMatrix()
	C1, _ := A.Times(B)
	C2, _ := A.TimesSparse(B)
	if !ApproxEquals(C1, Product(A, B), ε) {
		t.Fail()
	}
	if !ApproxEquals(C2, Product(A, B), ε) {
		t.Fail()
	}
}

func TestElementMult_Sparse(t *testing.T) {
	A := Normals(3, 3).SparseMatrix()
	B := Normals(3, 3).SparseMatrix()
	C1, _ := A.ElementMult(B)
	C2, _ := A.ElementMultSparse(B)
	D, _ := A.DenseMatrix().ElementMult(B)
	if !Equals(D, C1) {
		t.Fail()
	}
	if !Equals(D, C2) {
		t.Fail()
	}
}

func TestGetMatrix_Sparse(t *testing.T) {
	A := ZerosSparse(6, 6)
	for i := 0; i < 36; i++ {
		x := rand.Intn(6)
		y := rand.Intn(6)
		A.Set(y, x, 1)
	}
	B := A.GetMatrix(1, 1, 4, 4)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if B.Get(i, j) != A.Get(i+1, j+1) {
				t.Fail()
			}
		}
	}

}

func TestAugment_Sparse(t *testing.T) {
	var A, B, C *SparseMatrix
	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 4, 16)
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

	A = NormalsSparse(2, 2, 4)
	B = NormalsSparse(4, 4, 16)
	C, err := A.Augment(B)
	if err == nil {
		t.Fail()
	}

	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 2, 8)
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

func TestStack_Sparse(t *testing.T) {
	var A, B, C *SparseMatrix
	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(4, 4, 16)
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

	A = NormalsSparse(2, 2, 4)
	B = NormalsSparse(4, 4, 16)
	C, err := A.Stack(B)
	if err == nil {
		if verbose {
			fmt.Printf("%v\n", err)
		}
		t.Fail()
	}

	A = NormalsSparse(4, 4, 16)
	B = NormalsSparse(2, 4, 8)
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
}
