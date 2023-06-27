package main

import (

	"fmt"

	"os"

)

func main() {

	file, _ := os.Open("non-existing.txt")

	fmt.Println(file)

}
// در بالا ما خطای تابع Open را نادیده گرفتیم و مقدار file را چاپ کردیم و مقدار چاپ شده nil است چون تایپ خروجی با اشاره گر است و
//قطعا مقدار خالی بودش nil است و دقت کنید اگر این متغیر را به تابع دیگری پاس دهید قطعا با panic مواجه خواهید شد.