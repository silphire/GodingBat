package string1

func makeTags(tag string, word string) string {
	return "<" + tag + ">" + word + "</" + tag + ">";
}