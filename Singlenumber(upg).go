//upd function of searching single number in slice
func singleNumber(nums []int) int {
    f:=0
    d:=map[int]int{}
    for z:=len(nums)-1; z>=0; z--{
        d[nums[z]]++
    }
    for x,y:=range d{
        if y==1{
            f=x
        }
    }
return f
}
