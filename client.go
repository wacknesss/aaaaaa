package hedera

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/rand"
	"os"
)

// Default max fees and payments to 1 h-bar
const defaultMaxTransactionFee uint64 = 100_000_000
const defaultMaxQueryPayment uint64 = 100_000_000

type Client struct {
	maxTransactionFee uint64
	maxQueryPayment   uint64

	operator *operator

	networkNodes   map[AccountID]*node
	networkNodeIds []AccountID
}

type node struct {
	conn    *grpc.ClientConn
	id      AccountID
	address string
}

type Signer func(message []byte) []byte

type operator struct {
	accountID  AccountID
	privateKey *Ed25519PrivateKey
	publicKey  Ed25519PublicKey
	signer     Signer
}

var mainnetNodes = map[string]AccountID{
	"35.237.200.180:50211": AccountID{Account: 3},
	"35.186.191.247:50211": AccountID{Account: 4},
	"35.192.2.25:50211":    AccountID{Account: 5},
	"35.199.161.108:50211": AccountID{Account: 6},
	"35.203.82.240:50211":  AccountID{Account: 7},
	"35.236.5.219:50211":   AccountID{Account: 8},
	"35.197.192.225:50211": AccountID{Account: 9},
	"35.242.233.154:50211": AccountID{Account: 10},
	"35.240.118.96:50211":  AccountID{Account: 11},
	"35.204.86.32:50211":   AccountID{Account: 12},
}

var testnetNodes = map[string]AccountID{
	"1.testnet.hedera.com:50211": AccountID{Account: 3},
	"2.testnet.hedera.com:50211": AccountID{Account: 4},
	"3.testnet.hedera.com:50211": AccountID{Account: 5},
	"4.testnet.hedera.com:50211": AccountID{Account: 6},
}

func MainnetClient() *Client {
	return NewClient(mainnetNodes)
}

func TestnetClient() *Client {
	return NewClient(testnetNodes)
}

func NewClient(network map[string]AccountID) *Client {
	client := &Client{
		maxQueryPayment:   defaultMaxQueryPayment,
		maxTransactionFee: defaultMaxTransactionFee,
		networkNodes:      map[AccountID]*node{},
		networkNodeIds:    []AccountID{},
	}

	client.ReplaceNodes(network)

	return client
}

func FromFile(filename string) (*Client, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var networkStrings map[string]string
	network := map[string]AccountID{}

	bytes, _ := ioutil.ReadAll(file)

	json.Unmarshal([]byte(bytes), &network)

	for address, id := range networkStrings {
		account, err := AccountIDFromString(id)
		if err != nil {
			return nil, err
		}

		network[address] = account
	}

	return NewClient(network), nil

}

func (client *Client) Close() error {
	for _, node := range client.networkNodes {
		if node.conn != nil {
			err := node.conn.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (client *Client) ReplaceNodes(network map[string]AccountID) *Client {
	networkNodes := map[AccountID]*node{}
	var networkNodeIds []AccountID

	for address, id := range network {
		networkNodeIds = append(networkNodeIds, id)
		networkNodes[id] = &node{
			id:      id,
			address: address,
		}
	}

	return client
}

func (client *Client) SetOperator(accountID AccountID, privateKey Ed25519PrivateKey) *Client {
	client.operator = &operator{
		accountID:  accountID,
		privateKey: &privateKey,
		publicKey:  privateKey.publicKey,
		signer:     privateKey.Sign,
	}

	return client
}

func (client *Client) SetOperatorWith(accountID AccountID, publicKey Ed25519PublicKey, signer Signer) *Client {
	client.operator = &operator{
		accountID:  accountID,
		privateKey: nil,
		publicKey:  publicKey,
		signer:     signer,
	}

	return client
}

func (client *Client) SetMaxTransactionFee(tinyBars uint64) *Client {
	client.maxTransactionFee = tinyBars
	return client
}

func (client *Client) SetMaxQueryPayment(tinyBars uint64) *Client {
	client.maxTransactionFee = tinyBars
	return client
}

func (client *Client) GetAccountInfo() (AccountInfo, error) {
	return NewAccountInfoQuery().
		SetAccountID(client.operator.accountID).
		Execute(client)
}

func (client *Client) GetAccountBalance() (uint64, error) {
	return NewAccountBalanceQuery().
		SetAccountID(client.operator.accountID).
		Execute(client)
}

func (client *Client) randomNode() *node {
	nodeIndex := rand.Intn(len(client.networkNodeIds))
	nodeId := client.networkNodeIds[nodeIndex]

	return client.networkNodes[nodeId]
}

func (client *Client) node(id AccountID) *node {
	return client.networkNodes[id]
}

func (node *node) invoke(method string, in interface{}, out interface{}) error {
	if node.conn == nil {
		conn, err := grpc.Dial(node.address, grpc.WithInsecure())
		if err != nil {
			return err
		}

		node.conn = conn
	}

	return node.conn.Invoke(context.TODO(), method, in, out)
}
