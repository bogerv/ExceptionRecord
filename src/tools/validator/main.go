package main

// Usage : http://localhost:9000?web=bogerv.wang&age=0&zip=&dob=2017-09-22&agree=true&email=124609563@qq.com&phone=13199526280
// messages 中的内容可以覆盖,验证后的返回内容

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/govalidator/validator"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on prot: 9000")
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	rules := validator.MapData{
		"username": []string{"required", "between:3,8"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"numeric_between:18,56"},
		"zip":      []string{"numeric"},
		"phone":    []string{"digits:11"},
		"roles":    []string{"in:admin,manager,rider"},
		"names":    []string{"not_in:john,jane,mila"},
		"agree":    []string{"bool"},
		"dob":      []string{"date"},
	}
	messages := validator.MapData{
		"username": []string{"required:আপনাকে অবশ্যই ইউজারনেম দিতে হবে", "between:ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে।"},
		"zip":      []string{"numeric:জিপ কোড অবশ্যই নাম্বার হবে"},
	}
	opts := validator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be required
	}
	v := validator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "applciation/json")
	json.NewEncoder(w).Encode(err)
}
