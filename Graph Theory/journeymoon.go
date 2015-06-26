package main
import "fmt"

var astro map[int][]int
var found_country map[int]bool

func dfs(node int, astro map[int][]int, fn func (int)) {
    dfs_recur(node, fn)
}

func dfs_recur(node int, fn func (int)) {
    found_country[node] = true
    fn(node)
    for _, n := range astro[node] {
        if _, ok := found_country[n]; !ok {
            dfs_recur(n, fn)
        }
    }
}

func main() {
    var N,I int
    fmt.Scanf("%d %d\n",&N,&I)
    astro = make(map[int][]int,N)
    found_country = make(map[int]bool,N)
    for i:=0;i<I;i++{
        var in1, in2 int
        fmt.Scanf("%d %d\n",&in1,&in2)
        astro[in1] = append(astro[in1],in2)
        astro[in2] = append(astro[in2],in1)
    }
    
    var ways,sum int
    for i:=0;i<N;i++{ 
        visited := []int{}
        if(!found_country[i]){
            dfs(i, astro, func (node int) {
                visited = append(visited, node)
            })    
            ways += sum * len(visited)
            sum += len(visited)
        }
    }
    fmt.Println(ways)
}
