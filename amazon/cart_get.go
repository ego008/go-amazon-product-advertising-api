package amazon

// CartGetResponseGroup represents constants those are capable ResponseGroups parameter
type CartGetResponseGroup string

const (
	// CartGetResponseGroupCart is a constant for Cart response group
	CartGetResponseGroupCart CartGetResponseGroup = "Cart"
	// CartGetResponseGroupCartTopSellers is a constant for CartTopSellers response group
	CartGetResponseGroupCartTopSellers CartGetResponseGroup = "CartTopSellers"
	// CartGetResponseGroupCartSimilarities is a constant for CartSimilarities response group
	CartGetResponseGroupCartSimilarities CartGetResponseGroup = "CartSimilarities"
	// CartGetResponseGroupCartNewReleases is a constant for CartNewReleases response group
	CartGetResponseGroupCartNewReleases CartGetResponseGroup = "CartNewReleases"
)

// CartGetRequest represents request for CartGet operation
type CartGetRequest struct {
	ResponseGroups []CartGetResponseGroup
	Client         *Client
}

// CartGetResponse represents response for CartGet operation
type CartGetResponse struct {
	Error error
}

// Do send request for the API
func (req *CartGetRequest) Do() (*CartGetResponse, error) {
	return nil, nil
}

// CartGet returns new request for CartGet
func (client *Client) CartGet(responseGroups ...CartGetResponseGroup) *CartGetRequest {
	return &CartGetRequest{
		Client:         client,
		ResponseGroups: responseGroups,
	}
}
