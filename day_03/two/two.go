package two

// Direction represents a cardinal map direction
type Direction int

// Cardinal directions
const (
	North Direction = 0
	West  Direction = 1
	South Direction = 2
	East  Direction = 3
)

// Position tracks both the coordinate and the heading
type Position struct {
	coordinate Coordinate
	heading    Direction
}

// Coordinate identifies a single point on the grid
type Coordinate struct {
	x int
	y int
}

// Grid is the problem space for Day 03
type Grid map[Coordinate]int

// CountTo Counts up in a spiral grid to a certain value, returns the first value higher
func CountTo(a int) (r int) {
	grid := make(Grid)
	grid[Coordinate{0, 0}] = 1
	grid[Coordinate{1, 0}] = 1
	pos := Position{Coordinate{1, 0}, East}
	for {
		// Calculate Next Position
		pos = move(grid, pos)
		// Calculate Next Value
		r = calculateVal(grid, pos)
		// Store in Grid
		grid[pos.coordinate] = r
		// Return if greater than input
		if r > a {
			return
		}
	}
}

func move(grid Grid, pos Position) Position {
	switch pos.heading {
	case North:
		if grid[Coordinate{pos.coordinate.x - 1, pos.coordinate.y}] == 0 {
			return Position{Coordinate{pos.coordinate.x - 1, pos.coordinate.y}, West}
		}
		return Position{Coordinate{pos.coordinate.x, pos.coordinate.y + 1}, North}
	case West:
		if grid[Coordinate{pos.coordinate.x, pos.coordinate.y - 1}] == 0 {
			return Position{Coordinate{pos.coordinate.x, pos.coordinate.y - 1}, South}
		}
		return Position{Coordinate{pos.coordinate.x - 1, pos.coordinate.y}, West}
	case South:
		if grid[Coordinate{pos.coordinate.x + 1, pos.coordinate.y}] == 0 {
			return Position{Coordinate{pos.coordinate.x + 1, pos.coordinate.y}, East}
		}
		return Position{Coordinate{pos.coordinate.x, pos.coordinate.y - 1}, South}
	case East:
		if grid[Coordinate{pos.coordinate.x, pos.coordinate.y + 1}] == 0 {
			return Position{Coordinate{pos.coordinate.x, pos.coordinate.y + 1}, North}
		}
		return Position{Coordinate{pos.coordinate.x + 1, pos.coordinate.y}, East}
	}
	return pos
}

func calculateVal(grid Grid, pos Position) int {
	return grid[Coordinate{pos.coordinate.x - 1, pos.coordinate.y - 1}] +
		grid[Coordinate{pos.coordinate.x, pos.coordinate.y - 1}] +
		grid[Coordinate{pos.coordinate.x + 1, pos.coordinate.y - 1}] +
		grid[Coordinate{pos.coordinate.x - 1, pos.coordinate.y}] +
		grid[Coordinate{pos.coordinate.x, pos.coordinate.y}] +
		grid[Coordinate{pos.coordinate.x + 1, pos.coordinate.y}] +
		grid[Coordinate{pos.coordinate.x - 1, pos.coordinate.y + 1}] +
		grid[Coordinate{pos.coordinate.x, pos.coordinate.y + 1}] +
		grid[Coordinate{pos.coordinate.x + 1, pos.coordinate.y + 1}]
}
