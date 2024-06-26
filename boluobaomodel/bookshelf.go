package boluobaomodel

type ShelfData struct {
	AccountID  int    `json:"accountId"`
	PocketID   int    `json:"pocketId"`
	Name       string `json:"name"`
	TypeID     int    `json:"typeId"`
	CreateTime string `json:"createTime"`
	IsFull     bool   `json:"isFull"`
	CanModify  bool   `json:"canModify"`
	Expand     struct {
		Novels []BookInfoData `json:"novels"`
	} `json:"expand"`
}

type InfoData struct {
	Status `json:"status"`
	Data   []ShelfData `json:"data"`
}

func (shelf *InfoData) EachShelf(f func(int, ShelfData)) {
	if shelf.Data != nil {
		for i, book := range shelf.Data {
			f(i, book)
		}
	}
}
func (shelf *InfoData) GetShelf(index int) *ShelfData {
	if index >= 0 && index < len(shelf.Data) {
		return &shelf.Data[index]
	}
	return nil
}
func (bookshelf *ShelfData) EachBookshelf(f func(int, BookInfoData)) {
	for i, book := range bookshelf.Expand.Novels {
		f(i, book)
	}
}
func (bookshelf *ShelfData) GetBookshelf(index int) *BookInfoData {
	if bookshelf.Expand.Novels != nil && index < len(bookshelf.Expand.Novels) {
		return &bookshelf.Expand.Novels[index]
	}
	return nil
}

type ConsumeData struct {
	EntityId                int    `json:"entityId"`
	ConsumeFireMoney        int    `json:"consumeFireMoney"`
	ConsumeGeneralCouponNum int    `json:"consumeGeneralCouponNum"`
	ConsumeChapterNum       int    `json:"consumeChapterNum"`
	EntityName              string `json:"entityName"`
	EntityCover             string `json:"entityCover"`
	LastConsumeDate         string `json:"lastConsumeDate"`
	CouponNum               int    `json:"couponNum"`
	Novel                   struct {
		AuthorId       int     `json:"authorId"`
		LastUpdateTime string  `json:"lastUpdateTime"`
		MarkCount      int     `json:"markCount"`
		NovelCover     string  `json:"novelCover"`
		BgBanner       string  `json:"bgBanner"`
		NovelId        int     `json:"novelId"`
		NovelName      string  `json:"novelName"`
		Point          float64 `json:"point"`
		IsFinish       bool    `json:"isFinish"`
		AuthorName     string  `json:"authorName"`
		CharCount      int     `json:"charCount"`
		ViewTimes      int     `json:"viewTimes"`
		TypeId         int     `json:"typeId"`
		AllowDown      bool    `json:"allowDown"`
		AddTime        string  `json:"addTime"`
		IsSensitive    bool    `json:"isSensitive"`
		SignStatus     string  `json:"signStatus"`
		CategoryId     int     `json:"categoryId"`
	} `json:"novel"`
}
type Consume struct {
	Status struct {
		HttpCode  int         `json:"httpCode"`
		ErrorCode int         `json:"errorCode"`
		MsgType   int         `json:"msgType"`
		Msg       interface{} `json:"msg"`
	} `json:"status"`
	Data []ConsumeData `json:"data"`
}
