package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"goSmartContract/model"
	"goSmartContract/pkg/utils"
	"time"
)

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/22 17:22

 * @Note:

 **/




func SignContract(stub shim.ChaincodeStubInterface ,business *model.AssetAppraisalBusiness)error{
	if err := utils.WriteLedger(business, stub, model.ASSET_APPRAISAL_BUSINESS, []string{business.Id}); err != nil {
		return err
	}

	return nil
}



func CompleteEvaluation(stub shim.ChaincodeStubInterface ,id string,resultHash string)error{
	result,err:=utils.GetStateByPartialCompositeKeys(stub,model.ASSET_APPRAISAL_BUSINESS,[]string{id})
	if err!=nil{
		return err
	}
	if len(result)!=1{
		return fmt.Errorf("无法找到该业务")
	}
	var business model.AssetAppraisalBusiness
	if err = json.Unmarshal(result[0],business); err != nil {
		return fmt.Errorf("AssetAppraisalBusiness-反序列化出错: %s", err)
	}
	if business.State!=model.UNDER_EVALUATION{
		return fmt.Errorf("资产评估业务未处于评估中")
	}
	currTime,_:=stub.GetTxTimestamp()
	business.UpdateTime=time.Unix(int64(currTime.GetSeconds()), int64(currTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	business.ContractHash=resultHash
	business.State=model.THREE_LEVEL_AUDIT
	if err := utils.WriteLedger(business, stub, model.ASSET_APPRAISAL_BUSINESS, []string{business.Id}); err != nil {
		return err
	}
	return nil
}


func QueryBusiness(stub shim.ChaincodeStubInterface,id string)(business *model.AssetAppraisalBusiness,err error){
	result,err:=utils.GetStateByPartialCompositeKeys(stub,model.ASSET_APPRAISAL_BUSINESS,[]string{id})
	if err!=nil{
		return nil,err
	}
	if len(result)!=1{
		return nil,fmt.Errorf("无法找到该业务")
	}
	err = json.Unmarshal(result[0], business)
	if err != nil {
		return nil, err
	}
	return
}


func ThirdLevelAudit(stub shim.ChaincodeStubInterface ,id string)error{
	result,err:=utils.GetStateByPartialCompositeKeys(stub,model.ASSET_APPRAISAL_BUSINESS,[]string{id})
	if err!=nil{
		return err
	}
	if len(result)!=1{
		return fmt.Errorf("无法找到该业务")
	}
	var business model.AssetAppraisalBusiness
	if err = json.Unmarshal(result[0],business); err != nil {
		return fmt.Errorf("AssetAppraisalBusiness-反序列化出错: %s", err)
	}
	if business.State!=model.THREE_LEVEL_AUDIT{
		return fmt.Errorf("资产评估业务未处于三级审核中")
	}
	currTime,_:=stub.GetTxTimestamp()
	business.UpdateTime=time.Unix(int64(currTime.GetSeconds()), int64(currTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	business.State=model.WAIT_APPRAISE
	if err := utils.WriteLedger(business, stub, model.ASSET_APPRAISAL_BUSINESS, []string{business.Id}); err != nil {
		return err
	}
	return nil
}


func Score(stub shim.ChaincodeStubInterface ,id string,score int)error{
	result,err:=utils.GetStateByPartialCompositeKeys(stub,model.ASSET_APPRAISAL_BUSINESS,[]string{id})
	if err!=nil{
		return err
	}
	if len(result)!=1{
		return fmt.Errorf("无法找到该业务")
	}
	var business model.AssetAppraisalBusiness
	if err = json.Unmarshal(result[0],business); err != nil {
		return fmt.Errorf("AssetAppraisalBusiness-反序列化出错: %s", err)
	}
	if business.State!=model.WAIT_APPRAISE{
		return fmt.Errorf("资产评估业务未处于等待评价")
	}
	currTime,_:=stub.GetTxTimestamp()
	business.UpdateTime=time.Unix(int64(currTime.GetSeconds()), int64(currTime.GetNanos())).Local().Format("2006-01-02 15:04:05")
	business.State=model.FINISH
	business.Score=score
	if err := utils.WriteLedger(business, stub, model.ASSET_APPRAISAL_BUSINESS, []string{business.Id}); err != nil {
		return err
	}
	return nil
}