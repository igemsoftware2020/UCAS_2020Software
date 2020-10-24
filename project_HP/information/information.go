/*
指引用户使用
完成提示和导引任务
*/
package information

import(
	"fmt"
	"time"
	"go_code/project_HP/data"
)

//项目整体的说明与参数介绍
func ProjectInformation(){
	fmt.Println("")
	fmt.Println("\t\t\t--------------UCAS_iGEM2020_SHEEP--------------")
	fmt.Println("")
	fmt.Println("This program simulates the complanate process of Hp invading stomach and our modified germ(La) curing Hp relevant diseases")
	//fmt.Println("该项目模拟了幽门螺旋杆菌侵染胃体幽门与我们的工程菌治疗的平面过程")
	fmt.Println("All constants in this program are in \"data\" pack. Researchers can change them if need. These constants are below:")
	//fmt.Println("本程序的定量常数全部在data包中，研究者可以根据自己的需求修改。这些常量显示如下：")
	fmt.Println("\tLine number of Stomach Matrix:", data.RowNum)
	//fmt.Println("\t胃体矩阵的行数 = ", data.RowNum)
	fmt.Println("\tColumn number of Stomach Matrix:", data.ArrNum)
	//fmt.Println("\t胃体矩阵的列数 = ", data.ArrNum)
	fmt.Printf("\tAnti-LL37 mutational proportion in Hp = %.2f %%\n", data.MutationSuperHp)
	//fmt.Printf("\t幽门螺旋杆菌对LL37抗药性的突变体比例 = %.2f %%\n", data.MutationSuperHp)
	fmt.Printf("\tColonization proportion of La in stomach = %.2f %%\n", data.LaReproductionRate)
	//fmt.Printf("\t嗜酸乳杆菌在胃幽门与胃大弯处的定植率 = %.2f %%\n", data.LaReproductionRate)
	fmt.Printf("\tReproduce proportion of Hp = %.2f %%\n", data.HpReproductionRate)
	//fmt.Printf("\t幽门螺旋杆菌的繁殖率 = %.2f %%\n", data.HpReproductionRate)
	fmt.Printf("\tReproduce proportion of La = %.2f %%\n", data.LaReproductionRate)
	//fmt.Printf("\t嗜酸乳杆菌的繁殖率 = %.2f %%\n", data.LaReproductionRate)
	fmt.Printf("\tReproduce proportion of anti-LL37 Hp = %.2f %%\n", data.SuperHpReproductionRate)
	//fmt.Printf("\t抗药性幽门螺旋杆菌的繁殖率 = %.2f %%\n", data.SuperHpReproductionRate)
	fmt.Printf("\tLethal proportion of anti-LL37 Hp = %.2f %%\n", data.SuperHpKilledRate)
	//fmt.Printf("\tLL37对抗药性幽门螺旋菌的杀伤率 = %.2f %%\n", data.SuperHpKilledRate)
	fmt.Println("")
	fmt.Println("\t\t\t------Symbols in Project------\t\t")
	//fmt.Println("\t\t\t------项目中的符号说明如下------\t\t")
	fmt.Println("\tconcrete figure ：means pH in this point; pH > 7.0 means ammonium barrier")
	//fmt.Println("\t具体数值 ：代表胃体中该处的pH值，当 > 7.0时表示由幽门螺杆菌分解尿素产生的铵盐屏障")
	fmt.Println("\tSHEEP : stands for our modified germ(La)")
	//fmt.Println("\tSHEEP : 代表我们的工程嗜酸乳杆菌")
	fmt.Println("\tHp : stands for Hp")
	fmt.Println("\tSuperHp : stands for anti-LL37 Hp")
	fmt.Println("")
	fmt.Println("\t\t\t------Clarification in this Project------\t\t")
	fmt.Println("1. This program simulates the work of a doctor: several rounds of treatment for several patients, each time the doctor is required to observe the test results and write a prescription")
	//fmt.Println("1.本程序模拟医生的工作：对若干名患者进行若干轮治疗，每次需要医生观测检测结果并且开出处方")
	fmt.Println("2. Simulate the environment in the human body by operating the stomach matrix")
	//fmt.Println("2.通过对胃体矩阵的操作模拟人体中的环境")
	fmt.Println("3. The number of operating cycles is self-defined. If you define too much, you must either quit and restart or be patient")
	//fmt.Println("3.操作循环数是自行定义的，如果定义太多，要么退出重来要么耐心操作")
	fmt.Println("4. Random selection of the location of Hp and Sheep")
	//fmt.Println("4.Hp和Sheep的位置随机选择")
	fmt.Println("5. Sheep's movement logic is to randomly walk to any place until it is positioned at the ammonium salt barrier and no longer moves")
	//fmt.Println("5.Sheep的移动逻辑是随机行走到任意地点，直至定位到铵盐屏障处不再移动")
	fmt.Println("6. Regarding the logic of the expression intensity of LL37: 0 means low-intensity expression, which eliminates the Hp up, down, left and right of Sheep; 1 means high-intensity expression, which eliminates Hp in the nine-square grid centered on Sheep")
	//fmt.Println("6.关于LL37表达强度的逻辑：0为低强度表达，消灭Sheep的上下左右的Hp；1为高强度表达，消灭以Sheep为中心的九宫格内的Hp")
	fmt.Println("7. The logic of reproduction: using the Polya jar model, randomly selected, the selected individuals split and reproduce in random directions around")
	//fmt.Println("7.繁殖的逻辑：采用波利亚罐子模型，随机抽取，被抽取到的个体向四周随机方向进行分裂繁殖")
	fmt.Println("8. Sheep's colonization rate logic: only a part of Sheep will not be cleared by the stomach")
	//fmt.Println("8.Sheep的定植率逻辑：只有一部分的Sheep不会被胃体清除")
	fmt.Println("")
	fmt.Println("\t\t\tPlease follow the instructions below")
	//fmt.Println("\t\t\t请阁下按照下述的提示进行操作")
	fmt.Println("")
}


