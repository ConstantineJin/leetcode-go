package main

// 模拟
//type tsk struct {
//	name byte
//	cnt  int
//}
//
//func leastInterval(tasks []byte, n int) (ans int) {
//	var tsks, mp = make([]tsk, 26), make(map[byte]int)
//	for i := range tsks {
//		tsks[i].name = byte('A' + i)
//	}
//	for _, task := range tasks {
//		tsks[task-'A'].cnt++
//	}
//	sort.Slice(tsks, func(i, j int) bool { return tsks[i].cnt > tsks[j].cnt })
//	for tsks[0].cnt > 0 {
//		for i := 0; i < 26; i++ {
//			if tsks[i].cnt == 0 {
//				ans++
//				break
//			}
//			if v, ok := mp[tsks[i].name]; ok && v >= ans-n {
//				continue
//			}
//			mp[tsks[i].name] = ans
//			ans++
//			var temp = tsk{name: tsks[i].name, cnt: tsks[i].cnt - 1}
//			tsks = slices.Insert(append(tsks[:i], tsks[i+1:]...), sort.Search(25, func(i int) bool { return temp.cnt >= tsks[i].cnt }), temp)
//			break
//		}
//	}
//	return
//}

//func leastInterval(tasks []byte, n int) (minTime int) {
//	cnt := map[byte]int{}
//	for _, t := range tasks {
//		cnt[t]++
//	}
//
//	nextValid := make([]int, 0, len(cnt))
//	rest := make([]int, 0, len(cnt))
//	for _, c := range cnt {
//		nextValid = append(nextValid, 1)
//		rest = append(rest, c)
//	}
//
//	for range tasks {
//		minTime++
//		minNextValid := math.MaxInt64
//		for i, r := range rest {
//			if r > 0 && nextValid[i] < minNextValid {
//				minNextValid = nextValid[i]
//			}
//		}
//		minTime = max(minTime, minNextValid)
//		best := -1
//		for i, r := range rest {
//			if r > 0 && nextValid[i] <= minTime && (best == -1 || r > rest[best]) {
//				best = i
//			}
//		}
//		nextValid[best] = minTime + n + 1
//		rest[best]--
//	}
//	return
//}

// 构造
func leastInterval(tasks []byte, n int) int {
	var cnt = map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}
	var maxExec, maxExecCnt int
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}
	return max(len(tasks), (maxExec-1)*(n+1)+maxExecCnt)
}
