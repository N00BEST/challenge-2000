package databases

type DatabaseAccess struct {
	Username        string `json:"username"`
	Table           string `json:"table"`
	IsSensitiveInfo int8   `json:"is_sensitive_info"`
}

func (da DatabaseAccess) IsSensitive() bool {
	return da.IsSensitiveInfo == 1
}

func CountSensitiveDatabases(databases []DatabaseAccess) int {
	sensitiveDatabases := make(map[string]bool, 0)

	for _, database := range databases {
		if !database.IsSensitive() {
			continue
		}

		if _, exists := sensitiveDatabases[database.Table]; !exists {
			sensitiveDatabases[database.Table] = true
		}
	}

	return len(sensitiveDatabases)
}

func GetTotalSensitiveDatabases() int {
	databases := []DatabaseAccess{
		{
			Table:           "random",
			IsSensitiveInfo: 0,
		},
		{
			Table:           "sensitive",
			IsSensitiveInfo: 1,
		},
	}

	return CountSensitiveDatabases(databases)
}
