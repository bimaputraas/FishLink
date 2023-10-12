package main

type Abc interface {
	Error123(aasdasdasdasdasdasdasd string) string
}

type Tes struct{}

func InitTes() Abc {
	return Tes{}
}
func (t Tes) Error123(asdasdasd string) string {
	return "teserror"
}
func main() {

}