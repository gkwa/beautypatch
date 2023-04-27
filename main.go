package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func main() {
	paths := []string{
		"C:\\Windows\\System32\\control.exe",
		"C:\\Windows\\System32\\windowspowershell\\v1.0\\powershell.exe",
		"C:\\Windows\\System32\\windowspowershell\\v1.0\\powershell.exe.lnk",
		"C:\\Windows\\explorer.exe",
	}

	// Initialize logrus
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	xmlTemplate := `<?xml version="1.0" encoding="utf-8"?>
<LayoutModificationTemplate
    xmlns="http://schemas.microsoft.com/Start/2014/LayoutModification"
    xmlns:defaultlayout="http://schemas.microsoft.com/Start/2014/FullDefaultLayout"
    xmlns:start="http://schemas.microsoft.com/Start/2014/StartLayout"
    xmlns:taskbar="http://schemas.microsoft.com/Start/2014/TaskbarLayout"
    Version="1">
  <CustomTaskbarLayoutCollection PinListPlacement="Replace">
    <defaultlayout:TaskbarLayout>
      <taskbar:TaskbarPinList>
%s
      </taskbar:TaskbarPinList>
    </defaultlayout:TaskbarLayout>
  </CustomTaskbarLayoutCollection>
</LayoutModificationTemplate>`

	var taskbarApps string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			logger.Infof("Path exists: %s", path)
		} else {
			logger.Warnf("Path does not exist, but adding it to : %s", path)
		}
		taskbarApps += "        <taskbar:DesktopApp DesktopApplicationLinkPath=\"" + path + "\" />\n"
	}

	xmlContent := []byte(fmt.Sprintf(xmlTemplate, taskbarApps))

	// Write to file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	xmlPath := filepath.Join(wd, "Taskbar.xml")

	file, err := os.Create(xmlPath)
	if err != nil {
		logger.Errorf("Failed to create file: %s", err.Error())
		return
	}
	defer file.Close()

	if _, err := file.Write(xmlContent); err != nil {
		logger.Errorf("Failed to write to file: %s", err.Error())
		return
	}

	logger.Infof("LayoutModificationTemplate written to Taskbar.xml file")
}
