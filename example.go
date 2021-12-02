package main

import (
	"encoding/json"
	"fmt"
)

type TagUpdateEventMsgOld struct {
	MainOrderId int64           `json:"main_order_id"`
	BitTags     map[int64]int32 `json:"bit_tags"`
	Timestamp   int64           `json:"timestamp"`
}

type TagUpdateEventMsg struct {
	MainOrderId         int64                     `json:"main_order_id"`
	BitTags             map[int64]int32           `json:"bit_tags"`
	Timestamp           int64                     `json:"timestamp"`
	OrderLineId2BitTags map[int64]map[int64]int32 `json:"order_line_id_2_bit_tags"`
}

func main() {
	tag := map[int64]int32{6: 7}

	tagUpdateEventMsg := TagUpdateEventMsgOld{
		MainOrderId: 5,
		BitTags:     tag,
		Timestamp:   23232,
	}
	tagStr, err := json.Marshal(tagUpdateEventMsg)
	if err != nil {
		fmt.Printf("err:%+v not nil", err)
		return
	}
	fmt.Printf("str: %+v", string(tagStr))

	newTag := TagUpdateEventMsg{}
	err = json.Unmarshal(tagStr, &newTag)
	if err != nil {
		fmt.Printf("unmarshal err:%+v not nil", err)
		return
	}
	for k, v := range newTag.OrderLineId2BitTags {
		fmt.Printf("k: %+v, v:%+v", k, v)
	}
	fmt.Printf("newTag %+v", newTag)

}
