package main
import "fmt"

func main() {

		// Tipe Data Dasar di Go
		// 1. String
		// 2. Integer (int, int8, int16, int32, int64)
		// 3. Float (float32, float64)
		// 4. Boolean (bool)
		var nama string
		nama = "Boby"
		fmt.Println("Hello", nama)

		var usia int = 20
		fmt.Println("Usia saya adalah", usia)

		var tiggi float32 = 180.0
		fmt.Println("Tinggi saya adalah", tiggi)

		var benar bool = true
		fmt.Println("Apakah saya benar?", benar)

		kesimpulan := "Nama saya adalah " + fmt.Sprint(nama) + ", usia saya " + fmt.Sprint(usia) + " tahun, tinggi saya " + fmt.Sprint(tiggi) + " cm."
		fmt.Println(kesimpulan)

		// Zero value
		// Kalo variabel yang dideklari ga dikasih nilai awal, dia bakal punya nilai default, misal
		// var angka int
		// fmt.Println(angka) // output: 0

		// Konversi tipe data
		// misal:
		// var x int = 10
		// var y float64 = float64(x)
		// fmt.Println(y) // output: 10.0
}
