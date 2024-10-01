package cmd

import (
	log "github.com/sirupsen/logrus"
)

type node struct {
	name           string
	drawing        *drawing
	drawingCoord   *drawingCoord
	gridCoord      *gridCoord
	drawn          bool
	index          int // Index of the node in the graph.nodes slice
	styleClassName string
	styleClass     styleClass
}

func (n node) String() string {
	return n.name
}

func (n *node) setCoord(c *drawingCoord) {
	n.drawingCoord = c
}

func (n *node) setDrawing() *drawing {
	d := drawBox(n)
	n.drawing = d
	return d
}

func (g *graph) setColumnWidth(n *node) {
	// For every node there are three columns:
	// - 2 lines of border
	// - 1 line of text
	// - 2x padding
	// - 2x margin
	col1 := 1
	col2 := 2*boxBorderPadding + len(n.name)
	col3 := 1
	colsToBePlaced := []int{col1, col2, col3}
	rowsToBePlaced := []int{1, 1 + 2*boxBorderPadding, 1} // Border, padding + line, border

	for idx, col := range colsToBePlaced {
		// Set new width for column if the size increased
		xCoord := n.gridCoord.x + idx
		g.columnWidth[xCoord] = Max(g.columnWidth[xCoord], col)
	}

	for idx, row := range rowsToBePlaced {
		// Set new width for column if the size increased
		yCoord := n.gridCoord.y + idx
		g.rowHeight[yCoord] = Max(g.rowHeight[yCoord], row)
	}

	// Set padding before node
	if n.gridCoord.x > 0 {
		g.columnWidth[n.gridCoord.x-1] = paddingBetweenX // TODO: x2?
	}
	if n.gridCoord.y > 0 {
		g.rowHeight[n.gridCoord.y-1] = paddingBetweenY // TODO: x2?
	}
}

func (g *graph) reserveSpotInGrid(n *node, requestedCoord *gridCoord) *gridCoord {
	if g.grid[*requestedCoord] != nil {
		log.Debugf("Coord %d,%d is already taken", requestedCoord.x, requestedCoord.y)
		// Next column is 4 coords further. This is because every node is 3 coords wide + 1 coord inbetween.
		if graphDirection == "LR" {
			return g.reserveSpotInGrid(n, &gridCoord{x: requestedCoord.x, y: requestedCoord.y + 4})
		} else {
			return g.reserveSpotInGrid(n, &gridCoord{x: requestedCoord.x + 4, y: requestedCoord.y})
		}
	}
	// Reserve border + middle + border for node
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			reservedCoord := gridCoord{x: requestedCoord.x + x, y: requestedCoord.y + y}
			log.Debugf("Reserving coord %d,%d for node %v", reservedCoord.x, reservedCoord.y, n)
			g.grid[reservedCoord] = n
		}
	}
	n.gridCoord = requestedCoord
	return requestedCoord
}
