package leetcode207

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (54.75%)
 * Likes:    781
 * Dislikes: 0
 * Total Accepted:    105.8K
 * Total Submissions: 193.2K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
 *
 * 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi]
 * ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
 *
 *
 * 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
 *
 *
 * 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0]]
 * 输出：true
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
 *
 * 示例 2：
 *
 *
 * 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
 * 输出：false
 * 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * prerequisites[i].length == 2
 * 0 i, bi < numCourses
 * prerequisites[i] 中的所有课程对 互不相同
 *
 *
 */

// @lc code=start
// DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)
	valid := true

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	var dfs func(int)

	dfs = func(n int) {
		visited[n] = 1
		for _, v:= range graph[n] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
				// 形成环
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		// 这个n节点的所有叶子节点都遍历过了/ 或者其没有叶子节点
		visited[n] = 2
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid

}

// @lc code=end

// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)

	for _, v := range prerequisites {
		indegree[v[0]] += 1
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	queue := []int{}
	res := []int{}

	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		res = append(res, top)

		for _, v := range graph[top] {
			indegree[v] = indegree[v] - 1
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(res) == numCourses

}
