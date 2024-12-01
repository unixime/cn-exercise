package query

type OPERATOR string

const (
	EQUAL        OPERATOR = "EQ"
	NOTEQUAL     OPERATOR = "NE"
	LESSTHAN     OPERATOR = "LT"
	LESSEQUAL    OPERATOR = "LE"
	GREATERTHAN  OPERATOR = "GT"
	GREATEREQUAL OPERATOR = "GE"
	LIKE         OPERATOR = "LIKE"
)

type ORDER bool

const (
	ASC  ORDER = false
	DESC ORDER = true
)

type Constraint struct {
	Field    string      `json:"field"`
	Operator OPERATOR    `json:"operator"`
	Value    interface{} `json:"value"`
}

type FieldComparison struct {
	Constraints []Constraint
}

type OrderConstraint struct {
	Field string `json:"field"`
	Order ORDER  `json:"desc"`
}

type Expression struct {
	FieldComparisons []Constraint `json:"fieldComparisons"`
}

type Query struct {
	Expressions []Expression      `json:"expressions"`
	OrderBy     []OrderConstraint `json:"orderBy"`
	Limit       int               `json:"limit"`
}

type Search struct {
	SearchID string `json:"searchId,omitempty"`
	KeepOpen bool   `json:"keepOpen,omitempty"`
	Query    Query  `json:"query"`
	Page     int    `json:"page"`
	PerPage  int    `json:"perPage"`
}
