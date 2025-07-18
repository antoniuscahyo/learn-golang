package entity

import "encoding/json"

type SendEmailReq struct {
	Title         string `json:"title"`
	ReceiverEmail string `json:"receiver_email"`
	Message       string `json:"message"`
}

func (c *SendEmailReq) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}
