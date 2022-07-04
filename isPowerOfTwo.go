//find out if int == 2 to some power
func isPowerOfTwo(n int) bool {
    for {
        if n<=0 {
        return false}
        if n==1 {
        return true}
        if n%2!=0 {return false}
        n=n/2
    }
return true
}
