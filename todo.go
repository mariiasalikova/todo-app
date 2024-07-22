package todo


type TodoList struct {
	Id int `json:"-"`
	Title string `json:"title"`
	Description string `json:"description"`

}

type TodoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Done bool `json:"done"`

}

type UsersList struct {
	Id int
	UserId int
	ListId int
}

type ListsItem struct {
	Id int
	ItemId int
	ListId int
}