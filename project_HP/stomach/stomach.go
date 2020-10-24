/*
stomach produces gastric acid
胃体产生胃酸，根据进食调节
draw the matrix
绘制矩阵统计结果
*/
package stomach

import(
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"go_code/project_HP/data"
	"go_code/project_HP/information"
	"go_code/project_HP/hp"
	"go_code/project_HP/sheep"
)

//
func StomachBody(StomachMatrix *[data.RowNum][data.ArrNum]string){
	//简写名称
	//for abbreviation
	sm := StomachMatrix
	//dicide pH by the last time of meal, and print the result
	//根据饮食时间来判断胃体中酸度值，并且打印结果
	ProduceAcid(sm)  
	Draw(sm)         
	
	//Hp is going to invade, forming ammonium wall
	//Hp即将入侵,构建铵盐屏障，并且已经有抗药突变体
	for{
		if information.Infection() {
			hp.HPinvade(sm)
			hp.Mutant(sm)
			hp.ProduceN(sm)
			break
		}
	}
	Draw(sm)         

	//SHEEP come on
	//接下来欢迎我们的小绵羊上场吧
	sheepFlag, sheepNum := information.SheepImport()
	if sheepFlag {
		sheep.Sheep(sheepNum, sm)
	}
	Draw(sm)

	//SHEEP move to ammonium barrier
	//准备调用Move让嗜酸乳杆菌定位到铵盐屏障
	if information.SheepMove() {
		sheep.Move(sm)
		Draw(sm)
	}

	//SHEEP find and kill Hp
	//Sheep定位到Hp，并且将其打成<Hp>,表示抗菌肽作用到了Hp
	information.LL37Secretion()
	sheep.ProduceKiller(sm)
	Draw(sm)

	//Both SHEEP and Hp are going to reproduce
	//嗜酸乳杆菌和幽门螺旋杆菌开始繁殖，注意其中抗药菌的繁殖率低于野生型
	information.BreedInformation()
	data.Polya(data.StrHp, data.HpReproductionRate, sm)
	data.Polya(data.StrSuperHp, data.SuperHpReproductionRate, sm)
	data.Polya(data.StrSheep, data.LaReproductionRate, sm)
	Draw(sm)

	//SHEEP beging to commit suicide
	//嗜酸乳杆菌启动自杀程序，
	information.SuicideInformation()
	sheep.Suicide(sm)
	Draw(sm)
}

//gastric acid
//胃体产生胃酸，根据是否进食调节pH值
func ProduceAcid(StomachMatrix *[data.RowNum][data.ArrNum]string)(){
	//简写名称
	//for abbreviation
	sm := StomachMatrix
	var eatFlag string 
	for{
		fmt.Print("Please input \"yes/no\"，yes stands for patient has eaten，no stands for patient has not:")
		fmt.Scanln(&eatFlag)
		if eatFlag == "yes" || eatFlag == "no"{
			fmt.Println("\tthe first time of matrix initialization，Please wait for a moment")
			break
		} else {
			fmt.Println("You got wrong input, Please try again")
		}
	}
	if eatFlag == "no"{                                        //did not eat，pH = 3.2 +- o(1)
		for i:=0; i < len(sm); i++{
			for j:=0; j < len(sm[i]); j++{
				if sm[i][j] != "SHEEP" && sm[i][j] != "Hp" && sm[i][j] != "SuperHp"{
					rand.Seed(time.Now().UnixNano())       
					n, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",3.1 + rand.Float64()), 64)
					time.Sleep(1*time.Millisecond)  
					m := fmt.Sprintf("%2.2f", n)
					sm[i][j] = m
				}	
			}
		} 
	} else if eatFlag == "yes"{                                 //had eaten, pH = 0.9 + o(1)
		for i:=0; i < len(sm); i++{
			for j:=0; j < len(sm[i]); j++{
				if sm[i][j] != "SHEEP" && sm[i][j] != "Hp" && sm[i][j] != "SuperHp"{
					rand.Seed(time.Now().UnixNano())       
					n, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",0.9 + rand.Float64()), 64)
					time.Sleep(1*time.Millisecond)  
					m := fmt.Sprintf("%2.2f", n)
					sm[i][j] = m
				}
			}
		} 
	}
}

//count the number of germ community
//显示二维数组中的数值，并统计菌群
func Draw(StomachMatrix *[data.RowNum][data.ArrNum]string){
	fmt.Println("")
	fmt.Println("")
	fmt.Println("\t--------------------------SHEEP--------------------------")
	//遍历二维数组，是否打印矩阵可由用户自己选择
	fmt.Print("Please input 0 / 1: 0 means NOT printing the whole matrix, 1 is opposed(it will occupying the whole screen): ")
	//fmt.Print("请输入0或者1表示选择是否矩阵：0代表不打印，1表示打印（会占满屏幕的）: ")
	var matrixFlaf int = 0
	fmt.Scanln(&matrixFlaf)
	for {
		if matrixFlaf == 0 || matrixFlaf == 1{
			break
		} else {
			fmt.Println("Some inputs wrong, Please try again")
			//fmt.Println("请重新输入正确的数值")
			fmt.Scanln(&matrixFlaf)
		}
	}
	if matrixFlaf == 1 {
		for i:=0; i < len(StomachMatrix); i++{
			for j:=0; j < len(StomachMatrix[i]); j++{
					fmt.Printf("%v\t", StomachMatrix[i][j])
			}
			fmt.Println("")
			fmt.Println("")
		}
		fmt.Println("\t\t\tthe number here stands for pH in pylorus")
		//fmt.Println("\t\t\t其中的数字代表胃体幽门处的pH值")
		fmt.Println(" ")
	}
	//Counting SHEEP
	//遍历矩阵，统计Sheep定植量
	SheepNum := data.Count(StomachMatrix, data.StrSheep)
	fmt.Printf("\tnumber of SHEEP：%v\n", SheepNum)
	//Count Hp
	//遍历矩阵，统计Hp定植量
	HpNum := data.Count(StomachMatrix, data.StrHp)
	fmt.Printf("\tnumber of relic Hp：%v\n", HpNum)
	//Counting SuperHp
	//遍历矩阵，统计SuperHp的定植量，反应耐药性
	SuperHpNum := data.Count(StomachMatrix, data.StrSuperHp)
	fmt.Printf("\tnumber of relic Super Hp：%v\n", SuperHpNum)
	fmt.Println("")
}
