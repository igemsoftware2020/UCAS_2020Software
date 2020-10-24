/*
根据硬件探测的结果判断折合成矩阵内合适的幽门螺旋杆菌的数目
幽门螺旋菌中有不部分抗药突变体
幽门螺旋菌生成铵盐屏障
*/
package hp

import(
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"go_code/project_HP/data"
)

//幽门螺旋杆菌入侵，数目依靠硬件检测的反馈结果
func HPinvade(StomachMatrix *[data.RowNum][data.ArrNum]string)(){
	//fmt.Print("请输入硬件胶囊检测到的值（个/ml）（）：")
	
	var hpNum = data.RowNum * data.ArrNum / 10   //单位体积内个数
	for i := 0; i < hpNum; i++ {
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		randArr := rand.Intn(data.ArrNum)   //随机列
		time.Sleep(2*time.Millisecond)  //休眠一段时间，让结果更随机
		rand.Seed(time.Now().UnixNano())       //引入随机种子
		randRow := rand.Intn(data.RowNum)   //随机行
		//随机的设置hpNum个Hp在矩阵中
		if StomachMatrix[randRow][randArr] == data.StrHp{
			i--
		}else{
			StomachMatrix[randRow][randArr] = data.StrHp
		}
	}
}

//幽门螺旋杆菌有部分抗药突变体，LL37对其杀伤效果有限
func Mutant(StomachMatrix *[data.RowNum][data.ArrNum]string){
	for i:=0; i < len(StomachMatrix); i++{
		for j:=0; j < len(StomachMatrix[i]); j++{
			if StomachMatrix[i][j] == data.StrHp {
				mutationRate := data.Probability(data.MutationSuperHp)
				if !mutationRate {              //当mutationRate为false时，突变为耐药的SuperHp        
					StomachMatrix[i][j] = data.StrSuperHp
				}
			}
		}
	}
}

//将指定位点处的pH改成中性或者弱碱性
func basicity(a int, b int, StomachMatrix *[data.RowNum][data.ArrNum]string){
	if StomachMatrix[a][b] != data.StrHp && StomachMatrix[a][b] != data.StrSuperHp && StomachMatrix[a][b] != data.StrSheep{
		rand.Seed(time.Now().UnixNano())       //引入随机种子
			n, _ := strconv.ParseFloat(fmt.Sprintf("%.2f",6.9 + rand.Float64()), 64)
			m := fmt.Sprintf("%2.2f", n)
			StomachMatrix[a][b] = m
	}
}
//幽门螺旋杆菌产生的铵盐屏障，铵盐屏障为九宫格范围
func ProduceN(StomachMatrix *[data.RowNum][data.ArrNum]string){
	//定义几个变量代替下面的反复书写
	arrBoundary := data.ArrNum - 1
	rowBoundary := data.RowNum - 1
	hp := data.StrHp
	sh := data.StrSuperHp

	if StomachMatrix[0][0] == hp || StomachMatrix[0][0] == sh {                              //左上顶点
		for s:=0; s <= 1; s++{
			for t:=0; t <= 1; t++{
				basicity(s, t, StomachMatrix)
			}
		}
	}
	if StomachMatrix[0][arrBoundary] == hp ||  StomachMatrix[0][arrBoundary] == sh{                    //右上顶点
		for s:=0; s <= 1; s++{
			for t:=arrBoundary - 1; t <= arrBoundary; t++{
				basicity(s, t, StomachMatrix)
			}
		}
	}
	if StomachMatrix[rowBoundary][0] == hp || StomachMatrix[rowBoundary][0] == sh{                 //左下顶点
		for s:=rowBoundary - 1; s <= rowBoundary; s++{
			for t:=0; t <= 1; t++{
				basicity(s, t, StomachMatrix)
			}
		}
	}
	if StomachMatrix[rowBoundary][arrBoundary] == hp || StomachMatrix[rowBoundary][arrBoundary] == sh{             //右下顶点
		for s:=rowBoundary - 1; s <= rowBoundary; s++{
			for t := arrBoundary - 1; t <= arrBoundary; t++{
				basicity(s, t, StomachMatrix)
			}
		}
	}
	for k:=1; k <= rowBoundary - 1; k++{                   //左边界
		if StomachMatrix[k][0] == hp || StomachMatrix[k][0] == sh{
			for s:=k - 1; s <= k + 1;s++{
				for t:=0; t <= 1; t++{
					basicity(s, t, StomachMatrix)
				}
			}
		}
	}
	for k:=1; k <= arrBoundary - 1; k++{                   //上边界
		if StomachMatrix[0][k] == hp  || StomachMatrix[0][k] == sh{
			for s:=0; s <= 1;s++{
				for t:=k-1; t <= k+1; t++{
					basicity(s, t, StomachMatrix)
				}
			}
		}
	}
	for k:=1; k <= rowBoundary - 1; k++{                   //右边界
		if StomachMatrix[k][data.ArrNum - 1] == hp  || StomachMatrix[k][data.ArrNum - 1] == sh{
			for s:=k - 1; s <= k + 1;s++{
				for t:=arrBoundary - 1; t <= arrBoundary; t++{
					basicity(s, t, StomachMatrix)
				}
			}
		}
	}
	for k:=1; k <= arrBoundary -1; k++{                 //下边界
		if StomachMatrix[rowBoundary][k] == hp || StomachMatrix[rowBoundary][k] == sh{
			for s := rowBoundary - 1; s <= rowBoundary;s++{
				for t:=k - 1; t <= k + 1; t++{
					basicity(s, t, StomachMatrix)
				}
			}
		}   
	}
	for i:=1; i <= len(StomachMatrix) - 2; i++{                      //内点
		for j:=1; j < len(StomachMatrix[i]) - 1; j++{
			if StomachMatrix[i][j] == hp || StomachMatrix[i][j] == sh{
				for s:=i - 1; s <= i + 1; s++{
					for t:= j - 1; t <= j + 1; t++{
						basicity(s, t, StomachMatrix)
					}
				}
			}
		}
	}
}
