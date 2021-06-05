package main

/////////////////////////////////getwowstats/////////////////////////////
// getwowstats is a cmd-line utility written in GO that can interact    /
//            with the Blizzard World of Warcraft API.                  /
/////////////////////////////////////////////////////////////////////////
///////////////////Author: Gerald L. Hevener Jr. M.S.////////////////////
/////////////////////////////////getwowstats/////////////////////////////

// Import required packages
import (
	"fmt"
	"os"
    "encoding/json"
	"io/ioutil"
	"net/http"
	flag "github.com/ogier/pflag"
)

// Declare Variable types
var (
	region			      string
	character		      string
	realm			  	  string
	method		       	  string
	version				  string
)

// Declare program's main function
func main() {
	
	// If user does not supply flags print tool usage
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Uncomment below to debug user supplied cmd line options
	//fmt.Printf("Using API token for World of Warcraft: %s\n", wow_api_token)
	//fmt.Printf("Using two letter country code region : %s\n", region)
    //fmt.Printf("Using character name: %s\n", character)
	//fmt.Printf("Using realm: %s\n", realm)
	//fmt.Printf("Calling method: %s\n", method)

}

// Declare init function to parse user supplied cmd line options
func init() {

    flag.StringVarP(&region, "locale", "l", "", "Your two letter country code ( locale )")
	flag.StringVarP(&character, "character", "c", "", "Name of WoW character to lookup stats about")
	flag.StringVarP(&realm, "realm", "r", "", "Realm of WoW character to lookup stats about")
	flag.StringVarP(&method, "method", "m", "", "Method in getwowstats to call")
	flag.StringVarP(&version, "version", "v", "", "Current version of getwowstats")
   
	// Parse CMD line ARGS
	flag.Parse()

	// Process cmd line ARGS and call appropriate function
    if method == "GetProfileSummary" {
	  GetProfileSummary()
	}

	// Print software version of getwowstats
	if version != "" {
		GetWowStatsVersion()
	}

}

// Set version of getwowstats
func GetWowStatsVersion() {
 
	getwowstats_version := "1.0.0"
	fmt.Println("getwowstats: Version " + getwowstats_version)

}

// Retrieve Character Profile summary
func GetProfileSummary() {

    // Get WoW API token
	val, env_var := os.LookupEnv("WOW_API_TOKEN")
    if env_var {
	  // Uncomment to debug value of ENV VAR WOW_API_TOKEN	
	  //fmt.Println(val)
	  wow_api_endpoint := "https://us.api.blizzard.com/profile/wow/character/" + realm + "/" + character + "?namespace=profile-" + region + "&locale=en_US&access_token=" + val
	
	// Create an HTTP client
    client := &http.Client{}

	// Process API response and errors
    resp, err := client.Get(wow_api_endpoint)
	if err != nil {
		fmt.Println("Failed to call API!", err)
	} else {
		//Print API response
		data, _ := ioutil.ReadAll(resp.Body)

		// Uncomment below to print entire API json response
        // fmt.Println(string(data))

		// Unmarshall to decode API json response
	    var profile map[string]interface{}
		json.Unmarshal([]byte(data),&profile)

		// Print out character profile summary
		fmt.Println("Name:", profile["name"])
		fmt.Println("Gender:", profile["gender"].(map[string]interface{})["name"])
		fmt.Println("Faction:", profile["faction"].(map[string]interface{})["name"])
		fmt.Println("Race:", profile["race"].(map[string]interface{})["name"])
		fmt.Println("Class:", profile["character_class"].(map[string]interface{})["name"])
		fmt.Println("Realm:", profile["realm"].(map[string]interface{})["name"])
		fmt.Println("Level:", profile["level"])
		fmt.Println("Active Spec:", profile["active_spec"].(map[string]interface{})["name"])
		fmt.Println("Experience:", profile["experience"])
        fmt.Println("Achieveent Points:", profile["achievement_points"])
        fmt.Println("Last Login:", profile["last_login_timestamp"])
        fmt.Println("Avg. Item Lvl.:", profile["average_item_level"])
        fmt.Println("Equipped Item Lvl.:", profile["equipped_item_level"])

		// Close API connection
		defer resp.Body.Close()
	}   
	
	} 
}	
