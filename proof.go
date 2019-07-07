package bridge

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/incognitochain/bridge-eth/jsonresult"
)

const inst_max_path = 8
const pubkey_max_path = 3
const comm_size = 1 << pubkey_max_path
const pubkey_node = comm_size * pubkey_max_path
const inst_size = 1 << pubkey_max_path
const inst_length = 150

type getProofResult struct {
	Result jsonresult.GetInstructionProof
	Error  string
	Id     int
}

type decodedProof struct {
	instruction  []byte
	beaconHeight *big.Int
	bridgeHeight *big.Int

	beaconInstPath         [inst_max_path][32]byte
	beaconInstPathIsLeft   [inst_max_path]bool
	beaconInstPathLen      *big.Int
	beaconInstRoot         [32]byte
	beaconBlkData          [32]byte
	beaconBlkHash          [32]byte
	beaconSignerPubkeys    []byte
	beaconSignerCount      *big.Int
	beaconSignerSig        [32]byte
	beaconSignerPaths      [pubkey_node][32]byte
	beaconSignerPathIsLeft [pubkey_node]bool
	beaconSignerPathLen    *big.Int

	bridgeInstPath         [inst_max_path][32]byte
	bridgeInstPathIsLeft   [inst_max_path]bool
	bridgeInstPathLen      *big.Int
	bridgeInstRoot         [32]byte
	bridgeBlkData          [32]byte
	bridgeBlkHash          [32]byte
	bridgeSignerPubkeys    []byte
	bridgeSignerCount      *big.Int
	bridgeSignerSig        [32]byte
	bridgeSignerPaths      [pubkey_node][32]byte
	bridgeSignerPathIsLeft [pubkey_node]bool
	bridgeSignerPathLen    *big.Int
}

func getAndDecodeBurnProof(txID string) (*decodedProof, error) {
	body := getBurnProof(txID)
	if len(body) < 1 {
		return nil, fmt.Errorf("burn proof not found")
	}

	r := getProofResult{}
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		return nil, err
	}
	return decodeProof(&r), nil
}

