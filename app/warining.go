package app

import (
	"fmt"
	"os"
	"plotcarrier/service"

	"github.com/spf13/viper"
)

//It's still running if return true with error
//Stop running if return false
func pauseOrWarning(dst string) (bool, error) {
	pauseSize := viper.GetInt("pause_size")      //GB
	warningSize := viper.GetInt("warining_size") //GB
	if pauseSize <= 0 {
		//Do nothing is pauseSize doesn't set
		return true, fmt.Errorf("pause_size doesn't set")
	}
	//Check plot size
	currentPlotSize, err := service.CurrentPlotSize()
	if err != nil {
		return true, err
	}
	//Check disk size
	currentDiskSize, err := service.DiskSizeGB(dst)
	if err != nil {
		return true, err
	}

	currentFree := currentDiskSize - currentPlotSize

	if pauseSize > int(currentFree) {
		//Notify
		notifiyAndAlert("可用容量警告", fmt.Sprintf("%s 可用空间严重不足,已执行相关任务暂停操作,请在更换磁盘后再激活任务:", dst))
		//Pause plotman
		return false, fmt.Errorf("touch pause_size")
	}

	if warningSize > int(currentFree) {
		//Notify
		notifiyAndAlert("暂停任务警告", fmt.Sprintf("%s 可用空间即将不足,请立即查看.", dst))
		return true, fmt.Errorf("touch warning size")
	}

	return true, nil

}

func notifiyAndAlert(title, msg string) {
	hostname, _ := os.Hostname()
	title = fmt.Sprintf("[%s] %s", hostname, title)
	notifyType := viper.GetString("alerty_type")
	switch notifyType {
	case "email":
		service.SMTPMailTo(title, msg)
	case "dingtalk":
		service.NewDingTalk().SendMessage(service.DingTalkMessage{
			Type:    "text",
			Message: fmt.Sprintf("%s\n%s", title, msg),
		})
	}
}
