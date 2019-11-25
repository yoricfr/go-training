package main

import "fmt"

type Board struct {
	nailsNeeded int
	nailsDriven int
}

type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

type NailDriverPuller interface {
	NailDriver
	NailPuller
}

type Mallet struct {
}

func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.nailsDriven++
	fmt.Println("Mallet: pounded nail into the board.")
}

type Crowbar struct{}

func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.nailsDriven--
	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

type Contractor struct{}

func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
	for b.nailsDriven < b.nailsNeeded {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(d NailPuller, nailSupply *int, b *Board) {
	for b.nailsDriven > b.nailsNeeded {
		d.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp NailDriverPuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("Contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.nailsDriven < b.nailsNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.nailsDriven > b.nailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

type ToolBox struct {
	NailDriver
	NailPuller
	nails int
}

func main() {
	boards := []Board{
		{nailsDriven: 3},
		{nailsDriven: 1},
		{nailsDriven: 6},

		{nailsNeeded: 6},
		{nailsNeeded: 9},
		{nailsNeeded: 4},
	}

	tb := ToolBox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	displayState(&tb, boards)

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)

	displayState(&tb, boards)
}

func displayState(tb *ToolBox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}
	fmt.Println()
}
