func removeElement(nums []int, val int) int {
    for i:=len(nums)-1; i>=0; i--{
        if nums[i]==val{
            if i==len(nums)-1{
                nums=nums[:len(nums)-1]
            }else{
                copy(nums[i:], nums [i+1:])
                nums=nums[:len(nums)-1]
            }
        }

    }
    return(len(nums))
}
