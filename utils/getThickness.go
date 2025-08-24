package utils

func GetThickness(totalPage int) string {
	if totalPage >= 100 {
		return "tebal"
	}
	return "tipis"
}