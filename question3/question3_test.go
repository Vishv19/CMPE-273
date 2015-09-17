package question3

import ("testing"; "time")

func Test_Sleep(t *testing.T) {
    sleepTime := 1
    startTime := time.Now()
    sleep(sleepTime)
    endTime := time.Now()
    timeDifference := endTime.Second() - startTime.Second()
    if timeDifference!= sleepTime {
        t.Error("Expected sleeptime to be 1 second, got ", timeDifference)
    }
}