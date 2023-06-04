package cycle

func cycleSortV1(arr []int) int {
	writes := 0

	for cyclestart := 0; cyclestart < len(arr)-1; cyclestart++ {
		item := arr[cyclestart]

		pos := cyclestart

		for i := cyclestart + 1; i < len(arr); i++ {
			if arr[i] < item {
				pos++
			}
		}

		if pos == cyclestart {
			continue
		}

		for item == arr[pos] {
			pos++
		}

		arr[pos], item = item, arr[pos]

		writes++

		for pos != cyclestart {
			pos = cyclestart
			for i := cyclestart + 1; i < len(arr); i++ {
				if arr[i] < item {
					pos++
				}
			}

			for item == arr[pos] {
				pos++
			}

			arr[pos], item = item, arr[pos]
			writes++
		}
	}

	return writes
}
