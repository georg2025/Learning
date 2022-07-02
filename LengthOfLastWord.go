//this function give lenght of last word in sting of endlish letters and spaces.
func lengthOfLastWord(s string) int {
    r:=0
    x:=strings.TrimSpace(s)
    z:=strings.Split(x, "")
    for f:=len(z)-1; f>=0; f--{
        if z[f]!=" "{
            r+=1
        }else{break}

    }
return r
}
