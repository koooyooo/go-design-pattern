package factory

func GetStorage() Storage {
	// 何らかの具象を返却（例として diskStorageを選択）
	// 呼び出し元が具象を意識しないので、気づかれずに切り替えることが可能
	return &diskStorage{}
}
