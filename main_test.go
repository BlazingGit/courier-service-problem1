package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"example.com/courier-service/model"
)

func TestCalculateDeliveryCost(t *testing.T) {
	testDataList := getTestData()

	for i, testData := range testDataList {
		processTestData(testData)
		calculateDeliveryCost()

		fmt.Print("\n***Comparing output and expected***\n", i)
		for j, expected := range testData.ExpectedOutput {
			output := pkgDetailList[j].PkgId +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", pkgDetailList[j].Discount), "0"), ".") +
				" " +
				strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", pkgDetailList[j].DeliveryCost), "0"), ".")

			if expected != output {
				t.Errorf("Dataset %v output %v not equal to expected %v", i, output, expected)
			}
		}
	}
}

func processTestData(testData *model.TestData) {
	pkgDetailList = []*model.PackageDetail{}
	for i, input := range testData.Input {
		s := strings.Split(input, " ")
		if i == 0 {
			base, _ := strconv.Atoi(s[0])
			baseDeliveryCost = base
			noOfPkg, _ := strconv.Atoi(s[1])
			noOfPackages = noOfPkg

		} else {
			weight, _ := strconv.Atoi(s[1])
			distance, _ := strconv.Atoi(s[2])
			pkgDetailList = append(pkgDetailList, &model.PackageDetail{PkgId: s[0], PkgWeight: weight, Distance: distance, OfferCode: s[3]})
		}
	}
}

func getTestData() []*model.TestData {
	input := []string{
		"100 3",
		"PKG1 5 5 OFR001",
		"PKG2 15 5 OFR002",
		"PKG3 10 100 OFR003",
	}

	expectedOutput := []string{
		"PKG1 0 175",
		"PKG2 0 275",
		"PKG3 35 665",
	}

	testDataList := []*model.TestData{}
	testDataList = append(testDataList, &model.TestData{Input: input, ExpectedOutput: expectedOutput})
	return testDataList
}
