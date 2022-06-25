package model

type UserField struct {
    UserId   int    `json:"userId"`
    UserName string `json:"userName"`
    UserPwd  string `json:"userPwd"`
}
