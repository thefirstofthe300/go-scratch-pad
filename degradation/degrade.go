package degradation

func Degrade(str string, index1 int, index2 int) string {
	var s1, s2, s3 string

	if index1 == index2 {
		return str
	} else if index1 < index2 {
		s1 = str[:index1]
		// Adding one to the second index to make in inclusive
		s2 = reverseString(str[index1 : index2+1])
		s3 = str[index2+1:]
	} else {
		// Add one to the second index to make it inclusive
		s1 = str[:index2+1]
		s2 = str[index2+1 : index1]
		s3 = str[index1:]
		revd := reverseString(s3 + s1)
		s3 = revd[:len(s3)]
		s1 = revd[len(s3):]
	}
	return s1 + s2 + s3
}

func reverseString(str string) string {
	runes := []rune(str)
	var revd []rune
	for i, _ := range runes {
		// Subtract one to account for length being 1-indexed
		revd = append(revd, runes[len(runes)-i-1])
	}
	return string(revd)
}
