package Internal

func Likes(array []string) string {
	for i := 0; i < len(array)-1; i++ {
		if array[i] == "Like" && array[i+1] == "Like" {
			array[i] = ""
			array[i+1] = "Nothing"
		} else if array[i] == "Dislike" && array[i+1] == "Dislike" {
			array[i] = ""
			array[i+1] = "Nothing"
		} else if array[i] == "Like" && array[i+1] == "Dislike" {
			array[i] = ""
			array[i+1] = "Dislike"
		} else if array[i] == "Dislike" && array[i+1] == "Like" {
			array[i] = ""
			array[i+1] = "Like"
		} else if array[i] == "Dislike" && array[i+1] == "Like" {
			array[i] = ""
			array[i+1] = "Like"
		} else if array[i] == "Nothing" && array[i+1] == "Nothing" {
			array[i] = ""
			array[i+1] = "Like"
		} else if array[i] == "Nothing" && array[i+1] == "Like" {
			array[i] = ""
			array[i+1] = "Like"
		} else if array[i] == "Nothing" && array[i+1] == "Dislike" {
			array[i] = ""
			array[i+1] = "Dislike"
		}
	}
	return array[len(array)-1]
}
