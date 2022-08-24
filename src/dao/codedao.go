package dao

import (
	"cinema/model"
	"cinema/utils"
)

// AddCode 增加邀请码
func AddCode(code *model.Code) error {
	sql := "insert into code(code, root) VALUES (?,?)"
	_, err := utils.DB.Exec(sql, code.Code, code.Root)
	if err != nil {
		return err
	}
	return nil
}

// CheckCode 验证邀请码
func CheckCode(codeStr string) (*model.Code, error) {
	sql := "select * from code where code = ?;"
	var code model.Code
	err := utils.DB.QueryRow(sql, codeStr).Scan(&code.Code, &code.Root)
	if err != nil {
		return nil, err
	}
	return &code, nil
}

// DeleteAllCodes 删除所有邀请码
// func DeleteAllCodes() error {
//	sql := "delete from code;"
//	_, err := utils.DB.Exec(sql)
//	if err != nil {
//		return err
//	}
//	return nil
// }

// DeleteCode 删除单个邀请码
func DeleteCode(codeStr string) error {
	sql := "delete from code where code = ?"
	_, err := utils.DB.Exec(sql, codeStr)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCodes 获取所有codes
func GetAllCodes() ([]*model.Code, error) {
	sql := "select code, root from code;"
	rows, err := utils.DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var codes []*model.Code
	for rows.Next() {
		var code model.Code
		err := rows.Scan(&code.Code, &code.Root)
		if err != nil {
			return nil, err
		}
		codes = append(codes, &code)
	}
	return codes, err
}
