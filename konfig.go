package konfig

import "encoding/json"
import "log"
import "fmt"
import "io/ioutil"

// configLayer manages config file for pipe setup.
// Overlay read file on top of defaults resulting in at least read settings.
func ConfigLayer(n map[string]interface{}) map[string]interface{} {

	// set known defaults here .. 
	//
	m := map[string]interface{}{
		"Target": "http://127.0.0.1:9200",
		"Name":     "thevortex",
		"User":     "",
		"Password": "",
	}

	// map overlay defaults
	//
	for key, val := range n {
		m[key] = val
	}

	return m
}

// configRead reads local json config file.
// Then it returns a decoded map.
func ConfigRead(s string) (map[string]interface{}, error) {

	var m, n map[string]interface{} // zero value

	// Read config file from install directory
	//
	file, err := ioutil.ReadFile(s)
	if err != nil { // soft error
		log.Printf("Egopipe config Get file error #%v, Defaults used. ", err)
		return m, err

	}

	// json to map
	//
	err = json.Unmarshal(file, &n)
	if err != nil {
		return m, err
	}
	return n, nil

}

// getConf call to get config info from file and layer it with defaults
func GetConf(fpn string) map[string]interface{} {

	// Returns zero value or decode
	//
	read, er := ConfigRead(fpn)
	if er != nil {
		fmt.Println(er)
	}
	// Returns overlay
	return ConfigLayer(read)

}

// verify the passed parameter names
func RequiredKeys(c map[string]interface{}, r []string) bool {

	// are irequired keys in config map?
	//
	for _, s := range r {
		_, xist := c[s]
		if !xist {
			panic("Missing config parameter")
		}
	}
	return true
}
