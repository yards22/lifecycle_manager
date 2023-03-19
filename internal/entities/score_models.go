package entities

type IInnings struct {
	Score   int     `json:"score"`
	Wickets int     `json:"wickets"`
	RunRate float32 `json:"run_rate"`
	Overs   int     `json:"overs"`
	Balls   int     `json:"balls"`
}

type InningDetails struct {
	InningsId   int      `json:"innings_id"`
	Innings1    IInnings `json:"innings_1,omitempty"`
	Innings2    IInnings `json:"innings_2,omitempty"`
	BattingTeam string   `json:"batting_team"`
	BowlingTeam string   `json:"bowling_team"`
	Overs       int      `json:"overs"`
	Balls       int      `json:"balls"`
}

type IBatsman struct {
	Name  string `json:"name"`
	Runs  int    `json:"runs"`
	Balls int    `json:"balls"`
	Fours int    `json:"fours"`
	Sixes int    `json:"sixes"`
}

type IBowler struct {
	Name    string `json:"name"`
	Overs   int    `json:"overs"`
	Balls   int    `json:"balls"`
	Runs    int    `json:"runs"`
	Maiden  int    `json:"maiden"`
	Wickets int    `json:"wickets"`
}

type PlayersInAction struct {
	Bowler            IBowler  `json:"bowler"`
	StrikerBatsman    IBatsman `json:"striker_batsman"`
	NonStrikerBatsman IBatsman `json:"non_striker_batsman"`
}

type ExtraDetails struct {
	IsExtra   bool   `json:"is_extra"`
	ExtraType string `json:"extra_type,omitempty"`
}

type RunsDetails struct {
	RunsScored     int    `json:"runs_scored"`
	IsBoundary     bool   `json:"is_boundary"`
	BoundaryType   string `json:"boundary_type,omitempty"`
	ScoredBy       string `json:"scored_by"`
	WagonDirection string `json:"wagon_direction,omitempty"`
}

type WicketDetails struct {
	IsWicket   bool   `json:"is_wicket"`
	WicketType string `json:"wicket_type,omitempty"`
	WicketOf   string `json:"wicket_of,omitempty"`
	WicketBy   string `json:"wicket_by,omitempty"`
	IsFielder  bool   `json:"is_fielder,omitempty"`
	FieldedBy  string `json:"fielded_by,omitempty"`
}

type ScoreItem struct {
	MatchId         string          `json:"match_id"`
	OwnerId         int             `json:"owner_id"`
	InningsDetails  InningDetails   `json:"innings_details"`
	PlayersInAction PlayersInAction `json:"players_in_action"`
	ExtraDetails    ExtraDetails    `json:"extra_details"`
	RunsDetails     RunsDetails     `json:"runs_details"`
	WicketDetails   WicketDetails   `json:"wicket_details"`
	Undo            bool            `json:"undo"`
}

type ScoreSummary struct {
	StrikerBatsman    IBatsman     `json:"striker_batsman"`
	NonStrikerBatsman IBatsman     `json:"non_striker_batsman"`
	CurBowler         IBowler      `json:"cur_bowler"`
	PrevBowler        IBowler      `json:"Prev_bowler"`
	PrevOvers         []int        `json:"prev_overs"`
	Description       []CommentryF `json:"description"`
}
