/*
定义常数
统计群落成员
群落成员进行繁殖繁殖
*/
package data

import(
	"math/rand"
	"time"
	"fmt"
	"strconv"
)
/**/
const RowNum = 80                                         //胃体矩阵的行数
const ArrNum = 80                                        //胃体矩阵的列数
const MutationSuperHp = 18.2     //Hp中有18.2%的数目对LL-37有抗药性
const LaColonizationRate = 80.70     //嗜酸乳杆菌的定植率
const HpReproductionRate = 10.28     //Hp的繁殖率
const LaReproductionRate = 10.61	    //嗜酸乳杆菌的繁殖率
const SuperHpReproductionRate = 6.24 //超级耐药菌的繁殖率
const SuperHpKilledRate = 40.72//LL37对SuperHp的杀伤率
const DeviationRatio = 40.00   //数据误差超过这个就删除

const StrSheep = "SHEEP"
const StrHp = "Hp"
const StrSuperHp = "SuperHp"

/*计数函数*/
func Count(StomachMatrix *[RowNum][ArrNum]string, nameCount string)(countNum int){
	for i:=0; i < len(StomachMatrix); i++{
		for j:=0; j < len(StomachMatrix[i]); j++{
			if StomachMatrix[i][j] == nameCount{
				countNum++
			}
		}
	}
	return countNum
}

/*用随机数表达概率的函数 ： 超过输入的值cuttingPoint返回true*/
func Probability(cuttingPoint float64)(comparsionResult bool){                                   //用随机数的分布代替概率，皆以%结算
	for{
		if cuttingPoint >= 0 && cuttingPoint <= 100{       //代表概率分界点位为cuttiPoint%
			break
		}
		fmt.Println("Please input a decimal from 0 to 100:")
		//fmt.Println("请合法输入0~100之间的小数：")
	}
	rand.Seed(time.Now().UnixNano())  //引入随机种子
	time.Sleep(2*time.Millisecond)  //休眠一段时间，让结果更随机
	n, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", 100 * rand.Float64()), 64)
	if n >= cuttingPoint{
		comparsionResult = true                 //随机数大于判断点则返回true
	} else {
		comparsionResult = false               //随机数小于判断点则返回false
	}
	return comparsionResult
}

//检查该位点是否被占据，如果没有则分裂到此处
func checkAndProduce(a int, b int, StomachMatrix *[RowNum][ArrNum]string, name string)(){
	//简写名称
	sm := StomachMatrix
	if sm[a][b] != "Hp" && sm[a][b] != "SHEEP" && sm[a][b] != "SuperHp"{
		sm[a][b] = name
	}
}
//菌群的繁殖类似于波利亚罐子模型，每一个个体有一定的机率被选中，被选中的个体可以繁殖
func Polya(name string, reproduceRate float64 , StomachMatrix *[RowNum][ArrNum]string){
	//name = {Hp, SuperHp, Leon}
	for i:=0; i < len(StomachMatrix); i++{
		for j:=0; j < len(StomachMatrix[i]); j++{
			if StomachMatrix[i][j] == name {             
				reproduceFlag := Probability(reproduceRate)  //繁殖标杆，false时繁殖
				if !reproduceFlag {
					switch {
						//左上顶点
					case i == 0 && j == 0:                    
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(2)  //两个方向可供繁殖，任意选取一个
						if k == 0 {   //同行不同列
							checkAndProduce(0, 1, StomachMatrix, name)
						}else {  //同列不同行
							checkAndProduce(1, 0, StomachMatrix, name)
						}
						//右上顶点
					case i == 0 && j == ArrNum - 1:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(2)
						if k == 0{     //同行不同列
							checkAndProduce(0, ArrNum - 2, StomachMatrix, name)
						} else { //同列不同行
							checkAndProduce(1, ArrNum - 1, StomachMatrix, name)
						}
						//右下顶点
					case i == RowNum - 1 && j == ArrNum - 1:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(2)
						if k == 0{  //同行不同列
							checkAndProduce(RowNum - 1, ArrNum - 2, StomachMatrix, name)
						} else {//同列不同行
							checkAndProduce(RowNum - 2, ArrNum - 1, StomachMatrix, name)
						}
						//左下顶点
					case i == RowNum -1 && j == 0:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(2)
						if k == 0{      //同行不同列
							checkAndProduce(RowNum - 1, 1, StomachMatrix, name)
						} else {
							checkAndProduce(RowNum - 2, 0, StomachMatrix, name)
						}
						//左边界
					case i >= 1 && i <= RowNum - 2 && j == 0:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(4)       //生成0，1，2代表三个方位
						if k == 0 {        //上方
							checkAndProduce(i - 1, j, StomachMatrix, name)
						} else if k == 1{       //右方
							checkAndProduce(i, 1, StomachMatrix, name)
						} else {            //下方
							checkAndProduce(i + 1, 0, StomachMatrix, name)
						}
						//上边界
					case i == 0 && j >= 1 && j <= ArrNum -2:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(4)
						if k == 0{      //左方
							checkAndProduce(i, j - 1, StomachMatrix, name)
						} else if k == 1{       //下方
							checkAndProduce(i + 1, j, StomachMatrix, name)
						} else {                //右方
							checkAndProduce(i, j + 1, StomachMatrix, name)
						}
						//右边界
					case i >= 1 && i <= RowNum - 2 && j == ArrNum - 1:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(4)
						if k == 0{        //上方
							checkAndProduce(i - 1, j, StomachMatrix, name)
						} else if k == 1{    //左方
							checkAndProduce(i, j - 1, StomachMatrix, name)
						} else {        //下方
							checkAndProduce(i + 1, j, StomachMatrix, name)
						}
						//下边界
					case i == RowNum - 1 && j >= 1 && j <= ArrNum - 2:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(4)
						if k == 0{          //左方
							checkAndProduce(i, j - 1, StomachMatrix, name)
						} else if k == 1{       //上方
							checkAndProduce(i - 1, j, StomachMatrix, name)
						} else {              //右方
							checkAndProduce(i, j + 1, StomachMatrix, name)
						}
						//内点
					default:
						rand.Seed(time.Now().UnixNano())  //引入随机种子
						k := rand.Intn(5)   //四个方位
						if k == 0 {             //上方
							checkAndProduce(i - 1, j, StomachMatrix, name)
						} else if k  == 1{       //左方
							checkAndProduce(i, j - 1, StomachMatrix, name)
						} else if k == 2{        //下方
							checkAndProduce(i - 1, j, StomachMatrix, name)
						} else {    //右方
							checkAndProduce(i, j + 1, StomachMatrix, name)
						} 
					}
				}
			}
		}
	}
}
