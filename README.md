# go-game-of-life
## Overview
A simple project for playing with game of life in go, used to experiment with
testing and code implementation

http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

You start with a two dimensional grid of cells, where each cell is either alive or dead. When calcuating the next generation of the grid, follow these rules (NB a neighbour includes diagonals):

   1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
   2. Any live cell with more than three live neighbours dies, as if by overcrowding.
   3. Any live cell with two or three live neighbours lives on to the next generation.
   4. Any dead cell with exactly three live neighbours becomes a live cell.

Your job is to implement the rules above, with tests, and the play with the result. Assume that the edge of the board is a hard edge with no life beyond.

**Don't forget to write some tests :)**

## Running
The game includes a simple main and terminal based drawing to visualse the game.


Run the game using the commands below. Press escape to exit.

```
go run main.go
```

There are `-width` and `-height` options, with defaults, to allow the game size to change.


## Some life examples:

### Stable life, doesn't change
```
**
**
```

### Starting life thatc hanges then becomes stable
```
***   *   **    **
*    **   **   *  *
     *    **    **
```

### Starting life that changes then dies off
```
*     **   **    *    **
 ***  **     *    *   
       *   **    *
```

## Extensions

### Refactoring

Can you refactor the code to make it as clean as possible? Try alternative options for calculations etc.

### Wrapping game area

If you have the game working for a static game area, try allowing the game to wrap around the top/bottom and left/right sides.

### Alternative rendering

If you complete this quickly, because you have experience with Game Of Life, then try writing an alternative rendering engine, try https://ebiten.org/ which is a very simple game engine and will allow drawing of graphics.

Or use Fyne https://github.com/fyne-io/life/blob/main/main.go this is their game of life example, but you should use your own game of life implementation.
