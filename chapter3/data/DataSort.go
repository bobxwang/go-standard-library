package data

import (
	"fmt"
	"sort"
)

type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

// Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func DataSort() {

	stus := StuScores{
		{"abc", 89},
		{"def", 81},
		{"ghi", 100},
		{"jkl", 56},
	}
	fmt.Println("Default:\n\t", stus)
	sort.Sort(stus)
	fmt.Println("Is Sorted?\n\t", sort.IsSorted(stus))
	fmt.Println("Sorted:\n\t", stus)

	x := 11
	si := []int{3, 6, 9, 11, 45}
	// Search 本身要求列表排好序
	pos := sort.Search(len(si), func(i int) bool { return si[i] >= x })
	if pos < len(si) && si[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素", x)
	}

	var s string
	fmt.Printf("Pick an int from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is you number <= %d?", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your answer is %d.\n", answer)
}
