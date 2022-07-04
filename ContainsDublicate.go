//function search if slice contains dublicate
func containsDuplicate(nums []int) bool {
    mp:=map[int]int{}
    for i:=0; i<len(nums); i++{
        mp[nums[i]]++
    }
    for _,y:=range mp{
        if y>=2{
            return true
        }
    }
return false
}
