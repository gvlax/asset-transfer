package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var log *logging.Logger
var format logging.Formatter

//
// AssetTransferChaincode is a type of the chaincode
//
type AssetTransferChaincode struct {
}

func init() {
	log = logging.MustGetLogger("asset-transfer-chaincode")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶▶▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	logging.SetFormatter(format)
}

func main() {
	log.Debug("AssetTransferChaincode main() called ...")
	err := shim.Start(new(AssetTransferChaincode))
	if err != nil {
		log.Errorf("Error starting Simple chaincode: %s", err)
	}
	log.Info("The chaincode AssetTransferChaincode has been started.")
}

//
// Init method will be called during deployment (installation?)
//
func (t *AssetTransferChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Debug("Init() Chaincode...")
	/*
		_, args := stub.GetFunctionAndParameters()
		if len(args) != 0 {
			return shim.Error("Incorrect number of arguments. Expecting 0")
		}
	*/
	return shim.Success(nil)
}

//
// Query ....
//
func (t *AssetTransferChaincode) Query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	log.Debug("Query() called ...")
	return shim.Success(nil)
}

//
// Invoke ...
//
func (t *AssetTransferChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	log.Debug("Invoke() called ...")

	function, args := stub.GetFunctionAndParameters()
	log.Infof("# of args: %d, function name: %s", len(args), function)

	return shim.Success(nil)
}
