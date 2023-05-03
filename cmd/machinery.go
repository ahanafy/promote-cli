package cmd

import (
	"os"
	"strings"

	"github.com/ahanafy/promote-cli/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Environments struct {
	Order           int    `yaml:"order"`
	Name            string `yaml:"name"`
	Hash            string `yaml:"hash"`
	isDefaultBranch bool   `yaml:"isDefaultBranch"`
}

func start(_ *cobra.Command, _ []string) {
	if Check == "" {
		utils.Errorf("--check flag is required")
		os.Exit(0)
	}

	gitpath := "./"

	// if the user has specified a gitpath, use that instead of the default
	// open the git repo
	repo := openRepo(gitpath)

	// get the environments from the config file.
	environments := viper.GetStringSlice("environments")

	// get the resolved environments.
	resolvedEnvironments := getResolvedEnvironments(environments, repo)

	result := promotionSafety(Check, resolvedEnvironments)
	utils.Infof("Safe: %t", result)
	if !result {
		os.Exit(1)
	}
}

func getResolvedEnvironments(environments []string, repo *git.Repository) []Environments {
	resolvedEnvironments := make([]Environments, 0)

	for i, v := range environments {
		utils.Infof("Checking %s", v)
		revHash, err := repo.ResolveRevision(plumbing.Revision(v))
		utils.CheckIfError(err)

		ref, err := repo.Head()
		utils.CheckIfError(err)

		// check if the resolved hash is the same as the HEAD of the default branch.
		defaultBranchBool := false
		if *revHash == ref.Hash() {
			defaultBranchBool = true
		}
		// append the resolved environment to the slice and order it starting from 1.
		resolvedEnvironments = append(
			resolvedEnvironments,
			Environments{Order: i + 1, Name: v, Hash: revHash.String(), isDefaultBranch: defaultBranchBool},
		)
	}
	return resolvedEnvironments
}

func openRepo(gitpath string) *git.Repository {
	if viper.Get("gitpath") != "" {
		gitpath = viper.GetString("gitpath")
	}

	repo, err := git.PlainOpen(gitpath)
	errMsg := err.Error()
	if strings.Contains(errMsg, "repository does not exist") {
		utils.Infof("Not a git repository")
		os.Exit(1)
	}
	utils.CheckIfError(err)

	return repo
}

func promotionSafety(targetEnvironment string, orderedEnvironments []Environments) bool {
	for i, v := range orderedEnvironments {
		if v.Name == targetEnvironment && v.isDefaultBranch {
			utils.Infof("Target environment `%s` is already at HEAD of default branch", targetEnvironment)
			return false
		}
		if i > 0 {
			if v.Name == targetEnvironment && !v.isDefaultBranch && orderedEnvironments[i-1].isDefaultBranch {
				return true
			}
		} else {
			if v.Name == targetEnvironment && !v.isDefaultBranch {
				return true
			}
		}
	}

	return false
}
