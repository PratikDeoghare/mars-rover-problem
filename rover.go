package mars_rover_problem

import "fmt"

type vect struct {
	x, y int
}

func add(c1, c2 vect) vect {
	return vect{
		x: c1.x + c2.x,
		y: c1.y + c2.y,
	}
}

func mul(c1, c2 vect) vect {
	return vect{
		x: c1.x*c2.x - c1.y*c2.y,
		y: c1.x*c2.y + c1.y*c2.x,
	}
}

var (
	dirToVect = map[string]vect{
		"N": {0, 1},
		"E": {1, 0},
		"W": {-1, 0},
		"S": {0, -1},
	}

	vectToDir = map[vect]string{
		vect{0, 1}:  "N",
		vect{1, 0}:  "E",
		vect{-1, 0}: "W",
		vect{0, -1}: "S",
	}
)

type rover struct {
	pos vect
	dir vect

	gridlimit vect
}

func newRover(x, y, gridlimitx, gridlimity int, heading string) rover {
	r := rover{}
	r.pos = vect{x: x, y: y}
	r.gridlimit = vect{x: gridlimitx, y: gridlimity}
	r.dir = dirToVect[heading]
	return r
}

func (r rover) String() string {
	return fmt.Sprintf("%d %d %s", r.pos.x, r.pos.y, vectToDir[r.dir])
}

func (r *rover) L() {
	r.dir = mul(r.dir, vect{x: 0, y: 1})
}

func (r *rover) R() {
	r.dir = mul(r.dir, vect{x: 0, y: -1})
}

func (r *rover) M() error {
	nextPos := add(r.pos, r.dir)
	if nextPos.x > r.gridlimit.x || nextPos.y > r.gridlimit.y || nextPos.x < 0 || nextPos.y < 0 {
		return fmt.Errorf("not moving: fatal move %v -> %v", r.pos, nextPos)
	}
	r.pos = nextPos
	return nil
}

func runProg(p string, r rover) (rover, error) {
	for i, ch := range p {
		switch ch {
		case 'R':
			r.R()
		case 'L':
			r.L()
		case 'M':
			err := r.M()
			if err != nil {
				return r, fmt.Errorf("completed %d steps before: %w", i+1, err)
			}
		}
	}
	return r, nil
}
