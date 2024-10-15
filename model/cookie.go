package models

import (
	"fmt"
	"net/http"
	"strconv"
)

func CreateCookie(w http.ResponseWriter, r *http.Request, UserID int) {
	cookie := http.Cookie{
		Name:  "userID",
		Value: fmt.Sprint(UserID),
		Path:  "/",
	}
	http.SetCookie(w, &cookie)

	userID:=UserID

	// Debug UserID
	fmt.Println("Session has been created:", userID)
}

func ReceiveCookie(r *http.Request) int {
	cookie, err := r.Cookie("userID")
	if err != nil {
		return 0
	}

	userIDStr := cookie.Value
	userIDINT, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("Error converting userID to integer:", err)
		return 0
	}

	return userIDINT
}