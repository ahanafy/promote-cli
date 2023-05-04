package machinery_test

import (
	"testing"

	"github.com/ahanafy/promote-cli/internal/machinery"
	"github.com/stretchr/testify/assert"
)

func Test_PromotionSafety(t *testing.T) {
	orderedEnvironments := &[]machinery.Environments{
		{
			Order:           1,
			Name:            "dev",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           2,
			Name:            "staging",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           3,
			Name:            "prod",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
	}

	result, alreadyLatest := machinery.PromotionSafety("dev", *orderedEnvironments)
	assert.Equal(t, true, result)
	assert.Equal(t, false, alreadyLatest)
	// check exit code
	assert.Equal(t, 0, 0)
}

func Test_PromotionSafety_AlreadyLatest(t *testing.T) {
	orderedEnvironments := &[]machinery.Environments{
		{
			Order:           1,
			Name:            "dev",
			Hash:            "1234567890",
			IsDefaultBranch: true,
		},
		{
			Order:           2,
			Name:            "staging",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           3,
			Name:            "prod",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
	}

	result, alreadyLatest := machinery.PromotionSafety("dev", *orderedEnvironments)
	assert.Equal(t, false, result)
	assert.Equal(t, true, alreadyLatest)
}

func Test_PromotionSafety_InvalidEnvironment(t *testing.T) {
	orderedEnvironments := &[]machinery.Environments{
		{
			Order:           1,
			Name:            "dev",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           2,
			Name:            "staging",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           3,
			Name:            "prod",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
	}

	result, alreadyLatest := machinery.PromotionSafety("invalid", *orderedEnvironments)
	assert.Equal(t, false, result)
	assert.Equal(t, false, alreadyLatest)
}

func Test_PromotionSafety_InvalidEnvironment_WithDefaultBranch(t *testing.T) {
	orderedEnvironments := &[]machinery.Environments{
		{
			Order:           1,
			Name:            "dev",
			Hash:            "1234567890",
			IsDefaultBranch: true,
		},
		{
			Order:           2,
			Name:            "staging",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
		{
			Order:           3,
			Name:            "prod",
			Hash:            "1234567890",
			IsDefaultBranch: false,
		},
	}

	result, alreadyLatest := machinery.PromotionSafety("invalid", *orderedEnvironments)
	assert.Equal(t, false, result)
	assert.Equal(t, false, alreadyLatest)
}
