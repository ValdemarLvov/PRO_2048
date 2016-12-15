// game_2048 project main.go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func TestBuildGrid() {
	grid := Grid{Size: 6, StartCells: 3}
	grid.Build()

	//if grid.Size*grid.Size-len(grid.EmptyCells()) != 5 {
	if grid.Size*grid.Size-len(grid.EmptyCells()) != grid.StartCells {
		fmt.Printf("Fail")
	}
}

func TestNewTile() {
	grid := Grid{Size: 4, StartCells: 6}
	grid.Build()

	grid.MoveLeft()
	prev := len(grid.EmptyCells())

	//	grid.newTile()
	grid.newTile()
	next := len(grid.EmptyCells())

	if prev-next == 0 {
		fmt.Printf("NewTile not created!")
	} else if prev-next > 1 {
		fmt.Printf("Created more than one new Tile!")
	}
}

//Position structure
type pos struct {
	X int
	Y int
}

//Tiles track value and position on the grid
type Tile struct {
	Value int
}

func randVal() *rand.Rand {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r
}

func randTileVal() int {
	if randVal().Float64() > 0.9 {
		return 4
	} else {
		return 2
	}
}

type TileList []*Tile

//Grid tracks the Tiles
type Grid struct {
	Size       int
	StartCells int
	Tiles      TileList
	Cells      [][]*Cell
	Score      int
	maxScore   int
	GameOver   bool
}

type Cell struct {
	Tile *Tile
	Pos  pos
}

//Create the grid
func (g *Grid) Build() {
	g.Cells = make([][]*Cell, g.Size)
	g.Tiles = make(TileList, 0)

	for i := range g.Cells {
		g.Cells[i] = make([]*Cell, g.Size)
		for j := range g.Cells[i] {
			cell := new(Cell)

			cell.Pos = pos{
				X: i,
				Y: j,
			}

			g.Cells[i][j] = cell
		}
	}

	start := g.StartCells

	for start != 0 {
		start = start - 1
		g.newTile()
	}
}

func (g *Grid) newTile() {

	avail := g.EmptyCells()

	if len(avail) == 0 {
		return
	}
	//Two random values between 0 and Grid.Size
	i := randVal().Int() % len(avail)
	cell := avail[i]

	newTile := Tile{
		Value: randTileVal(),
		//    MergeHistory: make(TileList, 0),
		//    Current: cell.Pos,
		//    New: true,
	}

	g.Tiles = append(g.Tiles, &newTile)
	cell.Tile = &newTile

	return
}

func (g *Grid) EmptyCells() (ret []*Cell) {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			c := g.Cells[i][j]
			if c.Tile == nil {
				ret = append(ret, c)
			}
		}
	}

	return ret
}

func (g *Grid) PrintGrid() {
	for i := range g.Cells {
		for j := range g.Cells[i] {
			if g.Cells[i][j].Tile != nil {
				fmt.Printf("%4d|", g.Cells[i][j].Tile.Value)
			} else {
				fmt.Printf("%4s|", " ")
			}
		}
		fmt.Println("\n")
	}
	fmt.Println("-----------------------------------------------")
}

func (g *Grid) MoveLeft() {
	for i := range g.Cells {
		for j := 1; j < g.Size; {

			if g.Cells[i][j-1].Tile == nil {
				k := j
				for k < g.Size {
					if g.Cells[i][k].Tile == nil {
						k++
					} else {
						break
					}
				}

				if k == g.Size {
					j++
					continue
				}

				for l := 0; l+k < g.Size; l++ {
					g.Cells[i][j+l-1].Tile = g.Cells[i][k+l].Tile
				}
				g.Cells[i][g.Size-1].Tile = nil
				//				g.PrintGrid()

			} else {
				for k := j; k < g.Size; {

					if g.Cells[i][k].Tile == nil {
						k++
					} else if g.Cells[i][j-1].Tile.Value == g.Cells[i][k].Tile.Value {
						g.Cells[i][j-1].Tile.Value += g.Cells[i][k].Tile.Value
						g.Cells[i][k].Tile = nil
						break
					} else if j != k {
						g.Cells[i][j].Tile = g.Cells[i][k].Tile
						g.Cells[i][k].Tile = nil
						break
					} else {
						break
					}
				}
				//				g.PrintGrid()
				j++
			}
		}
	}
}

