package main

import "fmt"

func main() {
	// 	9. Dado un array de 5 x 5 booleanos determinar cual será el camino (solo para valientes) para llegar de un punto a otro de la matriz teniendo en cuenta que
	// la matriz representa un laberinto
	// el laberinto puede no tener salida
	// solo existe un posible camino (no hay bifurcaciones)
	// uno comienza en la posición 0,0 y debe llegar a la 4,4
	// uno puede moverse de a un casillero horizontalmente (izquierda o derecha) o verticalmente (arriba o abajo)
	// una posición true representa un camino libre por el que se puede ir
	// una posición false representa un camino bloqueado por el que no se puede ir
	// imprimir:

	// el camino es: n1,m1 -> n2,m2 -> n3,m3 -> n4,m4 -> etc
	// no existe un camino para llegar
	const cant int = 5
	var laby [cant][cant]bool = [cant][cant]bool{
		[cant]bool{true, true, true, true, true},
		[cant]bool{true, false, true, true, true},
		[cant]bool{true, false, true, false, true},
		[cant]bool{true, true, true, false, true},
		[cant]bool{true, false, true, false, true},
	}
	var y int
	var x int
	var escape bool
	var now [2]int
	var lastSpot [2]int
	var path [][2]int

	for escape = false; !escape; {
		path = append(path, now)
		if now == [2]int{4, 4} {
			escape = true
		} else if y < cant-1 && laby[y+1][x] && [2]int{y + 1, x} != lastSpot {
			lastSpot = now
			now = [2]int{y + 1, x}
			y++
		} else if x < cant-1 && laby[y][x+1] && [2]int{y, x + 1} != lastSpot {
			lastSpot = now
			now = [2]int{y, x + 1}
			x++
		} else if y > 0 && laby[y-1][x] && [2]int{y - 1, x} != lastSpot {
			lastSpot = now
			now = [2]int{y - 1, x}
			y--
		} else if x > 0 && laby[y][x-1] && [2]int{y, x - 1} != lastSpot {
			lastSpot = now
			now = [2]int{y, x - 1}
			x--
		} else {
			break
		}
	}

	if escape {
		fmt.Printf("El camino es: ")
		for y = 0; y < len(path); y++ {
			fmt.Printf("%d,%d", path[y][0], path[y][1])
			if y+1 != len(path) {
				fmt.Printf(" -> ")
			}
		}
	} else {
		fmt.Printf("No existe un camino para llegar.")
	}
}
