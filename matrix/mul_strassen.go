// Copyright 2012 Harry de Boer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "fmt"

// MulStrassen returns A * B.
//
// Original paper: Gaussian Elimination is not Optimal.
//                 Volker Strassen, 1969.
//
// This implementation is not optimized, it serves as a reference for testing.
func MulStrassen(A, B *SparseMatrix) *SparseMatrix {
	//return Zeros(A.height, B.width).MulStrassen(A, B)
    //fmt.Printf(">>>>>Strassen: create zeros matrix: %vx%v\n", A.rows, B.cols);
	Z := ZerosSparse(A.rows, B.cols)
    ////fmt.Printf(">>>>>Strassen: done zeros matrix: %vx%v\n", A.rows, B.cols);
    //return Z
    C := Z.MulStrassen(A, B)
    ////fmt.Printf("Strassen: return: C: %vx%v\n", C.rows, C.cols);
    return C
}

/*  */
// MulStrassen calculates C = A * B and returns C.
func (C *SparseMatrix) MulStrassen(A, B *SparseMatrix) *SparseMatrix {
    //A = A.Copy()
    //B = B.Copy()

    //fmt.Printf("\n\n\n==================STARTING C: %v %vx%v\n", C.rows, A.rows, B.rows);
    ////fmt.Printf(">>>>>222: Strassen: CALLING C: %vx%v\n", A.rows, B.cols);
	if A.cols < 2 {
        //fmt.Printf("DUMP Strassen: A:\n%v B: \n%v\n", A, B);
	    RETB := ZerosSparse(1, 1)
        //fmt.Printf("RETURNING Strassen: SET A, B: %v,%v\n", A.Get(0,0), B.Get(0,0));
		RETB.Set(0, 0,  A.Get(0,0)*B.Get(0,0))
        //fmt.Printf("RETURNING \n%v\n", RETB);
        return RETB
    }
    /*
	if A.cols < 80 || A.rows != A.cols || A.rows % 2 != 0 {
		return C.MulBLAS(A, B)
	}
    */

    //fmt.Printf(">>>>>000: Strassen: A: \n%v\n", A);

	m := A.rows / 2
	A11 := A.Copy().GetMatrix(0, 0, m, m)
    //fmt.Printf("AFTER>>>>>000: Strassen: A: \n%v\n", A);
    //fmt.Printf("AFTER>>>>>000: Strassen: A11: \n%v\n", A11);
	A12 := A.Copy().GetMatrix(0, m, m, m)
	A21 := A.Copy().GetMatrix(m, 0, m, m)
	A22 := A.Copy().GetMatrix(m, m, m, m)
	B11 := B.Copy().GetMatrix(0, 0, m, m)
	B12 := B.Copy().GetMatrix(0, m, m, m)
	B21 := B.Copy().GetMatrix(m, 0, m, m)
	B22 := B.Copy().GetMatrix(m, m, m, m)
	C11 := C.GetMatrix(0, 0, m, m)
	C12 := C.GetMatrix(0, m, m, m)
	C21 := C.GetMatrix(m, 0, m, m)
	C22 := C.GetMatrix(m, m, m, m)

    //fmt.Printf(">>>>>111: Strassen: C11: \n%v\n", C11);
    //fmt.Printf(">>>>>111: Strassen: C12: \n%v\n", C12);
    //fmt.Printf(">>>>>111: Strassen: C21: \n%v\n", C21);
    //fmt.Printf(">>>>>111: Strassen: C22: \n%v\n", C22);

    fmt.Printf(">>>>>XXX: Strassen: A11: \n%v\n", A11);
    fmt.Printf(">>>>>XXX: Strassen: A11: \n%v\n", A11.elements);
    fmt.Printf(">>>>>XXX: Strassen: A22: \n%v\n", A22);
    fmt.Printf(">>>>>XXX: Strassen: A22: \n%v\n", A22.elements);
    //fmt.Printf(">>>>>XXX: Strassen: B11: \n%v\n", B11);
    //fmt.Printf(">>>>>XXX: Strassen: B22: \n%v\n", B22);

	Z := A11.PlusSparseQuiet(A22)
	Z1 := B11.PlusSparseQuiet(B22)
    fmt.Printf(">>>>>ZZZ: Strassen: Z: \n%v\n", Z)
    fmt.Printf(">>>>>ZZZ: Strassen: Z1: \n%v\n", Z1)

	M1 := MulStrassen(Z, Z1)
    fmt.Printf("AFTER>>>>>000: Strassen: M1: \n%v\n", M1)
    //fmt.Printf("SHOULD NOT CHANGE AFTER>>>>>000: Strassen: A11: \n%v\n", A11);
    //fmt.Printf("SHOULD NOT CHANGE AFTER>>>>>000: Strassen: B11: \n%v\n", B11);
	M2 := MulStrassen(A21.PlusSparseQuiet(A22), B11)
    fmt.Printf("AFTER>>>>>000: Strassen: M2: \n%v\n", M2)
	M3 := MulStrassen(A11, B12.MinusSparseQuiet(B22))
    fmt.Printf("AFTER>>>>>000: Strassen: M3: \n%v\n", M3)

    //fmt.Printf("DEBUG1>>>>>000: Strassen: B21: \n%v\n", B21)
    //fmt.Printf("DEBUG1>>>>>000: Strassen: B21: \n%v\n", B21.elements)
    //fmt.Printf("DEBUG1>>>>>000: Strassen: B11: \n%v\n", B11)
    //fmt.Printf("DEBUG1>>>>>000: Strassen: B11: \n%v\n", B11.elements)
    //fmt.Printf("DEBUG1>>>>>000: Strassen: A22: \n%v\n", A22)
	Z2 := B21.MinusSparseQuiet(B11)
    //fmt.Printf("DEBUG1>>>>>000: Strassen: Z2: \n%v\n", Z2)
	M4 := MulStrassen(A22, Z2)
    fmt.Printf("AFTER>>>>>000: Strassen: M4: \n%v\n", M4)
	M5 := MulStrassen(A11.PlusSparseQuiet(A12), B22)
    fmt.Printf("AFTER>>>>>000: Strassen: M5: \n%v\n", M5)
	M6 := MulStrassen(A21.MinusSparseQuiet(A11), B11.PlusSparseQuiet(B12))
    fmt.Printf("AFTER>>>>>000: Strassen: M6: \n%v\n", M6)
	M7 := MulStrassen(A12.MinusSparseQuiet(A22), B21.PlusSparseQuiet(B22))
    fmt.Printf("AFTER>>>>>000: Strassen: M7: \n%v\n", M7)


	C11.AddSparse(M7)
    C11.AddSparse(M1)
    C11.AddSparse(M4)
    C11.SubtractSparse(M5)

	C12.AddSparse(M5)
    C12.AddSparse(M3)

	C21.AddSparse(M4)
    C21.AddSparse(M2)

	C22.AddSparse(M6)
    C22.AddSparse(M1)
    C22.SubtractSparse(M2)
    C22.AddSparse(M3)
    //fmt.Printf(">>>>>222: Strassen: C11: \n%v\n", C11);
    //fmt.Printf(">>>>>222: Strassen: C12: \n%v\n", C12);
    //fmt.Printf(">>>>>222: Strassen: C21: \n%v\n", C21);
    //fmt.Printf(">>>>>222: Strassen: C22: \n%v\n", C22);

    //fmt.Printf(">>>>>999: Strassen: C: \n%v\n", C);

	return C
}
/* */
