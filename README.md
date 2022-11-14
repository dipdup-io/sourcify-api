# Sourcify API
Golang wrapper over Sourcify API. Library implements interaction with endpoints are described in the [Sourcify docs](https://docs.sourcify.dev/docs/api/).

## Install

```go
go get github.com/dipdup-net/sourcify-api
```

## Usage

Create API structure

```go
api := sourcify.NewAPI("https://sourcify.dev/")
```

Call one of available methods

```go
// GetFile - gets the file from the repository server
func (api *API) GetFile(ctx context.Context, chainID, address, match, filename string) (*Metadata, error)

// GetFileTreeFullMatches - Returns repository URLs for every file in the source tree for the desired chain and address. Searches only for full matches.
func (api *API) GetFileTreeFullMatches(ctx context.Context, chainID, address string) ([]string, error)

// GetFileTree - returns repository URLs for every file in the source tree for the desired chain and address. Searches for full and partial matches.
func (api *API) GetFileTree(ctx context.Context, chainID, address string) (*FileTree, error)

// GetContractAddresses - Returns all verified contracts from the repository for the desired chain. Searches for full and partial matches.
func (api *API) GetContractAddresses(ctx context.Context, chainID string) (*ContractAddresses, error)

// CheckByAddresses - Checks if contract with the desired chain and address is verified and in the repository. It will only search for perfect matches.
func (api *API) CheckByAddresses(ctx context.Context, addresses []string, chainIds []string) ([]CheckStatus, error)

// CheckAllByAddresses - Checks if contract with the desired chain and address is verified and in the repository. It will search for both perfect and partial matches.
func (api *API) CheckAllByAddresses(ctx context.Context, addresses []string, chainIds []string) ([]CheckAllStatus, error)

// GetFiles - Returns all verified sources from the repository for the desired contract address and chain, including metadata.json. Searches for full and partial matches.
func (api *API) GetFiles(ctx context.Context, chainID, address string) ([]Sources, error)

// GetFilesFullMatch - Returns all verified sources from the repository for the desired contract address and chain, including metadata.json. Searches only for full matches.
func (api *API) GetFilesFullMatch(ctx context.Context, chainID, address string) ([]File, error)

// Chains - Returns the chains (networks) added to the Sourcify. Contains both supported, unsupported, monitored, unmonitored chains.
func (api *API) Chains(ctx context.Context) ([]Chain, error)

// Health - Ping the server and see if it is alive and ready for requests.
func (api *API) Health(ctx context.Context) (string, error)
```

Example of using library can be found [here](/example/main.go)

**Verification API is not realized yet in library**