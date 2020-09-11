package Test.Package

type FunctionRequest struct {
	name string
	page_number int
	result_per_page int
}

type FunctionResponse struct {
	result Test.Package2.Result[]
}

package Test.Package2

type Result struct {
	url string
	title string
	snippets string
}