package pkg

import "log/slog"

func CheckError(err error, msg string) bool {
	if err != nil {
		slog.Default().Error(msg, slog.String("error", err.Error()))
		return true
	}

	return false
}