func (g *Grid) MoveRight() {
	for i := range g.Cells {
		for j := g.Size - 2; j >= 0; {

			if g.Cells[i][j+1].Tile == nil {
				k := j
				for k >= 0 {
					if g.Cells[i][k].Tile == nil {
						k--
					} else {
						break
					}
				}

				if k == -1 {
					j--
					continue
				}

				for l := 0; k-l >= 0; l++ {
					g.Cells[i][j-l+1].Tile = g.Cells[i][k-l].Tile
				}
				g.Cells[i][0].Tile = nil
				//				g.PrintGrid()

			} else {
				for k := j; k >= 0; {

					if g.Cells[i][k].Tile == nil {
						k--
					} else if g.Cells[i][j+1].Tile.Value == g.Cells[i][k].Tile.Value {
						g.Cells[i][j+1].Tile.Value += g.Cells[i][k].Tile.Value
						g.Cells[i][k].Tile = nil
						break
					} else if j != k {
						g.Cells[i][j].Tile = g.Cells[i][k].Tile
						g.Cells[i][k].Tile = nil
						break
					} else {
						break
					}
				}
				//				g.PrintGrid()
				j--
			}
		}
	}
}

func (g *Grid) MoveUp() {
	for j := range g.Cells {
		for i := 1; i < g.Size; {

			if g.Cells[i-1][j].Tile == nil {
				k := i
				for k < g.Size {
					if g.Cells[k][j].Tile == nil {
						k++
					} else {
						break
					}
				}

				if k == g.Size {
					i++
					continue
				}

				for l := 0; l+k < g.Size; l++ {
					g.Cells[i+l-1][j].Tile = g.Cells[k+l][j].Tile
				}
				g.Cells[g.Size-1][j].Tile = nil
				//				g.PrintGrid()

			} else {
				for k := i; k < g.Size; {

					if g.Cells[k][j].Tile == nil {
						k++
					} else if g.Cells[i-1][j].Tile.Value == g.Cells[k][j].Tile.Value {
						g.Cells[i-1][j].Tile.Value += g.Cells[k][j].Tile.Value
						g.Cells[k][j].Tile = nil
						break
					} else if i != k {
						g.Cells[i][j].Tile = g.Cells[k][j].Tile
						g.Cells[k][j].Tile = nil
						break
					} else {
						break
					}
				}
				//				g.PrintGrid()
				i++
			}
		}
	}
}

func (g *Grid) MoveDown() {
	for j := range g.Cells {
		for i := g.Size - 2; i >= 0; {

			if g.Cells[i+1][j].Tile == nil {
				k := i
				for k >= 0 {
					if g.Cells[k][j].Tile == nil {
						k--
					} else {
						break
					}
				}

				if k == -1 {
					i--
					continue
				}

				for l := 0; k-l >= 0; l++ {
					g.Cells[i-l+1][j].Tile = g.Cells[k-l][j].Tile
				}
				g.Cells[0][j].Tile = nil
				//				g.PrintGrid()

			} else {
				for k := i; k >= 0; {

					if g.Cells[k][j].Tile == nil {
						k--
					} else if g.Cells[i+1][j].Tile.Value == g.Cells[k][j].Tile.Value {
						g.Cells[i+1][j].Tile.Value += g.Cells[k][j].Tile.Value
						g.Cells[k][j].Tile = nil
						break
					} else if i != k {
						g.Cells[i][j].Tile = g.Cells[k][j].Tile
						g.Cells[k][j].Tile = nil
						break
					} else {
						break
					}
				}
				//				g.PrintGrid()
				i--
			}
		}
	}
}

func CopyGrid(g *Grid) Grid {
	result := Grid{Size: g.Size}
	result.Build()

	for i := 0; i < result.Size; i++ {
		for j := 0; j < result.Size; j++ {
			result.Cells[i][j] = g.Cells[i][j]
		}
	}
	return result
}

func (g *Grid) IsOver() bool {
	tmp_grid := CopyGrid(g)
	tmp_grid.MoveLeft()
	if len(tmp_grid.EmptyCells()) != 0 {
		return false
	}

	tmp_grid = CopyGrid(g)
	tmp_grid.MoveDown()
	if len(tmp_grid.EmptyCells()) != 0 {
		return false
	}

	tmp_grid = CopyGrid(g)
	tmp_grid.MoveRight()
	if len(tmp_grid.EmptyCells()) != 0 {
		return false
	}

	tmp_grid = CopyGrid(g)
	tmp_grid.MoveUp()
	if len(tmp_grid.EmptyCells()) != 0 {
		return false
	}
	return true
}

func main() {
	TestBuildGrid()
	TestNewTile()
	grid := Grid{Size: 3, StartCells: 1}
	grid.Build()

	reader := bufio.NewReader(os.Stdin)

	for {
		ClearCMD()
		grid.newTile()
		grid.PrintGrid()

		if len(grid.EmptyCells()) == 0 {
			if grid.IsOver() {
				break
			}
		}

		text, _ := reader.ReadString('\n')

		if text == "q\n" {
			break
		}

		switch {
		case text == "w\n":
			grid.MoveUp()
			break
		case text == "a\n":
			grid.MoveLeft()
			break
		case text == "s\n":
			grid.MoveDown()
			break
		case text == "d\n":
			grid.MoveRight()
			break
		}
	}
}

func ClearCMD() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
