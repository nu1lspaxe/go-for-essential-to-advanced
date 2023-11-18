package main

// 品項定義
type ITEM struct {
	name   string
	sugar  string
	ice    string
	price  int
	number int
}

// 品項列表
type ITEMLIST struct {
	items []ITEM
}

// 從商品列表添加品項
func (list *ITEMLIST) itemAdd(it ITEM) []ITEM {
	list.items = append(list.items, it)
	return list.items
}

// 從購物列表刪除品項
func (list *ITEMLIST) itemRemove(index int) []ITEM {
	list.items = append(list.items[:index], list.items[index+1:]...)
	return list.items

}

// 從購物列表插入品項
func (list *ITEMLIST) itemInsert(it ITEM, index int) []ITEM {
	list.items = append(list.items[:index], append([]ITEM{it}, list.items[index:]...)...)
	return list.items
}
