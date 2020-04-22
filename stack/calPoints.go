package stack

import "strconv"

/*
682. 棒球比赛
https://leetcode-cn.com/problems/baseball-game/
你现在是棒球比赛记录员。
给定一个字符串列表，每个字符串可以是以下四种类型之一：
1.整数（一轮的得分）：直接表示您在本轮中获得的积分数。
2. "+"（一轮的得分）：表示本轮获得的得分是前两轮有效 回合得分的总和。
3. "D"（一轮的得分）：表示本轮获得的得分是前一轮有效 回合得分的两倍。
4. "C"（一个操作，这不是一个回合的分数）：表示您获得的最后一个有效 回合的分数是无效的，应该被移除。

每一轮的操作都是永久性的，可能会对前一轮和后一轮产生影响。
你需要返回你在所有回合中得分的总和。

示例 1:

输入: ["5","2","C","D","+"]
输出: 30
解释:
第1轮：你可以得到5分。总和是：5。
第2轮：你可以得到2分。总和是：7。
操作1：第2轮的数据无效。总和是：5。
第3轮：你可以得到10分（第2轮的数据已被删除）。总数是：15。
第4轮：你可以得到5 + 10 = 15分。总数是：30。
 */
func CalPoints(ops []string) int {
	length := len(ops)
	s := Stack{}
	for i := 0; i < length; i++ {
		switch ops[i] {
		case "+":
			s.Push(s.TopTwoSum())
		case "D":
			s.Push(s.Top()*2)
		case "C":
			s.Pop()
		default:
			temp, _ := strconv.Atoi(ops[i])
			s.Push(temp)
		}
	}
	result := 0
	for s.Len() != 0 {
		result += s.Pop()
	}
	return result
}

type Stack []int

func (s Stack) TopTwoSum() int {
	if len(s) - 2 >= 0 {
		return s[len(s)-1] + s[len(s)-2]
	} else if len(s) - 1 >= 0 {
		return s[len(s)-1]
	} else {
		return 0
	}
}

func (s Stack) Top() int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s)-1]
}

func (s Stack) Len() int {
	return len(s)
}

func (s *Stack) Push(i int) {
	old := *s
	*s = append(old, i)
}

func (s *Stack) Pop() int {
	old := *s
	*s = old[:len(old)-1]
	return old[len(old)-1]
}