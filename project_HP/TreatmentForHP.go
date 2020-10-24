/*
simulate biological logic via function
函数模拟生物的逻辑
the core of this project is operating matrix
核心思想在于对二维矩阵的操作
*/
package main

import(
	"go_code/project_HP/information"
	"go_code/project_HP/patient"
)

func main(){
	//information of this project
	//首先向用户呈现项目的介绍
	information.ProjectInformation()
	
	//invoke patients' cycle
	//调用患者大外层循环
	patient.PatientCycle()
	
	//exit this program
	//退出程序
	information.Exit()
}
