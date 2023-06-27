package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	var ss = make([]int, n+2)
	ss[1] = 1
	ss[2] = 2
	for i := 3; i <= n; i++ {
		ss[i] = ss[i-1] + ss[i-2]
	}
	return ss[n]
}
