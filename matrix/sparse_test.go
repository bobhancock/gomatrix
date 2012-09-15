// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	//"os"
	//"log"
	//"flag"
	"fmt"
	"math/rand"
	"testing"
    //"runtime/pprof"
)
//import pprof "net/http/pprof"

func TestMulStrassenRandom(t *testing.T) {
	n :=  16
		A := ZerosSparse(n, n)
		for i := 0; i < 36; i++ {
			x := rand.Intn(6)
			y := rand.Intn(6)
			v := rand.Float64()
			A.Set(y, x, v)
		}
		B := ZerosSparse(n, n)
		for i := 0; i < 36; i++ {
			x := rand.Intn(6)
			y := rand.Intn(6)
			v := rand.Float64()
			B.Set(y, x, v)
		}

	D := MulStrassen(A, B)
	E := MulSimple(A, B)
	//if !Equals(D, E) {
	if !ApproxEquals(D, E, ε) {
        fmt.Printf("D: \n", D)
        fmt.Printf("E: \n", E)
		t.Fail()
	}
}


func TestMulSimple(t *testing.T) {
	n := 8
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(0, i, 2)
		A.Set(i, i, 2)
		B.Set(i, i, 2)
	}

	E := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		E.Set(0, i, 4)
		E.Set(i, i, 4)
	}
    /*
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memprofile = flag.String("memprofile", "", "write memory profile to this file")
	flag.Parse()
	//Setup  profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		//return
	}
    */
	D := MulSimple(A, B)
	if !Equals(D, E) {
		t.Fail()
	}
}

func BenchmarkMulSimple(b *testing.B) {
	n := 8
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(0, i, 2)
		A.Set(i, i, 2)
		B.Set(i, i, 2)
	}

	E := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		E.Set(0, i, 4)
		E.Set(i, i, 4)
	}
    b.StartTimer()
	D := MulSimple(A, B)
    b.StopTimer()
	if !Equals(D, E) {
		b.Fail()
	}
}

func TestMulStrassenBigger2(t *testing.T) {
	n := 8
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(0, i, 2)
		A.Set(i, i, 2)
		B.Set(i, i, 2)
	}

	E := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		E.Set(0, i, 4)
		E.Set(i, i, 4)
	}

	////fmt.Printf("strassen bigger: A: %v\n", A)
	////fmt.Printf("strassen bigger: B: %v\n", B)
	D := MulStrassen(A, B)
	//fmt.Printf("strassen bigger: %v\n", D)
	if !Equals(D, E) {
		t.Fail()
	}
}

func TestGetMatrix_Sparse2(t *testing.T) {
	n := 2
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		A.Set(i, i, 2)
	}
	B := A.GetMatrix(1, 1, 1, 1)
	//B := A.GetMatrix(0, 0, 1, 1)
	C := Zeros(1, 1).SparseMatrix()
	for i := 0; i < 1; i++ {
		C.Set(i, i, 2)
	}
	if !Equals(B, C) {
		t.Fail()
	}
}

func TestPlusSparse2(t *testing.T) {
	n := 4
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		A.Set(i, i, 2)
	}
	A = A.GetMatrix(2, 2, 2, 2)
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		B.Set(i, i, 2)
	}
	B = B.GetMatrix(0, 0, 2, 2)
	n = 2
	C := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		C.Set(i, i, 4)
	}
	////fmt.Printf("A: %v\n",A)
	////fmt.Printf("B: %v\n",B)
	C2 := A.PlusSparseQuiet(B)
	////fmt.Printf("C2: %v\n",C2)

	if !Equals(C2, C) {
		t.Fail()
	}
}

func TestPlusSparse(t *testing.T) {
	n := 2
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		A.Set(i, i, 2)
	}
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		B.Set(i, i, 2)
	}
	C := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		C.Set(i, i, 4)
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
		A.Set(i, i, 2)
	}
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		B.Set(i, i, 2)
	}
	A.AddSparse(B)
}

func TestElementMult_Sparse2(t *testing.T) {
	n := 8
	A := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		A.Set(i, i, 2)
	}
	B := Zeros(n, n).SparseMatrix()
	for i := 0; i < n; i++ {
		B.Set(i, i, 2)
	}
	C1, _ := A.ElementMult(B)
	C2, _ := A.ElementMultSparse(B)
	D, _ := A.DenseMatrix().ElementMult(B)
	////fmt.Printf("simple: %v\n", C2)
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
    ////fmt.Printf("init1\n")
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(i, i, 1)
		B.Set(i, i, 1)
	}

    ////fmt.Printf("A: %v B: %v\n",A.cols,B.cols)
    ////fmt.Printf("A: \n%v \nB: \n%v\n",A,B)
	D := MulNaive(A, B)
    ////fmt.Printf("Naive: completed. D.width: %v\n",D.cols);
    //fmt.Printf("Naive: completed. D: \n%v\n",D);

}
*/

func TestMulStrassenOnly(t *testing.T) {
	// force out of memory
	//n := 2000
	//n := 8000
	//n := 2 
	//n := 1
	//n := 3
	//n := 4
	//n := 200
	//n := 16
	n := 8
	////fmt.Printf("init1\n")
	/*
		A := ZerosSparse(n, n)
		for i := 0; i < 36; i++ {
			x := rand.Intn(6)
			y := rand.Intn(6)
			A.Set(y, x, 1)
		}
	    ////fmt.Printf("init2\n")
		B := ZerosSparse(n, n)
		for i := 0; i < 36; i++ {
			x := rand.Intn(6)
			y := rand.Intn(6)
			B.Set(y, x, 1)
		}
	*/
	/**/
	A := ZerosSparse(n, n)
	B := ZerosSparse(n, n)
	for i := 0; i < n; i++ {
		A.Set(0, i, 2)
		A.Set(i, i, 2)
		B.Set(i, i, 2)
	}
	A1 := Zeros(n, n)
	B1 := Zeros(n, n)
	for i := 0; i < n; i++ {
		A1.Set(0, i, 2)
		A1.Set(i, i, 2)
		B1.Set(i, i, 2)
	}
	/**/

	/*
		E := ZerosSparse(n, n)
		for i := 0; i < n; i++ {
			E.Set(i, i, 4)
		}
	*/
	//fmt.Printf("A: %v\n", A)
	//fmt.Printf("B: %v\n", B)

	D := MulStrassen(A, B)
	//E, _ := A.ElementMultSparse(B)
	//E, _ := A.DenseMatrix().ElementMult(B)
	// element mult is not matrix multiplication!!!
	//E, _ := A1.ElementMult(B1)
	E := MulStrassen(A, B)
	////fmt.Printf("D: %v\n", D)
	////fmt.Printf("E: %v\n", E)
	if !Equals(D, E) {
		t.Fail()
	}
}

func TestMulStrassenBigger(t *testing.T) {
	n := 4
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

	////fmt.Printf("strassen bigger: A: %v\n", A)
	////fmt.Printf("strassen bigger: B: %v\n", B)
	D := MulStrassen(A, B)
	//fmt.Printf("strassen bigger: %v\n", D)
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
