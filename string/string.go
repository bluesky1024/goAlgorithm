package string

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

/*问题*/
/*
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。

不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。



示例 1：

输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
示例 2：

输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/
/*思路*/
/*
毫无亮点，无非就是第i位与第length-i-1位互换，注意别换重了
*/
func ReverseString(s []byte) {
	length := len(s)
	mid := length / 2
	for i := 0; i <= mid; i++ {
		if i >= length-i-1 {
			break
		}
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

/*问题*/
/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
*/
/*思路*/
/*
1.字符串转[]byte
2.找到空格位置,根据空格位置对原字符串进行拆分
3.[]byte转字符串
*/
func ReverseWords(s string) string {
	byteArr := []byte(s)
	spaceInd := -1
	for i, ch := range byteArr {
		if ch == ' ' {
			temp := byteArr[spaceInd+1 : i]
			ReverseString(temp)
			spaceInd = i
		}
	}
	temp := byteArr[spaceInd+1:]
	ReverseString(temp)
	return string(byteArr)
}

/*问题*/
/*
给定两个字符串，你需要从这两个字符串中找出最长的特殊序列。最长特殊序列定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。

子序列可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。

输入为两个字符串，输出最长特殊序列的长度。如果不存在，则返回 -1。

示例 :

输入: "aba", "cdc"
输出: 3
解析: 最长特殊序列可为 "aba" (或 "cdc")
说明:

两个字符串长度均小于100。
字符串中的字符仅含有 'a'~'z'。
*/
/*思路*/
/*
这道题我没做出来，看了别人的解题思路
感觉智商被按在地上摩擦了。。。
两个字符串不相同,就是特殊,返回长的；相同就没有，返回-1；
*/
func FindLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	lengthA := len(a)
	lengthB := len(b)
	if lengthA > lengthB {
		return lengthA
	}
	return lengthB
}

/*问题*/
/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。

示例 1:

输入: "abc"
输出: 3
解释: 三个回文子串: "a", "b", "c".
示例 2:

输入: "aaa"
输出: 6
说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
注意:

输入的字符串长度不会超过1000。
*/
/*思路*/
/*
1.两层for循环进行遍历检查
2.一层循环，以各个点为中心，考虑奇数个字符和偶数个字符
*/
func CheckPalindrome(s []byte) bool {
	length := len(s)
	mid := length / 2
	for i := 0; i <= mid; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func CountSubstrings_v1(s string) int {
	byteS := []byte(s)
	length := len(s)
	res := 0
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if CheckPalindrome(byteS[i : j+1]) {
				res++
			}
		}
	}
	return res
}

func CountSubstrings_v2(s string) int {
	byteS := []byte(s)
	length := len(s)
	res := 0
	//考虑奇数个字符
	for i := 0; i < length; i++ {
		right := i
		left := i
		for right >= 0 && left < length {
			if byteS[left] == byteS[right] {
				res++
				right--
				left++
			} else {
				break
			}
		}
	}
	//考虑偶数个字符
	for i := 0; i < length-1; i++ {
		right := i
		left := i + 1
		for right >= 0 && left < length {
			if byteS[left] == byteS[right] {
				res++
				right--
				left++
			} else {
				break
			}
		}
	}
	return res
}

/*问题*/
/*
给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：

() 得 1 分。
AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
(A) 得 2 * A 分，其中 A 是平衡括号字符串。


示例 1：

输入： "()"
输出： 1
示例 2：

输入： "(())"
输出： 2
示例 3：

输入： "()()"
输出： 2
示例 4：

输入： "( () ( () ) )" => ( 1 + 1*2 ）* 2 = 6
输出： 6


提示：

S 是平衡括号字符串，且只含有 ( 和 ) 。
2 <= S.length <= 50

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/score-of-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
级别平行括号，分数简单相加
括号套括号，分数翻倍

最好是能用分治方法来做，从小括号逐步扩展到大括号，尽量减少内嵌括号分数的重复计算


*/

