/*
introduce SHEEP
引入嗜酸乳杆菌
SHEEP move randomly until find the ammonium barrier
嗜酸乳杆菌通过随机行走的方式定位到铵盐屏障处
secrete LL-37 
分泌抗菌肽==>对幽门螺旋菌和变异体的杀伤效果不同
SHEEP abscond
嗜酸乳杆菌排除体外
*/
package sheep

import(
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"go_code/project_HP/data"
)

//introduce SHEEP
//引入一定数目的工程嗜酸乳杆菌
func Sheep(n int, StomachMatrix *[data.RowNum][data.ArrNum]string){
	for i:=0; i < n; i++{
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		randArr := rand.Intn(data.ArrNum)   //随机列
		time.Sleep(10*time.Millisecond)  //休眠一段时间，让结果更随机
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		randRow := rand.Intn(data.RowNum)   //随机行
		//随机的设置n个SHEEP在矩阵中
		if StomachMatrix[randRow][randArr] == "Hp"||StomachMatrix[randRow][randArr] == "SHEEP"||StomachMatrix[randRow][randArr] == "SuperHp" {
			i--
		}else{
			StomachMatrix[randRow][randArr] = "SHEEP"
		}
	}
}

//SHEEP move randomly in order to find Hp
//嗜酸乳杆菌随机行走，定位到铵盐屏障
func Move(StomachMatrix *[data.RowNum][data.ArrNum]string){      
	var slice_row []int = make([]int, 0)   //寄存Leon所在行的信息
	var slice_arr []int = make([]int, 0)   //寄存Leon所在列的信息                                               
	for i:=0; i < len(StomachMatrix); i ++{
		for j:=0; j < len(StomachMatrix[i]); j++{
			if StomachMatrix[i][j] == "SHEEP"{
				slice_row = append(slice_row, i)
				slice_arr = append(slice_arr, j)
			}
		}
	}
	for k:=0; k<len(slice_arr); k++{
		leon_i, leon_j := slice_row[k], slice_arr[k]  //标记Leon原来所处的位置
		       //Leon的新位置
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		fake_i :=  leon_i + rand.Intn(data.RowNum)    //伪行:加上的数在0 ~ RowNum - 1的闭区间内
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		fake_j := leon_j + rand.Intn(data.ArrNum)      //伪列
		i := fake_i % data.RowNum                    //Leon的新位置
		j := fake_j % data.ArrNum                    //取余，使得运算封闭
		//如果感受到铵盐屏障：Leon定植，原来位置改为酸性环境
		pH,_ := strconv.ParseFloat(StomachMatrix[i][j],64)
		if pH > 6.9 {
			StomachMatrix[i][j] = "SHEEP"
			rand.Seed(time.Now().UnixNano())       //引入随机种子
			k,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0.9 + rand.Float64()), 64)
			StomachMatrix[leon_i][leon_j] = fmt.Sprintf("%2.2f", k)
		} else {  //否则：回到原来的位置，当刚刚什么都没有发生过
			k--
		}
	}
}

