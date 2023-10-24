package main

import "sync"

func main() {
	m := make(map[string]int)
    m["ares"] = 1

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for i := 0; i < 1000; i++ {
            m["ares"]++
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            m["ares"]++
        }
        wg.Done()
    }()

    wg.Wait()
}