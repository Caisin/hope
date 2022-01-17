package command

import "os/exec"

// RunCommand 运行命令
//dir:运行目录
//命令名
//args 参数
func RunCommand(dir string, name string, args ...string) string {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	msg, _ := cmd.CombinedOutput()
	_ = cmd.Run()
	return string(msg)
}

func Ffmpeg(args ...string) string {
	return RunCommand(".", "ffmpeg", args...)
}

func GetDuration(url string) string {
	//ffmpeg -i video_url 2>&1 | grep Duration | awk '{print $2}
	return Ffmpeg("-i", url)
}
