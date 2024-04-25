package main

func distanceTraveled(mainTank int, additionalTank int) int {
	return (mainTank + min((mainTank-1)/4, additionalTank)) * 10 // 主油箱从副油箱总共获取了(mainTank-1)/4升油（最后如果主油箱不到4升油则无法获取）
}
