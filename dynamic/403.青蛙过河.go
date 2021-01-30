package leetcode403

/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 *
 * https://leetcode-cn.com/problems/frog-jump/description/
 *
 * algorithms
 * Hard (39.24%)
 * Likes:    138
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 31.2K
 * Testcase Example:  '[0,1,3,5,6,8,12,17]'
 *
 * 一只青蛙想要过河。 假定河流被等分为 x 个单元格，并且在每一个单元格内都有可能放有一石子（也有可能没有）。 青蛙可以跳上石头，但是不可以跳入水中。
 *
 * 给定石子的位置列表（用单元格序号升序表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一个石子上）。 开始时，
 * 青蛙默认已站在第一个石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格1跳至单元格2）。
 *
 * 如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1个单位。
 * 另请注意，青蛙只能向前方（终点的方向）跳跃。
 *
 * 请注意：
 *
 *
 * 石子的数量 ≥ 2 且 < 1100；
 * 每一个石子的位置序号都是一个非负整数，且其 < 2^31；
 * 第一个石子的位置永远是0。
 *
 *
 * 示例 1:
 *
 *
 * [0,1,3,5,6,8,12,17]
 *
 * 总共有8个石子。
 * 第一个石子处于序号为0的单元格的位置, 第二个石子处于序号为1的单元格的位置,
 * 第三个石子在序号为3的单元格的位置， 以此定义整个数组...
 * 最后一个石子处于序号为17的单元格的位置。
 *
 * 返回 true。即青蛙可以成功过河，按照如下方案跳跃：
 * 跳1个单位到第2块石子, 然后跳2个单位到第3块石子, 接着
 * 跳2个单位到第4块石子, 然后跳3个单位到第6块石子,
 * 跳4个单位到第7块石子, 最后，跳5个单位到第8个石子（即最后一块石子）。
 *
 *
 * 示例 2:
 *
 *
 * [0,1,2,3,4,8,9,11]
 *
 * 返回 false。青蛙没有办法过河。
 * 这是因为第5和第6个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。
 *
 *
 */

// @lc code=start
func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}
	m := make(map[int]map[int]bool)

	for _, v := range stones {
		// map key为stone的序号, vaule为一个[int]bool, key为青蛙在当前stone可以跳的步数
		m[v] = make(map[int]bool)
	}
	m[0][1] = true

	for i := 0; i < len(stones)-1; i++ {
		for step := range m[stones[i]] {
			// 可以跳step步
			reach := stones[i] + step
			if reach == stones[len(stones)-1] {
				return true
			}
			if _, ok := m[reach] ; ok {
				m[reach][step] = true
				if step -1 > 0 {
					m[reach][step-1] = true
				}
				m[reach][step+1] = true
			}
		}
	}
	return false

}

// @lc code=end

func canCross2(stones []int) bool {
	if stones[1] != 1 {
		return false
	}
	m := make(map[int]map[int]int)

	for _, v := range stones {
		// map key为stone的序号, vaule为一个map{k, v}
		// 序号k可以通过v步数到达当前位置序号
		m[v] = make(map[int]int)
	}
	m[1][1] = 1

	for i := 1; i < len(stones)-1; i++ {
		for _, k := range m[stones[i]] {
			// 可以通过k步到达位置i
			// 走j到达下一个位置
			for j := k - 1; j <= k+1; j++ {
				if j > 0 {
					if j+stones[i] == stones[len(stones)-1] {
						return true
					}
					// 走j步所到的位置 是存在的
					if _, ok := m[j+stones[i]]; ok {
						m[j+stones[i]][stones[i]] = j
					}
				}
			}
		}
	}
	return len(m[stones[len(stones)-1]]) > 0

}


func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}

	for i := 2; i < len(stones); i++ {
		// i位置最多跳i+1步
		// 0,1,3,6,10  这样i位置都跳了i+1步, 其他情况都不会超过i+1
		if stones[i] > stones[i-1] + i {
			return false
		}
	}
	position := []int{}
	jump := []int{}
	positionMap := make(map[int]struct{})

	for _, v := range stones {
		positionMap[v] = struct{}{}
	}

	position = append(position, stones[0])
	jump = append(jump, 0)

	for len(position) > 0 {
		// 这里一定要从尾部取, 这样能尽快遍历位置后面的石块
		pos := position[len(position)-1]
		position = position[:len(position)-1]

		step := jump[len(jump)-1]
		jump = jump[:len(jump)-1]

		for s := step - 1; s <= step+1; s++ {
			if s > 0 {
				reach := pos + s
				if reach == stones[len(stones)-1] {
					return true
				}
				if _, ok := positionMap[reach]; ok {
					position = append(position, reach)
					jump = append(jump, s)
				}
			}
		}
	}
	return false
}