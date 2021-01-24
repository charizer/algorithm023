package walking_robot_simulation

func RobotSim(commands []int, obstacles [][]int) int {
	// ans为结果，x，y为坐标轴上的位置
	ans, x, y := 0, 0, 0
	// 当前方向
	direct := 0
	// 北： 向Y+方向移动1格， 东：向X+方向移动1格， 南：向Y-方向移动1格， 西：向X-方向移动1格
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1,0}}
	// 记录障碍
	obstacleMap := make(map[[2]int]bool)
	for _, obstacle := range obstacles {
		obstacleMap[[2]int{obstacle[0], obstacle[1]}] = true
	}
	for _, command := range commands {
		switch command {
		// 右转90度
		case -1:
			direct++
			if direct >= 4 {
				direct = 0
			}
		// 左转90度
		case -2:
			direct--
			if direct < 0 {
				direct = 3
			}
		// 在当前方向直接前进
		default:
			// 在当前方向移动，直到遇到障碍
			for ; command > 0 && !obstacleMap[[2]int{x+dirs[direct][0], y+dirs[direct][1]}]; command-- {
				x += dirs[direct][0]
				y += dirs[direct][1]
			}
			// 更新最大值
			if ans < x * x + y * y {
				ans = x * x + y * y
			}
		}
	}
	return ans
}
