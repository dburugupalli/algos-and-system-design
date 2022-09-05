pkg main 

import "fmt"

func generateTriangle(numRows int) [][]int {
  res := [][]int{{1}}
  for i:=1; i< len(numRows); i++{
        arr := []int{1}
        for j:= 1; j < i; j++ {
            arr = append(arr, res[i - 1][j - 1] + res[i - 1][j])
        } 
        
        arr = append(arr, 1)
        res = append(res, arr)
  }
  return res 
}


func main(){
  fmt.Println("Generate Pascals Triangle")
  answer := generateTriangle(3)
  fmt.Println(answer)
}
