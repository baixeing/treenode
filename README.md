# treenode
The most annoying thing to do with solving leetcode's TreeNode problems is generating and displaying binary trees.
This package is set of functions is designed to help solve problems with binary trees. 

Example:
```golang
package main

import (
    "fmt"

    "github.com/baixeing/treenode"
)

func main() {
    t1 := treenode.NewRandomBST(5, -10, 100, 15)
    t2 := treenode.NewRandomFull(4, -10, 200)
    
    fmt.Println(t1)
    fmt.Println(t1.Invert())
    fmt.Println(t2)
    fmt.Println(t2.Invert())
}
```

```
                 ╭─ 24 ─────────────────────────╮                  
             ╭─ 12                   ╭───────── 61 ─────────────╮  
      ╭───── 1               ╭───── 37 ─╮                ╭───── 89 
  ╭─ -2 ─╮               ╭─ 31 ─╮       39 ─╮        ╭─ 81 ─╮      
 -4      -1             28      34          49      73      83     

                  ╭───────────────────────── 24 ─╮                 
  ╭───────────── 61 ─────────╮                   12 ─╮             
 89 ─────╮                ╭─ 37 ─────╮               1 ─────╮      
      ╭─ 81 ─╮        ╭─ 39       ╭─ 31 ─╮               ╭─ -2 ─╮  
     83      73      49          34      28             -1      -4 

               ╭─────────────── 155 ──────────────╮               
      ╭────── 19 ──────╮                   ╭───── 13 ──────╮      
  ╭─ 48 ──╮        ╭── 46 ──╮         ╭── 64 ─╮        ╭── 74 ─╮  
 44      158      184      143       161      17      194      10 

               ╭────────────── 155 ───────────────╮               
      ╭────── 13 ─────╮                   ╭────── 19 ──────╮      
  ╭─ 74 ──╮        ╭─ 64 ──╮         ╭── 46 ──╮        ╭── 48 ─╮  
 10      194      17      161       143      184      158      44 

```

