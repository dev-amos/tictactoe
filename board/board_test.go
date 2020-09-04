package board

import (
	"reflect"
	"testing"
)

var winConditionChecks = generateChecks()

func TestNewBoard(t *testing.T) {
	type args struct {
		winCount   int
		dimensions int
	}

	type want struct {
		err         error
		expectBoard *Board
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"returns error when dimension to create board is negative",
			args{
				3,
				-1,
			},
			want{
				ErrInvalidDimension,
				nil,
			},
		},
		{
			"returns error when dimension to create board is zero",
			args{
				3,
				0,
			},
			want{
				ErrInvalidDimension,
				nil,
			},
		},
		{
			"returns error when winCount to create board is negative",
			args{
				-1,
				3,
			},
			want{
				ErrInvalidWinCondition,
				nil,
			},
		},
		{
			"returns error when winCount to create board is 0",
			args{
				0,
				3,
			},
			want{
				ErrInvalidWinCondition,
				nil,
			},
		},
		{
			"creates an empty board of 3x3 dimension",
			args{
				3,
				3,
			},
			want{
				nil,
				&Board{
					3,
					[][]BoxContent{
						{E, E, E},
						{E, E, E},
						{E, E, E},
					},
					winConditionChecks,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newBoardParams := NewBoardParams{
				WinCount:   test.args.winCount,
				Dimensions: test.args.dimensions,
			}

			getBoard, err := NewBoard(newBoardParams)

			if !reflect.DeepEqual(err, test.want.err) {
				t.Errorf("unexpected error = %v, want %v", err, test.want.err)
			}

			if !reflect.DeepEqual(getBoard, test.want.expectBoard) {
				t.Errorf("unexpected Board = %v, want %v", getBoard, test.want.expectBoard)
			}

		})
	}

}

func TestSelectBox(t *testing.T) {
	type args struct {
		rowIdx  int
		colIdx  int
		content BoxContent
	}

	type fields struct {
		WinCount           int
		Boxes              [][]BoxContent
		WinConditionChecks []winConditionCheck
	}

	type want struct {
		err         error
		expectBoard Board
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   want
	}{
		{
			"returns no error when inserting content into unfilled box",
			args{
				0,
				0,
				X,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, E},
					{E, E, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				nil,
				Board{
					3,
					[][]BoxContent{
						{X, E, E},
						{E, E, E},
						{E, E, E},
					},
					winConditionChecks,
				},
			},
		},
		{
			"returns error when inserting content into filled box",
			args{
				0,
				0,
				X,
			},
			fields{
				3,
				[][]BoxContent{
					{X, E, E},
					{E, E, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				ErrBoxOccupied,
				Board{
					3,
					[][]BoxContent{
						{X, E, E},
						{E, E, E},
						{E, E, E},
					},
					winConditionChecks,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			testBoard := Board{
				WinCount:           test.fields.WinCount,
				Boxes:              test.fields.Boxes,
				WinConditionChecks: test.fields.WinConditionChecks,
			}

			insertBoxWithContentParams := InsertBoxWithContentParams{
				RowIdx:  test.args.rowIdx,
				ColIdx:  test.args.colIdx,
				Content: test.args.content,
			}

			if err := testBoard.SelectBox(insertBoxWithContentParams); !reflect.DeepEqual(err, test.want.err) {
				t.Errorf("unexpected error = %v, want %v", err, test.want.err)
			}

			if !reflect.DeepEqual(testBoard, test.want.expectBoard) {
				t.Errorf("unexpected Board = %v, want %v", testBoard, test.want.expectBoard)
			}

		})
	}

}

func TestCheckForWinner(t *testing.T) {

	type args struct {
		playerSymbol BoxContent
		rowIdx       int
		colIdx       int
	}

	type fields struct {
		WinCount           int
		Boxes              [][]BoxContent
		WinConditionChecks []winConditionCheck
	}

	type want struct {
		expectWin bool
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   want
	}{
		{
			"returns true when board reflects a win via a row",
			args{
				X,
				0,
				0,
			},
			fields{
				3,
				[][]BoxContent{
					{X, X, X},
					{E, E, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				true,
			},
		},
		{
			"returns false when board does not reflect a win via a row",
			args{
				X,
				0,
				0,
			},
			fields{
				3,
				[][]BoxContent{
					{X, X, E},
					{E, E, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				false,
			},
		},
		{
			"returns true when board reflects a win via a column",
			args{
				X,
				0,
				2,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, X},
					{E, E, X},
					{E, E, X},
				},
				winConditionChecks,
			},
			want{
				true,
			},
		},
		{
			"returns false when board does not reflect a win via a column",
			args{
				X,
				0,
				2,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, X},
					{E, E, X},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				false,
			},
		},
		{
			"returns true when board reflects a win via a diagonal",
			args{
				X,
				0,
				0,
			},
			fields{
				3,
				[][]BoxContent{
					{X, E, E},
					{E, X, E},
					{E, E, X},
				},
				winConditionChecks,
			},
			want{
				true,
			},
		},
		{
			"returns false when board does not reflect a win via a diagonal",
			args{
				X,
				0,
				0,
			},
			fields{
				3,
				[][]BoxContent{
					{X, E, E},
					{E, X, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				false,
			},
		},
		{
			"returns true when board reflects a win via a reverse diagonal",
			args{
				X,
				0,
				2,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, X},
					{E, X, E},
					{X, E, E},
				},
				winConditionChecks,
			},
			want{
				true,
			},
		},
		{
			"returns false when board does not reflect a win via a reverse diagonal",
			args{
				X,
				0,
				2,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, X},
					{E, X, E},
					{E, E, E},
				},
				winConditionChecks,
			},
			want{
				false,
			},
		},
		{
			"returns true when board reflects a win via a reverse diagonal on a 5*5 dimension",
			args{
				X,
				0,
				3,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, E, X, E},
					{E, E, X, E, E},
					{E, X, E, E, E},
					{X, E, E, E, E},
					{E, E, E, E, E},
				},
				winConditionChecks,
			},
			want{
				true,
			},
		},
		{
			"returns false when board does not reflect a win via a reverse diagonal on a 5*5 dimension",
			args{
				X,
				0,
				3,
			},
			fields{
				3,
				[][]BoxContent{
					{E, E, E, X, E},
					{E, E, X, E, E},
					{E, E, E, E, E},
					{X, E, E, E, E},
					{E, E, E, E, E},
				},
				winConditionChecks,
			},
			want{
				false,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testBoard := Board{
				WinCount:           test.fields.WinCount,
				Boxes:              test.fields.Boxes,
				WinConditionChecks: test.fields.WinConditionChecks,
			}

			checkForWinnerParams := CheckForWinnerParams{
				PlayerSymbol: test.args.playerSymbol,
				RowIdx:       test.args.rowIdx,
				ColIdx:       test.args.colIdx,
			}

			gotWin := testBoard.CheckForWinner(checkForWinnerParams)

			if gotWin != test.want.expectWin {
				t.Errorf("unexpected check result = %t, want %t", gotWin, test.want.expectWin)
			}

		})
	}

}
