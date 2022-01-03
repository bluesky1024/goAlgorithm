// 牛客网的测试需要自行处理输入输出模块
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// IOStr()
	IOForg()

	// io demo
	// fixedInput()
	// return
	// fullLineInput()
	// multiLineInput()
	// mixInput()
	// return

	// real instance
	// IOMaxDrinkCnt()
	// IOUniqAndSortForRand()
	// IOSixteenToTen()
	// IOTeacherOpts()
	// IOStatErrorFile()
}

func IOStr() {
	input := bufio.NewScanner(os.Stdin)
	for {
		if !input.Scan() || input.Text() == "" {
			break
		}
		strs := strings.Split(input.Text(), " ")
		newStr := splitStrs(strs[1:])
		mSort := &sortStr{
			s: newStr,
		}
		sort.Sort(mSort)
		// 打印结果
		for _, str := range mSort.s {
			fmt.Printf("%s ", str)
		}
	}
}

func splitStrs(strs []string) []string {
	res := make([]string, 0)
	for _, str := range strs {
		if len(str) == 8 {
			res = append(res, str)
			continue
		}
		if len(str) < 8 {
			res = append(res, str+strings.Repeat("0", 8-len(str)))
			continue
		}
		curInd := 0
		for {
			if curInd >= len(str) {
				break
			}
			tmp := make([]rune, 8)
			if curInd+8 <= len(str) {
				copy(tmp, []rune(str[curInd:curInd+8]))
				curInd = curInd + 8
				res = append(res, string(tmp))
				continue
			}
			for i := 0; i < 8; i++ {
				tmp[i] = '0'
			}
			copy(tmp, []rune(str[curInd:]))
			res = append(res, string(tmp))
			curInd = curInd + 8
		}
	}
	return res
}

type sortStr struct {
	s []string
}

func (s *sortStr) Len() int {
	return len(s.s)
}

func (s *sortStr) Less(i, j int) bool {
	for ind := 0; ind < 8; ind++ {
		if s.s[i][ind] < s.s[j][ind] {
			return true
		}
		if s.s[i][ind] > s.s[j][ind] {
			return false
		}
	}
	return false
}

func (s *sortStr) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func IOForg() {
	l := 0
	s, t, m := 0, 0, 0
	stones := make(map[int]bool, 0)
	input := bufio.NewScanner(os.Stdin)
	for {
		// 读取l
		if !input.Scan() || input.Text() == "" {
			break
		}
		l, _ = strconv.Atoi(input.Text())
		// 读取 s t m
		if !input.Scan() {
			break
		}
		stms := strings.Split(input.Text(), " ")
		if len(stms) != 3 {
			break
		}
		s, _ = strconv.Atoi(stms[0])
		t, _ = strconv.Atoi(stms[1])
		m, _ = strconv.Atoi(stms[2])
		// 读取 stones
		if !input.Scan() {
			break
		}
		stonesStr := strings.Split(input.Text(), " ")
		if len(stonesStr) != m {
			break
		}
		stones = make(map[int]bool, len(stonesStr))
		for _, stone := range stonesStr {
			tmp, _ := strconv.Atoi(stone)
			stones[tmp] = true
		}

		// 输出结果
		fmt.Println(Forg(l, s, t, stones))
	}
}

// 动态规划处理
// l >= 1
// 某些位置不让走，则这些位置设置为空即可,不考虑从这些点继续往后走
// steps[i] = min{ steps[i-s], steps[i-s+1], step[i-s+2],... step[i-t] }
func Forg(l, s, t int, stones map[int]bool) int {
	steps := make([]int, l+t)
	for i := range steps {
		steps[i] = -1
	}
	steps[0] = 0
	for i := 1; i <= l+t-1; i++ {
		minPos := -1
		for ii := s; ii <= t; ii++ {
			lastPos := i - ii
			// 若 lastPos < 0,说明该位置走不到
			if lastPos < 0 {
				continue
			}
			// 若 lastPos == 0，说明该位置一步就能到
			if lastPos == 0 {
				minPos = 0
				break
			}
			// 若lastPos > 0,且 steps[lastPos] == -1 ，说明该位置永远跳不到
			if lastPos > 0 && steps[lastPos] == -1 {
				continue
			}
			// 若 lastPos > 0, 则应该是从该位置用一步跳过去
			if minPos == -1 {
				minPos = lastPos
				continue
			}
			// 若本轮迭代的 lastPos 的步数更少，则用该轮的
			if steps[minPos] > steps[lastPos] {
				minPos = lastPos
			}
		}
		if minPos == -1 {
			continue
		}
		steps[i] = steps[minPos]
		if _, ok := stones[i]; ok {
			steps[i] = steps[i] + 1
		}
	}

	min := steps[l]
	for i := l + 1; i < l+t-1; i++ {
		if min > steps[i] {
			min = steps[i]
		}
	}

	return min
}

// 固定长度的输入处理
func fixedInput() {
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF {
			break
		}
		fmt.Println(a + b)
	}
}

// 逐行读取输入
func fullLineInput() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		nums := strings.Split(input.Text(), " ")
		res := 0
		for i := range nums {
			num, _ := strconv.Atoi(nums[i])
			res += num
		}
		fmt.Println(res)
	}
}

