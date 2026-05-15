package entity

/*
	catatan

1. huruf kafital di depan menandakan ini public property
2. singkatan biasanya kafital semua
3. parameter yang ketika namanya struct tag pemetaan saat diubah ke bentuk lain kalo gak di metakan akan gini {"ID": 1, "Name": "Jamal"}
*/
type User struct {
	ID    int    `json:id`
	Name  string `json:name`
	Email string `json:email`
}
