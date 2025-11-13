package kokkeneskoekken

type apiResponse struct {
	Offers map[string]offer
}

type offer struct {
	Items []item
}

type item struct {
	Name  string
	Dates map[int64]date
}

type date struct {
	Available bool
	Menu      struct {
		Name string
	}
}

type Menu map[string][]string

type Schedule map[string]Menu
