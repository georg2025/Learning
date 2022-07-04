//function return len of longest palindrome, that is possible to make of letters in string.
func longestPalindrome(s string) int {
    if s==""{
        return 0
    }
    P:=0
    x:=map[string]int{}
    S:=strings.Split(s, "")
    if len(S)==1 {
        return 1
    }
    for i:=0; i<len(S); i++{
        x[S[i]]++
    }
    for _,y:=range x{
        if y % 2 != 0 {
            P-=1
        }else if y > 2 {P+=y-1}

    }
    if P!=len(S) {P+=1}
    return P
}
