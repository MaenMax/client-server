package queue

import l4g "github.com/alecthomas/log4go"

type RequestType int // 1:AddItem  2:RemoveItem 3:GetItem 4:GetAllItems

type T_Item struct {
	Id   string `json:"id,omitempty"`
	Data string `json:"data,omitempty"`
}

type T_Request struct {
	ClientId string      `json:"client_id,omitempty"`
	Item     T_Item      `json:"item,omitempty"`
	Type     RequestType `json:"request_type,omitempty"`
}

type T_Queue struct {
	Items map[string][]T_Item
}

func New() *T_Queue {
	queue := &T_Queue{}
	queue.Items = make(map[string][]T_Item)
	return queue
}

func (q *T_Queue) AddItem(item T_Item, client_id string) {
	exists, _ := q.Exists(item, client_id)
	if !exists {
		q.Items[client_id] = append(q.Items[client_id], item)
	}
}
func (q *T_Queue) GetItem(item T_Item, client_id string) string {
	exists, index := q.Exists(item, client_id)
	if !exists {
		q.Items[client_id] = append(q.Items[client_id], item)
	}
	return q.Items[client_id][index].Data
}
func (q *T_Queue) GetAllItems(item T_Item, client_id string) []string {
	exists, _ := q.Exists(item, client_id)
	if !exists {
		q.Items[client_id] = append(q.Items[client_id], item)
	}
	var all_items_data []string
	for _, item := range q.Items[client_id] {
		all_items_data = append(all_items_data, item.Data)
	}
	return all_items_data
}

func (q *T_Queue) RemoveItem(client_id string, item T_Item) (removed bool) {
	exists, index := q.Exists(item, client_id)
	if !exists {
		return false
	}
	arr := q.Items[client_id]
	arr1 := arr[:index]
	arr2 := arr[index+1:]
	arr1 = append(arr1, arr2...)

	q.Items[client_id] = arr1
	removed = true
	return
}
func (q *T_Queue) Exists(item T_Item, client_id string) (exists bool, index int) {
	for i, q_item := range q.Items[client_id] {
		if q_item.Id == item.Id {
			exists = true
			index = i
			break
		}
	}
	return
}
func (q *T_Queue) Print() {
	for k, v := range q.Items {
		l4g.Info("Client ID:%s  Items: %s", k, v)
	}
}
