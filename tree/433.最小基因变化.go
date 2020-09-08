package leetcode_433

/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (52.19%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    7.7K
 * Total Submissions: 14.8K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
 *
 * 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
 *
 * 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
 *
 * 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
 *
 * 现在给定3个参数 — start, end,
 * bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回
 * -1。
 *
 * 注意:
 *
 *
 * 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
 * 所有的目标基因序列必须是合法的。
 * 假定起始基因序列与目标基因序列是不一样的。
 *
 *
 * 示例 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * 返回值: 1
 *
 *
 * 示例 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * 返回值: 2
 *
 *
 * 示例 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * 返回值: 3
 *
 *
 */

// @lc code=start
type GeneNode struct {
	Val  string
	Step int
}

func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}
	diffTwoString := func(one, two string) int {
		diff := 0
		for i, l := 0, len(one); i < l; i++ {
			if one[i] != two[i] {
				diff++
			}
		}
		return diff
	}
	queue := make([]*GeneNode, 0, 10)
	queue = append(queue, &GeneNode{start, 0})

	visited := make(map[int]bool)

	// 判断bank中是否含有end, 若没有直接return -1
	bankContainEnd := false
	// BFS
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		for i, b := range bank {
			if ok := visited[i]; ok {
				continue
			}
			if !bankContainEnd && b == end {
				bankContainEnd = true
			}
			// 找到只需要突变1次的节点
			if diffTwoString(head.Val, b) == 1 {
				// 这些节点后续不需要访问, 已经不是我们的目标节点
				visited[i] = true
				if b == end {
					return head.Step + 1
				}
				queue = append(queue, &GeneNode{b, head.Step + 1})
			}
		}
		if !bankContainEnd {
			return -1
		}
	}
	return -1
}

// @lc code=end
