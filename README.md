# go-game-of-life
Simple outline project for playing with game of life in go

http://en.wikipedia.org/wiki/Conway%27s_Game_of_Life for background.

You start with a two dimensional grid of cells, where each cell is either alive or dead. When calcuating the next generation of the grid, follow these rules:

   1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
   2. Any live cell with more than three live neighbours dies, as if by overcrowding.
   3. Any live cell with two or three live neighbours lives on to the next generation.
   4. Any dead cell with exactly three live neighbours becomes a live cell.

Your job is to implement the rules above, with tests, and the play with the result. For dealing with the edges you have 2 choices, either wrap the board, or assume that no life can exist outside of the board.

Don't forget to write some tests :)

Run the game using the commands below. Press escape to exit.

```
go run main.go
```

To start with a simple glider which will move across the grid, you can use
```
go run main.go -glider
```


Some life examples:

Stable life, doesn't change
```
**
**
```

Starting life thatc hanges then becomes stable
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

Starting life that changes then does off
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
