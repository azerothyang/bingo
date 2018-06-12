package model

import (
	"common/orm"
	"common/sqlbuilder"
	"fmt"
	"sync"
)

type NoticeModel struct {
	baseModel
}

type Notice struct {
	NoticeId      int64  `json:"NoticeId,omitempty"`
	NoticeCode    string `json:"NoticeCode,omitempty"`
	NoticeContent string `json:"NoticeContent,omitempty" valid:"Required;MaxSize(60000)"`
	NoticeType    string `json:"NoticeType,omitempty" valid:"Required;MaxSize(16)"`
	CreateTime    string `json:"CreateTime,omitempty"`
	Publish       string `json:"Publish,omitempty"`
}

const NoticeTable = "notice"

//新建
func (n *NoticeModel) Add(data *map[string]string) (int64, error) {
	cols, values := n.initMapData(data)
	sql := (&sqlbuilder.SqlBuilder{}).Insert(NoticeTable, cols).GetSql()
	db := orm.NewOrm()
	res, err := db.Raw(sql, *values...).Exec()
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (n *NoticeModel) GetById(noticeId string, cols string) (*Notice, error) {
	sql := (&sqlbuilder.SqlBuilder{}).
		Select(cols, NoticeTable).
		Where("notice_id", "=").
		GetSql()
	var notice = Notice{}
	db := orm.NewOrm()
	db.Using("slave")
	err := db.Raw(sql, noticeId).QueryRow(&notice)
	return &notice, err
}

//协程, 管道是本身就是引用, 切片传参默认是会新生成一个内存对应到之前切片对应的底层数组， 其实就是不同内存指向同一底层内存
func (n *NoticeModel) GetAllCount() map[string]int64 {
	mp := make(map[string]int64) //键值对, 分别存"tableName" : 20
	var wg sync.WaitGroup
	//等待两个任务
	wg.Add(2)
	go n.getCount(mp, NoticeTable, &wg)
	go n.getCount(mp, "sub_policy_result_detail", &wg)
	wg.Wait()
	return mp
}

func (n *NoticeModel) getCount(mp map[string]int64, table string, wg *sync.WaitGroup) {
	type Count struct {
		Total int64
	}
	sql := (&sqlbuilder.SqlBuilder{}).
		Select("count(1) as total", table).
		GetSql()
	var count = Count{}
	db := orm.NewOrm()
	db.Using("slave")
	err := db.Raw(sql).QueryRow(&count)
	if err != nil {
		fmt.Println("error")
	}
	mp[table] = count.Total
	wg.Done()
}
