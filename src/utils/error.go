package utils

import (
    "errors"
)

func HandleError(err error) {
    if err != nil {
        log.Println(err)
    }
}
