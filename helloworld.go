package main

func nama() (string, string) {
	return "mamang", "dopan"
}

func multiple() (first string, mid string, last string) {
	first = "mamang"
	mid = "dopan"
	last = "ganteng"

	return
}

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	//fmt.Println("hello world")
	//sepuluh := 10000
	//time.Sleep(time.Duration(sepuluh))
	//println(sepuluh)

	//var names [2]string
	//
	//names[0] = "mamang"
	//names[1] = "racing"
	//
	//println(names[0] + " " + names[1])
	//
	//var values = [3]int{
	//	90, 20, 209,
	//}
	//fmt.Println(values[0])
	//
	//balap := [4]string{
	//	"mamang",
	//	"racing",
	//	"balap",
	//	"supra",
	//}
	//
	//slice := balap[1:]
	//
	//fmt.Println(len(slice))
	//fmt.Println(cap(slice))
	//
	//orang := map[string]string{
	//	"name":  "mamang",
	//	"hobby": "balap",
	//}
	//
	//fmt.Println(orang)
	//fmt.Println(orang["name"])
	//fmt.Println(orang["hobby"])

	//objmap := make(map[string]string)
	//
	//objmap["nama"] = "mamang racing"
	//objmap["hobby"] = "balap"
	//objmap["salah"] = "salah"
	//fmt.Println(objmap)
	//
	//delete(objmap, "salah")
	//
	//fmt.Println(objmap)
	//fmt.Println(objmap["nama"])
	//fmt.Println(objmap["hobby"])

	//counter := 1
	//
	//for counter <= 20 {
	//	fmt.Println("perulangan ke ", counter)
	//
	//	counter++
	//}
	//
	//for counter := 1; counter <= 50; counter++ {
	//	fmt.Println("perulangan yang ke ", counter)
	//}
	//
	//slice := []string{
	//	"mamang",
	//	"racing",
	//	"hobby balap",
	//}
	//
	//for i := 0; i < len(slice); i++ {
	//	fmt.Println(slice[i])
	//}
	//
	//for _, v := range slice {
	//	fmt.Println("Value =", v)
	//}
	//
	//for i, _ := range slice {
	//	fmt.Println("index =", i)
	//}
	//
	//mamap := make(map[string]string)
	//
	//mamap["nama"] = "mamang"
	//mamap["julukan"] = "dopan"
	//
	//for key, val := range mamap {
	//	fmt.Println(key, val)
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println("perulangan ke ", i)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	if i%2 == 1 {
	//		continue
	//	}
	//	fmt.Println("perulangan ke ", i)
	//}

	//awal, akhir := nama()
	//first, _ := nama()
	//
	//fmt.Println(first)
	//fmt.Println(awal, akhir)

	//a, b, c := multiple()
	//
	//fmt.Println(a, b, c)

	//total := sumAll(10, 23, 43, 45, 64)
	//fmt.Println(total)
	//
	//slice := []int{10, 34, 123, 34, 54}
	//total = sumAll(slice...)
	//fmt.Println(total)
}
