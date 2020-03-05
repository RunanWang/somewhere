package utils

// func SortItemByScore(score msg.ScoreResp, plist []model.TProduct) ([]model.TProduct, error) {
// 	// 输入为每个item的得分，和所有product的信息
// 	// 输出为按照item得分由高到低排序之后的product list
// 	var ans []model.TProduct
// 	// fmt.Println(score.List)
// 	sort.Slice(score.List, func(i, j int) bool {
// 		return score.List[i].Score > score.List[j].Score
// 	})
// 	// fmt.Println(score.List)
// 	for _, item := range score.List {
// 		for _, detail := range plist {
// 			if item.ItemID == detail.ID.Hex() {
// 				ans = append(ans, detail)
// 				break
// 			}
// 		}
// 	}
// 	return ans, nil
// }
