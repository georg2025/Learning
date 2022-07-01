package palindrome

import (
"strconv"  
)


func isPalindrome(x int) bool {
    z:=strconv.Itoa(x)
	for i, _:=range z{
	    if z[i]!=z[len(z)-1-i] {
	    return false
	    }
	}
return true
}
