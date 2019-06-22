package string

import (
	"fmt"
	"strconv"
)

/*问题*/
/*
国际摩尔斯密码定义一种标准编码方式，将每个字母对应于一个由一系列点和短线组成的字符串， 比如: "a" 对应 ".-", "b" 对应 "-...", "c" 对应 "-.-.", 等等。

为了方便，所有26个英文字母对应摩尔斯密码表如下：

[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
给定一个单词列表，每个单词可以写成每个字母对应摩尔斯密码的组合。例如，"cab" 可以写成 "-.-..--..."，(即 "-.-." + "-..." + ".-"字符串的结合)。我们将这样一个连接过程称作单词翻译。

返回我们可以获得所有词不同单词翻译的数量。

例如:
输入: words = ["gin", "zen", "gig", "msg"]
输出: 2
解释:
各单词翻译如下:
"gin" -> "--...-."
"zen" -> "--...-."
"gig" -> "--...--."
"msg" -> "--...--."

共有 2 种不同翻译, "--...-." 和 "--...--.".


注意:

单词列表words 的长度不会超过 100。
每个单词 words[i]的长度范围为 [1, 12]。
每个单词 words[i]只包含小写字母。
*/

/*思路*/
/*
建立摩斯码和小写字母之间的映射数组
map计数
*/
var DictMorse = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func UniqueMorseRepresentations(words []string) int {
	transMap := make(map[string]bool)
	for _, str := range words {
		temp := ""
		for _, ch := range str {
			temp += DictMorse[ch-'a']
		}
		transMap[temp] = true
	}

	res := 0
	for range transMap {
		res++
	}
	return res
}

/*问题*/
/*
给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。

示例:

输入: "25525511135"
输出: ["255.255.11.135", "255.255.111.35"]
*/
/*思路*/
/*
ip的合法格式:每一段只能是[0,255]
a.递归
字符串s从第二位开始选择，是否插入'.',选择完之后，截取剩余部分s[i+1：]，进行下一轮迭代
迭代终止于插入完3个'.'。根据ip合法性判断本次迭代结果是否有效
b.三层循环
相当于将三个'.'插入字符串，保证插入之后各段均有效
*/
func RestoreIpAddresses(s string) []string {
	res := make([]string, 0)
	length := len(s)
	if length < 4 || length > 16 {
		return nil
	}
	//第一个'.'在k位之后
	for k := 0; k < length-3; k++ {
		if s[0] == '0' && k > 0 {
			continue
		}
		a := s[:k+1]

		aInt, _ := strconv.Atoi(a)
		if aInt > 255 {
			continue
		}
		//第二个'.'在l位之后(k后)
		for l := k + 1; l < length-2; l++ {
			if s[k+1] == '0' && l > k+1 {
				continue
			}

			b := s[k+1 : l+1]
			bInt, _ := strconv.Atoi(b)
			if bInt > 255 {
				continue
			}

			//第三个'.'在m位之后(l后)
			for m := l + 1; m < length-1; m++ {
				if s[l+1] == '0' && m > l+1 {
					continue
				}

				c := s[l+1 : m+1]
				cInt, _ := strconv.Atoi(c)
				if cInt > 255 {
					continue
				}

				d := s[m+1:]
				if d[0] == '0' && len(d) > 1 {
					continue
				}
				dInt, _ := strconv.Atoi(d)
				if dInt > 255 {
					continue
				}

				res = append(res, a+"."+b+"."+c+"."+d)
			}
		}
	}

	return res
}

/*问题*/
/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/
/*思路*/
/*
纯模仿竖式乘法
效率低的一笔。。。再想想怎么改
*/
func Add(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	len := len1
	if len2 > len1 {
		len = len2
	}
	var temp1 uint8
	var temp2 uint8
	var tempRes uint8
	isAddOne := false

	res := make([]byte, len)
	for i := 0; i < len; i++ {
		temp1 = 0
		temp2 = 0
		ind1 := len1 - i - 1
		if ind1 >= 0 {
			temp1 = num1[ind1] - '0'
		}
		ind2 := len2 - i - 1
		if ind2 >= 0 {
			temp2 = num2[ind2] - '0'
		}
		tempRes = temp1 + temp2

		if isAddOne {
			tempRes++
			isAddOne = false
		}
		if tempRes >= 10 {
			isAddOne = true
		}
		res[len-i-1] = tempRes%10 + '0'
	}
	str := string(res)
	if isAddOne {
		str = "1" + str
	}
	return str
}
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := "0"

	var temp1 uint8
	var temp2 uint8
	var temp string

	len1 := len(num1)
	len2 := len(num2)

	for i := 0; i < len1; i++ {
		temp1 = num1[len1-i-1] - '0'
		tempRes := "0"
		for j := 0; j < len2; j++ {
			temp2 = num2[len2-j-1] - '0'
			temp = strconv.Itoa(int(temp1 * temp2))
			for indj := j; indj > 0; indj-- {
				temp = temp + "0"
			}
			tempRes = Add(tempRes, temp)
		}
		for indi := i; indi > 0; indi-- {
			tempRes = tempRes + "0"
		}
		res = Add(tempRes, res)
	}

	return res
}

