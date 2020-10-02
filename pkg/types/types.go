package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.).
type Money int64

//PaymentCategory представляет собой категорию, в который был совершён платёж (авто, аптеки, рестораны и т.д.).
type PaymentCategory string

//PaymentStatus представляет собой статус платежа.
type PaymentStatus string

//Предопределённые статусы платежой.
const (
	StatusOk PaymentStatus = "OK"
	PaymentStatusFail PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

//Payment представляет информацию о платежа.
type Payment struct{
	AccountID              int64
	ID                     string
	Amount                 Money
	Category               PaymentCategory
	Status                 PaymentStatus
}

type Phone string

//Account представляет информацию о счёте пользователя.
type Account struct {
	ID        int64
	Phone     Phone
	Balance   Money
}