//This function search number in sorted array of int. If there is no such number - it search where it could be.
func searchInsert(nums []int, target int) int {
    p:=0
    for i,z:=range nums {
        if z==target {
            p=i
            break
        }else if z>target{
            p=i
            break
        }else{
            p=len(nums)
        }

    }
return p
}