func ScoreOfParentheses(S string) int {
	S = strings.Replace(S, "()", "1", -1)

	res := make([]int, 0)
	a := 0
	tempRes := 0
	for _, v := range S {
		if v == '(' {
			res = append(res, 0)
			if tempRes != 0 {
				if len(res) != 1 {
					res[len(res)-2] = res[len(res)-2] + tempRes
				} else {
					a = a + tempRes
				}
				tempRes = 0
			}
		} else if v == ')' {
			if len(res) != 1 {
				res[len(res)-2] = res[len(res)-2] + tempRes*2
				tempRes = res[len(res)-1] * 2
			} else {
				res[0] = res[0] + tempRes
				a = a + res[0]*2
				tempRes = 0
			}
			res = res[:len(res)-1]
		} else {
			tempRes++
		}
	}

	a = a + tempRes
	return a
}

/*问题*/
/*
给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。

所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。



示例 1：

输入：address = "1.1.1.1"
输出："1[.]1[.]1[.]1"
示例 2：

输入：address = "255.100.50.0"
输出："255[.]100[.]50[.]0"


提示：

给出的 address 是一个有效的 IPv4 地址

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/defanging-an-ip-address
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
用正则就行了
*/
func DefangIPaddr(address string) string {
	reg := regexp.MustCompile(`[\.]`)
	return reg.ReplaceAllString(address, "[.]")
}

/*问题*/
/*
给定两个字符串a，b，求取其最大公共子序列长度
此处子序列可不连续，但要保证前后顺序关系
例：
a:abcdefghi
b: bdgh

最大子序列为bdgh,长度为4
*/
/*思路*/
/*
首先，应该存在一定递归关系，a(i,j)应该只与a(i-1,j-1),a(i-1,j),a(i,j-1)这几个值有关
具体关系需要细看
*/

/*问题*/
/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/
/*思路*/
/*
很直观的递归方法，一个个字符处理，从前往后，从后往前都可以
*/
var digitMatchMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	tailSlice := LetterCombinations(digits[1:])

	firstNum := string(digits[0])
	chars, ok := digitMatchMap[firstNum]
	if !ok {
		return nil
	}
	res := make([]string, 0)
	for _, head := range chars {
		for _, tail := range tailSlice {
			res = append(res, head+tail)
		}
		if len(tailSlice) == 0 {
			res = append(res, head)
		}
	}

	return res
}

/*问题*/
/*
给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。

但是，你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。你的表达式不应该含有冗余的括号。

示例：

输入: [1000,100,10,2]
输出: "1000/(100/10/2)"
解释:
1000/(100/10/2) = 1000/((100/10)/2) = 200
但是，以下加粗的括号 "1000/((100/10)/2)" 是冗余的，
因为他们并不影响操作的优先级，所以你需要返回 "1000/(100/10/2)"。

其他用例:
1000/(100/10)/2 = 50
1000/(100/(10/2)) = 50
1000/100/10/2 = 0.5
1000/100/(10/2) = 2
说明:

输入数组的长度在 [1, 10] 之间。
数组中每个元素的大小都在 [2, 1000] 之间。
每个测试用例只有一个最优除法解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/optimal-division
*/
/*思路*/
/*
似乎想复杂了，第二个数一定是被除数，在第二个数之后的数因为都是正整数，只需要保证这些数都是被*进去的，那么结果一定最大

所以无论nums是什么形式
结果都是 a1 / (a2/a3/a4.../an) 的形式
*/
func OptimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.FormatInt(int64(nums[0]), 10)
	}
	if len(nums) == 2 {
		return strconv.FormatInt(int64(nums[0]), 10) + "/" + strconv.FormatInt(int64(nums[1]), 10)
	}

	res := strconv.FormatInt(int64(nums[0]), 10) + "/(" + strconv.FormatInt(int64(nums[1]), 10)
	for i := 2; i < len(nums); i++ {
		res += "/" + strconv.FormatInt(int64(nums[i]), 10)
	}

	res += ")"
	return res
}

