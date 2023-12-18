package leetcode

import "sort"

type chair struct {
	Left   *chair
	Right  *chair
	Number int
}

type chairs struct {
	Occupied *chair
}

func (c *chairs) Sit() int {
	if c.Occupied == nil {
		c.Occupied = &chair{
			Number: 0,
		}
		return 0
	}
	ch := c.Occupied
	var num int
	for num = 0; ch != nil; num++ {
		if num < ch.Number {
			newChair := &chair{
				Number: num,
				Right:  ch,
			}
			if ch.Left == nil {
				ch.Left = newChair
				c.Occupied = ch.Left
			} else {
				ch.Left.Right = newChair
			}
			break
		}
		ch = ch.Right
	}
	return num
}

func (c *chairs) Leave(number int) {
	if c.Occupied != nil {
		ch := c.Occupied
		for ch != nil {
			if ch.Number == number {
				if ch.Left != nil {
					ch.Left.Right = ch.Right
				} else {
					c.Occupied = ch.Right
				}
				break
			}
		}
	}
}

func SmallestChair(times [][]int, targetFriend int) int {
	timeSeries := make([][]int, 0)
	for friend, t := range times {
		timeSeries = append(timeSeries, []int{t[0], friend}, []int{-t[1], friend})
	}
	sort.Slice(timeSeries, func(i, j int) bool {
		ti := timeSeries[i][0]
		tj := timeSeries[j][0]
		if ti >= 0 && tj < 0 {
			return ti < -tj
		}
		if ti < 0 && tj >= 0 {
			return -ti < tj
		}
		if ti < 0 && tj < 0 {
			return -ti < -tj
		}
		return ti < tj
	})
	positions := &chairs{}
	friendSits := make(map[int]int, 0)
	for _, t := range timeSeries {
		friend := t[1]
		if t[0] >= 0 {
			possition := positions.Sit()
			friendSits[friend] = possition
			if friend == targetFriend {
				return possition
			}
		} else {
			positions.Leave(friendSits[friend])
		}
	}
	return 0
}
