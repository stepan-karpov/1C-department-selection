package ui

import (
	"fmt"
	common "main/common"
	db "main/server/db"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func OutputRegisteredStats() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	subscriptionsInfo, err := db.GetAllSubscriptions(common.DB_FILE_PATH)
	if err != nil {
		fmt.Println("Error retrieving subscriptions:", err)
		return
	}

	fmt.Println("Registered Subscriptions:")
	fmt.Printf("%-5s %-15s %-20s\n", "ID", "IP", "Experiment Number")
	fmt.Println(strings.Repeat("-", 42))

	for _, subscription := range subscriptionsInfo {
		fmt.Printf("%-5d %-15s %-20d\n", subscription.ID, subscription.IP, subscription.ExperimentNumber)
	}

	fmt.Printf("\nDo you want to start an experiment? (yes/no)\n\n")
	fmt.Printf("Your input is: ")
}