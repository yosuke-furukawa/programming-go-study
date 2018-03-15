package types

// Query query params struct
type Query struct {
	Stroke string `url:"s"`
	Fill   string `url:"f"`
	Width  int    `url:"w"`
	Height int    `url:"h"`
	Top    string `url:"t"`
	Bottom string `url:"b"`
	Cells  int    `url:"c"`
}
