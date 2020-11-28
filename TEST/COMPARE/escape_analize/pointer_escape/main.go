package main

type UserData struct {
	Name string
}

func main() {
	var info UserData
	info.Name = "WilburXu"
	// count := 100000
	// start := time.Now().UnixNano()
	// var infochange *UserData
	// for i := 0; i < count; i++ {
	// infochange = GetUserInfo(&info)
	// }
	// println(infochange)
	// col := time.Now().UnixNano() - start

	// start2 := time.Now().UnixNano()
	// for i := 0; i < count; i++ {
	// _ = UsualGetUserInfo(info)
	// }
	// col2 := time.Now().UnixNano() - start2
	// fmt.Printf("指针传递耗时:%#v, Copy传递再赋值耗时是：%#v\n", col, col2)
	_ = UsualGetUserInfo(info)
	_ = GetUserInfo(&info)
	_ = NotPointerRet(info)
	_ = PointerRet(&info)
}

func GetUserInfo(userInfo *UserData) *UserData { //不逃逸
	//较为优势，可以考虑指针传递，main函数内部定义，再指针传递修改值
	return userInfo
}

func UsualGetUserInfo(UserInfo UserData) *UserData { //逃逸
	return &UserInfo
}

func NotPointerRet(userInfo UserData) UserData { //不逃逸
	return userInfo
}

func PointerRet(userInfo *UserData) UserData { //不逃逸
	return *userInfo
}

// go run -gcflags '-m -l' main.go
// # command-line-arguments
// ./main.go:31:18: leaking param: userInfo to result ~r1 level=0
// ./main.go:36:23: moved to heap: UserInfo
// ./main.go:28:12: ... argument does not escape
// ./main.go:28:13: col escapes to heap
// ./main.go:28:13: col2 escapes to heap
// 指针传递耗时:128649, Copy传递再赋值耗时是：2871534
