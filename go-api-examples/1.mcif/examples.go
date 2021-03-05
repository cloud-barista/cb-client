// Example of Cloud-Barista Go API.
//
// * Cloud-Barista: https://github.com/cloud-barista

package main

import (
  "fmt"
  "time"
  "log"

  "github.com/cloud-barista/cb-spider/interface/api"
)


/**
  Global Variables & Struct Definition
*/


//  Data Structure to contain user configuration
type ConnectionInfoTestData struct {
  // driver
  DriverName string
  DriverLibFileName string  

  // credential
  CredentialName string
  KeyValueList []api.KeyValue

  // region, zone
  RegionName string
  Region string
  Zone string

  // connection config
  ConnectionConfigName string
}

var spiderServer string

/**
  Initialize
*/
func init() {  
  // 2K is Spider gRPC Daemon Port
  spiderServer = "localhost:2048"
}
  
  //----------------------------------------------------------- Set Test Data
  // 0: Variable Environments for AWS 
  // 1: Variable Environments for GCP

  var targetCSP = []string {
	"AWS",
	"GCP",
  }

  var connConfigNameList = []string {
    "aws-ohio-connection-config-01",  		// for AWS
    "gcp-ohio-region-connection-config-01",     // for GCP
  }

  // AWS credential info
  var AWSCredentialList = []api.KeyValue{
      api.KeyValue{Key: "ClientId", Value: "xxxxxxxxxx"},
      api.KeyValue{Key: "ClientSecret", Value: "xxxxxxxxxxxxxxxxxxxxxx"},
  }
  // GCP credential info
  var GCPCredentialList = []api.KeyValue{
      api.KeyValue{Key: "PrivateKey", Value: "-----BEGIN PRIVATE KEY-----\nxxxxxxxxxxx\n-----END PRIVATE KEY-----"},
      api.KeyValue{Key: "ProjectID", Value: "powerkimhub"},
      api.KeyValue{Key: "ClientEmail", Value: "powerkimhub@powerkimhub.iam.gserviceaccount.com"},
  }

  var connInfoTestData = []ConnectionInfoTestData {
    { // for AWS
      "aws-driver-01", "aws-driver-v1.0.so", // driver
      "aws-credential-01", AWSCredentialList, // credential
      "aws-(ohio)us-east-2", "us-east-2", "us-east-2a", // region, zone
      connConfigNameList[0], // connection config
    },

    { // for GCP
      "gcp-driver-01", "gcp-driver-v1.0.so", // driver
      "gcp-credential-01", GCPCredentialList, // credential
      "gcp-us-central1-us-central1-a", "us-central1", "us-central1-a", // region, zone
      connConfigNameList[1], // connection config
    },
  }

//----------------------------------------------------------- Set Test Data


/******************************************
  1. Cloud Connection Info Manager Test  
******************************************/
func main() {

  // 1. Cloud Connection Info Manager(CIM) Test
  fmt.Print("\n\n====================================================\n")
  fmt.Print("\n\n==== 1. Cloud Connection Info Manager(CIM) Test ====\n")
  fmt.Print("\n\n====================================================\n")
  CIM_Simple_Test()
  CIM_With_Config_Test()  

  CIM_Create_Info_Test(0) // AWS Test
  CIM_Create_Info_Test(1) // GCP Test
  CIM_IOFormat_Test()

}

/******************************************
  1. Create CloudInfoManager
  2. Setup env.  
  3. Open New Session
  4. Close (with defer)
  5. Call API
******************************************/
func CIM_Simple_Test() {
  fmt.Print("\n\n============= CloudInfoManager: Simple Test =============\n")

  // 1. Create CloudInfoManager
  cim := api.NewCloudInfoManager()
  err := cim.SetServerAddr(spiderServer)
  if err != nil {
    log.Fatal(err)
  }

  // 2. Setup env.
  err = cim.SetTimeout(90 * time.Second)
  if err != nil {
    log.Fatal(err)
  }

  /* skip in this examples
  err = cim.SetTLSCA(os.Getenv("CBSPIDER_ROOT") + "/certs/ca.crt")
  if err != nil {
    log.Fatal(err)
  }  
  err = cim.SetJWTToken("xxxxxxxxxxxxxxxxxxx")
  if err != nil {
    log.Fatal(err)
  }
  */

  // 3. Open New Session
  err = cim.Open()
  if err != nil {
          log.Fatal(err)
  }
  // 4. Close (with defer)
  defer cim.Close()

  // 5. Call API
  result, err := cim.ListCloudOS()
  if err != nil {
          log.Fatal(err)
  }

  fmt.Printf("\nresult :\n%s\n", result)
}

/******************************************
  1. Create CloudInfoManager
  2. Setup env. with a config file
  3. Open New Session
  4. Close (with defer)
  5. Call API
******************************************/ 
func CIM_With_Config_Test() {
  fmt.Print("\n\n============= CloudInfoManager: Simple Test with a Configuration file =============\n")

  // 1. Create CloudInfoManager
  cim := api.NewCloudInfoManager()

  // 2. Setup env. with a config file
  err := cim.SetConfigPath("./conf/server.yaml")
  if err != nil {
    log.Fatal(err)
  }

  // 3. Open New Session
  err = cim.Open()
  if err != nil {
    log.Fatal(err)
  }
  // 4. Close (with defer)
  defer cim.Close()

  // 5. Call API
  result, err := cim.ListCloudOS()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("\nresult :\n%s\n", result)
}

