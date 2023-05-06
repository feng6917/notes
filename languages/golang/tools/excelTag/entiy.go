package excelTag

// ig ignore 忽略
// rn reName 重命名 用于 表header
// fn funcName 方法名 用于复杂方法处理
// en enum 枚举

type Example struct {
	Id        uint64 `json:"id" ig:"1" `
	Name      string `json:"name" rn:"姓名"`
	NameSpell string `json:"nameSpell" rn:"姓名拼写" ig:""`
	Age       int    `json:"age" rn:"年龄" fn:"addTen"`
	Sex       string `json:"sex" rn:"性别" en:"0|保密;1|男;2|女"`
}
