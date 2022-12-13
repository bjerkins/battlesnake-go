package main

type safeMoves map[string]bool

func updateSafeMoves(moves safeMoves, head Coord, bodyPart Coord) safeMoves {
	// head is below a body part, don't move up
	if bodyPart.X == head.X && bodyPart.Y == head.Y+1 {
		moves["up"] = false
	}
	// head is above a body part, don't move down
	if bodyPart.X == head.X && bodyPart.Y == head.Y-1 {
		moves["down"] = false
	}

	// head is left of a body part, don't move left
	if bodyPart.X == head.X-1 && bodyPart.Y == head.Y {
		moves["left"] = false
	}

	// head is right of a body part, don't move right
	if bodyPart.X == head.X+1 && bodyPart.Y == head.Y {
		moves["right"] = false
	}
	return moves
}
