package main
import "fmt"
func main(){
	nums := []int{3,4,5,6,7}
	slice := []int{1}
    slice = append(slice, append(nums,1)...)
	fmt.Println(slice)
}