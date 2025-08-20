package api

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"goSmartContract/model"
	"goSmartContract/service"
	"strconv"
	"time"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/22 16:30

 * @Note:

 **/

// SignContract  签署合约
func SignContract(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//验证参数
	if len(args)!=3{
		return shim.Error("参数个数不满足")
	}
	user:=args[0] // 资产评估需求者(乙方）
	assetAppraiser:=args[1]	 //资产评估师(甲方）
	contractHash:=args[2]  //所签合同哈希

	if user!=""{
		return shim.Error("资产评估需求方(乙方）不能为空")
	}
	if assetAppraiser!=""{
		return shim.Error("资产评估师(甲方）不能为空")
	}
	if contractHash!=""{
		return shim.Error("所签合同哈希不能为空")
	}
	id:=stub.GetTxID()
	createTime,_:=stub.GetTxTimestamp()
	createTimeStr:=time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	updateTimeStr:=createTimeStr

	assetAppraisalBusiness:=&model.AssetAppraisalBusiness{
		Id: id,
		CreateTime: createTimeStr,
		UpdateTime: updateTimeStr,
		User: user,
		ContractHash: contractHash,
		State:model.UNDER_EVALUATION,
	}
	if err:=service.SignContract(stub,assetAppraisalBusiness);err!=nil{
		return shim.Success([]byte(id))
	}else{
		return shim.Error(err.Error())
	}
}

// CompleteEvaluation 完成评估
func CompleteEvaluation(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)==2{
		return shim.Error("参数个数不满足")
	}
	id:=args[0]
	evaluationResultHash:=args[1]
	if id==""{
		return shim.Error("资产评估业务ID不能为空")
	}
	if evaluationResultHash==""{
		return shim.Error("资产评估结果Hash不能为空")
	}
	if err:=service.CompleteEvaluation(stub,id,evaluationResultHash);err!=nil{
		return shim.Success([]byte("完成评估"))
	}else{
		return shim.Error(err.Error())
	}
}

// ThirdLevelAudit 三级审核
func ThirdLevelAudit(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)!=1{
		return shim.Error("参数个数不满足")
	}
	id:=args[0]
	if id==""{
		return shim.Error("资产评估业务ID不能为空")
	}
	if err:=service.ThirdLevelAudit(stub,id);err!=nil{
		return shim.Success([]byte(id))
	}else{
		return shim.Error(err.Error())
	}
}

// Score 完成评分
func Score(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)!=2{
		return shim.Error("参数个数不满足")
	}
	id:=args[0]
	scoreStr:=args[1]
	if id==""{
		return shim.Error("资产评估业务ID不能为空")
	}
	score,err:=strconv.Atoi(scoreStr)
	if err!=nil{
		return shim.Error("评分只能为数字")
	}
	if err:=service.Score(stub,id,score);err!=nil{
		return shim.Success([]byte(id))
	}else{
		return shim.Error(err.Error())
	}
}


func QueryBusiness(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)!=1{
		return shim.Error("参数个数不满足")
	}
	id:=args[0]
	if business,err:=service.QueryBusiness(stub,id);err!=nil{
		businessByte, _ := json.Marshal(business)
		return shim.Success(businessByte)
	}else{
		return shim.Error(err.Error())
	}
}

