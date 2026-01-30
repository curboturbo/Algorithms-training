package main
import "fmt"

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapStoT := make(map[byte]byte)
    mapTtoS := make(map[byte]byte)
    for i := 0; i < len(s); i++ {
        charS := s[i]
        charT := t[i]
        valS, existsS := mapStoT[charS]
        if existsS {
            if valS != charT {
                return false
            }
        } else {
            mapStoT[charS] = charT
        }
        valT, existsT := mapTtoS[charT]
        if existsT {
            if valT != charS {
                return false
            }
        } else {
            mapTtoS[charT] = charS
        }
    }
    return true
}
func main(){
	fmt.Print(isIsomorphic("egg","add"))
}