func getBurnProof(txID string) string {
	url := "http://127.0.0.1:9338"

	if len(txID) == 0 {
		txID = "2a0cf9a94a60a2aba82f17e331d4158695cd4943b36eab707bd774dffd6086a6"
	}
	payload := strings.NewReader(fmt.Sprintf("{\n    \"id\": 1,\n    \"jsonrpc\": \"1.0\",\n    \"method\": \"getburnproof\",\n    \"params\": [\n    \t\"%s\"\n    ]\n}", txID))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "127.0.0.1:9338")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func decodeProof(r *getProofResult) *decodedProof {
	inst := decode(r.Result.Instruction)
	fmt.Printf("inst: %d %x\n", len(inst), inst)

	// Block heights
	beaconHeight := big.NewInt(0).SetBytes(decode(r.Result.BeaconHeight))
	bridgeHeight := big.NewInt(0).SetBytes(decode(r.Result.BridgeHeight))

	beaconInstRoot := decode32(r.Result.BeaconInstRoot)
	beaconInstPath := [inst_max_path][32]byte{}
	beaconInstPathIsLeft := [inst_max_path]bool{}
	for i, path := range r.Result.BeaconInstPath {
		beaconInstPath[i] = decode32(path)
		beaconInstPathIsLeft[i] = r.Result.BeaconInstPathIsLeft[i]
	}
	beaconInstPathLen := big.NewInt(int64(len(r.Result.BeaconInstPath)))
	fmt.Printf("beaconInstRoot: %x\n", beaconInstRoot)

	beaconBlkData := toByte32(decode(r.Result.BeaconBlkData))
	beaconBlkHash := toByte32(decode(r.Result.BeaconBlkHash))
	fmt.Printf("expected beaconBlkHash: %x\n", keccak256(beaconBlkData[:], beaconInstRoot[:]))
	fmt.Printf("beaconBlkHash: %x\n\n", beaconBlkHash)

	beaconSignerPubkeys := []byte{}
	for _, signer := range r.Result.BeaconSignerPubkeys {
		beaconSignerPubkeys = append(beaconSignerPubkeys, decode(signer)...)
	}
	beaconSignerCount := big.NewInt(int64(len(r.Result.BeaconSignerPubkeys)))

	beaconSignerSig := toByte32(decode(r.Result.BeaconSignerSig))
	beaconSignerPaths := [pubkey_node][32]byte{}
	beaconSignerPathIsLeft := [pubkey_node]bool{}
	for i, fullPath := range r.Result.BeaconSignerPaths {
		for j, node := range fullPath {
			k := i*len(fullPath) + j
			beaconSignerPaths[k] = decode32(node)
			beaconSignerPathIsLeft[k] = r.Result.BeaconSignerPathIsLeft[i][j]
		}
	}
	beaconSignerPathLen := big.NewInt(int64(len(r.Result.BeaconSignerPaths[0])))

	// For bridge
	bridgeInstRoot := decode32(r.Result.BridgeInstRoot)
	bridgeInstPath := [inst_max_path][32]byte{}
	bridgeInstPathIsLeft := [inst_max_path]bool{}
	for i, path := range r.Result.BridgeInstPath {
		bridgeInstPath[i] = decode32(path)
		bridgeInstPathIsLeft[i] = r.Result.BridgeInstPathIsLeft[i]
	}
	bridgeInstPathLen := big.NewInt(int64(len(r.Result.BridgeInstPath)))
	// fmt.Printf("bridgeInstRoot: %x\n", bridgeInstRoot)

	bridgeBlkData := toByte32(decode(r.Result.BridgeBlkData))
	bridgeBlkHash := toByte32(decode(r.Result.BridgeBlkHash))
	// fmt.Printf("bridgeBlkHash: %x\n", bridgeBlkHash)

	bridgeSignerPubkeys := []byte{}
	for _, signer := range r.Result.BridgeSignerPubkeys {
		bridgeSignerPubkeys = append(bridgeSignerPubkeys, decode(signer)...)
	}
	bridgeSignerCount := big.NewInt(int64(len(r.Result.BridgeSignerPubkeys)))

	bridgeSignerSig := toByte32(decode(r.Result.BridgeSignerSig))
	bridgeSignerPaths := [pubkey_node][32]byte{}
	bridgeSignerPathIsLeft := [pubkey_node]bool{}
	for i, fullPath := range r.Result.BridgeSignerPaths {
		for j, node := range fullPath {
			k := i*len(fullPath) + j
			bridgeSignerPaths[k] = decode32(node)
			bridgeSignerPathIsLeft[k] = r.Result.BridgeSignerPathIsLeft[i][j]
		}
	}
	bridgeSignerPathLen := big.NewInt(int64(len(r.Result.BridgeSignerPaths[0])))
	return &decodedProof{
		instruction: inst,

		beaconHeight:           beaconHeight,
		beaconInstPath:         beaconInstPath,
		beaconInstPathIsLeft:   beaconInstPathIsLeft,
		beaconInstPathLen:      beaconInstPathLen,
		beaconInstRoot:         beaconInstRoot,
		beaconBlkData:          beaconBlkData,
		beaconBlkHash:          beaconBlkHash,
		beaconSignerPubkeys:    beaconSignerPubkeys,
		beaconSignerCount:      beaconSignerCount,
		beaconSignerSig:        beaconSignerSig,
		beaconSignerPaths:      beaconSignerPaths,
		beaconSignerPathIsLeft: beaconSignerPathIsLeft,
		beaconSignerPathLen:    beaconSignerPathLen,

		bridgeHeight:           bridgeHeight,
		bridgeInstPath:         bridgeInstPath,
		bridgeInstPathIsLeft:   bridgeInstPathIsLeft,
		bridgeInstPathLen:      bridgeInstPathLen,
		bridgeInstRoot:         bridgeInstRoot,
		bridgeBlkData:          bridgeBlkData,
		bridgeBlkHash:          bridgeBlkHash,
		bridgeSignerPubkeys:    bridgeSignerPubkeys,
		bridgeSignerCount:      bridgeSignerCount,
		bridgeSignerSig:        bridgeSignerSig,
		bridgeSignerPaths:      bridgeSignerPaths,
		bridgeSignerPathIsLeft: bridgeSignerPathIsLeft,
		bridgeSignerPathLen:    bridgeSignerPathLen,
	}
}

func toByte32(s []byte) [32]byte {
	a := [32]byte{}
	copy(a[:], s)
	return a
}

func decode(s string) []byte {
	d, _ := hex.DecodeString(s)
	return d
}

func decode32(s string) [32]byte {
	return toByte32(decode(s))
}

func keccak256(b ...[]byte) [32]byte {
	h := crypto.Keccak256(b...)
	r := [32]byte{}
	copy(r[:], h)
	return r
}
