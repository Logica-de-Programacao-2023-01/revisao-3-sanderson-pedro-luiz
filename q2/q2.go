package q2

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefixo := strs[0]
	for i := 1; i < len(strs); i++ {
		palavra := strs[i]
		j := 0
		for j < len(prefixo) && j < len(palavra) && prefixo[j] == palavra[j] {
			j++
		}
		prefixo = prefixo[:j]
		if prefixo == "" {
			return ""
		}
	}
	return prefixo
}
