/*
统计诊疗结果
发送分析报告
*/
package main

import(
	"fmt"
	"math"
	"strconv"
)
//告知用户
func information(){
	fmt.Println("\tThis project performs a fitting analysis on the two sets of data obtained in the experiment (Input and Output in the experiment)")
	//fmt.Println("\t本项目对实验中获得的两组数据（实验中的Input和Output）进行拟合分析")
	fmt.Println("\tThe regression model adopts a polynomial, and the user defines the highest power")
	//fmt.Println("\t回归模型采取多项式，由用户自行定义最高幂次")
	fmt.Println("\tThe user first enters the data size, that is, how many sets of data are obtained in the experiment for regression analysis")
	//fmt.Println("\t用户先输入数据规模，也就是在实验中获得了多少组数据用于回归分析")
}
//单层线性回归
func linearRegression(power float64, sliceInput []float64, sliceOutput []float64)(k float64, b float64){
	length := len(sliceInput)
	var sumX, sumY, sumXY, sumXX float64        //求和  
	var averageX, averageY float64              //平均值
	l := float64(length)    //将长度转换为浮点数，用于k斜率的计算
	//最小二乘法的线性回归
	//注意到：Input中的元素需要乘以power幂
	for i:=0; i < length; i++{
		xPower := math.Pow(sliceInput[i], power)            //Input中元素的power次幂
		sumX += xPower
		sumY += sliceOutput[i]
		sumXY += xPower * sliceOutput[i]
		sumXX += math.Pow(xPower, 2.0)
		averageX = sumX / float64(length)
		averageY = sumY / float64(length)
		k = (sumXY - l * averageX * averageY) / (sumXX - l * averageX * averageX)
		b = averageY - k * averageX
	}
	return k, b
}
//根据用户需求，选择是否需要打印系数为0的项:输入yes，返回true；no则返回false
func hideZero()(hideFlag bool){
	fmt.Println("The coefficient keeps two decimal places; please choose whether to print the item with a coefficient of 0 or not")
	//fmt.Println("系数保留小数点后两位；请根据选择是否打印系数为0的项")
	var omitFlag string
	for {
	fmt.Print("If you need to print all items, please enter yes, otherwise, entering no will omit items with a coefficient of 0:")
	//fmt.Print("如果需要打印所有项请输入 yes,否则输入no将会省略系数为0的项: ")
	fmt.Scanln(&omitFlag)
	if omitFlag == "no" || omitFlag == "yes"{
		break
	} else {
		fmt.Println("Someting wrong with your input, Please try again")
		//fmt.Println("您的输入有误，请重新输入")
	}
	}
	if omitFlag == "no"{
	hideFlag = false
	}
	if omitFlag == "yes"{
	hideFlag = true
	}
	return hideFlag
}
//多项式拟合
func MultinomialRegression(power float64, sliceInput []float64, sliceOutput []float64){
	fmt.Print("Please enter the highest coefficient (a positive integer) of the polynomial you want to fit:")
	//fmt.Print("请输入您希望拟合的多项式的最高项系数（一个正整数）：")
	for {
	fmt.Scanln(&power)
	if power > 0 {
		break
	} else {
		fmt.Println("Please enter a positive integer")
		//fmt.Println("请输入正整数")
	}
	}
	var constantTerm float64           //常数项
	if hideZero() {                    //打印所有项
	fmt.Print("OutputNum = ")
	for p := power; p >= 1; p--{
		k, b := linearRegression(p, sliceInput, sliceOutput)
		//Y = y - x ^ p
		for i:=0; i < len(sliceInput); i++{
			newY := sliceOutput[i] - math.Pow(sliceInput[i], p)
			sliceOutput[i] = newY
		}
		constantTerm += b
		fmt.Printf("(%.2f) * InputNum^%v + ", k, p)
	}	
	fmt.Printf("(%.4f)", constantTerm)
	} else {                            //省略掉系数为0的项
	fmt.Print("OutpurNum = ")
	for p:=power; p >= 1; p--{
		kOriginal, bOriginal := linearRegression(p, sliceInput, sliceOutput)
		kExcessive, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", kOriginal), 64)
		bExcessive, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", bOriginal), 64)
		k := float64(kExcessive)
		b := float64(bExcessive)
		//Y = y - x ^ p
		for i:=0; i < len(sliceInput); i++{
			newY := sliceOutput[i] - math.Pow(sliceInput[i], p)
			sliceOutput[i] = newY
		}
		constantTerm += b
		if k != 0 {
			fmt.Printf("(%.2f) * InputNum^%v + ", k, p)
		}
	}
	fmt.Printf("(%.4f)", constantTerm)
	}
}

func Exit(){
	fmt.Println("Please enter any character to exit the program")
	//fmt.Println("请输入任意字符退出程序")
	var exit string
	fmt.Scanln(&exit)
}

func main(){
	information()

	var InputSlice []float64 = make([]float64, 0)
	var OutputSlice []float64 = make([]float64, 0)

	var scale int
	fmt.Print("Please enter the size of the data you will enter (how many):")
	//fmt.Print("请输入您将输入的数据规模（个）：")
	fmt.Scanln(&scale)

	for i:=1; i <= scale; i++{
		var inputData float64
		fmt.Printf("Input Group %v is: ", i)
		//fmt.Printf("第 %v 组的Input值为：", i)
		fmt.Scanln(&inputData)
		InputSlice = append(InputSlice, inputData)
		var outputData float64
		fmt.Printf("Output Group %v is: ", i)
		//fmt.Printf("第 %v 组的Output值为：", i)
		fmt.Scanln(&outputData)
		OutputSlice = append(OutputSlice, outputData)
		fmt.Println()
	}

	// fmt.Println("The value entered is: ")
	// //fmt.Println("输入的值为：", InputSlice)
	// fmt.Println("The output value is: ")
	// //fmt.Println("输出的值为：", OutputSlice)
	 
	var power float64  //幂
	 MultinomialRegression(power, InputSlice, OutputSlice)
	 fmt.Println("")
	 Exit()
}
