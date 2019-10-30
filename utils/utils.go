package utils

func JsonResponse(status bool, message string, data interface{}) map[string]interface{} {
	return map[string]interface{} {"status": status, "message": message, "data": data}
}

