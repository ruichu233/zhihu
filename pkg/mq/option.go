package mq

import "encoding/json"

type MsgEntity struct {
	MsgID string
	Key   string
	Val   string
}

func (st *MsgEntity) TransStructToMap() (map[string]interface{}, error) {
	var res map[string]interface{}
	b, err := json.Marshal(st)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	for k, v := range res {
		switch v.(type) {
		case float64:
			temp := int64(v.(float64))
			if float64(temp) == v.(float64) {
				res[k] = temp
			}
			break
		}
	}
	return res, nil
}
