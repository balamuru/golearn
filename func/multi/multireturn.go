package multi

// import "fmt"

//export func
func Vals() (int, int) {
	return 3, 7
}

func NowTime() string{
	return time.Now().String()
}