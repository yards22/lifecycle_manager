package entities

import "strconv"

// Wicket >> Boundaries >> Runs >> Extras

type CommentryF struct {
	Overs       string `json:"overs"`
	Result      Result `json:"result"`
	Description string `json:"description"`
}

type Result struct {
	IsWicket   bool   `json:"is_wicket"`
	RunsScored int    `json:"runs_scored"`
	IsExtra    bool   `json:"is_extra"`
	ExtraType  string `json:"extra_type,omitempty"`
}

type Wicket struct {
	Overs             int      `json:"overs"`
	Balls             int      `json:"balls"`
	WicketType        string   `json:"wicket_type"`
	StrikerBatsman    IBatsman `json:"striker_batsman"`
	NonStrikerBatsman IBatsman `json:"non_striker_batsman"`
	Bowler            string   `json:"bowler"`
	RunsScored        int      `json:"runs_scored"`
	WagonDirection    string   `json:"wagon_direction,omitempty"`
	IsFielder         bool     `json:"is_fielder,omitempty"`
	FieldedBy         string   `json:"fielded_by,omitempty"`
	IsExtra           bool     `json:"is_extra"`
	ExtraType         string   `json:"extra_type,omitempty"`
}

func (W *Wicket) WicketC() CommentryF {
	var c CommentryF

	c.Overs = strconv.Itoa(W.Overs) + "." + strconv.Itoa(W.Balls)
	c.Result.IsWicket = true
	c.Result.IsExtra = W.IsExtra
	c.Result.ExtraType = W.ExtraType
	c.Result.RunsScored = W.RunsScored

	if W.WicketType == "B" {

	}
	if W.WicketType == "C" {

	}
	if W.WicketType == "R" {

	}
	if W.WicketType == "S" {

	}
	if W.WicketType == "L" {

	}
	return c
}

// its not a wicket for sure ..

type Runs struct {
	Overs             int      `json:"overs"`
	Balls             int      `json:"balls"`
	RunsScored        int      `json:"runs_scored"`
	StrikerBatsman    IBatsman `json:"striker_batsman"`
	NonStrikerBatsman IBatsman `json:"non_striker_batsman"`
	Bowler            string   `json:"bowler"`
	IsBoundary        bool     `json:"is_boundary"`
	BoundaryType      string   `json:"boundary_type,omitempty"`
	WagonDirection    string   `json:"wagon_direction,omitempty"`
	IsExtra           bool     `json:"is_extra"`
	ExtraType         string   `json:"extra_type,omitempty"`
}

func (R *Runs) RunsC() CommentryF {
	var c CommentryF

	c.Overs = strconv.Itoa(R.Overs) + "." + strconv.Itoa(R.Balls)
	c.Result.IsWicket = false
	c.Result.IsExtra = R.IsExtra
	c.Result.ExtraType = R.ExtraType
	c.Result.RunsScored = R.RunsScored

	// add description for runs ...

	return c
}

// its not a run, not a wicket ...

type Extra struct {
	Overs             int      `json:"overs"`
	Balls             int      `json:"balls"`
	ExtraType         string   `json:"extra_type,omitempty"`
	StrikerBatsman    IBatsman `json:"striker_batsman"`
	NonStrikerBatsman IBatsman `json:"non_striker_batsman"`
	Bowler            string   `json:"bowler"`
}

func (E *Extra) ExtraC() CommentryF {
	var c CommentryF

	c.Overs = strconv.Itoa(E.Overs) + "." + strconv.Itoa(E.Balls)
	c.Result.IsWicket = false
	c.Result.IsExtra = true
	c.Result.ExtraType = E.ExtraType
	c.Result.RunsScored = 0

	// add description for runs ...

	return c
}