/******************************************
  1. Create CloudInfoManager
  2. Setup env. with a config file
  3. Open New Session
  4. Close (with defer)
  5. Call API
    (1) create driver info
    (2) create credential info
    (3) create region info
    (4) create Connection config
******************************************/ 
func CIM_Create_Info_Test(csp int) {
  fmt.Print("\n\n============= CloudInfoManager: Create Driver/Credential/Region/ConnConfig Info Test =============\n")

  // 1. Create CloudInfoManager
  cim := api.NewCloudInfoManager()

  // 2. Setup env. with a config file
  err := cim.SetConfigPath("./conf/server.yaml")
  if err != nil {
    log.Fatal(err)
  }

  // 3. Open New Session
  err = cim.Open()
  if err != nil {
    log.Fatal(err)
  }
  // 4. Close (with defer)
  defer cim.Close()

  // 5. Call API

  // (1) create driver info
  fmt.Print("\n\n\t============= CloudInfoManager: (1) create driver info Test =============\n")
  reqCloudDriver := &api.CloudDriverReq{
    DriverName:        connInfoTestData[csp].DriverName,
    ProviderName:      targetCSP[csp],
    DriverLibFileName: connInfoTestData[csp].DriverLibFileName,
  }  
  result, err := cim.CreateCloudDriverByParam(reqCloudDriver)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\nresult :\n%s\n", result)

  // (2) create credential info
  fmt.Print("\n\n\t============= CloudInfoManager: (2) create credential info Test =============\n")
  reqCredential := &api.CredentialReq{
    CredentialName: connInfoTestData[csp].CredentialName,
    ProviderName:   targetCSP[csp],
    KeyValueInfoList: connInfoTestData[csp].KeyValueList,
  }
  result, err = cim.CreateCredentialByParam(reqCredential)
  if err != nil {
          log.Fatal(err)
  }
  fmt.Printf("\nresult :\n%s\n", result)

  // (3) create region info    
  fmt.Print("\n\n\t============= CloudInfoManager: (3) create region info Test =============\n")
  reqRegion := &api.RegionReq{
    RegionName:   connInfoTestData[csp].RegionName,
    ProviderName: targetCSP[csp],
    KeyValueInfoList: []api.KeyValue{
            api.KeyValue{Key: "Region", Value: connInfoTestData[csp].Region},
            api.KeyValue{Key: "Zone", Value: connInfoTestData[csp].Zone},
    },
  }
  result, err = cim.CreateRegionByParam(reqRegion)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\nresult :\n%s\n", result)

  // (4) create Connection config
  fmt.Print("\n\n\t============= CloudInfoManager: (4) create connection config info Test =============\n")
  reqConnectionConfig := &api.ConnectionConfigReq{
    ConfigName:     connInfoTestData[csp].ConnectionConfigName,
    ProviderName:   targetCSP[csp],
    DriverName:     connInfoTestData[csp].DriverName,
    CredentialName: connInfoTestData[csp].CredentialName,
    RegionName:     connInfoTestData[csp].RegionName,
  }
  result, err = cim.CreateConnectionConfigByParam(reqConnectionConfig)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("\nresult :\n%s\n", result)
}

/******************************************
  1. Create CloudInfoManager
  2. Setup env. with a config file
  3. Open New Session
  4. Close (with defer)
  5. Call API
    (1) input: json, output: json
    (2) input: json, output: yaml
    (3) input: yaml, output: yaml
    (4) input: param, output: json
******************************************/
func CIM_IOFormat_Test() {
  fmt.Print("\n\n============= CloudInfoManager: API IO Format Test =============\n")

  // 1. Create CloudInfoManager
  cim := api.NewCloudInfoManager()

  // 2. Setup env. with a config file
  err := cim.SetConfigPath("./conf/server.yaml")
  if err != nil {
          log.Fatal(err)
  }

  // 3. Open New Session
  err = cim.Open()
  if err != nil {
          log.Fatal(err)
  }
  // 4. Close (with defer)
  defer cim.Close()


  // (1) input: json, output: json
  fmt.Print("\n\n\t============= CloudInfoManager: (1) API IO Format Test <input: json, output: json> =============\n")
  err = cim.SetInType("json")
  if err != nil {
    log.Fatal(err)
  }
  err = cim.SetOutType("json")
  if err != nil {
    log.Fatal(err)
  }
  doc := `{
          "DriverName":"aws-driver-01"
  }`
  result, err := cim.GetCloudDriver(doc)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\njson result :\n%s\n", result)


  // (2) input: json, output: yaml
  fmt.Print("\n\n\t============= CloudInfoManager: (2) API IO Format Test <input: json, output: yaml> =============\n")
  err = cim.SetOutType("yaml")
  if err != nil {
    log.Fatal(err)
  }
  result, err = cim.GetCloudDriver(doc)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\nyaml result :\n%s\n", result)


  // (3) input: yaml, output: yaml
  fmt.Print("\n\n\t============= CloudInfoManager: (3) API IO Format Test <input: yaml, output: yaml> =============\n")
  err = cim.SetInType("yaml")
  if err != nil {
    log.Fatal(err)
  }
  doc = `
          DriverName: aws-driver-01
        `
  result, err = cim.GetCloudDriver(doc)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\nyaml result :\n%s\n", result)


  // (4) input: param, output: json
  fmt.Print("\n\n\t============= CloudInfoManager: (4) API IO Format Test <input: param, output: yaml> =============\n")
  err = cim.SetOutType("json")
  if err != nil {
    log.Fatal(err)
  }
  result, err = cim.GetCloudDriverByParam("aws-driver-01")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("\njson result :\n%s\n", result)
}
