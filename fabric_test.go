package fabric_relayer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/polynetwork/fabric-relayer/internal/github.com/hyperledger/fabric/protoutil"
	"os"
	"testing"
	"time"
)

func newFabSdk() *fabsdk.FabricSDK {
	sdk, err := fabsdk.New(config.FromFile("./config/config_e2e.yaml"))
	if err != nil {
		panic(err)
	}
	return sdk
}

func newResMgmt(sdk *fabsdk.FabricSDK) *resmgmt.Client {
	rcp := sdk.Context(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	rc, err := resmgmt.New(rcp)
	if err != nil {
		panic(err)
	}
	return rc
}

func newChannelClient(sdk *fabsdk.FabricSDK) *channel.Client {
	ccp := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	cc, err := channel.New(ccp)
	if err != nil {
		panic(err)
	}
	return cc
}

func newEventClient(sdk *fabsdk.FabricSDK) *event.Client {
	ccp := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	eventClient, err := event.New(ccp, event.WithBlockEvents())
	if err != nil {
		panic(err)
	}
	return eventClient
}

func newLedger(sdk *fabsdk.FabricSDK) *ledger.Client {
	ccp := sdk.ChannelContext("mychannel", fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))
	ledgerClient, err := ledger.New(ccp)
	if err != nil {
		panic(err)
	}
	return ledgerClient
}

func packArgs(args []string) [][]byte {
	ret := make([][]byte, 0)
	for _, arg := range args {
		ret = append(ret, []byte(arg))
	}
	return ret
}

func TestCCQuery(t *testing.T) {
	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	req := channel.Request{
		ChaincodeID: "basic",
		Fcn: "GetAllAssets",
		Args: packArgs([]string{}),
	}
	response, err := channelClient.Query(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %s\n", string(response.Payload))
}

func TestCCInvoke(t *testing.T) {
	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	req := channel.Request{
		ChaincodeID: "basic",
		Fcn: "TransferAsset",
		Args: packArgs([]string{"asset6","Christopher"}),
	}
	response, err := channelClient.Execute(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %v\n", string(response.TransactionID))
}

func TestCCEvent(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("startServer - get current work directory failed!")
		return
	}
	os.Setenv("FABRIC_RELAYER_PATH", dir)

	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	eventClient := newEventClient(sdk)

	eventID := ".*"
	reg, notifier, err := eventClient.RegisterChaincodeEvent("mycc", eventID)
	if err != nil {
		panic(err)
	}
	defer eventClient.Unregister(reg)

	req := channel.Request{
		ChaincodeID: "mycc",
		Fcn: "query",
		Args: packArgs([]string{"a"}),
	}
	response, err := channelClient.Execute(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %s\n", string(response.TransactionID))

	select {
	case ccEvent := <- notifier:
		fmt.Printf("receive cc event:%v\n", ccEvent)
	case <- time.After(time.Second * 60):
		fmt.Printf("not receive cc event!")
	}
}

func TestCCEvent1(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("startServer - get current work directory failed!")
		return
	}
	os.Setenv("FABRIC_RELAYER_PATH", dir)

	sdk := newFabSdk()
	channelClient := newChannelClient(sdk)
	eventClient := newEventClient(sdk)

	eventID := "to.*"
	reg, notifier, err := eventClient.RegisterChaincodeEvent("ccm1", eventID)
	if err != nil {
		panic(err)
	}
	defer eventClient.Unregister(reg)

	req := channel.Request{
		ChaincodeID: "peth",
		Fcn: "lock",
		Args: packArgs([]string{"2", "BC8F34783742ea552C7e8823a2A9e8f58052B4D4", "11"}),
	}
	response, err := channelClient.Execute(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		panic(err)
	}
	fmt.Printf("response: %s\n", string(response.TransactionID))

	select {
	case ccEvent := <- notifier:
		fmt.Printf("receive cc event:%v\n", ccEvent)
	case <- time.After(time.Second * 600):
		fmt.Printf("not receive cc event!")
	}
}

func TestTransaction(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("startServer - get current work directory failed!")
		return
	}
	os.Setenv("FABRIC_RELAYER_PATH", dir)

	sdk := newFabSdk()
	ledgerClient := newLedger(sdk)
	tx, err := ledgerClient.QueryTransaction("5c69313e45b78a951a5ea01ad66de45ed11b198eeb3cd8f06bc968c0ff8e0cc9")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("transaction: %s\n", string(tx.TransactionEnvelope.Payload))

	pl := &common.Payload{}
	err = proto.Unmarshal(tx.TransactionEnvelope.Payload, pl)
	if err != nil {
		t.Fatal(err)
	}

	txn := &peer.Transaction{}
	err = proto.Unmarshal(pl.Data, txn)
	if err != nil {
		t.Fatal(err)
	}

	ac := &peer.TransactionAction{}
	err = proto.Unmarshal(txn.Actions[0].Payload, ac)
	if err != nil {
		t.Fatal(err)
	}

	capl := &peer.ChaincodeActionPayload{}
	err = proto.Unmarshal(ac.Payload, capl)
	if err != nil {
		t.Fatal(err)
	}

	hdr := &common.ChannelHeader{}
	err = proto.Unmarshal(pl.Header.ChannelHeader, hdr)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInfo(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("startServer - get current work directory failed!")
		return
	}
	os.Setenv("FABRIC_RELAYER_PATH", dir)

	sdk := newFabSdk()
	ledgerClient := newLedger(sdk)

	info, err := ledgerClient.QueryInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("height: %d\n", info.BCI.Height)
}


func TestBlock(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("startServer - get current work directory failed!")
		return
	}
	os.Setenv("FABRIC_RELAYER_PATH", dir)

	sdk := newFabSdk()
	ledgerClient := newLedger(sdk)
	for i := uint64(3); i < 50; i++ {
		fmt.Println(i)
		block, err := ledgerClient.QueryBlock(i)
		if err != nil {
			panic(err)
		}
		for _, v := range block.Data.Data {
			xx, err := protoutil.GetEnvelopeFromBlock(v)
			if err != nil {
				t.Fatal(err)
			}
			//cas, err := protoutil.GetActionsFromEnvelope(v)
			cas, err := protoutil.GetActionsFromEnvelopeMsg(xx)
			if err != nil {
				t.Fatal(err)
			}

			for _, e := range cas {
				chaincodeEvent := &peer.ChaincodeEvent{}
				err = proto.Unmarshal(e.Events, chaincodeEvent)
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println(chaincodeEvent.String())
			}
		}
	}
}

type TransferEvent struct {
	From   []byte `json:"from"`
	To     []byte `json:"to"`
	Amount []byte `json:"amount"`
}
