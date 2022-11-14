package sourcify

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// API -
type API struct {
	baseURL string
}

// NewAPI -
func NewAPI(baseURL string) *API {
	return &API{baseURL: baseURL}
}

func getClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 10
	t.MaxConnsPerHost = 10
	t.MaxIdleConnsPerHost = 10

	return &http.Client{
		Transport: t,
	}
}

func (api *API) get(ctx context.Context, link *url.URL, output any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link.String(), nil)
	if err != nil {
		return err
	}

	response, err := getClient().Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK, http.StatusNotModified:
		return json.NewDecoder(response.Body).Decode(output)
	case http.StatusNotFound:
		return ErrNotFound
	default:
		data, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return errors.Errorf("sourcify invalid status code %d: %s", response.StatusCode, data)
	}
}

func (api *API) getString(ctx context.Context, link *url.URL) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link.String(), nil)
	if err != nil {
		return "", err
	}

	response, err := getClient().Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusNotModified:
		return string(data), nil
	default:

		return string(data), errors.Errorf("sourcify invalid status code %d: %s", response.StatusCode, data)
	}
}

// GetFile - gets the file from the repository server
func (api *API) GetFile(ctx context.Context, chainID, address, match, filename string) (*Metadata, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/repository/contracts")
	link = link.JoinPath(match)
	link = link.JoinPath(chainID)
	link = link.JoinPath(address)
	link = link.JoinPath(filename)

	var metadata Metadata
	err = api.get(ctx, link, &metadata)
	return &metadata, err
}

// GetFileTreeFullMatches - Returns repository URLs for every file in the source tree for the desired chain and address. Searches only for full matches.
func (api *API) GetFileTreeFullMatches(ctx context.Context, chainID, address string) ([]string, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/files/tree")
	link = link.JoinPath(chainID)
	link = link.JoinPath(address)

	var fileTree []string
	err = api.get(ctx, link, &fileTree)
	return fileTree, err
}

// GetFileTree - returns repository URLs for every file in the source tree for the desired chain and address. Searches for full and partial matches.
func (api *API) GetFileTree(ctx context.Context, chainID, address string) (*FileTree, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/files/tree/any")
	link = link.JoinPath(chainID)
	link = link.JoinPath(address)

	var fileTree FileTree
	err = api.get(ctx, link, &fileTree)
	return &fileTree, err
}

// GetContractAddresses - Returns all verified contracts from the repository for the desired chain. Searches for full and partial matches.
func (api *API) GetContractAddresses(ctx context.Context, chainID string) (*ContractAddresses, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/files/contracts")
	link = link.JoinPath(chainID)

	var addresses ContractAddresses
	err = api.get(ctx, link, &addresses)
	return &addresses, err
}

// CheckByAddresses - Checks if contract with the desired chain and address is verified and in the repository. It will only search for perfect matches.
func (api *API) CheckByAddresses(ctx context.Context, addresses []string, chainIds []string) ([]CheckStatus, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/check-by-addresses")
	values := link.Query()
	if len(addresses) > 0 {
		values.Add("addresses", strings.Join(addresses, ","))
	}
	if len(chainIds) > 0 {
		values.Add("chainIds", strings.Join(chainIds, ","))
	}
	link.RawQuery = values.Encode()

	var response []CheckStatus
	err = api.get(ctx, link, &response)
	return response, err
}

// CheckAllByAddresses - Checks if contract with the desired chain and address is verified and in the repository. It will search for both perfect and partial matches.
func (api *API) CheckAllByAddresses(ctx context.Context, addresses []string, chainIds []string) ([]CheckAllStatus, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/check-all-by-addresses")
	values := link.Query()
	if len(addresses) > 0 {
		values.Add("addresses", strings.Join(addresses, ","))
	}
	if len(chainIds) > 0 {
		values.Add("chainIds", strings.Join(chainIds, ","))
	}
	link.RawQuery = values.Encode()

	var response []CheckAllStatus
	err = api.get(ctx, link, &response)
	return response, err
}

// GetFiles - Returns all verified sources from the repository for the desired contract address and chain, including metadata.json. Searches for full and partial matches.
func (api *API) GetFiles(ctx context.Context, chainID, address string) (*Sources, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/files/any")
	link = link.JoinPath(chainID)
	link = link.JoinPath(address)

	var response Sources
	err = api.get(ctx, link, &response)
	return &response, err
}

// GetFilesFullMatch - Returns all verified sources from the repository for the desired contract address and chain, including metadata.json. Searches only for full matches.
func (api *API) GetFilesFullMatch(ctx context.Context, chainID, address string) ([]File, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/files")
	link = link.JoinPath(chainID)
	link = link.JoinPath(address)

	var response []File
	err = api.get(ctx, link, &response)
	return response, err
}

// Chains - Returns the chains (networks) added to the Sourcify. Contains both supported, unsupported, monitored, unmonitored chains.
func (api *API) Chains(ctx context.Context) ([]Chain, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return nil, err
	}
	link = link.JoinPath("server/chains")
	var response []Chain
	err = api.get(ctx, link, &response)
	return response, err
}

// Health - Ping the server and see if it is alive and ready for requests.
func (api *API) Health(ctx context.Context) (string, error) {
	link, err := url.Parse(api.baseURL)
	if err != nil {
		return "", err
	}
	link = link.JoinPath("server/health")
	return api.getString(ctx, link)
}