//定义一个全局的双重循环： 外层循环患者数，内层循环患者的治疗数
//患者大外循环
func PatientCycleInformation() (string, int){
	var patientNum int       //用户输入的患者数目
	fmt.Print("Please enter the number of patients to be tested and treated: ")
	//fmt.Print("请输入待检测治疗的患者数目： ")
	fmt.Scanln(&patientNum)
	fmt.Printf("Therer %v patients in total are needy for treatment, Please check\n", patientNum)
	//fmt.Printf("共有 %v 名患者需要检测治疗，请核查", patientNum)
	//提供修改的机会
	var modifyFlag string          //标记用户是否需要修改数据
	for {
		fmt.Print("Do you need to modify the number of patients?\n\"No need\"Please input no, \"need\" Please input yes: ")
		//fmt.Print("请问需要修改患者数吗？\n\"不需要修改\"请输入no，\"需要修改\"请输入yes： ")
		fmt.Scanln(&modifyFlag)
		fmt.Println()
		if modifyFlag == "yes" || modifyFlag == "no" {
			break
		} else {
			fmt.Println("")
			fmt.Println("Something wrong with your input, Please try again")
			//fmt.Println("您的输入有误，请重新输入")
		}
	}
	return modifyFlag, patientNum
}
func ExtraDiagnosisInformation()(continueFlag bool){
	//前一个参数记录用户输入，后一个参数返回是否进行进一步诊疗：返回true表示继续，false表示停止
	var extraFlag string
	for{
		fmt.Println("Enter yes to continue treatment, no to exit treatment")
		//fmt.Println("输入yes表示继续进行治疗，no表示退出治疗")
		fmt.Scanln(&extraFlag)
		if extraFlag == "no" || extraFlag == "yes"{
			break
		} else {
			fmt.Println("Something wrong with your input, Please try again")
			//fmt.Println("您的输入有误，请重新输入")
		}
	}
	if extraFlag == "no"{
		continueFlag = false
	}
	if extraFlag == "yes"{
		continueFlag = true
	}
	return continueFlag
}


//判断是否选择Hp入侵
func Infection()(invadeYes bool){
	fmt.Println("\t\tHP is going to invade! Choose yes and meet the challenge!")
	//fmt.Println("\t\tHP要入侵啦！选择yes吧，迎接挑战！")
	for{
		var invadeFlag string
		fmt.Print("Please input yes: meet the chllenge：")
		//fmt.Print("请输入yes，yes表示接收挑战: ")
		fmt.Scanln(&invadeFlag)
		fmt.Println("")
		if invadeFlag == "yes" {
			fmt.Println("Please wait for a moment")
			//fmt.Println("需要耐心等待一会儿的")
			invadeYes = true
			break
		} else {
			fmt.Println("\t\tPlease enter yes, don’t be a deserter~")
			//fmt.Println("\t\t请输入yes，不要当逃兵哟~")
		}
	}
	return invadeYes
}

