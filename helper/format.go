package helper

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NullStringToString(i sql.NullString) string {
	if i.Valid {
		return i.String
	}
	return ""
}

func StringToNullString(i string) sql.NullString {
	return sql.NullString{
		String: i,
		Valid:  true,
	}
}

func NullTimeToString(i mysql.NullTime, format string) string {
	if i.Valid {
		return i.Time.Format(format)
	}
	return ""
}

func TimeToNullTime(t time.Time) mysql.NullTime {
	return mysql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
