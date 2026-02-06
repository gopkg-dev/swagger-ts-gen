package generator

type Param struct {
	Name        string
	VarName     string
	Type        string
	Description string
	Required    bool
}

type QueryInfo struct {
	TypeName string
	Optional bool
}

type BodyInfo struct {
	TypeName string
	Optional bool
	IsForm   bool
}

type ReturnInfo struct {
	Type           string
	IsVoid         bool
	UsesPageResult bool
}

type Operation struct {
	Name       string
	Summary    string
	Method     string
	Path       string
	Group      string
	PathParams []Param
	Query      *QueryInfo
	Body       *BodyInfo
	Return     ReturnInfo
	ErrorText  string
}
