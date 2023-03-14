package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VipUser struct {
	UserInfo User
	VipLevel int
	Price    int
}

type VipUser2 struct {
	User
	VipLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}
	vip := VipUser{
		User{"화랑", "hwarang", 40},
		3,
		250,
	}
	vip2 := VipUser2{
		User{"화랑2", "hwarang2", 42},
		3,
		250,
	}

	fmt.Printf("유저:%s ID:%s 나이:%d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d VIP레벨:%d 가격:%d만원\n",
		vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age,
		vip.VipLevel, vip.Price)
	fmt.Printf("VIP유저:%s ID:%s 나이:%d VIP레벨:%d 가격:%d만원\n",
		vip2.Name, vip2.ID, vip2.Age, vip.VipLevel, vip.Price)
}

// 유저:송하나 ID:hana 나이:23
// VIP유저:화랑 ID:hwarang 나이:40 VIP레벨:3 가격:250만원
// VIP유저:화랑2 ID:hwarang2 나이:42 VIP레벨:3 가격:250만원