/*问题*/
/*
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
回文串，从中间位置，往两边遍历都是一样的

最粗暴的方法，逐个数遍历
存在几种情况
a.回文串长度为奇数，当前位置为中点
b.回文串长度为偶数，ind与ind+1之间为中心
*/
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	res := ""

	// 奇数
	for ind := range s {
		leftPos := ind - 1
		rightPos := ind + 1
		tempLength := 1
		for {
			if leftPos < 0 || rightPos >= len(s) {
				break
			}
			if s[leftPos] != s[rightPos] {
				break
			}
			tempLength += 2
			leftPos--
			rightPos++
		}
		if tempLength > len(res) {
			res = s[leftPos+1 : rightPos]
		}
	}

	// 偶数
	for ind := 0; ind < len(s)-1; ind++ {
		leftPos := ind
		rightPos := ind + 1
		tempLength := 0
		for {
			if leftPos < 0 || rightPos >= len(s) {
				break
			}
			if s[leftPos] != s[rightPos] {
				break
			}
			tempLength += 2
			leftPos--
			rightPos++
		}
		if tempLength > len(res) {
			res = s[leftPos+1 : rightPos]
		}
	}

	return res
}

/*问题*/
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

有效括号组合需满足：左括号必须以正确的顺序闭合。



示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/generate-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
逐个字符处理，只能有两种选择，"("  or ")"
当左侧括号为0的时候只能是左侧括号
否则可以是左侧括号或者右侧括号
*/

func generateParenthesis(n int) []string {
	p := &parenthesisGenerator{
		n:    n,
		list: make([]string, 0),
	}
	p.genParenthesis("", 0, 0)
	return p.list
}

type parenthesisGenerator struct {
	n    int
	list []string
}

func (p *parenthesisGenerator) genParenthesis(curStr string, leftCnt int, rightCnt int) {
	if leftCnt == p.n && rightCnt == p.n {
		p.list = append(p.list, curStr)
		return
	}
	if leftCnt-rightCnt == 0 {
		p.genParenthesis(curStr+"(", leftCnt+1, rightCnt)
		return
	}
	if leftCnt-rightCnt > 0 {
		p.genParenthesis(curStr+")", leftCnt, rightCnt+1)
		if leftCnt+1 <= p.n {
			p.genParenthesis(curStr+"(", leftCnt+1, rightCnt)
		}
	}
}

/*问题*/
/*
给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

countAndSay(1) = "1"
countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
/*思路*/
/*
一会看看优化点，此处解法就是粗暴的从前往后迭代求结果
每一次迭代中遍历字符串，并计算cnt
*/
func countAndSaySub(s string) string {
	res := ""
	cur := 'a'
	cnt := 0
	r := 'a'
	for _, r = range s {
		if cur != r {
			if cnt == 0 {
				cnt = 1
				cur = r
				continue
			}
			res += fmt.Sprintf("%d%c", cnt, cur)
			cur = r
			cnt = 1
			continue
		}
		cnt++
	}
	if cnt != 0 {
		res += fmt.Sprintf("%d%c", cnt, cur)
	}

	return res
}

func countAndSay(n int) string {
	cur := "1"
	for i := 1; i < n; i++ {
		cur = countAndSaySub(cur)
	}

	return cur
}

/*问题*/
/*
简化路径
*/
/*思路*/
/*
手动维护一个栈，遍历path元素
发现.就继续，..就pop，其他的非空元素就push
*/
func simplifyPath(path string) string {
	resPath := make([]string, 0)
	subPaths := strings.Split(path, "/")
	for _, subPath := range subPaths {
		if subPath == "." || subPath == "" {
			continue
		}

		if subPath == ".." {
			resPath = resPath[:len(resPath)-1]
			continue
		}

		resPath = append(resPath, subPath)
	}
	return "/" + strings.Join(resPath, "/")
}