/*问题*/
/*
我们有一些二维坐标，如 "(1, 3)" 或 "(2, 0.5)"，然后我们移除所有逗号，小数点和空格，得到一个字符串S。返回所有可能的原始字符串到一个列表中。

原始的坐标表示法不会存在多余的零，所以不会出现类似于"00", "0.0", "0.00", "1.0", "001", "00.01"或一些其他更小的数来表示坐标。此外，一个小数点前至少存在一个数，所以也不会出现“.1”形式的数字。

最后返回的列表可以是任意顺序的。而且注意返回的两个数字中间（逗号之后）都有一个空格。



示例 1:
输入: "(123)"
输出: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
示例 2:
输入: "(00011)"
输出:  ["(0.001, 1)", "(0, 0.011)"]
解释:
0.0, 00, 0001 或 00.01 是不被允许的。
示例 3:
输入: "(0123)"
输出: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
示例 4:
输入: "(100)"
输出: [(10, 0)]
解释:
1.0 是不被允许的。


提示:

4 <= S.length <= 12.
S[0] = "(", S[S.length - 1] = ")", 且字符串 S 中的其他元素都是数字。
*/
/*思路*/
/*
1.逗号的位置拆分为两节，没变至少一个数字字符
2.拆分开后，各自循环，可能的小数点的位置
3.两节组合
*/
func checkValidPoint(str string, ind int) bool {
	//检验整数部分
	if ind >= 1 && str[0] == '0' {
		return false
	}
	//检验小数部分
	if len(str)-ind > 1 && str[len(str)-1] == '0' {
		return false
	}
	return true
}
func AmbiguousCoordinates(s string) []string {
	S := s[1 : len(s)-1]
	length := len(S)
	if length < 2 {
		return nil
	}
	res := make([]string, 0)
	for i := 1; i < length; i++ {
		temp1 := S[:i]
		temp2 := S[i:]
		resA := make([]string, 0)
		resB := make([]string, 0)
		for ii := 0; ii < i; ii++ {
			if checkValidPoint(temp1, ii) {
				if ii == i-1 {
					resA = append(resA, temp1[:ii+1])
				} else {
					resA = append(resA, temp1[:ii+1]+"."+temp1[ii+1:])
				}
			}
		}
		for ii := 0; ii < length-i; ii++ {
			if checkValidPoint(temp2, ii) {
				if ii == length-i-1 {
					resB = append(resB, temp2[:ii+1])
				} else {
					resB = append(resB, temp2[:ii+1]+"."+temp2[ii+1:])
				}
			}
			fmt.Println(temp2, temp2[:ii+1], temp2[ii+1:], ii, resB)
		}

		for _, str1 := range resA {
			for _, str2 := range resB {
				res = append(res, "("+str1+", "+str2+")")
			}
		}
	}
	return res
}

/*问题*/
/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/
/*思路*/
/*
n=4

0         6       12
1     5   7    11 13
2  4      8 10    14
3         9       15

首位两行
v(i+1) = v(i) + 2(n-1)

中间行
偶数行
v(i+1) = v(i) + 2(n-r)
奇数行
v(i+1) = v(i) + 2r


*/
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	newB := make([]byte, length)
	var i int
	ind := 0
	for i = 0; i <= (numRows - 1); i++ {
		temp := i
		cnt := 1
		for temp < length {
			newB[ind] = s[temp]
			ind++
			if i == 0 || i == (numRows-1) {
				temp = temp + 2*(numRows-1)
			} else {
				if cnt == 1 {
					temp = temp + 2*(numRows-i) - 2
					cnt = 2
				} else {
					temp = temp + 2*i
					cnt = 1
				}
			}
		}
	}

	return string(newB)
}
