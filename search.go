package kkbox

// FetchSearchData Search for various objects..
func (b *Box) FetchSearchData(key string, params ...Param) (*SearchData, error) {
	resp := new(SearchData)

	if len(params) == 0 {
		params = append(params, Param{
			Q: key,
		})
	} else {
		params[0].Q = key
	}

	err := b.fetchData(SearchURL, resp, params...)

	return resp, err
}
