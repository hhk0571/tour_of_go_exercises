/* https://tour.golang.org/moretypes/23

Task: Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
*/

package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    str := strings.Fields(s) //it's slice
    m := map[string]int{}

    for _, word := range str {
        _, ok := m[word]
        switch {
        case !ok:
            m[word] = 1
        case ok:
            m[word]++
        }
    }

    return m
}

func main() {
    wc.Test(WordCount)
}

