//function to count square root int (it cant be float, so 2,464345 is 2)
func mySqrt(x int) int {
    d:=0
    code:
    for i:=0; i<=x; i++ {

        if i*i>x{
            break code
        }
        d=i
    }

return d
}
