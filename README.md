# go-game-of-life
Simple outline project for playing with game of life in go

http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life for background.

You start with a two dimensional grid of cells, where each cell is either alive or dead. In this version of the problem, the grid is finite, and no life can exist off the edges. When calcuating the next generation of the grid, follow these rules:

   1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
   2. Any live cell with more than three live neighbours dies, as if by overcrowding.
   3. Any live cell with two or three live neighbours lives on to the next generation.
   4. Any dead cell with exactly three live neighbours becomes a live cell.

Your job is to implement the rules above, with tests, and the play with the result.

There is a simple main that can be used to run and display the game of line in a terminal. This relies on `github.com/nsf/termbox-go` so make sure you run `go get` before running the exercise.

Some life examples:

Stable life, doesn't change
```
**
**
```

Changes then becomes stable
```
***
*
```
```
 *
**
*
```
```
**
**
**
```
```
 ** 
*  *
 **
```

Changes then Dies
```
*
 ***
```
```
**
**
 *
```
```
**
  *
**
```
```
*
 *
*
```
```
**
```
```
Death
```
