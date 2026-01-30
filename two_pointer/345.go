package main






func isVowel(r rune) bool {
	vowels := map[rune]bool{
    'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
    return vowels[r]
}

func reverseVowels(s string) string {
    a:=[]rune(s)
	left,right:=0,len(a)-1
	for left < right{
		if isVowel(a[left]){
			for left < right{
				if isVowel(a[right]){break}else{right--}
			}
			if left < right{
			temp:=a[right]
			a[right] = a[left]
			a[left] = temp
			left++
			right--
			}else{break}
		}else{
			left++
		}
	}
	return string(a)
}
