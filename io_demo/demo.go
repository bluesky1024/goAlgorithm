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
	// io demo
	// fixedInput()
	// fullLineInput()
	// multiLineInput()

	// real instance
	// IOMaxDrinkCnt()
	// IOUniqAndSortForRand()
	// IOSixteenToTen()
	//IOTeacherOpts()
	IOStatErrorFile()
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
