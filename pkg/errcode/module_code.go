package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "獲取標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")

	ErrorGetFishFail    = NewError(20020001, "獲取單個魚隻失敗")
	ErrorGetFishsFail   = NewError(20020002, "獲取多個魚隻失敗")
	ErrorCreateFishFail = NewError(20020003, "建立魚隻失敗")
	ErrorUpdateFishFail = NewError(20020004, "更新魚隻失敗")
	ErrorDeleteFishFail = NewError(20020005, "刪除魚隻失敗")

	ErrorUploadFileFail = NewError(20030001, "上傳檔案失敗")

	ErrorGetFishpowerFail    = NewError(20040001, "獲取單個魚隻活動力失敗")
	ErrorGetFishpowersFail   = NewError(20040002, "獲取多個魚隻活動力失敗")
	ErrorCreateFishpowerFail = NewError(20040003, "建立魚隻活動力失敗")
	ErrorUpdateFishpowerFail = NewError(20040004, "更新魚隻活動力失敗")
	ErrorDeleteFishpowerFail = NewError(20040005, "刪除魚隻活動力失敗")
)
