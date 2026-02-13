package app

import (
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/terminalapi"
)

// Func newGrin defines the grid layout of the app
func newGrid(b *buttonSet, w *widgets, t terminalapi.Terminal) (*container.Container, error) {
	builder := grid.New()
	//Add first row builder
	builder.Add(

		//Add first row
		grid.RowHeightPerc(30,

			//Add the first column in first row
			grid.ColWidthPercWithOpts(30,
				[]container.Option{
					container.Border(linestyle.Light),
					container.BorderTitle("Press q to quit"),
				},
				//Add the first row in the column
				grid.RowHeightPerc(80,
					grid.Widget(w.donTimer),
				), //End of the first row of the first column
			), //End of the first column of the first row
			//Add the second row in the column
			grid.RowHeightPercWithOpts(20,
				[]container.Option{
					container.AlignHorizontal(align.HorizontalCenter),
				},
				grid.Widget(w.txtTimer,
					container.AlignHorizontal(align.HorizontalCenter),
					container.AlignVertical(align.VerticalMiddle),
					container.PaddingLeftPercent(49),
				),
			), //End of the second row of the
			//Add the second column in the first row
			grid.ColWidthPerc(70,
				//Add the first row in the second coulnm in the first row
				grid.RowHeightPerc(80,
					grid.Widget(w.disType,
						container.Border(linestyle.Light),
					),
				),
				//Add the second row in the second colunm in the first row
				grid.RowHeightPerc(20,
					grid.Widget(w.txtInfo,
						container.Border(linestyle.Light),
					),
				), //End ot the second row of the second column of the first row
			), //End of the second column of the first row
		), //End of the first row
	) //End of builder.Add() for first row

	//Add second row builder
	builder.Add(
		//Add the second row
		grid.RowHeightPerc(10,
			//Add the first column of the second row
			grid.ColWidthPerc(50,
				grid.Widget(b.btnStart),
			), //End of the first column of the second row
			//Add the second column of the second row
			grid.ColWidthPerc(50,
				grid.Widget(b.btnPause),
			), //End of the second column of the second row
		), //End of the second row
	) //End of the builder.Add() of the second row

	//Add third row builder
	builder.Add(
		//Add the third row
		grid.RowHeightPerc(60), // End of the third row
	) // End of the builder.Add() of the third row

	gridOpts, err := builder.Build()
	if err != nil {
		return nil, err
	}

	cont, err := container.New(t, gridOpts...)
	if err != nil {
		return nil, err
	}
	return cont, nil
}
