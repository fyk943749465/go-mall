// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type OrderReq struct {
	Id string `path:"id"`
}

type OrderReply struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}
