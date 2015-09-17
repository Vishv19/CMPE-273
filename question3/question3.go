package question3 

import "time"

func sleep(seconds int) {
    <-time.After(time.Second * time.Duration(seconds))
}