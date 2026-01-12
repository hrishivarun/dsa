// Online Go compiler to run Golang program online
// Print "Hello, World!" message

package main
import "fmt"


// 1)
// [3.5]
// [1,2]
// 2)
// [3.5]
// [1,4]
// 3)
// [3,6]
// [4,5]
// 4)
// [3,6]
// [4,7]
// 5)
// [3,5]
// [6,7]
func findCommonTiming(timingList [][][2]int) [][2]int{
    var pointer []int = make([]int, len(timingList))
    var common [][2]int = make([][2]int)

    mostAhead := 0
    for {
        currMostAhead := mostAhead
        start := timingList[mostAhead][pointers[mostAhead]][0]
        end := timingList[mostAhead][pointers[mostAhead]][1]

        latestStart := start
        breakInner := false
        finish := false
        for i:=0; i<len(timingList); i++ {
            // 1 )
            for (timingList[i][pointers[i]][1]<=start) {
                if(pointer[i]==len(timingList[i])) {
                    return common
                }
                pointer[i]++
            }
            // 2 )
            if(timingList[i][pointer[i]][0]>=end) {
                currMostAhead  = i
                breakInner = true
                break
            }
            // 3 )
            if(timingList[i][pointers[i]][0]<end) {
                if(timingList[i][pointers[i]][0]>start) {
                    start = timingList[i][pointers[i]][0]
                    currMostAhead = i
                }
                // 4)
                if(timingList[i][pointers[i]][1]<=end){
                    end = timingList[i][pointers[i]][1]
                    pointers[i]++
                    if(pointer[i]==len(timingList[i])) {
                        finish = true
                    }
                }
            }
        }
        
        mostAhead = currMostAhead
        if(breakInner) {
            continue
        }
        if(start<end) {
            common = append(common, []int{start, end})
        }
        if(finish) {
            return common
        }
    }

    return common
}

func main() {
  fmt.Println("Hello, World!")
}