//引入SHEEP
func SheepImport()(sheepFlag bool, sheepNum int){
	fmt.Println("\t\tSo next, please let our SHEEP debut~")
	//fmt.Println("\t\t那么接下来，有请我们的SHEEP登场~")
	fmt.Println("Tip: The number of SHEEP does not exceed the number of matrix units / 16")
	//fmt.Println("Tip : SHEEP的个数不超过矩阵单元数 / 16")
	var importNum int
	for {
		fmt.Printf("Please enter a positive integer not exceeding %v, representing the number of imported Sheep:", data.ArrNum * data.RowNum / 16)
		//fmt.Printf("请输入一个不超过%v的正整数，代表引入的Sheep的数量：", data.ArrNum * data.RowNum / 16)
		fmt.Scanln(&importNum)
		fmt.Println("")
	if importNum >= 1 && importNum <= data.ArrNum * data.RowNum / 16 {
		sheepFlag = true
		break
	}else{
		fmt.Println("\t\tPlease try again")
		//fmt.Println("\t\t请进行合法输入")
		}
	}
	sheepNum = importNum
	return sheepFlag, sheepNum
}

//调用Sheep的随机行走定位到铵盐屏障
func SheepMove()(moveFlag bool){
	var moveCommand string
	for{
		fmt.Print("Please input \"move\"Let Sheep locate the ammonium salt barrier by walking randomly:")
		//fmt.Print("请输入\"move\"让Sheep通过随机行走定位到铵盐屏障处：")
		fmt.Scanln(&moveCommand)
		if moveCommand == "move"{
			fmt.Println("\t\t")
			fmt.Println("\t\tSHEEP starts to walk randomly until it is positioned at the ammonium salt barrier (pH> 7.0)")
			//fmt.Println("\t\tSHEEP开始随机行走，直到定位在铵盐屏障(pH > 7.0)处")
			moveFlag = true
			break
		}else{
			fmt.Println("")
		}
	}
	return moveFlag
}

//嗜酸乳杆菌移动到铵盐屏障处后向周围分泌LL37
func LL37Secretion(){
	fmt.Println("\t\tTranscription and translation of antimicrobial peptides, <Hp> means Hp is acted on by antimicrobial peptides")
	//fmt.Println("\t\t抗菌肽的转录翻译,<Hp>表示Hp被抗菌肽所作用")
}

//幽门螺旋杆菌和嗜酸乳杆菌进行繁殖
func BreedInformation(){
	fmt.Println("")
	for i:=3; i >= 0; i--{
		fmt.Printf("The flora in the stomach will complete a round of reproduction in %v seconds \n", i)
		//fmt.Printf("胃体中的菌群将在 %v 秒内完成一轮繁殖 \n", i)
		time.Sleep(1*time.Second)
	}

}

//嗜酸乳杆菌的自杀程序
func SuicideInformation(){
	fmt.Println("\t\tLactobacillus acidophilus initiates suicide program")
	//fmt.Println("\t\t嗜酸乳杆菌启动自杀程序")
	fmt.Println("Lactobacillus acidophilus turn on the suicide switch after leaving the human environment (below 37°C)")
	//fmt.Println("工程菌需要在脱离人体环境（低于37℃）后开启自杀开关")
	for i:=5; i >= 0; i--{
		fmt.Printf("Engineered Lactobacillus acidophilus will be excluded from the body after %v seconds...\n", i)
		//fmt.Printf("工程嗜酸乳杆菌将在 %v 秒后排除体外...\n", i)
		time.Sleep(1*time.Second)
	}
	fmt.Println("The results of the treatment are as follows")
	//fmt.Println("治疗结果如下")
}

//退出程序
func Exit(){
	fmt.Println("Please enter any character to exit the program")
	//fmt.Println("请输入任意字符退出程序")
	var exit string
	fmt.Scanln(&exit)
}
