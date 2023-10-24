# two-sum

```golang
func twoSum(nums []int, targe)[]int {
    tmp := make(map[int]int)
    for i, v := range nums{
        if n, ok := tmp[targe-v];ok{
            return []int{n,i}
        }
        tmp[v] = i
    }
    return nil
}
```
