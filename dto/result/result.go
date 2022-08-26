package dto

type SuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"users"`
}
type SuccessCart struct {
	Status string      `json:"status"`
	Data   interface{} `json:"carts"`
}
type SuccessRegister struct {
	Status string      `json:"status"`
	Data   interface{} `json:"user"`
}
type SuccessDeleteUser struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ErrorResultData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type SuccesFindProducts struct {
	Status string      `json:"status"`
	Data   interface{} `json:"products"`
}
type SuccessGetProduct struct {
	Status string      `json:"status"`
	Data   interface{} `json:"product"`
}
type SuccessDeleteProduct struct {
	Status string      `json:"status`
	Data   interface{} `json:"data"`
}
type SuccesFindTopings struct {
	Status string      `json:"status"`
	Data   interface{} `json:"toppings"`
}
type SuccesGetToping struct {
	Status string      `json:"status"`
	Data   interface{} `json:"topping"`
}
type SuccessDeleteToping struct {
	Status string      `json:"status`
	Data   interface{} `json:"data"`
}
type SuccessFindTransactions struct {
	Status string      `json:"status`
	Data   interface{} `json:"Transactions"`
}
type SuccessGetTransaction struct {
	Status string      `json:"status`
	Data   interface{} `json:"Transaction"`
}
type CreateTransaction struct {
	Status string      `json:"status`
	Data   interface{} `json:"Products"`
}
type EditTransaction struct {
	Status string      `json:"status`
	Data   interface{} `json:"Transaction"`
}
type DeleteTransaction struct {
	Status string      `json:"status`
	Data   interface{} `json:"data"`
}
