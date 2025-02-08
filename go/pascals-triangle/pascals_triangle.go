package pascal

func Triangle(n int) (rows [][]int) {
	for i := 0; i < n; i++ {

		row := []int{}
		if i == 0 {
			row = append(row, 1)
		} else {
			prevRow := rows[i-1]
			row = append(row, 1)
			for i := 0; i < len(prevRow)-1; i++ {
				l := prevRow[i]
				r := prevRow[i+1]
				row = append(row, l+r)
			}
			row = append(row, 1)
		}
		rows = append(rows, row)
	}
	return rows
}
