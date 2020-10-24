/*
统计患者规模
对数据的回归分析
*/
package patient

import(
	"fmt"
	"go_code/project_HP/data"
	"go_code/project_HP/stomach"
	"go_code/project_HP/information"
)

//大外层：患者循环，关键在要初始化胃体矩阵
func PatientCycle()(){
	modifyFlag, patientNum := information.PatientCycleInformation()
	if modifyFlag == "no"{
		for PatientRange := 1; PatientRange <= patientNum; PatientRange++{
			fmt.Printf("\tPatient %v 's information is below:\n\n", PatientRange)
			//fmt.Printf("\t第 %v 号患者的数据如下 : \n\n", PatientRange)
			//初始化胃矩阵
			var StomachMatrix [data.RowNum][data.ArrNum]string
			//操作的主体
			treatmentCycle(&StomachMatrix)
		}
		fmt.Printf("Tere are %v patients in total finishing treatment, statistical result for whole is below: \n", patientNum)
		//fmt.Printf("总共 %v 名患者诊疗完毕，对全体的统计结果见下述分析：  \n", patientNum)
	} else if modifyFlag == "yes"{
		PatientCycle()
	}
}

//pre-treatment : tree times
//三次先干预诊疗
func triDiagnosis(StomachMatrix *[data.RowNum][data.ArrNum]string){
	for preDiagnosis := 1; preDiagnosis <= 3; preDiagnosis++{
		fmt.Printf("\t")
		fmt.Printf("\tTreatment cycle : %v\n ", preDiagnosis)
		//fmt.Printf("\t该患者的第 %v 次治疗\n", preDiagnosis)
		stomach.StomachBody(StomachMatrix)
	}
	fmt.Println("\tbasic treatment for patient has ended, Please decide continue or not")
	//fmt.Println("\t该患者的保底治疗结束，请研究者判断是否还要继续进行治疗")
}
//研究者自己决定是否进行的诊疗
func extraDiagnosis(StomachMatrix *[data.RowNum][data.ArrNum]string){
	fmt.Println("\tthis patient's condition is above, Please decide additional treatment or not")
	//fmt.Println("\t该患者的目前状态如上所示，请研究者人工判断是否需要进行额外的治疗")
	for {
		var treatmentCount int = 3
		if information.ExtraDiagnosisInformation() {    //继续治疗
			treatmentCount++
			fmt.Printf("Treatment cycle: %v is Running\n", treatmentCount)
			//fmt.Printf("目前正在进行的是第 %v 个疗程\n", treatmentCount)
			stomach.StomachBody(StomachMatrix)
		} else {
			fmt.Println("Treatment end, Please check the comprehensive report")
			//fmt.Println("疗程结束，请查看综合分析报告")
			break
		}
	}
}
//内层疗程循环，关键在不初始化胃体矩阵
func treatmentCycle(StomachMatrix *[data.RowNum][data.ArrNum]string){
	fmt.Println("Each patient will have at least tree treatment, with an evaluation attached; Let researchers decide additional treatment")
	//fmt.Println("每个患者至少进行三次的疗程，会返回一次评估；由研究者决定是否继续进行诊疗")
	fmt.Println("After thar every treatment will feedback an evaluation")
	///fmt.Println("此后的每一次疗程结束都会有系统对平衡态的评估，依旧交给研究者决定是否继续进行诊疗")
	//三次先干预诊疗
	triDiagnosis(StomachMatrix)
	//后续研究者自行决定进行的诊疗
	extraDiagnosis(StomachMatrix)
}
