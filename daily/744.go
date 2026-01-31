package main

func nextGreatestLetter(letters []byte, target byte) byte {
	left:=0
	right:=len(letters)-1
	for left < right{
		mid:=(left+right)/2
		if target>letters[mid]{
			left = mid+1
		}else if target < letters[mid]{
			right = mid
		}else if target == letters[mid]{
			left = mid+1
		}
    }
    if letters[left] > target{return letters[left]}
    return letters[0]
}
