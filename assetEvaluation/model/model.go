package model

/**

 * @Author: AloneAtWar

 * @Date:   2022/10/19 13:10

 * @Note:

 **/

// AssetAppraisalBusiness  资产评估业务
type AssetAppraisalBusiness struct {
	Id string `json:"id"`  // 资产评估业务Id
	User string `json:"user"`  // 资产评估需求者(乙方）
	AssetAppraiser string `json:"assetAppraiser"` //资产评估师(甲方）
	CreateTime string `json:"createTime"`  //创建时间
	UpdateTime string `json:"updateTime"` //更新时间
	State string `json:"state"`	//资产评估业务状态
	ContractHash string `json:"contractHash"`  //所签合同哈希
	EvaluationResultHash string `json:"evaluationResultHash"` //评估结果哈希
	Score int `json:"score"` //评分
}



const (
	ASSET_APPRAISAL_BUSINESS="asset-appraisal-business"
)


// 资产评估状态
const (
	UNDER_EVALUATION="评估中"
	THREE_LEVEL_AUDIT ="三级审核中"
	WAIT_APPRAISE="等待评价"
	FINISH="评估结束"
)