// 多种不同数据
func multiLineInput() {
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	cnt, _ := strconv.Atoi(input.Text())
	for cnt > 0 {
		if !input.Scan() {
			return
		}
		p1 := input.Text()
		if !input.Scan() {
			return
		}
		p2 := input.Text()
		fmt.Println(len(p1) + len(p2))
		cnt--
	}
}

func mixInput() {
	cnt := 0
	fmt.Scanf("%d", &cnt)
	fmt.Printf("get cnt:%d\n", cnt)
	input := bufio.NewScanner(os.Stdin)
	if !input.Scan() {
		return
	}
	line2 := input.Text()
	fmt.Printf("2nd line:%s\n", line2)
	if !input.Scan() {
		return
	}
	line3 := input.Text()
	fmt.Printf("3rd line:%s\n", line3)
}

func IOMaxDrinkCnt() {
	cnt := 0
	for {
		_, err := fmt.Scan(&cnt)
		if err == io.EOF || cnt == 0 {
			break
		}
		fmt.Println(maxDrinkCnt(0, cnt))
	}
}

func maxDrinkCnt(curDrink int, bottleCnt int) int {
	if bottleCnt < 2 {
		return curDrink
	}
	if bottleCnt == 2 {
		return curDrink + 1
	}
	return maxDrinkCnt(curDrink+bottleCnt/3, bottleCnt%3+bottleCnt/3)
}

func IOUniqAndSortForRand() {
	for {
		cnt := 0
		_, err := fmt.Scan(&cnt)
		if err == io.EOF {
			return
		}

		inputs := make([]bool, 1000)
		for i := 0; i < cnt; i++ {
			tmp := 0
			_, err := fmt.Scan(&tmp)
			if err == io.EOF {
				return
			}
			inputs[tmp-1] = true
		}

		for i, v := range inputs {
			if v {
				fmt.Println(i + 1)
			}
		}
	}
}

func IOSixteenToTen() {
	chToNum := func(c rune) float64 {
		if c >= '0' && c <= '9' {
			return float64(c - '0')
		}
		return float64(c - 'A' + 10)
	}
	//fmt.Println(chToNum('A'))
	//fmt.Println(chToNum('1'))
	//fmt.Println(chToNum('0'))
	//fmt.Println(chToNum('E'))
	//return
	input := ""
	for {
		_, err := fmt.Scan(&input)
		if err == io.EOF {
			return
		}
		input = strings.ToUpper(input)
		numsPart := input[2:]
		res := 0
		for i, v := range numsPart {
			res += int(chToNum(v) * math.Pow(16, float64(len(numsPart)-1-i)))
		}
		fmt.Println(res)
	}
}

func IOTeacherOpts() {
	findMax := func(grades []int) int {
		max := 0
		for _, grade := range grades {
			if max < grade {
				max = grade
			}
		}
		return max
	}

	for {
		sCnt := 0
		optCnt := 0
		_, err := fmt.Scan(&sCnt, &optCnt)
		if err == io.EOF {
			return
		}

		// 初试成绩
		input := bufio.NewScanner(os.Stdin)
		if !input.Scan() {
			return
		}
		originGradesStrSlice := strings.Split(input.Text(), " ")
		sGrades := make([]int, len(originGradesStrSlice))
		for i := 0; i < len(originGradesStrSlice); i++ {
			sGrades[i], _ = strconv.Atoi(originGradesStrSlice[i])
		}

		// 遍历操作
		opt := ""
		a := 0
		b := 0
		for i := 0; i < optCnt; i++ {
			_, err := fmt.Scan(&opt, &a, &b)
			if err == io.EOF {
				return
			}
			switch opt {
			case "Q":
				if a > b {
					a, b = b, a
				}
				fmt.Println(findMax(sGrades[a-1 : b]))
			case "U":
				sGrades[a-1] = b
			default:
				return
			}
		}
	}
}

type fileRecord struct {
	name string
	line int
	cnt  int
}

func IOStatErrorFile() {
	records := make([]fileRecord, 0)
	recordMap := make(map[string]map[int]int)
	fullPath := ""
	line := 0
	for {
		_, err := fmt.Scan(&fullPath, &line)
		if err == io.EOF {
			break
		}
		// 去除磁盘命名
		pathsWithDisk := strings.Split(fullPath, `:`)
		// 去除前置路径
		paths := strings.Split(pathsWithDisk[len(pathsWithDisk)-1], `\`)
		fileName := paths[len(paths)-1]
		if _, ok := recordMap[fileName]; !ok {
			recordMap[fileName] = make(map[int]int)
		}
		if _, ok := recordMap[fileName][line]; !ok {
			records = append(records, fileRecord{
				name: fileName,
				line: line,
				cnt:  1,
			})
			recordMap[fileName][line] = len(records) - 1
		} else {
			records[recordMap[fileName][line]].cnt = records[recordMap[fileName][line]].cnt + 1
		}
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].cnt < records[j].cnt
	})

	for i, record := range records {
		if i >= 8 {
			break
		}
		if len(record.name) > 16 {
			record.name = record.name[len(record.name)-16:]
		}
		fmt.Printf("%s %d %d\n", record.name, record.line, record.cnt)
	}
}
