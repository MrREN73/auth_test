package common

type BaseID = int64

func IsEmptyID(b BaseID) bool {
	return b <= 0
}
