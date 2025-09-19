package intermediate

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64,error){
	if x<0{
		return 0,errors.New("math errors:square root of negative number")
	}
	return 1,nil //compute the square root
}

func process(data []byte) error{
	if len(data)==0{
		return errors.New("empty data")
	}
	// Process data
	return nil
}

func main() {
	// result,err:=sqrt(16)
	// if err!=nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// result1,err1:=sqrt(-16)
	// if err1!=nil{
	// 	fmt.Println(err1)
	// 	return
	// }
	// fmt.Println(result1)

	// data:=[]byte{}
	// if err:=process(data); err!=nil{
	// 	fmt.Println("Error:",err)
	// 	return
	// }
	// fmt.Println("Data Processed Successfully!")

	//error interface of builtin package
	// if err1:=eprocess(); err1!=nil{
	// 	fmt.Println(err1)
	// 	return
	// }
	if err:=readData();err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data Read Successfully")
}

type myError struct{
	message string
}

func (m *myError) Error() string{
	return fmt.Sprintf("Error:%s",m.message)
}

func eprocess() error{
	return &myError{message: "Custom error message"}
}

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
// type error interface {
// 	Error() string
// }

func readData() error{
	err:=readConfig()
	if err!=nil{
		return fmt.Errorf("readData:%w",err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}