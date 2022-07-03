//find element that appears in slice more than half times
func majorityElement(nums []int) int {
    u:=0
    n:=map[int]int{}
    for i:=0; i<len(nums); i++{
        n[nums[i]]++
    }
    for x,y:=range n {
        if y>len(nums)/2 {
            u=x
        }
    }
return u
}
