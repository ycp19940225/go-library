package utils

func Diff[T comparable](src []T, dst []T) (res []T) {
OuterLoop:
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(dst); j++ {
			if src[i] == dst[j] {
				continue OuterLoop
			}
		}
		res = append(res, src[i])
	}
	return
}
