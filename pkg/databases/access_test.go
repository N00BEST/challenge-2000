package databases_test

import (
	"testing"

	"github.com/n00best/challenge-2000/pkg/databases"
)

func TestIsSensitive(t *testing.T) {
	sensitiveAccess := databases.DatabaseAccess{
		IsSensitiveInfo: 1,
	}

	nonSensitiveAccess := databases.DatabaseAccess{
		IsSensitiveInfo: 0,
	}

	if sensitiveAccess.IsSensitive() != true {
		t.Error("The sensitive database access is not marked as one")
	}

	if nonSensitiveAccess.IsSensitive() == true {
		t.Error("The non-sensitive database access is marked as sensitive")
	}
}

func TestCountSentitiveDatabases(t *testing.T) {
	expectedTotal := 2

	countDatabases := []databases.DatabaseAccess{
		{
			Table:           "random",
			IsSensitiveInfo: 0,
		},
		{
			Table:           "sensitive",
			IsSensitiveInfo: 1,
		},
		{
			Table:           "another sensitive",
			IsSensitiveInfo: 1,
		},
		{
			Table:           "non-sensitive",
			IsSensitiveInfo: 0,
		},
	}

	total := databases.CountSensitiveDatabases(countDatabases)

	if total != expectedTotal {
		t.Errorf("Got wrong total of sensitive databases, expected %d, got %d", expectedTotal, total)
	}

	expectedDuplicatedTotal := 1

	duplicatedDatabases := []databases.DatabaseAccess{
		{
			Table:           "random",
			IsSensitiveInfo: 0,
		},
		{
			Table:           "sensitive",
			IsSensitiveInfo: 1,
		},
		{
			Table:           "sensitive",
			IsSensitiveInfo: 1,
		},
	}

	duplicatedTotal := databases.CountSensitiveDatabases(duplicatedDatabases)

	if duplicatedTotal != expectedDuplicatedTotal {
		t.Errorf("Got wrong total of duplicated sensitive databases, expected %d, got %d", expectedDuplicatedTotal, duplicatedTotal)
	}
}

func TestGetTotalSentitiveDatabases(t *testing.T) {
	expectedTotal := 1

	total := databases.GetTotalSensitiveDatabases()

	if total != expectedTotal {
		t.Errorf("Got wrong total of full sensitive databases, expected %d, got %d", expectedTotal, total)
	}
}
