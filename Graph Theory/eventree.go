package main
import (
    "fmt"
    "os"
    "bufio"
)

func countVertice_rec(v map[int][]int,p int)int{
    res := 0
    for _,x := range v[p]{
        res = countVertice_rec(v,x) + res + 1
    }
    return res
}
func removeEdge(v map[int][]int)int{
    res := 0
    for _,p := range v{
        for _,cp := range p{
            nv := countVertice_rec(v,cp) + 1
            if(nv%2==0){
                res++
            }
        }
    }
    return res
}
func main() {
    io := bufio.NewReader(os.Stdin)
    var N,M int
    fmt.Fscan(io,&N,&M)
    vertice := make(map[int][]int,0)
    var child,parent int
    for i:= 1;i<=M;i++{
        fmt.Fscan(io,&child,&parent)
        vertice[parent] = append(vertice[parent],child)
    }
    
    fmt.Println(removeEdge(vertice))
    
}
