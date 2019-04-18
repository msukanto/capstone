/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"fmt"
	"bytes"
	 "strconv"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the vehicle structure.  Structure tags are used by encoding/json library
type Vehicle struct {
	VehicleId string  `json:"vehicleid"`
        Owner string `json:"owner"`
	ChasisNumber   string `json:"chasisnumber"`
	EngineNumber  string `json:"enginenumber"`
	VehicleModel  string `json:"vehiclemodel"`
	VehicleMake  string `json:"vehiclemake"`
	YearOfManufacturing string `json:"yearofmanufacturing"`
	Colour string `json:"colour"`
	SeatingCapacity int `json:"seatingcapacity"`
	VehicleInitialValue int `json:"vehicleinitialvalue"`
}

type Customer struct {
	CustomerFirstName   string `json:"custfname"`
	CustomerLastName   string `json:"custlname"`
	CustomerIdType  string `json:"custidtype"`
	CustomerId string `json:"custid"`
	CustomerAddress  string `json:"custaddress"`
	CustomerEmail  string `json:"custemail"`
	CustomerPhone  string `json:"custphone"`
}

type Dealer struct {
	DealerName   string `json:"dealername"`
	DealerId string `json:"dealerid"`
	DealerAddress  string `json:"dealeraddress"`
	DealerEmail  string `json:"dealeremail"`
	DealerPhone  string `json:"dealerphone"`
}

type Reseller struct {
	ResellerName   string `json:"resellername"`
	ResellerId string `json:"resellerid"`
	ResellerAddress  string `json:"reselleraddress"`
	ResellerEmail  string `json:"reselleremail"`
	ResellerPhone  string `json:"resellerphone"`
}

type InsuranceDetails struct {
	PolicyNumber string `json:"policynumber"`
	VehicleRegistrationNumber   string `json:"vehicleregistrationnumber"`
	CustomerId string `json:"custid"`
	StartDate  string `json:"startdate"`
	EndDate  string `json:"enddate"`
	PlanCode string `json:"plancode"`
	CoverageIds string `json:"coverageids"`
	InsuranceAgentId string `json:"insuranceagentid"`
	CoverageAmount int `json:"coverageamount"`
	PremiumAmount int `json:"premiumamount"`
}

type InsuranceClaimDetails struct {
	PolicyNumber   string `json:"policynumber"`
	VehicleRegistrationNumber   string `json:"vehicleregistrationnumber"`
	VehicleId string  `json:"vehicleid"`
	ClaimId string `json:"claimid"`
	ClaimAmount int `json:"claimamount"`
	ClaimDate string `json:"claimdate"`
	AccidentDate string `json:"accidentdate"`
	AccidentDescription string `json:"accidentdescription"`
	SurveyorDetails string `json:"surveyordetails"`
	EstimatedRepairCost int `json:"estimatedrepaircost"`
	ActualRepairCost int `json:"actualrepaircost"`
	ApprovedClaimAmount int `json:"ApprovedClaimAmount"`
	CashlessIndicator bool `json:"cashlessindicator"`
	ClaimStatus string `json:"claimstatus"`
}

type ServiceDetails struct {
	VehicleRegistrationNumber   string `json:"vehicleregistrationnumber"`
	VehicleId string  `json:"vehicleid"`
	CustomerId   string `json:"customerid"`
	ServiceId string `json:"serviceid"`
	ServiceCenterId string `json:"servicecenterid"`
	ServiceTechnician   string `json:"servicetechnician"`
	JobCardNumber string `json:"jobcardnumber"`
	JobsPerformed string `json:"jobsperformed"`
	PartsRepaired string `json:"partsrepaired"`
	PartsReplaced string `json:"partsreplaced"`
	ServiceDate string `json:"servicedate"`
}

type RegistrationDetails struct {
	VehicleRegistrationNumber   string `json:"vehicleregistrationnumber"`
	VehicleId string  `json:"vehicleid"`
	CustomerId   string `json:"customerid"`
	RegistrationDate   string `json:"registrationdate"`
	RegistrationStatus   string `json:"registrationstatus"`
	VehicleValue int `json:"vehiclevalue"`
	RegisteringRTA   string `json:"registeringrta"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addVehicle" {
		return s.addVehicle(APIstub, args)
	} else if function == "getVehicle" {
		return s.getVehicle(APIstub,args)
	} else if function == "transferVehicle" {
		return s.transferVehicle(APIstub, args)
	} else if function == "getVehicleHistory" {
		return s.getVehicleHistory(APIstub, args)
	}
	/*
	else if function == "addInsurace" {
		return s.addInsurace(APIstub, args)
	} else if function == "addRTANumber" { 
		return s.addRTANumber(APIstub, args)
	} else if function == "changeOwnership" {
		return s.changeOwnership(APIstub, args)
	} else if function == "addService" {
		return s.addService(APIstub, args)
	} else if function == "addnewClaim" {
		return s.addnewClaim(APIstub, args)
	}
	*/

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) addVehicle(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 9 inputs for vehicle details")
	}
	vehicleId := args[2] + "" + args[1]
	seatingCapacity, _ :=strconv.Atoi(args[6])
	vehicleInitialValue, _:=strconv.Atoi(args[7])
	var vehicle = Vehicle{VehicleId:vehicleId, Owner:"Audi Dealer",ChasisNumber: args[0], EngineNumber: args[1], VehicleModel: args[2], VehicleMake: args[3], YearOfManufacturing: args[4], Colour: args[5], SeatingCapacity:seatingCapacity,VehicleInitialValue:vehicleInitialValue}

	vehicleAsBytes, _ := json.Marshal(vehicle)
	APIstub.PutState(vehicleId, vehicleAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) getVehicle(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	   vehicleId := args[0];

        vehicleAsBytes, _ := APIstub.GetState(vehicleId)
	/* var vehicle Vehicle
        err := json.Unmarshal(vehicleAsBytes, &vehicle)
        if err != nil {
                return shim.Error("Issue with vehicle json unmarshaling")
        } */

	return shim.Success(vehicleAsBytes)
}

/*TransferVehicle expects 2 args, 1st arg=vehicleId, 2nd arg=new owner (customer) */
func (s *SmartContract) transferVehicle(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	vehicleAsBytes, _ := APIstub.GetState(args[0])
	vehicle := Vehicle{}
	json.Unmarshal(vehicleAsBytes, &vehicle)

	vehicle.Owner = args[1] 

	vehicleAsBytes, _ = json.Marshal(vehicle)
	APIstub.PutState(args[0], vehicleAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getVehicleHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	vehicleId := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(vehicleId)
	if err != nil {
		return shim.Error("Error retrieving Vehicle history.")
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("getVehicleHistory:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

