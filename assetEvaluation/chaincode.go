package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"goSmartContract/api"
	"time"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/19 13:22

 * @Note:

 **/

type BlockChainAssetEvaluation struct {

}


func (t *BlockChainAssetEvaluation) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化成功")
	return shim.Success(nil)
}


func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainAssetEvaluation))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *BlockChainAssetEvaluation) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
		case "signContract":   //签署合同
			return api.SignContract(stub,args)
		case "completeEvaluation":	// 完成评估
			return api.CompleteEvaluation(stub,args)
		case "thirdLevelAudit":	//三级审核
			return api.ThirdLevelAudit(stub,args)
		case "score"	:	//评价
			return api.Score(stub,args)
		case "queryBusiness": 	//查询业务
			return api.QueryBusiness(stub,args)
		default:
			return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}