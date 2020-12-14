package day3

var (
	tree    = '#'
	notTree = '.'
)

type slope struct {
	x int
	y int
}

func calculateTrees(input []string, s slope) int {
	x, y := 0, 0
	height, width := len(input), len(input[0])
	var treeCount int
	for y < height-1 {
		if x+s.x >= width {
			x = x - width
		}
		x, y = x+s.x, y+s.y
		if rune(input[y][x]) == tree {
			treeCount++
		}
	}
	return treeCount
}