/*问题*/
/*
输入一个整数，将这个整数以字符串的形式逆序输出
程序不考虑负数的情况，若数字含有0，则逆序形式也含有0，如输入为100，则输出为001


数据范围：
输入描述：
输入一个int整数

输出描述：
将这个整数以字符串的形式逆序输出

示例1
输入：
1516000
复制
输出：
0006151
复制
示例2
输入：
0
复制
输出：
0
复制
https://www.nowcoder.com/practice/ae809795fca34687a48b172186e3dafe?tpId=37&&tqId=21234&rp=1&ru=/ta/huawei&qru=/ta/huawei/question-ranking

*/
/*思路*/
/*
逆转数字，十进制逐个处理
*/
func ReverseNum(num int) string {
	res := ""
	for {
		left := num % 10
		num = num / 10
		res += fmt.Sprintf("%d", left)
		if num == 0 {
			break
		}
	}
	return res
}

/*问题*/
/*
使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
如果字符串的长度为 1 ，算法停止
如果字符串的长度 > 1 ，执行下述步骤：
在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
给你两个 长度相等 的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。如果是，返回 true ；否则，返回 false 。



示例 1：

输入：s1 = "great", s2 = "rgeat"
输出：true
解释：s1 上可能发生的一种情形是：
"great" --> "gr/eat" // 在一个随机下标处分割得到两个子字符串
"gr/eat" --> "gr/eat" // 随机决定：「保持这两个子字符串的顺序不变」
"gr/eat" --> "g/r / e/at" // 在子字符串上递归执行此算法。两个子字符串分别在随机下标处进行一轮分割
"g/r / e/at" --> "r/g / e/at" // 随机决定：第一组「交换两个子字符串」，第二组「保持这两个子字符串的顺序不变」
"r/g / e/at" --> "r/g / e/ a/t" // 继续递归执行此算法，将 "at" 分割得到 "a/t"
"r/g / e/ a/t" --> "r/g / e/ a/t" // 随机决定：「保持这两个子字符串的顺序不变」
算法终止，结果字符串和 s2 相同，都是 "rgeat"
这是一种能够扰乱 s1 得到 s2 的情形，可以认为 s2 是 s1 的扰乱字符串，返回 true
示例 2：

输入：s1 = "abcde", s2 = "caebd"
输出：false
示例 3：

输入：s1 = "a", s2 = "a"
输出：true


提示：

s1.length == s2.length
1 <= s1.length <= 30
s1 和 s2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/scramble-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*思路*/
/*
great rgeat
找到一个边界，满足能将完整数组分成两部分，保证左右包含的内容相同, 边界不能为0或者len(s1)
然后迭代进行
*/
func isHalfEqual(a, b string) bool {
	aa := []rune(a)
	bb := []rune(b)
	sort.Slice(aa, func(i, j int) bool {
		return aa[i] < aa[j]
	})
	sort.Slice(bb, func(i, j int) bool {
		return bb[i] < bb[j]
	})
	for i, r := range aa {
		if r != bb[i] {
			return false
		}
	}
	return true
}

func isScramble(s1 string, s2 string) bool {
	if len(s1) == 1 {
		return s1 == s2
	}
	length := len(s1)
	for i := 1; i < len(s1); i++ {
		if isHalfEqual(s1[0:i], s2[0:i]) &&
			isHalfEqual(s1[i:], s2[i:]) &&
			isScramble(s1[0:i], s2[0:i]) &&
			isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isHalfEqual(s1[0:i], s2[length-i:]) &&
			isHalfEqual(s1[i:], s2[0:length-i]) &&
			isScramble(s1[0:i], s2[length-i:]) &&
			isScramble(s1[i:], s2[0:length-i]) {
			return true
		}
	}
	return false
}
