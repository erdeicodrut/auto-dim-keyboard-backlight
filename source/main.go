package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	lit := true
	b4Value := []byte("2")
	file, err:= os.OpenFile("/sys/devices/platform/asus-nb-wmi/leds/asus::kbd_backlight/brightness", int(666), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	file.Read(b4Value)
	file.Close()	

	event, err := os.Open("/dev/input/event0")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 24)

	action := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-action :
				if !lit {
					file, err:= os.OpenFile("/sys/devices/platform/asus-nb-wmi/leds/asus::kbd_backlight/brightness", int(666), os.ModePerm)
					if err != nil {
						fmt.Println(err)
						return
					}
					file.Write(b4Value)
					
					file.Close()
				}
				
				lit = true
				
			case <-time.After(time.Minute):
				fmt.Println("keyboard not active ", lit)
				if lit {
					file, err:= os.OpenFile("/sys/devices/platform/asus-nb-wmi/leds/asus::kbd_backlight/brightness", int(666), os.ModePerm)
					if err != nil {
						fmt.Println(err)
						return
					}
					
					file.Read(b4Value)
					fmt.Println(string(b4Value))

					file.Write([]byte("0"))
					
					file.Close()
				}

				lit = false
			}
		}
	}()

	for {
		_, _ = event.Read(buf)
		action<-true
		fmt.Println(true)
	}
}
