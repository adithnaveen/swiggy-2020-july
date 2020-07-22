package stringutil

func FullName(f, l string) (string, int) {
	full := f + " " + l
	lenght := len(full)
	return full, lenght
}

func FullNameNakedReturn(f, l string) (full string, lenght int) {
	full = f + " " + l
	lenght = len(full)
	return
}
