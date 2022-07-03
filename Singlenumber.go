//find single number in slice of int. 
func singleNumber(nums []int) int {
    f:=0
    k:=1
    one:
    for i, n:=range nums{
        two:
        for y,x:=range nums{
            k=0
            if i==y{continue}
            if x==n{k=1
                break two}
        }
        if k==0{f=n
              break one}
    }
return f
}
