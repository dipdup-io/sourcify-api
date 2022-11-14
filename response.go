package sourcify

import (
	stdJSON "encoding/json"
)

// Metadata -
type Metadata struct {
	Compiler Compiler          `json:"compiler"`
	Language string            `json:"language"`
	Output   Output            `json:"output"`
	Settings Settings          `json:"settings"`
	Sources  map[string]Source `json:"sources"`
	Version  int               `json:"version"`
}

// Compiler -
type Compiler struct {
	Version string `json:"version"`
}

// Output -
type Output struct {
	ABI     stdJSON.RawMessage `json:"abi"`
	DevDoc  Doc                `json:"devdoc"`
	UserDoc Doc                `json:"userdoc"`
}

// Doc -
type Doc struct {
	Kind    string                        `json:"kind"`
	Methods map[string]stdJSON.RawMessage `json:"methods"`
	Version int                           `json:"version"`
}

// Settings -
type Settings struct {
	CompilationTarget map[string]string `json:"compilationTarget"`
	EvmVersion        string            `json:"evmVersion"`
	Libraries         struct{}          `json:"libraries"`
	Metadata          struct {
		BytecodeHash string `json:"bytecodeHash"`
	} `json:"metadata"`
	Optimizer struct {
		Enabled bool `json:"enabled"`
		Runs    int  `json:"runs"`
	} `json:"optimizer"`
	Remappings []interface{} `json:"remappings"`
}

// Source -
type Source struct {
	Keccak256 string   `json:"keccak256"`
	License   string   `json:"license"`
	Urls      []string `json:"urls"`
}

// ContractAddresses -
type ContractAddresses struct {
	Full    []string `json:"full"`
	Partial []string `json:"partial"`
}

// FileTree -
type FileTree struct {
	Status string   `json:"status"`
	Files  []string `json:"files"`
}

// Sources -
type Sources struct {
	Status string `json:"status"`
	Files  []File `json:"files"`
}

// File -
type File struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Content string `json:"content"`
}

// CheckStatus -
type CheckStatus struct {
	Address  string   `json:"address"`
	Status   string   `json:"status"`
	ChainIds []string `json:"chainIds"`
}

// CheckAllStatus -
type CheckAllStatus struct {
	Address  string          `json:"address"`
	ChainIds []ChainIdStatus `json:"chainIds"`
}

// ChainIdStatus -
type ChainIdStatus struct {
	ChainID string `json:"chainId"`
	Status  string `json:"status"`
}

// Error -
type Error struct {
	Error string `json:"error"`
}

// Chain -
type Chain struct {
	Name           string         `json:"name"`
	Chain          string         `json:"chain"`
	Network        string         `json:"network"`
	Icon           string         `json:"icon"`
	RPC            []string       `json:"rpc"`
	Faucets        []string       `json:"faucets"`
	NativeCurrency NativeCurrency `json:"nativeCurrency"`
	InfoURL        string         `json:"infoURL"`
	ShortName      string         `json:"shortName"`
	ChainID        int64          `json:"chainId"`
	NetworkID      int64          `json:"networkId"`
	Slip44         int64          `json:"slip44"`
	Ens            struct {
		Registry string `json:"registry"`
	} `json:"ens"`
	Explorers            []Explorer `json:"explorers"`
	Supported            bool       `json:"supported"`
	Monitored            bool       `json:"monitored"`
	ContractFetchAddress string     `json:"contractFetchAddress"`
	TxRegex              string     `json:"txRegex"`
}

// Explorer -
type Explorer struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Standard string `json:"standard"`
}

// NativeCurrency -
type NativeCurrency struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}
