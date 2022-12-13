package main

// Welcome to
// __________         __    __  .__                               __
// \______   \_____ _/  |__/  |_|  |   ____   ______ ____ _____  |  | __ ____
//  |    |  _/\__  \\   __\   __\  | _/ __ \ /  ___//    \\__  \ |  |/ // __ \
//  |    |   \ / __ \|  |  |  | |  |_\  ___/ \___ \|   |  \/ __ \|    <\  ___/
//  |________/(______/__|  |__| |____/\_____>______>___|__(______/__|__\\_____>
//
// This file can be a nice home for your Battlesnake logic and helper functions.
//
// To get you started we've included code to prevent your Battlesnake from moving backwards.
// For more info see docs.battlesnake.com

import (
	"log"
	"math"
	"math/rand"
)

// info is called when you create your Battlesnake on play.battlesnake.com
// and controls your Battlesnake's appearance
// TIP: If you open your Battlesnake URL in a browser you should see this data
func info() BattlesnakeInfoResponse {
	log.Println("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Bjarki",  // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}

// start is called when your Battlesnake begins a game
func start(state GameState) {
	log.Println("GAME START")
}

// end is called when your Battlesnake finishes a game
func end(state GameState) {
	log.Printf("GAME OVER\n\n")
}

// move is called on every turn and returns your next move
// Valid moves are "up", "down", "left", or "right"
// See https://docs.battlesnake.com/api/example-move for available data
func move(state GameState) BattlesnakeMoveResponse {

	isMoveSafe := safeMoves{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// We've included code to prevent your Battlesnake from moving backwards
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of your "neck"

	if myNeck.X < myHead.X { // Neck is left of head, don't move left
		isMoveSafe["left"] = false

	} else if myNeck.X > myHead.X { // Neck is right of head, don't move right
		isMoveSafe["right"] = false

	} else if myNeck.Y < myHead.Y { // Neck is below head, don't move down
		isMoveSafe["down"] = false

	} else if myNeck.Y > myHead.Y { // Neck is above head, don't move up
		isMoveSafe["up"] = false
	}

	// TODO: Step 1 - Prevent your Battlesnake from moving out of bounds
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	// if head is at the leftmost edge, don't move left
	if myHead.X == 0 {
		isMoveSafe["left"] = false
	}

	// if head is at the rightmost edge, don't move right
	if myHead.X == boardWidth-1 {
		isMoveSafe["right"] = false
	}

	// if head is at the bottom edge, don't move down
	if myHead.Y == 0 {
		isMoveSafe["down"] = false
	}

	// if head is at the top edge, don't move up
	if myHead.Y == boardHeight-1 {
		isMoveSafe["up"] = false
	}

	// TODO: Step 2 - Prevent your Battlesnake from colliding with itself
	mybody := state.You.Body

	for _, bodyPart := range mybody {
		updateSafeMoves(isMoveSafe, myHead, bodyPart)
	}

	// TODO: Step 3 - Prevent your Battlesnake from colliding with other Battlesnakes
	opponents := state.Board.Snakes

	for _, opponent := range opponents {
		for _, bodyPart := range opponent.Body {
			updateSafeMoves(isMoveSafe, myHead, bodyPart)
		}
	}

	// Are there any safe moves left?
	safeMoves := []string{}
	for move, isSafe := range isMoveSafe {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	log.Printf("safeMoves: %v\n", safeMoves)

	if len(safeMoves) == 0 {
		log.Printf("MOVE %d: No safe moves detected! Moving down\n", state.Turn)
		return BattlesnakeMoveResponse{Move: "down"}
	}

	// TODO: Step 4 - Move towards food instead of random, to regain health and survive longer
	// declare a mutable variable to hold the next move
	var nextMove string

	if state.You.Health < 30 {
		// declare a variable to hold the closest food
		var closestFood Coord
		// declare a variable to hold the shortest distance
		var shortestDistance float64
		// loop through all the food
		for _, food := range state.Board.Food {
			// calculate the distance between the food and the head
			distance := math.Abs(float64(myHead.X-food.X)) + math.Abs(float64(myHead.Y-food.Y))
			// if the distance is shorter than the shortest distance
			if distance < shortestDistance || shortestDistance == 0 {
				// update the shortest distance
				shortestDistance = distance
				// update the closest food
				closestFood = food
			}
		}

		log.Printf("im hungry! grabbing: %v", closestFood)

		if myHead.X < closestFood.X && isMoveSafe["right"] {
			nextMove = "right"
		} else if myHead.X > closestFood.X && isMoveSafe["left"] {
			nextMove = "left"
		} else if myHead.Y < closestFood.Y && isMoveSafe["up"] {
			nextMove = "up"
		} else if myHead.Y > closestFood.Y && isMoveSafe["down"] {
			nextMove = "down"
		}
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	log.Printf("MOVE %d: %s\n", state.Turn, nextMove)
	return BattlesnakeMoveResponse{Move: nextMove}
}

func main() {
	RunServer()
}
