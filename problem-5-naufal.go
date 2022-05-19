package main
//problem 5 menghitung luas permukaan tabung naufal
import "fmt"

func main() {
	//input
	var T float64 = 20
	var r float64 = 4
	var phi float64 = 22 / 7

	//kode disini
	var luasTabung float64 = 2 * phi * r * (r + T)
	fmt.Println(luasTabung)
}
