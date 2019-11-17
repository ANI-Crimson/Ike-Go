// 8. Dada una matriz de 9 x 9 números enteros, imprimir la matriz de la siguiente forma
// + - - - + - - - + - - - +
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// + - - - + - - - + - - - +
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// + - - - + - - - + - - - +
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// | n n n | n n n | n n n |
// + - - - + - - - + - - - +
// e indicar si representa una solución válida para el SUDOKU según las siguientes reglas
// las filas deben contener solo números del 1 al 9
// las filas no deben contener números repetidos
// las columnas deben contener solo números del 1 al 9
// las columnas no debe contener números repetidos
// cada cuadrante de 3 x 3 debe contener solo números del 1 al 9
// cada cuadrante de 3 x 3 no debe contener números repetidos

package main

import "fmt"

func main() {
	var sudoku [9][9]int = [9][9]int{
		[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[9]int{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[9]int{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[9]int{9, 1, 2, 3, 4, 5, 6, 7, 8},
		[9]int{6, 7, 8, 9, 1, 2, 3, 4, 5},
		[9]int{3, 4, 5, 6, 7, 8, 9, 1, 2},
		[9]int{8, 9, 1, 2, 3, 4, 5, 6, 7},
		[9]int{5, 6, 7, 8, 9, 1, 2, 3, 4},
		[9]int{2, 3, 4, 5, 6, 7, 8, 9, 1},
	}
	var i int
	var j int
	const line string = "+ - - - + - - - + - - - +"
	var valido bool = true
	var acumuladoresX [9]int
	var acumuladoresY [9]int
	var cuadAcum [3][3][9]int = [3][3][9]int{
		[3][9]int{[9]int{}, [9]int{}, [9]int{}},
		[3][9]int{[9]int{}, [9]int{}, [9]int{}},
		[3][9]int{[9]int{}, [9]int{}, [9]int{}},
	}

	// ------------------------------ Dibujito ----------------------------
	for i = 0; i < len(sudoku); i++ {
		if i%3 == 0 {
			fmt.Printf("%s\n", line)
		}
		for j = 0; j < len(sudoku[i]); j++ {
			if j%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", sudoku[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("%s\n", line)

	// ------------------------------ Reglas ------------------------------
	for i = 0; i < len(sudoku); i++ {
		for j = 0; j < len(sudoku[i]); j++ {
			if sudoku[i][j] < 1 || sudoku[i][j] > 9 { // Entre 1 y 9 -----
				valido = false
				break
			}
			acumuladoresX[sudoku[i][j]-1]++ // Filas no se repite ---------
			acumuladoresY[sudoku[j][i]-1]++ // Columnas no se repite ------
			if acumuladoresX[sudoku[i][j]-1] > 1 || acumuladoresY[sudoku[j][i]-1] > 1 {
				valido = false
				break
			}
			cuadAcum[i/3][j/3][sudoku[i][j]-1]++ // Cuadrante no se repite-
			if cuadAcum[i/3][j/3][sudoku[i][j]-1] > 1 {
				valido = false
				break
			}
		}
		if !valido {
			break
		}
		acumuladoresX = [9]int{}
		acumuladoresY = [9]int{}
	}

	// ------------------------------ Print ------------------------------
	if valido {
		fmt.Println("Es válido.")
	} else {
		fmt.Println("No es válido.")
	}

}