//KILL ==> LL-37
//本包使用的kill函数将指定位点的幽门螺旋杆菌杀灭，对抗药菌的杀灭效果有限
func kill(a int, b int, StomachMatrix *[data.RowNum][data.ArrNum]string){
	if StomachMatrix[a][b] == "Hp"{
		StomachMatrix[a][b] = "<>"
	}
	if StomachMatrix[a][b] == "SuperHp"{
		k := data.Probability(data.SuperHpKilledRate)     //为false时SuperHp被杀
		if !k {
			StomachMatrix[a][b] = "<>"
		}
	}
}
//users choose the expression intensity of LL-37
//1 means : kill the Hp in nine patch field
//0 means : kill the Hp in cross field
//抗菌肽表达强度供用户选择：1表示向周围九宫格范围内分泌抗菌肽，0表示向周围十字架范围内分泌抗菌肽
func ProduceKiller( StomachMatrix *[data.RowNum][data.ArrNum]string){                                        //生产抗菌肽
	var expressFlag int
	
	for {
		fmt.Print("Please input 0 / 1, 0 means basal expression intensity, 1 means high level:")
		//fmt.Print("请输入0或者1，代表抗菌肽的表达强度：0表示基础表达，1表示增强表达:")
		fmt.Scanln(&expressFlag)
		if expressFlag == 0 || expressFlag == 1{
			break
		}else{
			fmt.Println("Some input wrong, Please try again")
		}
	}

	//定义名称变量方便下述输入
	arrBoundary := data.ArrNum - 1
	rowBoundary := data.RowNum - 1
	s := "SHEEP"
	sm := StomachMatrix

	if expressFlag == 1{                                   //强化表达九宫格
		if sm[0][0] == s{                    //左上顶点
			kill(0, 1, sm)
			kill(1, 0, sm)
			kill(1, 1, sm)
		}
		if sm[0][arrBoundary] == s{                    //右上顶点
			kill(0, arrBoundary - 1, sm)
			kill(1, arrBoundary, sm)
			kill(1, arrBoundary - 1, sm)
		}
		if sm[rowBoundary][arrBoundary] == s{                    //右下顶点
			kill(rowBoundary - 1, arrBoundary - 1, sm)
			kill(rowBoundary - 1, arrBoundary, sm)
			kill(rowBoundary, arrBoundary - 1, sm)
		}
		if sm[rowBoundary][0] == s{                    //左下顶点
			kill(rowBoundary - 1, 0, sm)
			kill(rowBoundary - 1, 1, sm)
			kill(rowBoundary, 1 , sm)
		}
		for k:=1; k <= rowBoundary - 1; k++{                  //左边界
				if sm[k][0] == s{
					kill(k -1, 0, sm)
					kill(k - 1, 1, sm)
					kill(k, 1, sm)
					kill(k + 1, 0, sm)
					kill(k + 1, 1, sm)
				}
			}
		for k:=1; k <=arrBoundary - 1; k++{                  //上边界
			if sm[0][k] == s{
				kill(0, k - 1, sm)
				kill(0, k + 1, sm)
				kill(1, k - 1, sm)
				kill(1, k, sm)
				kill(1, k + 1, sm)
			}
		}
		for k:=1; k <= rowBoundary - 1; k++{                    //右边界
			if sm[k][arrBoundary] == s{
				kill(k - 1, arrBoundary - 1, sm)
				kill(k - 1,arrBoundary, sm)
				kill(k, arrBoundary - 1, sm)
				kill(k + 1, arrBoundary - 1, sm)
				kill(k + 1, arrBoundary - 1, sm)
			}
		}
        for k:=1; k <= arrBoundary - 1; k++{                     //下边界
			if sm[rowBoundary][k] == s{
				kill(rowBoundary - 1, k - 1, sm)
				kill(rowBoundary - 1, k, sm)
				kill(rowBoundary - 1, k + 1, sm)
				kill(rowBoundary, k + 1, sm)
				kill(rowBoundary, k - 1, sm)
			}
		}
		for i:=1; i <= arrBoundary - 1; i++{                      //内点
			for j:=1; j <= rowBoundary - 1;  j++{
				if sm[i][j] == s{
					for s:= i - 1; s <= i + 1; s++{
						for t:=j - 1; t <= j + 1; t++{
							kill(s, t, sm)
						}
					}
				}
			}
		}
	}
	if expressFlag == 0 {                                 //基础表达十字架
		if sm[0][0] == s{                    //左上顶点
			kill(0, 1, sm)
			kill(1, 0, sm)
		}
		if sm[0][arrBoundary] == s{                    //右上顶点
			kill(0, arrBoundary - 1, sm)
			kill(1, arrBoundary, sm)
		}
		if sm[rowBoundary][arrBoundary] == s{                    //右下顶点
			kill(rowBoundary, arrBoundary - 1, sm)
			kill(rowBoundary - 1, arrBoundary, sm)
		}
		if sm[rowBoundary][0] == s{                    //左下顶点
			kill(rowBoundary - 1, 0, sm)
			kill(rowBoundary, 1,sm)
		}
		for i:= 1; i <= rowBoundary - 1; i++{                              //左边界
			if sm[i][0] == s{
				kill(i - 1, 0, sm)
				kill(i + 1, 0, sm)
				kill(i, 1, sm)
			}
		}
		for i:= 1; i <= arrBoundary - 1; i++{                              //上边界
			if sm[0][i] == s{
				kill(0, i - 1, sm)
				kill(0, i + 1, sm)
				kill(1, i, sm)
			}
		}
		for i:= 1; i <= rowBoundary - 1; i++{                              //右边界
			if sm[i][arrBoundary] == s{
				kill(i, arrBoundary - 1, sm)
				kill(i - 1, arrBoundary, sm)
				kill(i + 1, arrBoundary, sm)
			}
		}
		for i:= 1; i <= arrBoundary - 1; i++{                              //下边界
			if sm[rowBoundary][i] == s{
				kill(rowBoundary - 1, i, sm)
				kill(rowBoundary, i - 1, sm)
				kill(rowBoundary, i + 1, sm)
			}
		}
		for i:=1; i <= arrBoundary - 1; i++{
			for j:=1; j<= rowBoundary - 1; j++{
				if sm[i][j] == s{
					kill(i, j - 1, sm)
					kill(i - 1, j, sm)
					kill(i + 1, j, sm)
					kill(i, j + 1, sm)
				}
			}
		}
	}
	for i:=0; i < len(sm); i++{
		for j:=0; j < len(sm[i]); j++{
			if sm[i][j] == "<>"{
				rand.Seed(time.Now().UnixNano())       //引入随机种子
				n, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",0.9 + rand.Float64()), 64)
				m := fmt.Sprintf("%2.2f", n)
				sm[i][j] = m
			}
		}
	}
}

//some SHEEP discharge out of boby
//根据定植率数据，部分嗜酸乳杆菌排除体外，在原处留下酸性环境
func Suicide(StomachMatrix *[data.RowNum][data.ArrNum]string){                                              //自杀开关
	for i:=0; i < len(StomachMatrix); i++{
		for j:=0; j < len(StomachMatrix[i]); j++{
			if StomachMatrix[i][j] == "SHEEP"{
				laColonization := data.Probability(data.LaColonizationRate) //true意味着超出定植率，被排出
				if laColonization {
					rand.Seed(time.Now().UnixNano())       //引入随机种子
					k,_ := strconv.ParseFloat(fmt.Sprintf("%.2f", 0.9 + rand.Float64()), 64)
					StomachMatrix[i][j] = fmt.Sprintf("%2.2f", k)
				}
			}
		}
	}
}
