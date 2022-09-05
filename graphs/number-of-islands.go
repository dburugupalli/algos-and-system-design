pkg main

import fmt

func numberOfIslands(arr int[][]) int {
  count := 0; 
  for i=0; i<len(arr); i++{
    for j=0; j< len(arr[0]); j++{
      if arr[i][j] == '1' {
        count = count +1
        callBFS(grid, i, j)
      }
    }
  }
  return count;
}

func callBFS(arr int[][], i int, j int) {
  if i<0 || i >= len(arr) || j<=0 || j>=len(arr[0]) || grid[i][j] == '0' {
    return
  }
    // to cover up, down, right and left. 
    callBFS(arr, i, j+1)
    callBFS(arr, i-1, j)
    callBFS(arr, i+1, j)
    callBFS(arr, i, j-1)
  }
}



func main() {
   /* an array with 5 rows and 2 columns*/
   var a = [5][2]int{ {0,0}, {1,0}, {1,1}, {0,0},{1,1}}
   var i, j int

   /* output each array element's value */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }
}
