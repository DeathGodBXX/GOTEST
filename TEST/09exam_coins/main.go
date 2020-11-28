/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int { //user  coins
	for _, user := range users {
		//字符串数组第一个元素是索引，第二个是字符串
		count := 0
		for _, i := range []byte(user) {
			//单个英文字符的比较，需要转换类型，第一个是索引，第二个是字符
			switch i {
			case 'e', 'E':
				count++
			case 'i', 'I':
				count += 2
			case 'o', 'O':
				count += 3
			case 'u', 'U':
				count += 4
			}
		}
		distribution[user] = count
		if coins >= 0 {
			coins -= count
			fmt.Printf("%s has got %d coins\n", user, count)
		} else {
			fmt.Printf("coins isn't enough for %s,stop!!!!\n", user)
			break
		}
	}
	return coins
}
