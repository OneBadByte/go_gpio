package go_gpio


import(

	"io/ioutil"
	//"fmt"
	"strconv"
      )

func    ChangeGpioValue(gpioPin int, gpioValue int){
	gpio := strconv.Itoa(gpioPin)
	var filePath string = "/sys/class/gpio/gpio"+ gpio + "/value"
	gpioVal := strconv.Itoa(gpioValue)
	ioutil.WriteFile(filePath, []byte(gpioVal), 777)
}

func ReadGpioValue(gpioPin int) string{
	gpio := strconv.Itoa(gpioPin)
	var filePath string = "/sys/class/gpio/gpio"+ gpio + "/value"
	value, err := ioutil.ReadFile(filePath)
	if err != nil{
		return "-1"
	}
	return string(value)
	

}

func ExportGpio(gpioPin int, gpioDirection string){

	gpio := strconv.Itoa(gpioPin)
	var filePath string = "/sys/class/gpio/export"
	var directionPath string = "/sys/class/gpio/gpio" + gpio + "/direction"
	ioutil.WriteFile(filePath, []byte(gpio), 777)
	ioutil.WriteFile(directionPath, []byte(gpioDirection), 777)
	
}
