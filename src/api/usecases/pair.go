package usecases

import (
	entity "pairs/src/entities"
	lib "pairs/src/libs"
)

// FindPair handle req
type FindPair struct {
	req     *lib.Request
	payload entity.Payload
}

// NewFindPair represent new FindPair
func NewFindPair(req *lib.Request, payload entity.Payload) *FindPair {
	return &FindPair{
		req:     req,
		payload: payload,
	}
}

// GetPairs return all possible target indices
func (fp *FindPair) GetPairs() {
	data := make(map[string][][]int)
	resultPair := make([][]int, 0)
	fp.req.AddResponsePayload("findPair", fp.payload)
	totalLength := len(fp.payload.Numbers)
	for i := 0; i < totalLength; i++ {
		for j := i + 1; j < totalLength; j++ {
			if fp.payload.Numbers[i]+fp.payload.Numbers[j] == fp.payload.Target {
				resultPair = append(resultPair, []int{i, j})
			}
		}
	}
	data["solutions"] = resultPair
	fp.req.AddResponsePayload("findPair", resultPair)
}
