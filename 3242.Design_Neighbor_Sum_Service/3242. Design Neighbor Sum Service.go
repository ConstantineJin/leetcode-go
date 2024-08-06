package main

type neighborSum struct {
	grid [][]int
}

func Constructor(grid [][]int) neighborSum {
	return neighborSum{grid}
}

func (this *neighborSum) AdjacentSum(value int) (ans int) {
	for i, row := range this.grid {
		for j, v := range row {
			if v == value {
				if i > 0 {
					ans += this.grid[i-1][j]
				}
				if i < len(this.grid)-1 {
					ans += this.grid[i+1][j]
				}
				if j > 0 {
					ans += this.grid[i][j-1]
				}
				if j < len(this.grid)-1 {
					ans += this.grid[i][j+1]
				}
				return
			}
		}
	}
	return
}

func (this *neighborSum) DiagonalSum(value int) (ans int) {
	for i, row := range this.grid {
		for j, v := range row {
			if v == value {
				if i > 0 && j > 0 {
					ans += this.grid[i-1][j-1]
				}
				if i < len(this.grid)-1 && j > 0 {
					ans += this.grid[i+1][j-1]
				}
				if i > 0 && j < len(this.grid)-1 {
					ans += this.grid[i-1][j+1]
				}
				if i < len(this.grid)-1 && j < len(this.grid)-1 {
					ans += this.grid[i+1][j+1]
				}
				return
			}
		}
	}
	return
}

/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
