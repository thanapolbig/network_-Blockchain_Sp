package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Init function for start network in step instantiate first block
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	functionName := "[Init]"
	println("=======================" + functionName + "=======================")
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

/*
	=============================================================
	================== Handle Map Function Name =================
	=============================================================
*/
// Invoke function request form API server is SDK
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println(function)
	if function == "IssueGarden" {
		return t.IssueGarden(stub, args)
	} else if function == "registerConsumer" {
		return t.registerConsumer(stub, args)
	} else if function == "IssueProduct" {
		return t.IssueProduct(stub, args)
	} else if function == "addPoint" {
		return t.addPoint(stub, args)
	} else if function == "usePoint" {
		return t.usePoint(stub, args)
	} else if function == "IssuePlanYearModel" {
		return t.IssuePlanYearModel(stub, args)
	} else if function == "IssuePlanting" {
		return t.IssuePlanting(stub, args)
	} else if function == "IssueManagePlanting" {
		return t.IssueManagePlanting(stub, args)
	} else if function == "IssueHarvest" {
		return t.IssueHarvest(stub, args)
		// } else if function == "AddSelling" {
		//  	return t.IssueSelling(stub, args)
	} else if function == "Verify" {
		return t.Verify(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	} else if function == "IssueStock" {
		return t.IssueStock(stub, args)
	} else if function == "IssuePrepareStock" {
		return t.IssuePrepareStock(stub, args)
	} else if function == "AddSelling" {
		return t.IssueSellingDoc(stub, args)
	} else if function == "queryPlantHistory" {
		return t.PlantingHistory(stub, args)
	} else if function == "queryMainpage" {
		return t.queryMainpage(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting " +
		"\"IssueGarden\" " +
		"\"registerConsumer\" " +
		"\"addPoint\" " +
		"\"usePoint\" " +
		"\"IssuePlanYearModel\" " +
		"\"IssueSelling\" " +
		"\"IssuePlanting\" " +
		"\"IssueStock\" " +
		"\"IssueManagePlanting\" " +
		"\"IssueHarvest\" " +
		"\"IssueSelling\" " +
		"\"Verify\" " +
		"\"query\" " +
		"\"IssuePrepareStock\" ")
}

/*
	=========================================================
	================== Structure Dictionary =================
	=========================================================
*/
type Stock_Model struct {
	Owner          string           `json:"Owner"`
	Location       string           `json:"Location"`
	CreateDate     string           `json:"CreateDate"`
	Stock_material []Stock_material `json:"Stock_material"`
}

type Stock_material struct {
	Stock_material_id       int64  `json:"id"`
	product_unit_id         int64  `json:"product_unit_id"`
	product_unit_name       string `json:"product_unit_name"`
	Stock_material_name     string `json:"name"`
	Stock_material_is_diy   string `json:"is_diy"`
	Stock_material_buy_from string `json:"buy_from"`
	Stock_material_price    string `json:"price"`
	Stock_material_quantity string `json:"quantity"`
	Stock_material_image    string `json:"image"`
}

type Consumer_Model struct {
	App_User string `json:"app_user"`
	Platform string `json:"Platform"`
	Point    int64  `json:"Point"`
}

type Garden_Model struct {
	GardenName              string                 `json:"gardenName"`
	Garden                  string                 `json:"garden"`
	Owner                   string                 `json:"owner"`
	Areas                   string                 `json:"areas"`
	Date_final_use_chemical string                 `json:"date_final_use_chemical"`
	History_use_chemical    []History_use_chemical `json:"history_use_chemical"`
	Status                  string                 `json:"status"`
	PathImage               []string               `json:"pathimage"`
	App_User                string                 `json:"app_user"`
	Details                 string                 `json:"details"`
}

type History_use_chemical struct {
	Name_use_chemical string `json:"name_use_chemical"`
}

type PlanYearModel struct {
	Doc_ref          string     `json:"doc_ref"`
	Update_date      string     `json:"update_date"`
	Username         string     `json:"username"`
	Name             string     `json:"name"`
	Group            string     `json:"group"`
	Planting         []Planting `json:"planting_model"`
	Garden           string     `json:"garden"`
	Agri_standard    string     `json:"agri_standard"`
	Register_appuser string     `json:"register_appuser"`
}
type Planting struct {
	Plant_id    string `json:"plant_id"`
	Status      string `json:"status"`
	Plant_name  string `json:"plant_name"`
	Path_images string `json:"path_images"`
}

type Plant_model struct {
	Doc_ref               string                  `json:"doc_ref"`
	Verify_ref		      []string				  `json:"verify_ref	"`		
	Plant_Id              string                  `json:"plan_id"`
	Plant_date            string                  `json:"plant_date"`
	Plant_name            string                  `json:"plant_name"`
	Seed_type             string                  `json:"seed_type"`
	Reproduction_type     string                  `json:"reproduction_type"`
	Seed_marketplace      string                  `json:"seed_marketplace"`
	Predict_harvest       string                  `json:"predict_harvest"`
	Predict_quantity      int64                   `json:"predict_quantity"`
	Production_activities []Production_activities `json:"production_activities"`
	Harvest               []string                `json:"harvest"`
	Product_grade_a       int64                   `json:"product_grade_a"`
	Product_grade_b       int64                   `json:"product_grade_b"`
	Product_grade_c       int64                   `json:"product_grade_c"`
	Product_grade_d       int64                   `json:"product_grade_d"`
	Product_grade_e       int64                   `json:"product_grade_e"`
	Product_total_good    int64                   `json:"product_total_best"`
	Product_total_bad     int64                   `json:"product_total_bad"`
	Product_total         int64                   `json:"product_total"`
	Unit                  string                  `json:"unit"`
	Selling               []string                `json:"selling"`
	Sold_grade_a          int64                   `json:"sold_grade_a"`
	Sold_grade_b          int64                   `json:"sold_grade_b"`
	Sold_grade_c          int64                   `json:"sold_grade_c"`
	Sold_grade_d          int64                   `json:"sold_grade_d"`
	Sold_grade_e          int64                   `json:"sold_grade_e"`
	Sold_total            int64                   `json:"sold_total"`
	// Buyer		            string                  `json:"buyer"` ไม่จำเป็นต้องแสดง
	// Sold_date                string                  `json:"sold_date"`
	// App_user	               string                  `json:"app_user"`
	// Lot_no                 string                  `json:"lot_no"`
	// SellingList			  []SellingList			  `json:"sellinglist"`
	Path_images string `json:"path_images"`
	// Process_image		string 						`json:"process_image"`
	// Product_image		string 						`json:"product_image"`
	Update_date string `json:"update_date"`
}

type Production_activities struct {
	App_user          string `json:"app_user"`
	Production_id     int64  `json:"production_id"`
	Production_name   string `json:"production_name"`
	Production_date   string `json:"production_date"`
	Activities_detail string `json:"activities_detail"`
	Production_factor string `json:"production_factor"`
}

type Planting_History_model struct {
	// Production_activities []Production_activities `json:"production_activities"`
	Production_activities []Production_activities_type `json:"production_activities"`
	Harvest               []Harvest_activities_type    `json:"harvest"`
	Selling               []Selling_activities_type    `json:"selling"`
}

type Production_activities_type struct {
	Activities_type string `json:"activities_type"` // การดูแล
	Activities_date string `json:"activities_date"`
	Activities_name string `json:"activities_name"`
	Activities_tool string `json:"activities_tool"` //tool
}
type Harvest_activities_type struct {
	Activities_type string `json:"activities_type"` // การเก็บเกี่ยว
	Activities_date string `json:"activities_date"`
	Activities_name string `json:"activities_name"` //เก็บเกี่ยว
	Activities_tool string `json:"activities_tool"` //ชุดอุปกรณเก็บเกี่ยว
}
type Selling_activities_type struct {
	Activities_type string `json:"activities_type"` // การขาย
	Activities_date string `json:"activities_date"`
	Activities_name string `json:"activities_name"` //การขาย
	Activities_tool string `json:"activities_tool"` //กระสอบ

}

//type Harvest struct {
//	App_user                     string     `	json:"app_user"`
//	Harvest_date                 string     `	json:"harvest_date"`
//	Harvest_transform_check      string     `json:"harvest_transform_check"`
//	Harvesting_product_date_data string     `json:"harvesting_product_date_data"`
//	Quantity                     []Quantity `json:"quantity"`
//	Total                        []Total    `json:"total"`
//	Lote_no                      string     `json:"lote_no"`
//}

type Harvest_Model struct {
	Plant_document_ref      string `json:"plant_document_ref"`
	Harvest_date            string `json:"harvest_date"`
	App_user                string `json:"app_user"`
	Harvest_transform_check string `json:"harvest_transform_check"`
	Product_grade_a         int64  `json:"product_grade_a"`
	Product_grade_b         int64  `json:"product_grade_b"`
	Product_grade_c         int64  `json:"product_grade_c"`
	Product_grade_d         int64  `json:"product_grade_d"`
	Product_grade_e         int64  `json:"product_grade_e"`
	Product_total_good      int64  `json:"product_total_good"`
	Product_total_bad       int64  `json:"product_total_bad"`
	Product_total           int64  `json:"product_total"`
	Unit                    string `json:"unit"`
	Process_image           string `json:"process_image"`
	Product_image           string `json:"product_image"`
	Harvest_status          string `json:"harvest_status"`
	Tools                   string `json:"tools"`
}

//type Quantity struct {
//	Quantity_grade       string `json:"quantity_grade"`
//	Quantity_amount      int64  `json:"quantity_amount"`
//	Quantity_amount_sell int64  `json:"quantity_amount_sell"`
//}
//type Total struct {
//	Total_grade       string `json:"total_grade"`
//	Total_amount      int64  `json:"total_amount"`
//	Total_amount_sell int64  `json:"total_amount_sell"`
//}
type Selling struct {
	Plant_document_ref string        `json:"plant_document_ref"`
	Selling_list       []SellingList `json:"selling_list"`
}
type SellingList struct {
	Sold_grade_A int64  `json:"sold_grade_a"`
	Sold_grade_B int64  `json:"sold_grade_b"`
	Sold_grade_C int64  `json:"sold_grade_c"`
	Sold_grade_D int64  `json:"sold_grade_d"`
	Sold_grade_E int64  `json:"sold_grade_e"`
	Sold_total   int64  `json:"sold_total"`
	Buyer        string `json:"buyer"`
	Sold_date    string `json:"sold_date"`
	App_user     string `json:"app_user"`
	Lot_no       string `json:"lot_no"`
}

// type Harvest_doc struct {
// 	App_user                     string     `	json:"app_user"`
// 	Harvest_date                 string     `	json:"harvest_date"`
// 	Harvest_transform_check      string     `json:"harvest_transform_check"`
// 	Harvesting_product_date_data string     `json:"harvesting_product_date_data"`
// 	Quantity                     []Quantity `json:"quantity"`
// 	Total                        []Total    `json:"total"`
// 	Lote_no                      string     `json:"lote_no"`
// }

type VerifyModel struct {
	Id                    string                `json:"id"`
	Company_id            string                `json:"company_id"`
	User_id               string                `json:"user_id"`
	Garden_id             string                `json:"garden_id"`
	Is_pass               string                `json:"is_pass"`
	Is_approve            string                `json:"is_approve"`
	Is_draft           	  string                `json:"is_draft"`
	Inspect_date          string                `json:"inspect_date"`
	Inspect_start_time    string                `json:"inspect_start_time"`
	Inspect_end_time      string                `json:"inspect_end_time"`
	Approve_start_date    string                `json:"approve_start_date"`
	Approve_end_date      string                `json:"approve_end_date"`
	Remark                string                `json:"remark"`
	Data                  Data                  `json:"data"`
	User_fullname         string                `json:"user_fullname"`
	User_farmer_code      string                `json:"user_farmer_code"`
	User_address          string                `json:"user_address"`
	User_latitude         string                `json:"user_latitude"`
	User_longitude        string                `json:"user_longitude"`
	Zip_code              string                `json:"zip_code"`
	District_name         string                `json:"district_name"`
	Amphur_name           string                `json:"amphur_name"`
	Province_name         string                `json:"province_name"`
	User_phone            string                `json:"user_phone"`
	Group_id              string                `json:"group_id"`
	Group_name            string                `json:"group_name"`
	Problem_images        []Problem_images      `json:"problem_images"`
	Report_images         []Report_images       `json:"report_images"`
	}
type Data struct {
	Inspect_type 					string `json:"inspect_type"`
	Standard_type  					string `json:"standard_type"`
	Consideration_for_approve 		[]consideration_for_approve `json:"consideration_for_approve"`
	Farm_document  					farm_document `json:"farm_document"`
	Planting_data 					string `json:"planting_data"`
	Inspection_scope  				inspection_scope `json:"inspection_scope"`
	User_sign_data 					user_sign_data `json:"user_sign_data"`
}
type consideration_for_approve struct {
	Plan_product_support_standard 		[]plan_product_support_standard `json:"plan_product_support_standard"`
	Approvement_and_condition		  	string `json:"approvement_and_condition"`
	Period_of_change 					string `json:"period_of_change"`
	Plan_product_unsupport_standard		[]plan_product_unsupport_standard `json:"plan_product_unsupport_standard"`
	Next_inspection 					string `json:"next_inspection"`
}
type plan_product_support_standard struct {
	Product_id 		string `json:"product_id"`
	Name 			string `json:"name"`
}
type plan_product_unsupport_standard struct {
	Product_id 		string `json:"product_id"`
	Name 			string `json:"name"`
}
type farm_document struct {
	Prepare_plating 						string `json:"prepare_plating"`
	Farm_plan_document 						string `json:"farm_plan_document"`
	Organic_standard 						string `json:"organic_standard"`
	Prepare_plating_is_completelyption4 	string `json:"prepare_plating_is_completelyption4"`
	Prepare_plating_describeption5 			string `json:"prepare_plating_describeption5"`
	Farm_plan_document_is_completely 		string `json:"farm_plan_document_is_completely"`
	Farm_plan_document_describe 			string `json:"farm_plan_document_describe"`
	Organic_standard_describe 				string `json:"organic_standard_describe"`
}
type inspection_scope struct {
	Is_all_oranic_garden 	string `json:"is_all_oranic_garden"`
	Remark 					string `json:"remark"`
}
type user_sign_data struct {
	Farmer   		farmer `json:"farmer"`
	Inspector 		inspector `json:"inspector"`
	Endorser 		endorser `json:"endorser"`
}
type farmer struct {
	Sign_image   		string `json:"sign_image"`
	Sign_date 			string `json:"sign_date"`
}
type inspector struct {
	Sign_image   		string `json:"sign_image"`
	Sign_date 			string `json:"sign_date"`
	Sign_image_upload 	string `json:"sign_image_upload"`
}
type endorser struct {
	Sign_image   		string `json:"sign_image"`
	Sign_date 			string `json:"sign_date"`
	Sign_image_upload 	string `json:"sign_image_upload"`
}
type Problem_images struct {
	Id   					string `json:"id"`
	Inspect_verifier_id 	string `json:"inspect_verifier_id"`
	Image_path 				string `json:"image_path"`
}
type Report_images struct {
	Id   					string `json:"id"`
	Inspect_verifier_id 	string `json:"inspect_verifier_id"`
	Image_path 				string `json:"image_path"`
}

type querymainpageModel struct {
	//harvest5555555
	Plant_name							string `json:"plant_name"`
	Lots 								string `json:"lots"`
	Garden_name	 						string `json:"garden_name"`
	Location							string `json:"location"`
	Total_amount 						int64  `json:"total_amount"`
	Harvest_date 						string `json:"harvest_date"`
	Garden_status_product 				string `json:"garden_status_product"`
	Endosrer 							string `json:"endosrer"`
	Product_image 						string `json:"product_image"`
	//ativity_date
	Join_date							string `json:"join_date"`
	Startplanting_date					string `json:"startplanting_date"`
	Checklast_date						string `json:"checklast_date"`
	Harvestlast_date					string `json:"harvestlast_date"`
	Harvestlast_producttotal			int64  `json:"Harvestlast_producttotal"`
	Harvestlast_productunit       		string `json:"harvestlast_productunit"`
	Transport_date						string `json:"transport_date"`
	Soldlast_date						string `json:"sold_date"`
	Soldlast_producttotal			 	int64  `json:"Soldlast_producttotal"`
	//detail_farme
	Farmer_name    string `json:"farmer_name"`
	Group_fammer   string `json:"group_fammer"`
	Status_organic string `json:"status_organic"`
	Endorser       string `json:"endorser"`
	//plant
	Product_total_good         int64           `json:"product_total_good"`
	Product_total              int64           `json:"product_totald"`
	Product_unit               string          `json:"product_unit"`
	Production_seccess_percent int64           `json:"production_seccess_percent"`
	Source                     string          `json:"source"`
	Planting_type              string          `json:"planting_type"`
	Harvest_prediction         string          `json:"harvest_prediction"`
	Plant_date                 string          `json:"plant_date"`
	Process_harvest_images     string          `json:"process_harvest_images"`
	Planting_images            string          `json:"planting_images"`
	Harvestdetail              []Harvestdetail `json:"harvest"`
	Sellingdetail              []Sellingdetail `json:"selling"`
	Garden_hash                string          `json:"garden_hash"`
	Planyear_hash              string          `json:"planyesr_hash"`
	Planting_hash              string          `json:"planting_hash"`
	// Selling_hash						string `json:"selling_hash"`
	Harvest_hash string `json:"harvest_hash"`
}
type Harvestdetail struct {
	Harvest_date   string `json:"harvest_date"`
	Harvest_amount int64  `json:"harvest_amount"`
	Harvest_status string `json:"harvest_status"`
}
type Sellingdetail struct {
	Selling_market_place string `json:"selling_market_place"`
	Selling_amount       int64  `json:"selling_amount"`
	Selling_unit         string `json:"selling_unit"`
	Selling_grade        string `json:"selling_grade"`
	Selling_date         string `json:"selling_date"`
}

/*
	=========================================================
	===================== Smart Contract  ===================
	=========================================================
*/
type ProductModel struct {
	Hash_planting []string `json:"hash_planting"`
	Product_name  string   `json:"product_name"`
	LotNo         string   `json:"lotNo"`
	Product_image string   `json:"product_images"`
}

func (t *SmartContract) IssueProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueProduct]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 5 {
		println("Incorrect number of arguments. Expecting 5")
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	Hash_planting, err := getHashPlanting(args[1])
	if err != nil {
		println("getHashPlanting" + err.Error())
		return shim.Error("PlantgetHashPlantingingId" + err.Error())
	}
	//mapping args array to product model
	ProductModel := ProductModel{
		Hash_planting: Hash_planting,
		LotNo:         args[2],
		Product_name:  args[3],
		Product_image: args[4],
	}
	ProductKey := "ProductDoc|" + args[0]
	for i := 0; i < len(Hash_planting); i++ {
		PlantingDocKey := "PlantDoc|" + Hash_planting[i] //hashPlanting
		//validate document
		println("validate document :" + PlantingDocKey)
		plantingDocAsBytes, err := stub.GetState(PlantingDocKey)
		if err != nil {
			println("GetState is error" + err.Error())
			return shim.Error("GetState is error" + err.Error())
		}
		if plantingDocAsBytes == nil {
			println("PlantingDocKey " + PlantingDocKey + " Not Found in state Blockchain")
			return shim.Error("PlantingDocKey " + PlantingDocKey + " Not Found in state Blockchain")
		}
		println("validate document Not Found successfully")
		//validate document has already exist
		ProductDocExist, err := stub.GetState(ProductKey)
		if err != nil {
			println("GetState is error" + err.Error())
			return shim.Error("GetState is error" + err.Error())
		}
		if ProductDocExist != nil {
			println("stockDocKey " + ProductKey + " has Already Exist in state Blockchain")
			return shim.Error("stockDocKey " + ProductKey + " has Already Exist in state Blockchain")
		}
	}
	println("validate document has already exist successfully")

	//parser productModel as JSON to ByteArray
	ProductDocAsBytes, err := json.Marshal(ProductModel)
	if err != nil {
		println("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(ProductKey, ProductDocAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func getHashPlanting(get string) ([]string, error) {
	functionName := "[getHashPlanting]"
	println(functionName)
	var HashPlantingAsStruct []string
	var jsonData = []byte(get)
	// ByteArray to json
	err := json.Unmarshal(jsonData, &HashPlantingAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return HashPlantingAsStruct, nil
}

func (t *SmartContract) IssueStock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueStock]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 4 {
		println("Incorrect number of arguments. Expecting 4")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//mapping args array to struct model
	stockModel := Stock_Model{
		Owner:      args[1],
		Location:   args[2],
		CreateDate: args[3],
	}
	stockDocKey := "StockDoc|" + args[0]

	//validate document has already exist
	println("validate document has already exist :" + stockDocKey)
	stockDocExist, err := stub.GetState(stockDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if stockDocExist != nil {
		println("stockDocKey " + stockDocKey + " has Already Exist in state Blockchain")
		return shim.Error("stockDocKey " + stockDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser stockModel as JSON to ByteArray
	stockDocAsBytes, err := json.Marshal(stockModel)
	if err != nil {
		println("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(stockDocKey, stockDocAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func (t *SmartContract) IssuePrepareStock(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssuePrepareStock]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 10 {
		println("Incorrect number of arguments. Expecting 10")
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	//parse some args type string to new type
	Stock_material_id, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_unit_id, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}

	//mapping args array to struct model
	PrepareStock := Stock_material{
		Stock_material_id:       Stock_material_id,
		product_unit_id:         product_unit_id,
		product_unit_name:       args[3],
		Stock_material_name:     args[4],
		Stock_material_is_diy:   args[5],
		Stock_material_buy_from: args[6],
		Stock_material_price:    args[7],
		Stock_material_quantity: args[8],
		Stock_material_image:    args[9],
	}
	stockDocKey := "StockDoc|" + args[0]

	//validate document find Not Found
	println("validate document find Not Found :" + stockDocKey)
	stockDocAsBytes, err := stub.GetState(stockDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if stockDocAsBytes == nil {
		println("stockDocKey " + stockDocKey + " find Not Found in state Blockchain")
		return shim.Error("stockDocKey " + stockDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//parser stockModel as ByteArray to JSON
	stockModel := Stock_Model{}
	errUnmarshal := json.Unmarshal(stockDocAsBytes, &stockModel)
	if errUnmarshal != nil {
		println(" Error " + functionName + " unmarshaling stockModel : " + errUnmarshal.Error())
		return shim.Error(" Error " + functionName + "  unmarshaling stockModel : " + errUnmarshal.Error())
	}

	//add material to array Stock_material for stockDocument
	stockModel.Stock_material = append(stockModel.Stock_material, PrepareStock)

	//parser stockModel as JSON to ByteArray
	stockDocAfterUpdateAsBytes, err := json.Marshal(stockModel)
	if err != nil {
		println("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser stockModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(stockDocKey, stockDocAfterUpdateAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//======================================================== Consumer ======================================================================

func (t *SmartContract) registerConsumer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[registerConsumer]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 3 {
		println("Incorrect number of arguments. Expecting 3")
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//mapping args array to struct model
	consumerModel := Consumer_Model{
		App_User: args[1],
		Platform: args[2],
	}
	consumerDocKey := args[2] + "|" + args[0]

	//validate document has already exist
	println("validate document has already exist :" + consumerDocKey)
	consumerDocExist, err := stub.GetState(consumerDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if consumerDocExist != nil {
		println("consumerDocKey " + consumerDocKey + " has Already Exist in state Blockchain")
		return shim.Error("consumerDocKey " + consumerDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser consumerModel as JSON to ByteArray
	consumerDocAsBytes, err := json.Marshal(consumerModel)
	if err != nil {
		println("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(consumerDocKey, consumerDocAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func (t *SmartContract) addPoint(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[addPoint]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 4 {
		println("Incorrect number of arguments. Expecting 4")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//parse some args type string to new type
	Point, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}

	addPoint := Consumer_Model{
		Point: Point,
	}
	consumerDocKey := args[2] + "|" + args[0]

	//validate document find Not Found
	println("validate document find Not Found :" + consumerDocKey)
	consumerDocAsBytes, err := stub.GetState(consumerDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if consumerDocAsBytes == nil {
		println("consumerDocKey " + consumerDocKey + " find Not Found in state Blockchain")
		return shim.Error("consumerDocKey " + consumerDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//parser consumerModel as ByteArray to JSON
	consumerModel := Consumer_Model{}
	errUnmarshal := json.Unmarshal(consumerDocAsBytes, &consumerModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling consumerModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling consumerModel : " + errUnmarshal.Error())
	}

	//update add Point in consumer Document
	consumerModel.Point = consumerModel.Point + addPoint.Point

	//parser consumerModel as JSON to ByteArray
	consumerDocAfterUpdateAsBytes, err := json.Marshal(consumerModel)
	if err != nil {
		println("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(consumerDocKey, consumerDocAfterUpdateAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func (t *SmartContract) usePoint(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[usePoint]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 4 {
		println("Incorrect number of arguments. Expecting 4")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//parse some args type string to new type
	valueUsePoint, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	consumerDocKey := args[2] + "|" + args[0]

	//validate document find Not Found
	println("validate document find Not Found :" + consumerDocKey)
	consumerDocAsBytes, err := stub.GetState(consumerDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if consumerDocAsBytes == nil {
		println("consumerDocKey:" + consumerDocKey + " find Not Found in state Blockchain")
		return shim.Error("consumerDocKey:" + consumerDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//parser consumerModel as ByteArray to JSON
	consumerModel := Consumer_Model{}
	errUnmarshal := json.Unmarshal(consumerDocAsBytes, &consumerModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling consumerModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling consumerModel : " + errUnmarshal.Error())
	}

	//validate now balance point for consumer can use
	if consumerModel.Point < valueUsePoint {
		println("Your Points're " + strconv.FormatInt(consumerModel.Point, 10) + " not enough.request use point : " + strconv.FormatInt(valueUsePoint, 10))
		return shim.Error("Your Points're " + strconv.FormatInt(consumerModel.Point, 10) + " not enough.request use point : " + strconv.FormatInt(valueUsePoint, 10))
	}

	//update use deduct Point in consumer Document
	consumerModel.Point = consumerModel.Point - valueUsePoint

	//parser consumerModel as JSON to ByteArray
	consumerDocAfterUpdateAsBytes, err := json.Marshal(consumerModel)
	if err != nil {
		println("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser consumerModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(consumerDocKey, consumerDocAfterUpdateAsBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//======================================================== GardenInfo ======================================================================

func (t *SmartContract) IssueGarden(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueGarden]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 11 {
		println("Incorrect number of arguments. Expecting 10")
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	//mapping args array to struct model
	history_use_chemical, err := getHistory_use_chemical(args[6])
	if err != nil {
		println("History_use_chemical" + err.Error())
		return shim.Error("History_use_chemical" + err.Error())
	}
	gardenModel := Garden_Model{
		GardenName: args[1],
		Garden:     args[2],
		Owner:      args[3],
		Areas:      args[4],
		Date_final_use_chemical: args[5],
		History_use_chemical:    history_use_chemical,
		Status:                  args[7],
		App_User:                args[9],
		Details:                 args[10],
	}
	gardenModel.PathImage = append(gardenModel.PathImage, args[8])

	gardenDocKey := "GardenDoc|" + args[0]
	//validate document has already exist
	println("validate document has already exist :" + gardenDocKey)
	gardenDocExist, err := stub.GetState(gardenDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if gardenDocExist != nil {
		println("gardenDocKey " + gardenDocKey + " has Already Exist in state Blockchain")
		return shim.Error("gardenDocKey " + gardenDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser gardenModel as JSON to ByteArray
	gardenDocBytes, err := json.Marshal(gardenModel)
	if err != nil {
		println("Marshal parser gardenModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser gardenModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(gardenDocKey, gardenDocBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}
func getHistory_use_chemical(get string) ([]History_use_chemical, error) {
	functionName := "[sub] getHistory_use_chemical"
	println(functionName)
	var History_use_chemicalAsStruct []History_use_chemical
	var jsonData = []byte(get)
	//parser consumerModel.getHistory_use_chemical as ByteArray to JSON
	err := json.Unmarshal(jsonData, &History_use_chemicalAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return History_use_chemicalAsStruct, nil
}

//========================================================  PlanYearModel  ======================================================================

func (t *SmartContract) IssuePlanYearModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssuePlanYearModel]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 9 {
		println("Incorrect number of arguments. Expecting 8")
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}
	gardenDocKey := "GardenDoc|" + args[1]
	planYearDocKey := "PlanYearDoc|" + args[0]

	//mapping args array to struct model
	planYearModel := PlanYearModel{
		Doc_ref:          args[1],
		Update_date:      args[2],
		Username:         args[3],
		Name:             args[4],
		Group:            args[5],
		Garden:           args[6],
		Agri_standard:    args[7],
		Register_appuser: args[8],
	}

	//validate document find Not Found
	println("validate document find Not Found :" + gardenDocKey)
	gardenDocAsByte, err := stub.GetState(gardenDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if gardenDocAsByte == nil {
		println("gardenDocKey " + gardenDocKey + " find Not Found in state Blockchain")
		return shim.Error("gardenDocKey " + gardenDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//validate document has already exist
	println("validate document has already exist :" + planYearDocKey)
	planYearDocExist, err := stub.GetState(planYearDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if planYearDocExist != nil {
		println("planYearDocKey " + planYearDocKey + " has Already Exist in state Blockchain")
		return shim.Error("planYearDocKey " + planYearDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser gardenModel as JSON to ByteArray
	planYearDocBytes, err := json.Marshal(planYearModel)
	if err != nil {
		println("Marshal parser planYearModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser planYearModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(planYearDocKey, planYearDocBytes)
	if err != nil {
		println("PutState is error" + err.Error())
		return shim.Error("PutState is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//=================================================== Planting======================================================================

func (t *SmartContract) IssuePlanting(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssuePlanting]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 12 {
		println("Incorrect number of arguments. Expecting 12")
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}

	//parse some args type string to new type
	predict_quantity, err := strconv.ParseInt(args[11], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}

	//mapping args array to struct model
	plantModel := Plant_model{
		Doc_ref:           args[1],
		Update_date:       args[4],
		Plant_Id:          args[3],
		Plant_date:        args[4],
		Plant_name:        args[5],
		Seed_type:         args[6],
		Reproduction_type: args[7],
		Seed_marketplace:  args[8],
		Path_images:       args[9],
		Predict_harvest:   args[10],
		Predict_quantity:  predict_quantity,
	}
	plantInPlanYear := Planting{
		Plant_id:    args[0],
		Status:      "ดำเนินการปลูก",
		Plant_name:  args[5],
		Path_images: args[7],
	}
	gardenDocKey := "GardenDoc|" + args[2]
	planYearDocKey := "PlanYearDoc|" + args[1]
	plantDocKey := "PlantDoc|" + args[0]

	//validate document find Not Found
	println("validate document find Not Found :" + planYearDocKey)
	planYearDocAsByte, err := stub.GetState(planYearDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if planYearDocAsByte == nil {
		println("planYearDocKey " + planYearDocKey + " find Not Found in state Blockchain")
		return shim.Error("planYearDocKey " + planYearDocKey + " find Not Found in state Blockchain")
	}

	gardenDocExist, err := stub.GetState(gardenDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if gardenDocExist == nil {
		println("gardenDocKey" + gardenDocKey + " find Not Found in state Blockchain")
		return shim.Error("gardenDocKey " + gardenDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//validate document has already exist
	println("validate document has already exist :" + plantDocKey)

	plantDocExist, err := stub.GetState(plantDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if plantDocExist != nil {
		println("plantDocKey" + plantDocKey + " has Already Exist in state Blockchain")
		return shim.Error("plantDocKey " + plantDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser planYearModel as ByteArray to JSON
	planYearModel := PlanYearModel{}
	errUnmarshal := json.Unmarshal(planYearDocAsByte, &planYearModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling planYearModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling planYearModel : " + errUnmarshal.Error())
	}

	//update action date to PlanYearModel.Update_date
	planYearModel.Update_date = plantModel.Plant_date
	//update plantDoc(Hash) to PlanYearModel.Planting
	planYearModel.Planting = append(planYearModel.Planting, plantInPlanYear)

	//parser planYearModel as JSON to ByteArray
	planYearDocAfterUpdateAsByte, err := json.Marshal(planYearModel)
	if err != nil {
		println("Marshal parser planYearModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser planYearModel as JSON to ByteArray is error" + err.Error())
	}

	//parser planModel as JSON to ByteArray
	plantDocBytes, err := json.Marshal(plantModel)
	if err != nil {
		println("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(plantDocKey, plantDocBytes)
	if err != nil {
		println("PutState plantDoc is error" + err.Error())
		return shim.Error("PutState plantDoc is error " + err.Error())
	}
	err = stub.PutState(planYearDocKey, planYearDocAfterUpdateAsByte)
	if err != nil {
		println("PutState planYearDoc is error" + err.Error())
		return shim.Error("PutState planYearDoc is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//======================================================== IssueManagePlanting ======================================================================
func (t *SmartContract) IssueManagePlanting(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueManagePlanting]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 7 {
		println("Incorrect number of arguments. Expecting 7")
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	//parse some args type string to new type
	production_id, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}

	//mapping args array to struct model
	plantActivities := Production_activities{
		App_user:          args[1],
		Production_id:     production_id,
		Production_name:   args[3],
		Production_date:   args[4],
		Activities_detail: args[5],
		Production_factor: args[6],
	}
	plantDocKey := "PlantDoc|" + args[0]

	//validate document find Not Found
	println("validate document find Not Found :" + plantDocKey)
	plantDocAsByte, err := stub.GetState(plantDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if plantDocAsByte == nil {
		println("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
		return shim.Error("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//parser plantModel as ByteArray to JSON
	plantModel := Plant_model{}
	errUnmarshal := json.Unmarshal(plantDocAsByte, &plantModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
	}

	//update activities in planting to plantDocument.Production_activities
	plantModel.Production_activities = append(plantModel.Production_activities, plantActivities)
	//record update date
	plantModel.Update_date = args[4]

	//parser plantModel as JSON to ByteArray
	plantDocAfterUpdateAsByte, err := json.Marshal(plantModel)
	if err != nil {
		println("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(plantDocKey, plantDocAfterUpdateAsByte)
	if err != nil {
		println("PutState plantDoc is error" + err.Error())
		return shim.Error("PutState plantDoc is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//======================================================== IssueHarvest ======================================================================
func (t *SmartContract) IssueHarvest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueHarvest]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 16 {
		println("Incorrect number of arguments. Expecting 13")
		return shim.Error("Incorrect number of arguments. Expecting 13")
	}
	//parse some args type string to new type

	//Quantity, err := getQuantity(args[5])
	//if err != nil {
	//	println("Quantity" + err.Error())
	//	return shim.Error("Quantity" + err.Error())
	//}
	//Total, err := getTotal(args[6])
	//if err != nil {
	//	println("Total" + err.Error())
	//	return shim.Error("Total" + err.Error())
	//}
	product_grade_a, err := strconv.ParseInt(args[5], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_grade_b, err := strconv.ParseInt(args[6], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_grade_c, err := strconv.ParseInt(args[7], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_grade_d, err := strconv.ParseInt(args[8], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_grade_e, err := strconv.ParseInt(args[9], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	product_total_bad, err := strconv.ParseInt(args[10], 10, 64)
	if err != nil {
		println("ParseInt is error" + err.Error())
		return shim.Error("ParseInt is error" + err.Error())
	}
	//total Product_total_good
	product_total_good := product_grade_a + product_grade_b + product_grade_c + product_grade_d + product_grade_e
	//total Product_total
	product_total := product_total_bad + product_total_good

	//TODO  จับคู่ข้อมูลที่ parse api ส่งเข้ามาให้ตรงกับ ข้อมูลแต่ละฟิลให้ทีครับ
	harvestModel := Harvest_Model{
		Plant_document_ref:      args[1], //เก๋บhashHarvest
		Harvest_date:            args[2],
		App_user:                args[3],
		Harvest_transform_check: args[4],
		Product_grade_a:         product_grade_a,
		Product_grade_b:         product_grade_b,
		Product_grade_c:         product_grade_c,
		Product_grade_d:         product_grade_d,
		Product_grade_e:         product_grade_e,
		Product_total_good:      product_total_good,
		Product_total_bad:       product_total_bad,
		Product_total:           product_total,
		Unit:                    args[11],
		Process_image:           args[12],
		Product_image:           args[13],
		Harvest_status:          args[14],
		Tools:                   args[15],
	}
	plantDocKey := "PlantDoc|" + args[1]  //hashPlanting
	harvestDocKey := "Harvest|" + args[0] //hashHarvest

	// // check key not found??
	// println("[Check PlantKey not found]")
	// // PlantKey := "" + args[0]
	// PlantKey := "PlantDoc|" + args[0]
	// PlantAsByte, _ := stub.GetState(PlantKey)
	// if PlantAsByte == nil {
	// 	println("ManageKey " + PlantKey + " not found")
	// 	return shim.Error("ManageKey " + PlantKey + " not found")
	// }
	// // PlantYearKey := "" + args[1]
	// PlantYearKey := "PlanYearDoc|" + args[1]
	// PlantYearAsByte, _ := stub.GetState(PlantYearKey)
	// if PlantYearAsByte == nil {
	// 	println("PlantYearKey " + PlantYearKey + " not found")
	// 	return shim.Error("PlantYearKey " + PlantYearKey + " not found")
	// }
	// println("CheckKeyExist successflly")

	//validate document find Not Found
	println("validate document find Not Found :" + plantDocKey)
	plantDocAsByte, err := stub.GetState(plantDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if plantDocAsByte == nil {
		println("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
		return shim.Error("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//validate document has already exist
	println("validate document has already exist :" + harvestDocKey)
	harvestDocExist, err := stub.GetState(harvestDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if harvestDocExist != nil {
		println("harvestDocKey" + harvestDocKey + " has Already Exist in state Blockchain")
		return shim.Error("harvestDocKey " + harvestDocKey + " has Already Exist in state Blockchain")
	}

	//parser plantModel as ByteArray to JSON
	plantModel := Plant_model{}
	errUnmarshal := json.Unmarshal(plantDocAsByte, &plantModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
	}

	//update Harvest_Document(Hash) in array Plant_Document.Harvest
	plantModel.Harvest = append(plantModel.Harvest, args[0])
	plantModel.Product_grade_a = plantModel.Product_grade_a + harvestModel.Product_grade_a
	plantModel.Product_grade_b = plantModel.Product_grade_b + harvestModel.Product_grade_b
	plantModel.Product_grade_c = plantModel.Product_grade_c + harvestModel.Product_grade_c
	plantModel.Product_grade_d = plantModel.Product_grade_d + harvestModel.Product_grade_d
	plantModel.Product_grade_e = plantModel.Product_grade_e + harvestModel.Product_grade_e
	plantModel.Product_total_good = plantModel.Product_total_good + harvestModel.Product_total_good
	plantModel.Product_total_bad = plantModel.Product_total_bad + harvestModel.Product_total_bad
	plantModel.Product_total = plantModel.Product_total + harvestModel.Product_total
	plantModel.Update_date = harvestModel.Harvest_date
	plantModel.Unit = harvestModel.Unit
	// plantModel.Process_image = harvestModel.Process_image
	// plantModel.Process_image = harvestModel.Process_image
	// plantModel.Product_image = harvestModel.Product_image

	// //ByteArray to Json
	// PlantModel := Plant_model{}
	// Plan_year_Model := PlanYearModel{}
	// errUnmarshalplanting := json.Unmarshal(PlantAsByte, &PlantModel)
	// errUnmarshalplanyear := json.Unmarshal(PlantYearAsByte, &Plan_year_Model)

	// // Validate CCR Document State
	// if errUnmarshalplanting != nil {
	// 	//error unmarshaling
	// 	println(functionName + " Error unmarshaling CCRDocument:" + errUnmarshalplanting.Error())
	// 	return shim.Error(functionName + " Error unmarshaling CCRDocument:" + errUnmarshalplanting.Error())
	// }
	// if errUnmarshalplanyear != nil {
	// 	//error errUnmarshalplanyear
	// 	println(functionName + " Error errUnmarshalplanyear CCRDocument:" + errUnmarshalplanyear.Error())
	// 	return shim.Error(functionName + " Error errUnmarshalplanyear CCRDocument:" + errUnmarshalplanyear.Error())
	// }
	/////////////////////updateplanting
	// PlantModel.Harvest = append(PlantModel.Harvest, harvest)
	// for i := 0; i < len(Plan_year_Model.Planting); i++ {
	// 	if Plan_year_Model.Planting[i].Plant_id == args[0] { //args[0] = plantid
	// 		Plan_year_Model.Planting[i].Status = "เก็บเกี่ยวแล้ว"
	// 	}
	// 	println(Plan_year_Model.Planting[i].Plant_id + "=====" + args[0])
	// }
	// for i := 0; i < len(PlantModel.Harvest); i++ {
	// 	if PlantModel.Harvest[i].Harvesting_product_date_data == args[3] {
	// 		PlantModel.Update_date = args[3]
	// 	}
	// }
	// PlantYearAsByte, err = json.Marshal(Plan_year_Model)
	// if err != nil {
	// 	println("Marshal is error" + err.Error())
	// 	return shim.Error("Marshal is error" + err.Error())
	// }
	// //PutState
	// err = stub.PutState(PlantYearKey, PlantYearAsByte)
	// if err != nil {
	// 	println("PutState is error" + err.Error())
	// 	return shim.Error("PutState is error" + err.Error())
	// }

	//==============================================checktotal=================================================================================
	// var total_waste int64 = 0
	// var total int64 = 0
	// for j := 0; j <= len(Quantity); j++ {
	// 	total = total + harvest_doc.Quantity[j].Quantity_amount
	// }
	// total_waste = harvest_doc.Quantity[5].Quantity_amount
	// if total != harvest_doc.Total[0].Total_amount {
	// 	return shim.Error(functionName + "Overload number")
	// } else if total_waste != harvest_doc.Total[1].Total_amount {
	// 	return shim.Error(functionName + "Overload number")
	// }

	// // Json to ByteArray
	// PlantAsByte, err = json.Marshal(PlantModel)
	// if err != nil {
	// 	println("Marshal is error" + err.Error())
	// 	return shim.Error("Marshal is error" + err.Error())
	// }
	//PutState
	// err = stub.PutState(PlantKey, PlantAsByte)
	// if err != nil {
	// 	println("PutState is error" + err.Error())
	// 	return shim.Error("PutState is error" + err.Error())
	// }

	//parser plantModel as JSON to ByteArray
	plantDocAfterUpdateAsByte, err := json.Marshal(plantModel)
	if err != nil {
		println("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
	}

	harvestDocBytes, err := json.Marshal(harvestModel)
	if err != nil {
		println("Marshal is error" + err.Error())
		return shim.Error("Marshal is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(plantDocKey, plantDocAfterUpdateAsByte)
	if err != nil {
		println("PutState plantDoc is error" + err.Error())
		return shim.Error("PutState plantDoc is error" + err.Error())
	}
	err = stub.PutState(harvestDocKey, harvestDocBytes)
	if err != nil {
		println("PutState harvestDoc is error" + err.Error())
		return shim.Error("PutState harvestDoc is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

//func getQuantity(get string) ([]Quantity, error) {
//	functionName := "[getQuantity]"
//	println(functionName)
//	var QuantityAsStruct []Quantity
//	var jsonData = []byte(get)
//	// ByteArray to json
//	err := json.Unmarshal(jsonData, &QuantityAsStruct)
//	if err != nil {
//		fmt.Printf("There was an error decoding the json. err = %s", err)
//	}
//	println(functionName + " successfully")
//	return QuantityAsStruct, nil
//}
//func getTotal(get string) ([]Total, error) {
//	functionName := "[getTotal]"
//	println(functionName)
//	var TotalAsStruct []Total
//	var jsonData = []byte(get)
//	// ByteArray to json
//	err := json.Unmarshal(jsonData, &TotalAsStruct)
//	if err != nil {
//		fmt.Printf("There was an error decoding the json. err = %s", err)
//	}
//	println(functionName + " successfully")
//	return TotalAsStruct, nil
//}

func (t *SmartContract) IssueSellingDoc(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[IssueSellingDoc]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 3 {
		println("Incorrect number of arguments. Expecting 3")
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	sellinglist, err := getSellinglist(args[2])
	if err != nil {
		println("sellinglist" + err.Error())
		return shim.Error("sellinglist" + err.Error())
	}

	var total int64 = 0
	var total_grade_A int64 = 0
	var total_grade_B int64 = 0
	var total_grade_C int64 = 0
	var total_grade_D int64 = 0
	var total_grade_E int64 = 0

	SellingListModel := Selling{
		Plant_document_ref: args[0],
		// SellingList    	   :		listInModel,
	}

	for j := 0; j < len(sellinglist); j++ {
		total_grade_A = total_grade_A + sellinglist[j].Sold_grade_A
		total_grade_B = total_grade_B + sellinglist[j].Sold_grade_B
		total_grade_C = total_grade_C + sellinglist[j].Sold_grade_C
		total_grade_D = total_grade_D + sellinglist[j].Sold_grade_D
		total_grade_E = total_grade_E + sellinglist[j].Sold_grade_E
		// sellinglist[j].Sold_total = sellinglist[j].Sold_grade_A + sellinglist[j].Sold_grade_B + sellinglist[j].Sold_grade_C + sellinglist[j].Sold_grade_D + sellinglist[j].Sold_grade_E

		// total = total+sellinglist[j].sold_total
		// Sold_date = sellinglist[j].Sold_date
		total = total_grade_A + total_grade_B + total_grade_C + total_grade_D + total_grade_E
		SellingListModel.Selling_list = append(SellingListModel.Selling_list, sellinglist[j])

	} // loop j
	total = total_grade_A + total_grade_B + total_grade_C + total_grade_D + total_grade_E
	// buyer = sellinglist

	// SellingListModel.Selling_list = append(SellingListModel.Selling_list,sellinglist)

	plantDocKey := "PlantDoc|" + args[1]  //hashPlanting
	sellingDocKey := "Selling|" + args[0] //hashSellingDoc

	//validate document find Not Found
	println("validate document find Not Found :" + plantDocKey)
	plantDocAsByte, err := stub.GetState(plantDocKey)
	if err != nil {
		println("getState is error" + err.Error())
		return shim.Error("getState is error" + err.Error())
	}
	if plantDocAsByte == nil {
		println("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
		return shim.Error("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")

	//validate document has already exist
	println("validate document has already exist :" + sellingDocKey)
	sellingDocExist, err := stub.GetState(sellingDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if sellingDocExist != nil {
		println("sellingDocKey" + sellingDocKey + " has Already Exist in state Blockchain")
		return shim.Error("sellingDocKey " + sellingDocKey + " has Already Exist in state Blockchain")
	}

	//parser plantModel as ByteArray to JSON
	plantModel := Plant_model{}
	errUnmarshal := json.Unmarshal(plantDocAsByte, &plantModel)
	if errUnmarshal != nil {
		println("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
		return shim.Error("Error " + functionName + " unmarshaling plantModel : " + errUnmarshal.Error())
	}

	//update Selling_Document(Hash) in array Plant_Document.Selling
	plantModel.Selling = append(plantModel.Selling, args[0])
	plantModel.Sold_grade_a = plantModel.Sold_grade_a + total_grade_A
	plantModel.Sold_grade_b = plantModel.Sold_grade_b + total_grade_B
	plantModel.Sold_grade_c = plantModel.Sold_grade_c + total_grade_C
	plantModel.Sold_grade_d = plantModel.Sold_grade_d + total_grade_D
	plantModel.Sold_grade_e = plantModel.Sold_grade_e + total_grade_E
	plantModel.Sold_total = plantModel.Sold_total + total
	// plantModel.Buyer		 = plantModel.Buyer
	// plantModel.Sold_date    = plantModel.Sold_date
	// plantModel.App_user	   = plantModel.App_user
	// plantModel.Lot_no	   = plantModel.Lot_no

	//parser plantModel as JSON to ByteArray
	plantDocAfterUpdateAsByte, err := json.Marshal(plantModel)
	if err != nil {
		println("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser plantModel as JSON to ByteArray is error" + err.Error())
	}

	sellingDocBytes, err := json.Marshal(SellingListModel)
	if err != nil {
		println("Marshal is error" + err.Error())
		return shim.Error("Marshal is error" + err.Error())
	}
	//byteArray put to state blockchain
	err = stub.PutState(plantDocKey, plantDocAfterUpdateAsByte)
	if err != nil {
		println("PutState plantDoc is error" + err.Error())
		return shim.Error("PutState plantDoc is error" + err.Error())
	}
	err = stub.PutState(sellingDocKey, sellingDocBytes)
	if err != nil {
		println("PutState sellingDoc is error" + err.Error())
		return shim.Error("PutState sellingDoc is error" + err.Error())
	}
	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func getSellinglist(get string) ([]SellingList, error) {
	functionName := "[getSellinglist]"
	println(functionName)
	var SellinglistAsStruct []SellingList
	var jsonData = []byte(get)
	// ByteArray to json
	err := json.Unmarshal(jsonData, &SellinglistAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return SellinglistAsStruct, nil
}

// //========================================================  SELLING  ======================================================================
// //TODO เนื่องจาก harvest เปลี่ยนไป แก้ไข function ให้เน้นทำงานได้ ยังไม่เน้นเงื่อนไขตรวจยอดการเก็บเกี่ยว
// // รองรับการทำงานในกรณีที่ มีการขาย หลายตลาด (รายการขายที่เข้ามาจะเป็น array)
// func (t *SmartContract) IssueSelling(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	functionName := "[IssueSelling]"
// 	println("=======================" + functionName + "=======================")
// 	println("Input: " + args[0])

// 	//parse args as string to array split by |
// 	justString := strings.Join(args, "")
// 	args = strings.Split(justString, "|")

// 	//validate length array for args
// 	if len(args) != 6 {
// 		println("Incorrect number of arguments. Expecting 6")
// 		return shim.Error("Incorrect number of arguments. Expecting 6")
// 	}
// 	// data, err := getData(args[1])
// 	// if err != nil {
// 	// 	println("data: " + err.Error())
// 	// 	return shim.Error("data: " + err.Error())
// 	// }
// 	// DocSelling := SellingModel{
// 	// 	Data: data,
// 	// }
// 	// // check key has already exist??
// 	// println("[CheckKeyExist]")
// 	// SellingKey := "SellingDoc|" + args[0]
// 	// SellingKeycheck, err := stub.GetState(SellingKey)
// 	// if SellingKeycheck == nil {
// 	// 	// Json to ByteArray
// 	// 	Sellingbytes, err := json.Marshal(DocSelling)
// 	// 	if err != nil {
// 	// 		println("Marshal is error" + err.Error())
// 	// 		return shim.Error("Marshal is error" + err.Error())
// 	// 	}
// 	// 	err = stub.PutState(SellingKey, Sellingbytes)
// 	// 	if err != nil {d
// 	// 		println("PutState is error" + err.Error())
// 	// 		return shim.Error("PutState is error" + err.Error())
// 	// 	}
// 	// } else {
// 	// 	println("SellingKey " + SellingKey + " has Already Exist")
// 	// 	return shim.Error("SellingKey " + SellingKey + " has Already Exist")
// 	// }
// 	Sell, err := getSell(args[2])
// 	if err != nil {
// 		println("Sell" + err.Error())
// 		return shim.Error("Sell" + err.Error())
// 	}
// 	DocSelling := Selling{
// 		Sell:         Sell,
// 		Selling_date: args[3],
// 		App_user:     args[4],
// 		Lot_no:       args[5],
// 	}
// 	// check key has already exist??
// 	println("[CheckKeyExist]")
// 	// check Plantkey
// 	PlantKey := "PlantDoc|" + args[0]
// 	// get Planting AsByteArray
// 	Planting, err := stub.GetState(PlantKey)
// 	if Planting == nil {
// 		return shim.Error("PlantKey not found ")
// 	}
// 	if err != nil {
// 		println("PlantKey is error " + err.Error())
// 		return shim.Error("PlantKey is error " + err.Error())
// 	}
// 	//check Plantyear
// 	PlanyearKey := "PlanYearDoc|" + args[1]
// 	// get Plantyear AsByteArray
// 	PlanyearKeyAsByte, err := stub.GetState(PlanyearKey)
// 	if PlanyearKeyAsByte == nil {
// 		return shim.Error("PlantyearKey not found ")
// 	}
// 	if err != nil {
// 		println("PlanyearKey is error " + err.Error())
// 		return shim.Error("PlanyearKey is error " + err.Error())
// 	}

// 	functionName1 := "[Update_planting]"
// 	PlantingModel := Plant_model{}
// 	errUnmarshal := json.Unmarshal(Planting, &PlantingModel)
// 	if errUnmarshal != nil {
// 		//error unmarshaling
// 		println(functionName1 + " Error unmarshaling CCRDocument:" + errUnmarshal.Error())
// 		return shim.Error(functionName1 + " Error unmarshaling CCRDocument:" + errUnmarshal.Error())
// 	}
// 	// update Planting
// 	PlantingModel.Selling = append(PlantingModel.Selling, DocSelling)
// 	// var arrayharvest_date = 0
// 	// var arrayquantity_grade = 0

// 	//find_harvest
// 	// for i := 0; i < len(PlantingModel.Harvest); i++ {
// 	// 	//find harvest_date
// 	// 	if PlantingModel.Harvest[i].Harvest_date == args[5] {
// 	// 		arrayharvest_date = i
// 	// 	}
// 	// 	println(PlantingModel.Harvest[arrayharvest_date].Harvest_date + "=====" + args[5])
// 	// }
// 	// println(arrayharvest_date)

// 	//find_quantity_grade //////////////////////////////////////////////////////////////////////////////////////////////////
// 	// for j := 0; j < len(PlantingModel.Selling[arrayharvest_date].Sell); j++ {
// 	// 	for i := 0; i < len(PlantingModel.Harvest[arrayharvest_date].Quantity); i++ {
// 	// 		if PlantingModel.Harvest[arrayharvest_date].Quantity[i].Quantity_grade == PlantingModel.Selling[arrayharvest_date].Sell[j].Selling_grade {
// 	// 			arrayquantity_grade = i
// 	// 			println(PlantingModel.Harvest[arrayharvest_date].Quantity[i].Quantity_grade + "===" + PlantingModel.Selling[arrayharvest_date].Sell[j].Selling_grade)
// 	// 			break;
// 	// 		}
// 	// 	}
// 	// 	//check_overloadnumber
// 	// 	if PlantingModel.Harvest[arrayharvest_date].Quantity[arrayquantity_grade].Quantity_amount_sell < Sell[j].Selling_amount {
// 	// 		println(functionName + "Overload number of quantity_amount")
// 	// 		return shim.Error(functionName + "Overload number of quantity_amount")
// 	// 	}
// 	// 	//cal_selling
// 	// 	harvest_quantity := PlantingModel.Harvest[arrayharvest_date].Quantity[arrayquantity_grade].Quantity_amount_sell
// 	// 	sell_amount := PlantingModel.Selling[arrayharvest_date].Sell[j].Selling_amount
// 	// 	PlantingModel.Harvest[arrayharvest_date].Quantity[arrayquantity_grade].Quantity_amount_sell = harvest_quantity - sell_amount
// 	// }
// 	//cal_total_selling
// 	// var total int64 = 0
// 	// for j := 0; j < (len(PlantingModel.Harvest[arrayharvest_date].Quantity) - 1); j++ {
// 	// 	total = total + PlantingModel.Harvest[arrayharvest_date].Quantity[j].Quantity_amount_sell
// 	// }
// 	// PlantingModel.Harvest[arrayharvest_date].Total[0].Total_amount_sell = total

// 	// Json to ByteArray
// 	Plantingbytes, err := json.Marshal(PlantingModel)
// 	if err != nil {
// 		println("Marshal is error" + err.Error())
// 		return shim.Error("Marshal is error" + err.Error())
// 	}
// 	err = stub.PutState(PlantKey, Plantingbytes)
// 	if err != nil {
// 		println("PutState is error" + err.Error())
// 		return shim.Error("PutState is error" + err.Error())
// 	}
// 	println(functionName1 + " successfully")

// 	functionName2 := "[Update_planyear]"
// 	PlanYearModel := PlanYearModel{}
// 	// ByteArray to Json
// 	errUnmarshalPlanyear := json.Unmarshal(PlanyearKeyAsByte, &PlanYearModel)
// 	// Validate CCR Document State
// 	if errUnmarshalPlanyear != nil {
// 		//error unmarshaling
// 		println(functionName1 + " Error unmarshaling CCRDocument:" + errUnmarshalPlanyear.Error())
// 		return shim.Error(functionName1 + " Error unmarshaling CCRDocument:" + errUnmarshalPlanyear.Error())
// 	}
// 	//updateplanting
// 	// PlanYearModel.Planting = append(PlanYearModel.Planting, Planting)
// 	for i := 0; i < len(PlanYearModel.Planting); i++ {
// 		if PlanYearModel.Planting[i].Plant_id == args[0] { //args[0] = plantid
// 			PlanYearModel.Planting[i].Status = "ขายแล้ว"
// 		}
// 	}
// 	// Json to ByteArray
// 	PlanyearKeyAsByte, err = json.Marshal(PlanYearModel)
// 	if err != nil {
// 		println("Marshal is error" + err.Error())
// 		return shim.Error("Marshal is error" + err.Error())
// 	}
// 	//PutState
// 	err = stub.PutState(PlanyearKey, PlanyearKeyAsByte)
// 	if err != nil {
// 		println("PutState is error" + err.Error())
// 		return shim.Error("PutState is error" + err.Error())
// 	}
// 	println(functionName2 + " successfully")

// 	println("CheckKeyExist successflly")
// 	println(functionName + " successfully")
// 	println("===================================================================================" + functionName + "=============================================================================")
// 	return shim.Success(nil)
// }
// func getSell(get string) ([]Sell, error) {
// 	functionName := "[getSell]"
// 	println(functionName)
// 	var SellAsStruct []Sell
// 	var jsonData = []byte(get)
// 	// ByteArray to json
// 	err := json.Unmarshal(jsonData, &SellAsStruct)
// 	if err != nil {
// 		fmt.Printf("There was an error decoding the json. err = %s", err)
// 	}
// 	println(functionName + " successfully")
// 	return SellAsStruct, nil
// }

//========================================================  Verify  ======================================================================

func (t *SmartContract) Verify(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[Verify]"
	println("=======================" + functionName + "=======================")
	println("Input: " + args[0])

	//parse args as string to array split by |
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	//validate length array for args
	if len(args) != 32 {
		println("Incorrect number of arguments. Expecting 32")
		return shim.Error("Incorrect number of arguments. Expecting 32")
	}
	//mapping args array to struct model
	Data, err := getData(args[16])
	if err != nil {
		println("Data" + err.Error())
		return shim.Error("Data" + err.Error())
	}
	Problem_images, err := getProblem_images(args[29])
	if err != nil {
		println("Problem_images" + err.Error())
		return shim.Error("Problem_images" + err.Error())
	}
	Report_images, err := getReport_images(args[30])
	if err != nil {
		println("Report_images" + err.Error())
		return shim.Error("Report_images" + err.Error())
	}	
	println("[getHash_plantyear]")
	var Hash_plantyear []string
	var jsonDataHash_plantyear = []byte(args[0])
	// ByteArray to json
	errHash_plantyear := json.Unmarshal(jsonDataHash_plantyear, &Hash_plantyear)
	if errHash_plantyear != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}

	println("[getHash_planting]")
	var Hash_planting []string
	var jsonDataHash_planting = []byte(args[31])
	// ByteArray to json
	errHash_planting := json.Unmarshal(jsonDataHash_planting, &Hash_planting)
	if errHash_planting != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}

	verifyModel := VerifyModel{
		Id                    :         args[3],
		Company_id            :         args[4],
		User_id               :         args[5],
		Garden_id             :         args[6],
		Is_pass               :         args[7],
		Is_approve            :         args[8],
		Is_draft           	  :         args[9],
		Inspect_date          :         args[10],
		Inspect_start_time    :         args[11],
		Inspect_end_time      :         args[12],
		Approve_start_date    :         args[13],
		Approve_end_date      :         args[14],
		Remark                :         args[15],
		Data                  :         Data,
		User_fullname         :         args[17],
		User_farmer_code      :         args[18],
		User_address          :         args[19],
		User_latitude         :         args[20],
		User_longitude        :         args[21],
		Zip_code              :         args[22],
		District_name         :         args[23],
		Amphur_name			  :         args[24],
		Province_name         :         args[25],
		User_phone            :         args[26],
		Group_id              :         args[27],
		Group_name            :         args[28],
		Problem_images        :         Problem_images,
		Report_images         :         Report_images,
	}		
	verifyDocKey := "VerifyDoc|" + args[1]
	for j := 0; j < len(Hash_plantyear); j++ {
		planYearDocKey := "PlanYearDoc|" + Hash_plantyear[j]
		plantDocKey := "PlantDoc|" + Hash_planting[j]
		//validate document find Not Found
		plantDoc, err := stub.GetState(plantDocKey)
		if err != nil {
			println("GetState is error" + err.Error())
			return shim.Error("GetState is error" + err.Error())
		}
		if plantDoc == nil {
			println("plantDocKey" + plantDocKey + " find Not Found in state Blockchain")
			return shim.Error("plantDocKey " + plantDocKey + " find Not Found in state Blockchain")
		}
		//parser planYearModel as ByteArray to JSON
		Plant_model := Plant_model{}
		errUnmarshal := json.Unmarshal(plantDoc, &Plant_model)
		if errUnmarshal != nil {
			println("Error " + functionName + " unmarshaling Plant_model : " + errUnmarshal.Error())
			return shim.Error("Error " + functionName + " unmarshaling Plant_model : " + errUnmarshal.Error())
		}
		//update plantDoc(Hash) to PlanYearModel.Planting
		Plant_model.Verify_ref = append(Plant_model.Verify_ref,args[1])

		//parser planYearModel as JSON to ByteArray
		Plant_modelAfterUpdate_verify, err := json.Marshal(Plant_model)
		if err != nil {
			println("Marshal parser Plant_model as JSON to ByteArray is error" + err.Error())
			return shim.Error("Marshal parser Plant_model as JSON to ByteArray is error" + err.Error())
		}
		err = stub.PutState(plantDocKey, Plant_modelAfterUpdate_verify)
		if err != nil {
			println("PutState Plant_modelAfterUpdate_verify is error" + err.Error())
			return shim.Error("PutState Plant_modelAfterUpdate_verify is error" + err.Error())
		}
	
	//validate document find Not Found
	println("validate document find Not Found :" + planYearDocKey)
	planYearDocAsByte, err := stub.GetState(planYearDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if planYearDocAsByte == nil {
		println("planYearDocKey " + planYearDocKey + " find Not Found in state Blockchain")
		return shim.Error("planYearDocKey " + planYearDocKey + " find Not Found in state Blockchain")
	}
	println("validate document find Not Found successfully")
	//validate document has already exist
	// println("validate document has already exist :" + plantDocKey)
	}
	verifyDocExist, err := stub.GetState(verifyDocKey)
	if err != nil {
		println("GetState is error" + err.Error())
		return shim.Error("GetState is error" + err.Error())
	}
	if verifyDocExist != nil {
		println("verifyDocKey " + verifyDocKey + " has Already Exist in state Blockchain")
		return shim.Error("verifyDocKey " + verifyDocKey + " has Already Exist in state Blockchain")
	}
	println("validate document has already exist successfully")

	//parser verifyModel as JSON to ByteArray
	verifyDocAsBytes, err := json.Marshal(verifyModel)
	if err != nil {
		println("Marshal parser verifyModel as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser verifyModel as JSON to ByteArray is error" + err.Error())
	}

	//byteArray put to state blockchain
	err = stub.PutState(verifyDocKey, verifyDocAsBytes)
	if err != nil {
		println("PutState verifyModel is error" + err.Error())
		return shim.Error("PutState verifyModel is error" + err.Error())
	}

	//pass validate smart contract
	println(functionName + " successfully")
	println("=======================" + functionName + "=======================")
	return shim.Success(nil)
}

func getData(get string) (Data, error) {
	functionName := "[getData]"
	println(functionName)
	var DataAsStruct Data
	var jsonData = []byte(get)
	// ByteArray to json
	err := json.Unmarshal(jsonData, &DataAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return DataAsStruct, nil
}
func getProblem_images(get string) ([]Problem_images, error) {
	functionName := "[getProblem_images]"
	println(functionName)
	var Problem_imagesAsStruct []Problem_images
	var jsonData = []byte(get)
	// ByteArray to json
	err := json.Unmarshal(jsonData, &Problem_imagesAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return Problem_imagesAsStruct, nil
}
func getReport_images(get string) ([]Report_images, error) {
	functionName := "[getReport_images]"
	println(functionName)
	var Report_imageAsStruct []Report_images
	var jsonData = []byte(get)
	// ByteArray to json
	err := json.Unmarshal(jsonData, &Report_imageAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	println(functionName + " successfully")
	return Report_imageAsStruct, nil
}

//========================================================  query  ======================================================================
func (t *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}
	documentKey := args[0]
	// Get the state from the ledger
	documentAsBytes, err := stub.GetState(documentKey)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + documentKey + "\"}"
		return shim.Error(jsonResp)
	}
	if documentAsBytes == nil {
		jsonResp := documentKey
		return shim.Error(jsonResp)
	}
	jsonResp := string(documentAsBytes)
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(documentAsBytes)
}

//========================================================  query queryHistory  ======================================================================
func (t *SmartContract) queryHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	keyname := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", keyname)

	resultsIterator, err := stub.GetHistoryForKey(keyname)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//============================================== queryPlantingHistory ================================================

func (t *SmartContract) PlantingHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[queryPlantingHistory]"
	print("==========================================" + functionName + "====================================================")
	println("Input: " + args[0])
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	if len(args) != 1 {
		println("Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// println("[CheckGardenKeyExist]")
	// GardenKey := "GardenDoc|" + args[0]
	// GardenDoc, err := stub.GetState(GardenKey)
	// if GardenDoc == nil {
	// println("GardenKey " + GardenKey + " is not defined")
	// return shim.Error("GardenKey " + GardenKey + " is not defined")
	// }
	// if err != nil {
	// println("getStateGardenDoc is error" + err.Error())
	// return shim.Error("getStateGardenDoc is error" + err.Error())
	// }

	// println("[CheckPlanYearkeyExist]")
	// PlanYearkey := "PlanYearDoc|" + args[1]
	// PlanYearDoc, err := stub.GetState(PlanYearkey)
	// if PlanYearDoc == nil {
	// println("PlanYearkey " + PlanYearkey + " is not defined")
	// return shim.Error("PlanYearkey " + PlanYearkey + " is not defined")
	// }
	// if err != nil {
	// println("getStatePlanYearDoc is error" + err.Error())
	// return shim.Error("getStatePlanYearDoc is error" + err.Error())
	// }

	println("[CheckPlantingkeyNotFound]")
	Plantingkey := "PlantDoc|" + args[0]
	PlantDoc, err := stub.GetState(Plantingkey)
	if PlantDoc == nil {
		println("Plantingkey " + Plantingkey + " is not found")
		return shim.Error("Plantingkey " + Plantingkey + " is not found")
	}
	if err != nil {
		println("getStatePlantDoc is error" + err.Error())
		return shim.Error("getStatePlantDoc is error" + err.Error())
	}
	// plantHistorykey := "plantHistoryDoc|" + args[0]

	//parser Plant_model as ByteArray to JSON
	var PlantDocAsStruct Plant_model
	error := json.Unmarshal(PlantDoc, &PlantDocAsStruct)
	if error != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}

	var Production_Activitie_AsStruct []Production_activities_type
	for i := 0; i < len(PlantDocAsStruct.Production_activities); i++ {
		// Production_Activitie_AsStruct[i].activities_type = "การดูแล"
		// Production_Activitie_AsStruct[i].activities_date = PlantDocAsStruct.Production_activities[i].Production_date
		// Production_Activitie_AsStruct[i].activities_name = PlantDocAsStruct.Production_activities[i].Production_name
		// Production_Activitie_AsStruct[i].activities_tool = PlantDocAsStruct.Production_activities[i].Production_factor

		production_activities_data := Production_activities_type{
			Activities_type: "การดูแล",
			Activities_date: PlantDocAsStruct.Production_activities[i].Production_date,
			Activities_name: PlantDocAsStruct.Production_activities[i].Production_name,
			Activities_tool: PlantDocAsStruct.Production_activities[i].Production_factor,
		}
		Production_Activitie_AsStruct = append(Production_Activitie_AsStruct, production_activities_data)
	}

	var Harvest_Activitie_AsStruct []Harvest_activities_type
	for i := 0; i < len(PlantDocAsStruct.Harvest); i++ {

		println("[CheckHarvestkeyNotFound]")

		HarvestKey := "Harvest|" + PlantDocAsStruct.Harvest[i]
		HarvestDoc, err := stub.GetState(HarvestKey)

		// if HarvestDoc == nil {
		// println("HarvestKey " + HarvestKey + " is not found")
		// return shim.Error("HarvestKey " + HarvestKey + " is not found")
		// }

		if err != nil {
			println("getStateHarvestDoc is error" + err.Error())
			return shim.Error("getStateHarvestDoc is error" + err.Error())
		}

		// var Harvest_Activitie_AsStruct []Harvest_activities_type

		//parser Plant_model as ByteArray to JSON
		var HarvestDocAsStruct Harvest_Model
		error := json.Unmarshal(HarvestDoc, &HarvestDocAsStruct)
		if error != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
		}
		harvest_activities_data := Harvest_activities_type{
			Activities_type: "การเก็บเกี่ยว",
			Activities_date: HarvestDocAsStruct.Harvest_date, //ปัญหาที่ array sellinglist
			Activities_name: "เก็บเกี่ยว",
			Activities_tool: HarvestDocAsStruct.Tools,
		}
		Harvest_Activitie_AsStruct = append(Harvest_Activitie_AsStruct, harvest_activities_data)

		// Harvest_Activitie_AsStruct[i].activities_type = "การเก็บเกี่ยว"   // การเก็บเกี่ยว
		// Harvest_Activitie_AsStruct[i].activities_date = harvestModel.Harvest_date
		// Harvest_Activitie_AsStruct[i].activities_name = "เก็บเกี่ยว" 	   //เก็บเกี่ยว
		// Harvest_Activitie_AsStruct[i].activities_tool = "ชุดอุปกรณ์เก็บเกี่ยว"
	} //harvest

	var Selling_Activitie_AsStruct []Selling_activities_type
	// var Selling_AsStruct Selling_activities_type
	for i := 0; i < len(PlantDocAsStruct.Selling); i++ {

		println("[CheckSellingkeyNotfound]")

		SellingKey := "Selling|" + PlantDocAsStruct.Selling[i]
		SellingDoc, err := stub.GetState(SellingKey) //[]byte
		fmt.Println(SellingDoc)

		if SellingDoc == nil {
			println("SellingKey " + SellingKey + " is not found")
			return shim.Error("SellingKey " + SellingKey + " is not found")
		}

		if err != nil {
			println("getStateSellingDoc is error" + err.Error())
			return shim.Error("getStateSellingDoc is error" + err.Error())
		}
		// var Selling_Activitie_AsStruct []Selling_activities_type

		//parser Plant_model as ByteArray to JSON
		var SellingDocAsStruct Selling
		error := json.Unmarshal(SellingDoc, &SellingDocAsStruct)
		if error != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
		}
		fmt.Printf("%+v", SellingDocAsStruct)
		// var SellAsStruct   SellingList
		// error2 := json.Unmarshal(SellingDocAsStruct.SellingList, &SellAsStruct)
		// 	if error2 != nil {
		// 	fmt.Printf("There was an error decoding the json. err = %s", error2)
		// }
		// for j := 0; j < len(SellingDocAsStruct); j++ {

		Selling_activities_data := Selling_activities_type{
			Activities_type: "การขาย",                                     // การขาย
			Activities_date: SellingDocAsStruct.Selling_list[0].Sold_date, //SellingDocAsStruct[0].Sold_date,
			Activities_name: "ขาย",                                        //ขาย
			Activities_tool: "กระสอบ",
		}
		// Selling_AsStruct = Selling_activities_data
		Selling_Activitie_AsStruct = append(Selling_Activitie_AsStruct, Selling_activities_data)

		// }//loopj
		// Selling_Activitie_AsStruct = append(Selling_Activitie_AsStruct,Selling_Activitie_AsStruct)

	} //loopi

	Docplanthistory := Planting_History_model{
		Production_activities: Production_Activitie_AsStruct,
		Harvest:               Harvest_Activitie_AsStruct,
		Selling:               Selling_Activitie_AsStruct,
	}

	//parser Planting_History_model as JSON to ByteArray
	planthistoryDocAsBytes, err := json.Marshal(Docplanthistory)
	if err != nil {
		println("Marshal parser Planting_History_model as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser Planting_History_model as JSON to ByteArray is error" + err.Error())
	}
	// //byteArray put to state blockchain
	// err = stub.PutState(plantHistorykey, Docplanthistory)
	// if err != nil {
	// 	println("PutState plantDoc is error" + err.Error())
	// 	return shim.Error("PutState plantDoc is error" + err.Error())
	// }

	println("PutState is Successfully ")
	println(functionName + " successfully")
	println("===================================================================================" + functionName + "=============================================================================")
	return shim.Success(planthistoryDocAsBytes)
}

//============================================== queryMainpage ================================================

func (t *SmartContract) queryMainpage(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	functionName := "[queryMainpage]"
	print("==========================================" + functionName + "====================================================")
	println("Input: " + args[0])
	justString := strings.Join(args, "")
	args = strings.Split(justString, "|")

	if len(args) != 1 {
		println("Incorrect number of arguments. Expecting 1")
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	println("[CheckHarvestKeyExist]")

	HarvestKey := "Harvest|" + args[0]
	HarvestDoc, err := stub.GetState(HarvestKey)
	if HarvestDoc == nil {
		println("HarvestKey " + HarvestKey + " is not defined")
		return shim.Error("HarvestKey " + HarvestKey + " is not defined")
	}
	if err != nil {
		println("getStateHarvestDoc is error" + err.Error())
		return shim.Error("getStateHarvestDoc is error" + err.Error())
	}
	println(HarvestDoc)

	HarvestModel := Harvest_Model{}
	errUnmarshalHarvestModel := json.Unmarshal(HarvestDoc, &HarvestModel)
	if errUnmarshalHarvestModel != nil {
		//error unmarshaling
		println("HarvestModel Error unmarshaling CCRDocument:" + errUnmarshalHarvestModel.Error())
		return shim.Error("HarvestModel Error unmarshaling CCRDocument:" + errUnmarshalHarvestModel.Error())
	}
	println(HarvestModel.Harvest_transform_check)

	println("get hash_planting" + HarvestModel.Plant_document_ref)
	println("[getPlanting on blockchain]")
	Plantingkey := "PlantDoc|" + HarvestModel.Plant_document_ref
	PlantDoc, err := stub.GetState(Plantingkey)
	if PlantDoc == nil {
		println("Plantingkey " + Plantingkey + " is not defined")
		return shim.Error("Plantingkey " + Plantingkey + " is not defined")
	}
	if err != nil {
		println("getStatePlantDoc is error" + err.Error())
		return shim.Error("getStatePlantDoc is error" + err.Error())
	}
	Plantmodel := Plant_model{}
	errUnmarshalPlantmodel := json.Unmarshal(PlantDoc, &Plantmodel)
	if errUnmarshalPlantmodel != nil {
		//error unmarshaling
		println("Plantmodel Error unmarshaling CCRDocument:" + errUnmarshalPlantmodel.Error())
		return shim.Error("Plantmodel Error unmarshaling CCRDocument:" + errUnmarshalPlantmodel.Error())
	}

	println("get hash_plantyear" + Plantmodel.Doc_ref)
	println("[getplantyear on blockchain]")
	planYearDocKey := "PlanYearDoc|" + Plantmodel.Doc_ref
	PlanYearDoc, err := stub.GetState(planYearDocKey)
	if PlanYearDoc == nil {
		println("PlanYearkey " + planYearDocKey + " is not defined")
		return shim.Error("PlanYearkey " + planYearDocKey + " is not defined")
	}
	if err != nil {
		println("getStatePlanYearDoc is error" + err.Error())
		return shim.Error("getStatePlanYearDoc is error" + err.Error())
	}
	PlanYearModel := PlanYearModel{}
	errUnmarshalPlanYearModel := json.Unmarshal(PlanYearDoc, &PlanYearModel)
	if errUnmarshalPlanYearModel != nil {
		//error unmarshaling
		println("PlanYearModel Error unmarshaling CCRDocument:" + errUnmarshalPlanYearModel.Error())
		return shim.Error("PlanYearModel Error unmarshaling CCRDocument:" + errUnmarshalPlanYearModel.Error())
	}
	println("get hash_GardenKey" + PlanYearModel.Doc_ref)
	println("[gethash_GardenKeyr on blockchain]")
	GardenKey := "GardenDoc|" + PlanYearModel.Doc_ref
	GardenDoc, err := stub.GetState(GardenKey)
	if GardenDoc == nil {
		println("GardenKey " + GardenKey + " is not defined")
		return shim.Error("GardenKey " + GardenKey + " is not defined")
	}
	if err != nil {
		println("getStateGardenDoc is error" + err.Error())
		return shim.Error("getStateGardenDoc is error" + err.Error())
	}
	GardenModel := Garden_Model{}
	errUnmarshalGardenModel := json.Unmarshal(GardenDoc, &GardenModel)
	if errUnmarshalGardenModel != nil {
		//error unmarshaling
		println("GardenModel Error unmarshaling CCRDocument:" + errUnmarshalGardenModel.Error())
		return shim.Error("GardenModel Error unmarshaling CCRDocument:" + errUnmarshalGardenModel.Error())
	}
	//getharvest_last
	println("[last_harvest on blockchain]")
	HarvestKeylast := "Harvest|" + Plantmodel.Harvest[len(Plantmodel.Harvest)-1]
	println("get last_harvest" + Plantmodel.Harvest[len(Plantmodel.Harvest)-1])
	HarvestDoclast, err := stub.GetState(HarvestKeylast)
	if HarvestDoclast == nil {
		println("HarvestKey " + HarvestKeylast + " is not defined")
		return shim.Error("HarvestKey " + HarvestKeylast + " is not defined")
	}
	if err != nil {
		println("getStateHarvestDoc is error" + err.Error())
		return shim.Error("getStateHarvestDoc is error" + err.Error())
	}
	HarvestModellast := Harvest_Model{}
	errUnmarshallHarvestModellast := json.Unmarshal(HarvestDoclast, &HarvestModellast)
	if errUnmarshallHarvestModellast != nil {
		//error unmarshaling
		println("HarvestModellast Error unmarshaling CCRDocument:" + errUnmarshallHarvestModellast.Error())
		return shim.Error("HarvestModellast Error unmarshaling CCRDocument:" + errUnmarshallHarvestModellast.Error())
	}
	//getsoldlast
	println("[last_selling on blockchain]")

	println("get last_selling" + Plantmodel.Selling[len(Plantmodel.Selling)-1])
	sellingDocKeylast := "Selling|" + Plantmodel.Selling[len(Plantmodel.Selling)-1]

	SellingDoclast, err := stub.GetState(sellingDocKeylast)
	if SellingDoclast == nil {
		println("SellingKey " + sellingDocKeylast + " is not defined")
		return shim.Error("SellingKey " + sellingDocKeylast + " is not defined")
	}
	if err != nil {
		println("getStateSellingDoclast is error" + err.Error())
		return shim.Error("getStateSellingDoclast is error" + err.Error())
	}
	SellingModellast := Selling{}
	errUnmarshallSellingModellast := json.Unmarshal(SellingDoclast, &SellingModellast)
	if errUnmarshallSellingModellast != nil {
		//error unmarshaling
		println("SellingModellast Error unmarshaling CCRDocument:" + errUnmarshallSellingModellast.Error())
		return shim.Error("SellingModellast Error unmarshaling CCRDocument:" + errUnmarshallSellingModellast.Error())
	}
	//cal_tatal
	Production_seccess_percent := (Plantmodel.Product_total_good / Plantmodel.Product_total) * 100

	//get_harvestdetail
	var Harvestdetailmodel []Harvestdetail
	for i := 0; i < len(Plantmodel.Harvest); i++ {
		println("[last_harvest on blockchain]")
		HarvestKey := "Harvest|" + Plantmodel.Harvest[i]
		println("get last_harvest" + Plantmodel.Harvest[i])
		HarvestDoc, err := stub.GetState(HarvestKey)
		if HarvestDoclast == nil {
			println("HarvestKey " + HarvestKeylast + " is not defined")
			return shim.Error("HarvestKey " + HarvestKeylast + " is not defined")
		}
		if err != nil {
			println("getStateHarvestDoc is error" + err.Error())
			return shim.Error("getStateHarvestDoc is error" + err.Error())
		}
		HarvestModel := Harvest_Model{}
		errUnmarshallHarvestModel := json.Unmarshal(HarvestDoc, &HarvestModel)
		if errUnmarshallHarvestModel != nil {
			//error unmarshaling
			println("HarvestModel Error unmarshaling CCRDocument:" + errUnmarshallHarvestModel.Error())
			return shim.Error("HarvestModel Error unmarshaling CCRDocument:" + errUnmarshallHarvestModel.Error())
		}
		Harvestdetail_data := Harvestdetail{
			Harvest_date:   HarvestModel.Harvest_date,
			Harvest_amount: HarvestModel.Product_total,
			Harvest_status: HarvestModel.Harvest_status,
		}
		// Selling_AsStruct = Selling_activities_data
		Harvestdetailmodel = append(Harvestdetailmodel, Harvestdetail_data)
	}
	//get_sellingdetail
	var Sellingdetailmodel []Sellingdetail
	for i := 0; i < len(Plantmodel.Selling); i++ {
		SellingKey := "Selling|" + Plantmodel.Selling[i]
		println("get SellingKey" + Plantmodel.Selling[i])
		SellingDoc, err := stub.GetState(SellingKey)
		if SellingDoc == nil {
			println("SellingKey " + SellingKey + " is not defined")
			return shim.Error("SellingKey " + SellingKey + " is not defined")
		}
		if err != nil {
			println("getStateSellingDoc is error" + err.Error())
			return shim.Error("getStateSellingDoc is error" + err.Error())
		}
		Sellingmodel := Selling{}
		errUnmarshallSellingmodel := json.Unmarshal(SellingDoc, &Sellingmodel)
		if errUnmarshallSellingmodel != nil {
			//error unmarshaling
			println("Sellingmodel Error unmarshaling CCRDocument:" + errUnmarshallSellingmodel.Error())
			return shim.Error("Sellingmodel Error unmarshaling CCRDocument:" + errUnmarshallSellingmodel.Error())
		}
		println("Sellingmodel.Selling_list[i].Sold_grade_A")
		println(Sellingmodel.Selling_list[0].Sold_grade_A)

		for i := 0; i < len(Sellingmodel.Selling_list); i++ {
			println("Sellingmodel.Sellinglist[i].Sold_grade_A")

			println(Sellingmodel.Selling_list[i].Sold_grade_A)

			if Sellingmodel.Selling_list[i].Sold_grade_A != 0 {
				println("--------------------------------------------------")
				Sellingdetail_data := Sellingdetail{
					Selling_market_place: Sellingmodel.Selling_list[i].Buyer,
					Selling_date:         Sellingmodel.Selling_list[i].Sold_date,
					Selling_amount:       Sellingmodel.Selling_list[i].Sold_grade_A,
					Selling_unit:         "nil",
					Selling_grade:        "sold_grade_a",
				}
				// Selling_AsStruct = Selling_activities_data
				Sellingdetailmodel = append(Sellingdetailmodel, Sellingdetail_data)
			}
			if Sellingmodel.Selling_list[i].Sold_grade_B != 0 {
				Sellingdetail_data := Sellingdetail{
					Selling_market_place: Sellingmodel.Selling_list[i].Buyer,
					Selling_date:         Sellingmodel.Selling_list[i].Sold_date,
					Selling_amount:       Sellingmodel.Selling_list[i].Sold_grade_B,
					Selling_unit:         "nil",
					Selling_grade:        "sold_grade_b",
				}
				// Selling_AsStruct = Selling_activities_data
				Sellingdetailmodel = append(Sellingdetailmodel, Sellingdetail_data)
			}
			if Sellingmodel.Selling_list[i].Sold_grade_C != 0 {
				Sellingdetail_data := Sellingdetail{
					Selling_market_place: Sellingmodel.Selling_list[i].Buyer,
					Selling_date:         Sellingmodel.Selling_list[i].Sold_date,
					Selling_amount:       Sellingmodel.Selling_list[i].Sold_grade_C,
					Selling_unit:         "nil",
					Selling_grade:        "sold_grade_c",
				}
				// Selling_AsStruct = Selling_activities_data
				Sellingdetailmodel = append(Sellingdetailmodel, Sellingdetail_data)
			}
			if Sellingmodel.Selling_list[i].Sold_grade_D != 0 {
				Sellingdetail_data := Sellingdetail{
					Selling_market_place: Sellingmodel.Selling_list[i].Buyer,
					Selling_date:         Sellingmodel.Selling_list[i].Sold_date,
					Selling_amount:       Sellingmodel.Selling_list[i].Sold_grade_D,
					Selling_unit:         "nil",
					Selling_grade:        "sold_grade_d", //TODO แก้เป็นD
				}
				// Selling_AsStruct = Selling_activities_data
				Sellingdetailmodel = append(Sellingdetailmodel, Sellingdetail_data)
			}
			if Sellingmodel.Selling_list[i].Sold_grade_E != 0 {
				Sellingdetail_data := Sellingdetail{
					Selling_market_place: Sellingmodel.Selling_list[i].Buyer,
					Selling_date:         Sellingmodel.Selling_list[i].Sold_date,
					Selling_amount:       Sellingmodel.Selling_list[i].Sold_grade_E,
					Selling_unit:         "nil",
					Selling_grade:        "sold_grade_e",
				}
				// Selling_AsStruct = Selling_activities_data
				Sellingdetailmodel = append(Sellingdetailmodel, Sellingdetail_data)
			}

		}
	}
	verifyDocKey := "VerifyDoc|" + Plantmodel.Verify_ref[0]
	verifyDoc, err := stub.GetState(verifyDocKey)
	if verifyDoc == nil {
		println("verifyDocKey " + verifyDocKey + " is not defined")
		return shim.Error("verifyDocKey " + verifyDocKey + " is not defined")
	}
	if err != nil {
		println("getStateverifyDoc is error" + err.Error())
		return shim.Error("getStateverifyDoc is error" + err.Error())
	}
	VerifyModel := VerifyModel{}
	errUnmarshalverifyDoc := json.Unmarshal(verifyDoc, &VerifyModel)
	if errUnmarshalverifyDoc != nil {
		//error unmarshaling
		println("VerifyModel Error unmarshaling CCRDocument:" + errUnmarshalverifyDoc.Error())
		return shim.Error("VerifyModel Error unmarshaling CCRDocument:" + errUnmarshalverifyDoc.Error())
	}

	Docquerymainpage := querymainpageModel{
		//harvest
		Plant_name						:	Plantmodel.Plant_name,
		Lots 							:	"nil",			
		Garden_name	 					:	GardenModel.GardenName,//แหล่งที่ปลูก
		Location						:	VerifyModel.District_name + VerifyModel.Amphur_name+ VerifyModel.Province_name,
		Total_amount 					:	Plantmodel.Product_total,//จำนวนที่ผลิตได้ทั้งหมด
		Harvest_date 					:	HarvestModel.Harvest_date,//วันที่เก็บเกี่ยว
		Garden_status_product 			:	VerifyModel.Data.Farm_document.Organic_standard,
		Endosrer 						:	"nil",
		Product_image 					:   "nil",
		//ativity_date
		Join_date						:	"nil",
		Startplanting_date				:	Plantmodel.Plant_date,
		Checklast_date					:	VerifyModel.Inspect_date,
		Harvestlast_date				:	HarvestModellast.Harvest_date,
		Harvestlast_producttotal        :	HarvestModellast.Product_total,
		Harvestlast_productunit       	:	HarvestModellast.Unit,
		Transport_date					:	"nil",
		Soldlast_date					:	SellingModellast.Selling_list[len(SellingModellast.Selling_list)-1].Sold_date,
		Soldlast_producttotal			:	SellingModellast.Selling_list[len(SellingModellast.Selling_list)-1].Sold_total,
		//detail_farmer
		Farmer_name:    GardenModel.Owner,
		Group_fammer:   PlanYearModel.Group,
		Status_organic: GardenModel.Status,
		Endorser:       "nil",
		//plant
		Product_total_good:         Plantmodel.Product_total_good,
		Product_total:              Plantmodel.Product_total,
		Product_unit:               Plantmodel.Unit,
		Production_seccess_percent: Production_seccess_percent,
		Source:                 "nil",
		Planting_type:          "nil",
		Harvest_prediction:     Plantmodel.Predict_harvest,
		Plant_date:             Plantmodel.Plant_date,
		Process_harvest_images: HarvestModel.Process_image,
		Planting_images:        Plantmodel.Path_images,
		Harvestdetail:          Harvestdetailmodel,
		Sellingdetail:          Sellingdetailmodel,
		Garden_hash:            PlanYearModel.Doc_ref,
		Planyear_hash:          Plantmodel.Doc_ref,
		Planting_hash:          HarvestModel.Plant_document_ref,
		// Selling_hash						:sellingDocKeylast
		Harvest_hash: args[0],
	}
	DocquerymainpageDocAsBytes, err := json.Marshal(Docquerymainpage)
	if err != nil {
		println("Marshal parser Planting_History_model as JSON to ByteArray is error" + err.Error())
		return shim.Error("Marshal parser Planting_History_model as JSON to ByteArray is error" + err.Error())
	}

	println("quary is Successfully ")
	println(functionName + " successfully")
	println("===================================================================================" + functionName + "=============================================================================")
	return shim.Success(DocquerymainpageDocAsBytes)
}

type SmartContract struct {
}

func main() {
	// // Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
