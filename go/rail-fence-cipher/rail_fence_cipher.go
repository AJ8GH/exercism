package railfence

func Encode(message string, rails int) (out string) {
	j := 0
	fun := add

	chunks := []string{}
	for i := 0; i < rails; i++ {
		chunks = append(chunks, "")
	}

	for _, v := range message {
		chunks[j] += string(v)
		j, fun = getNewJ(j, rails, fun)
	}
	for _, v := range chunks {
		out += v
	}

	return out
}

func Decode(message string, rails int) (out string) {
	chunkSizes := []int{}
	for i := 0; i < rails; i++ {
		chunkSizes = append(chunkSizes, 0)
	}

	j := 0
	fun := add
	for range message {
		chunkSizes[j]++
		j, fun = getNewJ(j, rails, fun)
	}

	chunks := []string{}
	i := 0
	for _, v := range chunkSizes {
		v += i
		chunks = append(chunks, message[i:v])
		i = v
	}

	j = 0
	fun = add
	chunkIndices := []int{}
	for range chunks {
		chunkIndices = append(chunkIndices, 0)
	}
	for i := 0; i < len(message); i++ {
		out += string(chunks[j][chunkIndices[j]])
		chunkIndices[j]++
		j, fun = getNewJ(j, rails, fun)
	}

	return out
}

func getNewJ(j, rails int, fun func(int) int) (int, func(int) int) {
	switch j {
	case 0:
		fun = add
	case rails - 1:
		fun = sub
	}
	return fun(j), fun
}

func add(j int) int {
	return j + 1
}

func sub(j int) int {
	return j - 1
}
