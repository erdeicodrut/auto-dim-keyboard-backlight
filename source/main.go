package main

import(
	"os"
	"fmt"
)

func main() {
	file, err:= os.OpenFile("/sys/devices/platform/asus-nb-wmi/leds/asus::kbd_backlight/brightness", int(666), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	shit := []byte("0")

	p, err := file.Write(shit)
	
	fmt.Println(p, err)

	fmt.Println(string(shit))

	file.Close()
}
