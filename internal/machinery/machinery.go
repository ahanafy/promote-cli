package machinery

import (
	"os"
	"strings"

	"github.com/ahanafy/promote-cli/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/viper"
)

type Environments struct {
	Order           int    `yaml:"order"`
	Name            string `yaml:"name"`
	Hash            string `yaml:"hash"`
	isDefaultBranch bool   `yaml:"isDefaultBranch"`
}

func Start(check string) {
	if check == "" {
		utils.ConsoleOutputf("--check flag is required")
		utils.Debugf("--check flag is required")
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

	result := promotionSafety(check, resolvedEnvironments)
	if result {
		utils.Debugf("Safe to Promote: %s", check)
		utils.ConsoleOutputf("Safe to Promote: %s", check)
	} else {
		utils.Debugf("Not Safe to Promote: %s", check)
		utils.ConsoleOutputf("Not Safe to Promote: %s", check)
	}

	if !result {
		os.Exit(1)
	}
}

func View() {
	gitpath := "./"

	// if the user has specified a gitpath, use that instead of the default
	// open the git repo
	repo := openRepo(gitpath)

	// get the environments from the config file.
	environments := viper.GetStringSlice("environments")
	resolvedEnvironments := getResolvedEnvironments(environments, repo)
	printCurrentState(resolvedEnvironments)
}

func getResolvedEnvironments(environments []string, repo *git.Repository) []Environments {
	resolvedEnvironments := make([]Environments, 0)

	for i, v := range environments {
		utils.ConsoleOutputf("Checking %s", v)
		utils.Debugf("Checking %s", v)
		revHash, err := repo.ResolveRevision(plumbing.Revision(v))
		if err != nil {
			utils.ConsoleOutputf("Could not find Git Tag for Environment: '%s'", v)
			utils.Debugf("Could not find Git Tag for Environment: '%s'", v)
			os.Exit(0)
		}

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

	// if err is not nil check if the error is a "not a git repository" error.

	if err != nil {
		if strings.Contains(err.Error(), "repository does not exist") {
			utils.Errorf("No git repository found at %s", gitpath)
			os.Exit(0)
		}
	}

	utils.CheckIfError(err)

	return repo
}

func promotionSafety(targetEnvironment string, orderedEnvironments []Environments) bool {
	for i, v := range orderedEnvironments {
		if v.Name == targetEnvironment && v.isDefaultBranch {
			utils.ConsoleOutputf("Target environment `%s` is already at HEAD of default branch", targetEnvironment)
			utils.Debugf("Target environment `%s` is already at HEAD of default branch", targetEnvironment)
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

func printCurrentState(orderedEnvironments []Environments) {
	environmentProgressionLine := make([]string, 0)

	// variable for the furthest environment order position.
	furthestEnvironment := -1

	for _, v := range orderedEnvironments {
		if v.isDefaultBranch {
			furthestEnvironment = v.Order - 1
		}
		environmentProgressionLine = append(environmentProgressionLine, v.Name)
	}

	if furthestEnvironment >= 0 {
		environmentProgressionLine[furthestEnvironment] = "[" + environmentProgressionLine[furthestEnvironment] + "]"
	}

	// join environmentProgressionLine with arrows => and print it.
	utils.ConsoleOutputf(strings.Join(environmentProgressionLine, " => "))
	utils.Debugf(strings.Join(environmentProgressionLine, " => "))
}
