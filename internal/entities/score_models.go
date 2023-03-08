package entities

type InningDetails struct {
	InningsId   int    `json:"innings_id"`
	Overs       int    `json:"overs"`
	Balls       int    `json:"balls"`
	BattingTeam string `json:"batting_team"`
	BowlingTeam string `json:"bowling_team"`
}

type IBatsman struct {
	Name  string `json:"name"`
	Runs  int    `json:"runs"`
	Balls int    `json:"balls"`
	Fours int    `json:"fours"`
	Sixes int    `json:"sixes"`
}

type PlayersInAction struct {
	Bowler            string   `json:"bowler"`
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
}
