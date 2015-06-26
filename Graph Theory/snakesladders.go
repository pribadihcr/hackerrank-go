//thanks to https://www.youtube.com/watch?v=9QV6QpyhN0o for excellent tutorial of floyd Warshall algorithm
package main
import (
    "fmt"
    "os"
    "bufio"
)

func floydWarshall(m [][]int, board int)[][]int{    
    for k:=0;k<board;k++{
        for i:=0;i<board;i++{
            if m[i][k]!=101 && i!=k{ 
                for j:=0;j<board;j++{
                    if i!=j && m[k][j]!=101 {
                        if m[i][j]>m[i][k]+m[k][j]{
                            m[i][j]=m[i][k]+m[k][j]
                        }
                    }
                }
            }
        }
    }
    return m
}

func initialize(board int, rangeroll int)[][]int{
    m := make([][]int,board)
    for i:=0;i<board;i++{
        m[i] = make([]int,board)
        for j:=0;j<board;j++{
            if i==j{
                m[i][j]=0
            }else if j>i && j<i+rangeroll+1{
                m[i][j]=1
            }else{
                m[i][j]=101
            }
        }
    }
    return m
}
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var T,N,M,p1,p2 int
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io,&T)
    for T>0{
        adjm := initialize(100,6)
        fmt.Fscan(io,&N)
        for i:=0;i<N;i++{
            fmt.Fscan(io,&p1,&p2)
            adjm[p1-1][p2-1]=0
        }
        fmt.Fscan(io,&M)
        for i:=0;i<M;i++{
            fmt.Fscan(io,&p1,&p2)
            adjm[p1-1][p2-1]=0
            for j:=1;j<=6;j++{
                if (p1-1+j)>100-1 { break }
                adjm[p1-1][p1-1+j]=101
            }
        }
        dist := floydWarshall(adjm,100)
        if dist[0][99]==101{
            fmt.Println("-1")
        }else{
            fmt.Println(dist[0][99])
        }
        T--
    }
}


