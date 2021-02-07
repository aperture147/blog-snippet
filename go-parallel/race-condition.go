package main

func main() {
	a := []int{1, 2, 3}
	
	for i := 0; i < 5; i++ {
		go func(i int) {
			a = append(a, i)
		}(i)
	}
}
