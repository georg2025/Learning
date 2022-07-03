//you have slice of int. You need to imagine if it was one number (ex [9,9,9] is 999) and return this number + 1 as a new slice.
func plusOne(digits []int) []int {
    v:=digits[len(digits)-1]
    if v!=9{
        digits[len(digits)-1]+=1
        return digits
    }

    for x:=len(digits)-1; x>=0;x--{
        if digits[x]!=9 {
            digits[x]+=1
            for z:=x+1; z<len(digits); z++{
                digits[z]=0}
                return digits


        }
    }
        var new_digits []int
        new_digits = append(new_digits, 1)
        for n:=1; n<=len(digits); n++{
        new_digits = append(new_digits, 0)}
        return new_digits

return digits
}
