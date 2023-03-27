package main

import (
	"fmt"

	"example.com/courier-service/model"
)

var baseDeliveryCost, noOfPackages int
var pkgDetailList = []*model.PackageDetail{}
var couponMap map[string]*model.Coupon = getCouponMap()

func main() {
	getInitialInput()       //Get the base delivery cost and number of package
	getPkgInputList()       //Get the list of package detail
	calculateDeliveryCost() //For each package detail, calculate the deliveryCost and discount

	var anyKey string
	fmt.Println("Type any character and enter to close the program...")
	fmt.Scan(&anyKey)
}

func getInitialInput() {
	fmt.Println("Please enter Base Delivery Cost and Number of Package separated by space: ")
	_, err := fmt.Scan(&baseDeliveryCost, &noOfPackages)

	if err != nil {
		fmt.Print("Base Delivery Cost and Number of Package must be a number...\n\n")
		err = nil
		getInitialInput()
	}
}

func getPkgInputList() {
	var pkgId, offerCode string
	var pkgWeight, distance int
	fmt.Println("Please enter", noOfPackages, "lines of package details:")

	for i := 0; i < noOfPackages; i++ {
		_, err := fmt.Scan(&pkgId, &pkgWeight, &distance, &offerCode)

		if err != nil {
			fmt.Print("Package detail not in \"string int int string\" format, please enter the list again...\n\n")
			pkgDetailList = []*model.PackageDetail{}
			i = -1
			err = nil
		} else {
			pkgDetailList = append(pkgDetailList, &model.PackageDetail{PkgId: pkgId, PkgWeight: pkgWeight, Distance: distance, OfferCode: offerCode})
			fmt.Printf("Saved %v...\n", pkgId)
		}
	}
}

func calculateDeliveryCost() {
	fmt.Print("\nCalculating Delivery Cost...\n\n")
	for _, val := range pkgDetailList {
		var discount, deliveryCost float64
		deliveryCost = float64(baseDeliveryCost) + (float64(val.PkgWeight) * 10) + (float64(val.Distance) * 5)
		discount = calculateDiscount(deliveryCost, val)
		val.Discount = discount
		deliveryCost -= discount
		val.DeliveryCost = deliveryCost
	}

	fmt.Print("\n*****Final Result*****\n")
	for _, pkg := range pkgDetailList {
		fmt.Println(pkg.PkgId, pkg.Discount, pkg.DeliveryCost, pkg.DeliveryTime)
	}
}

func calculateDiscount(deliveryCost float64, pkgDetail *model.PackageDetail) (result float64) {
	coupon, couponExist := couponMap[pkgDetail.OfferCode]
	if couponExist {
		if pkgDetail.PkgWeight <= coupon.MaxWeight && pkgDetail.PkgWeight >= coupon.MinWeight && pkgDetail.Distance <= coupon.MaxDistance && pkgDetail.Distance >= coupon.MinDistance {
			result = deliveryCost * float64(coupon.DiscountPerc) / 100
			fmt.Printf("%v: Calculated discount is %v.\n", pkgDetail.PkgId, result)
		} else {
			fmt.Println(pkgDetail.PkgId, ": Weight or distance does not meet coupon", pkgDetail.OfferCode, "criteria.")
		}
	} else {
		fmt.Println(pkgDetail.PkgId, ": Coupon with offer code", pkgDetail.OfferCode, "does not exist.")
	}
	return result
}

func getCouponMap() map[string]*model.Coupon {
	var couponMap = make(map[string]*model.Coupon)

	couponMap["OFR001"] = &model.Coupon{OfferCode: "OFR001", DiscountPerc: 10, MinDistance: 0, MaxDistance: 199, MinWeight: 70, MaxWeight: 200}
	couponMap["OFR002"] = &model.Coupon{OfferCode: "OFR002", DiscountPerc: 7, MinDistance: 50, MaxDistance: 150, MinWeight: 100, MaxWeight: 250}
	couponMap["OFR003"] = &model.Coupon{OfferCode: "OFR003", DiscountPerc: 5, MinDistance: 50, MaxDistance: 250, MinWeight: 10, MaxWeight: 150}

	return couponMap
}
