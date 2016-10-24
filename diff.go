package ginsa

import (
	"fmt"
	"github.com/fatih/color"
	"sort"
)

type Diff interface {
	String() string
}

type AllDiff struct {
	NewBanks map[string]*Bank
	Banks    []Diff
	Branches map[string][]Diff
}

func (a *AllDiff) Out() {
	cyan := color.New(color.FgCyan, color.Bold).PrintfFunc()
	blue := color.New(color.FgBlue, color.Bold).PrintfFunc()
	cyan("=> Banks diff(%d)\n", len(a.Banks))

	for _, d := range a.Banks {
		fmt.Println("   " + d.String())
	}

	cyan("=> Branches diff\n")

	dCodes := []string{}
	for code, _ := range a.Branches {
		dCodes = append(dCodes, code)
	}

	sort.Strings(dCodes)

	for _, code := range dCodes {
		diffs := a.Branches[code]
		bank := a.NewBanks[code]
		blue("===> %s(%s)'s branches diff\n", bank.Name, bank.Code)

		for _, d := range diffs {
			fmt.Println("     " + d.String())
		}
	}
}

type AddBankDiff struct {
	Bank *Bank
}

func (d *AddBankDiff) String() string {
	c := color.New(color.FgGreen, color.Bold).SprintFunc()
	return c("+   Bank") + ": " + d.Bank.Name + "(" + d.Bank.Code + ")"
}

type RemoveBankDiff struct {
	Bank *Bank
}

func (d *RemoveBankDiff) String() string {
	c := color.New(color.FgRed, color.Bold).SprintFunc()
	return c("-   Bank") + ": " + d.Bank.Name + "(" + d.Bank.Code + ")"
}

type ChangeBankDiff struct {
	OldBank *Bank
	NewBank *Bank
}

func (d *ChangeBankDiff) String() string {
	c := color.New(color.FgYellow, color.Bold).SprintFunc()
	return c("+/- Bank") + ": " + d.OldBank.Name + " -> " + d.NewBank.Name + "(" + d.NewBank.Code + ")"
}

type AddBranchDiff struct {
	Branch *Branch
}

func (d *AddBranchDiff) String() string {
	c := color.New(color.FgGreen, color.Bold).SprintFunc()
	return c("+   Branch") + ": " + d.Branch.Name + "(" + d.Branch.Code + ")"
}

type RemoveBranchDiff struct {
	Branch *Branch
}

func (d *RemoveBranchDiff) String() string {
	c := color.New(color.FgRed, color.Bold).SprintFunc()
	return c("-   Branch") + ": " + d.Branch.Name + "(" + d.Branch.Code + ")"
}

type ChangeBranchDiff struct {
	OldBranch *Branch
	NewBranch *Branch
}

func (d *ChangeBranchDiff) String() string {
	c := color.New(color.FgYellow, color.Bold).SprintFunc()
	return c("+/- Branch") + ": " + d.OldBranch.Name + " -> " + d.NewBranch.Name + "(" + d.NewBranch.Code + ")"
}

func DiffSourceData(old *SourceData, now *SourceData) *AllDiff {
	allDiff := &AllDiff{
		NewBanks: now.Banks,
		Branches: map[string][]Diff{},
	}

	allDiff.Banks = DiffBanks(old.Banks, now.Banks)

	for code, branches := range now.Branches {
		if oldBranches, ok := old.Branches[code]; ok {
			bdiff := DiffBranches(oldBranches, branches)
			if len(bdiff) > 0 {
				allDiff.Branches[code] = bdiff
			}
		}
	}

	return allDiff
}

func DiffBanks(old map[string]*Bank, now map[string]*Bank) []Diff {
	found := map[string]bool{}
	diffs := []Diff{}

	for code, bank := range old {
		found[code] = true
		if nowBank, ok := now[code]; ok {
			if nowBank.Name != bank.Name {
				diffs = append(diffs, &ChangeBankDiff{OldBank: bank, NewBank: nowBank})
			}
		} else {
			diffs = append(diffs, &RemoveBankDiff{Bank: bank})
		}
	}

	for code, bank := range now {
		if _, ok := found[code]; !ok {
			diffs = append(diffs, &AddBankDiff{Bank: bank})
		}
	}

	return diffs
}

func DiffBranches(old map[string]*Branch, now map[string]*Branch) []Diff {
	found := map[string]bool{}
	diffs := []Diff{}

	for code, bank := range old {
		found[code] = true
		if nowBranch, ok := now[code]; ok {
			if nowBranch.Name != bank.Name {
				diffs = append(diffs, &ChangeBranchDiff{OldBranch: bank, NewBranch: nowBranch})
			}
		} else {
			diffs = append(diffs, &RemoveBranchDiff{Branch: bank})
		}
	}

	for code, bank := range now {
		if _, ok := found[code]; !ok {
			diffs = append(diffs, &AddBranchDiff{Branch: bank})
		}
	}

	return diffs
}
