package bit

import "fmt"

/*问题*/
/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。

给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:

输入: x = 1, y = 4

输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置。
*/
/*思路*/
/*
x%2 与 y%2比较，不等汉明距离加一
x=x/2 y=y/2，进行下一位的比较
*/
func HammingDistance(x int, y int) int {
	res := 0
	modX := 0
	modY := 0
	for {
		modX = x % 2
		modY = y % 2
		if modX != modY {
			res++
		}

		x = x / 2
		y = y / 2

		if x == 0 && y == 0 {
			break
		}
	}
	return res
}

/*问题*/
/*
给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。

注意:

给定的整数保证在32位带符号整数的范围内。
你可以假定二进制数不包含前导零位。
示例 1:

输入: 5
输出: 2
解释: 5的二进制表示为101（没有前导零位），其补数为010。所以你需要输出2。
示例 2:

输入: 1
输出: 0
解释: 1的二进制表示为1（没有前导零位），其补数为0。所以你需要输出0。
*/
/*思路*/
/*
逐位处理
*/
func FindComplement(num int) int {
	var res int = 0
	ind := 0
	for {
		//不考虑num==0
		if num == 0 {
			break
		}
		temp := (1 - num%2)
		for i := 0; i < ind; i++ {
			temp = temp * 2
		}

		res = res + temp
		ind++
		fmt.Println(num%2, num, res)
		num = num / 2
	}
	return res
}

/*问题*/
/*
所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找 DNA 分子中所有出现超多一次的10个字母长的序列（子串）。

示例:

输入: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

输出: ["AAAAACCCCC", "CCCCCAAAAA"]
*/
/*思路*/
/*
10个字母一组计算hash值
以字母序列为key，是否进入res为value,构建map
*/
func FindRepeatedDnaSequences(s string) []string {
	var res []string
	length := len(s)
	if length <= 10 {
		return res
	}
	tempMap := make(map[string]bool, length-10)
	for i := 0; i <= length-10; i++ {
		if v, ok := tempMap[s[i:i+10]]; ok {
			if v {
				continue
			}
			res = append(res, string(s[i:i+10]))
			tempMap[s[i:i+10]] = true
			continue
		}
		tempMap[s[i:i+10]] = false
	}
	return res
}
