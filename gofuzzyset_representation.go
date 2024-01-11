package gofuzzyset

type FuzzySetRepresentation struct {
	ItemsByGramSize map[int][]item
	MatchDict       map[string][]uint16
	ExactSet        map[string]string
	UseLevenshtein  bool
	GramSizeLower   int
	GramSizeUpper   int
	MinScore        float64
}

type ItemRepresentation struct {
	NormalizedValue string
	VectorNormal    float64
}

func NewFuzzySetRepresentation() FuzzySetRepresentation {
	return FuzzySetRepresentation{
		ItemsByGramSize: make(map[int][]item),
		MatchDict:       make(map[string][]uint16),
		ExactSet:        make(map[string]string),
	}
}

func (f FuzzySet) Export() FuzzySetRepresentation {
	return FuzzySetRepresentation{
		ItemsByGramSize: f.itemsByGramSize,
		MatchDict:       f.matchDict,
		ExactSet:        f.exactSet,
		UseLevenshtein:  f.useLevenshtein,
		GramSizeLower:   f.gramSizeLower,
		GramSizeUpper:   f.gramSizeUpper,
		MinScore:        f.minScore,
	}
}

func (f *FuzzySet) Import(r FuzzySetRepresentation) {
	f.itemsByGramSize = r.ItemsByGramSize
	f.matchDict = r.MatchDict
	f.exactSet = r.ExactSet
	f.useLevenshtein = r.UseLevenshtein
	f.gramSizeLower = r.GramSizeLower
	f.gramSizeUpper = r.GramSizeUpper
	f.minScore = r.MinScore
}

func (i item) Export() ItemRepresentation {
	return ItemRepresentation{
		NormalizedValue: i.normalizedValue,
		VectorNormal:    i.vectorNormal,
	}
}

func (i *item) Import(r ItemRepresentation) {
	i.normalizedValue = r.NormalizedValue
	i.vectorNormal = r.VectorNormal
}
