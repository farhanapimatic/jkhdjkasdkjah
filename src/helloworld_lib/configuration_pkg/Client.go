/*
 * helloworld_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package configuration_pkg

import(
	"helloworld_lib/apihelper_pkg"
)
/* Setting up enums for Environments and Servers 
*/
type Environments int

type Servers int

// Environment Enums
const (
     PRODUCTION Environments = 1 + iota
)

// Servers Enums
const (
 	 ENUM_DEFAULT Servers = 1 + iota
)

type CONFIGURATION_IMPL struct {
    /** Replace the value of defaulthost with SetDefaultHost function */
    defaulthost *string
}

/*
 * Getter function returning defaulthost
 */
func (me *CONFIGURATION_IMPL) DefaultHost() *string{
    return me.defaulthost
}

/*
 * Setter function setting up defaulthost
 */
func (me *CONFIGURATION_IMPL) SetDefaultHost(defaultHost *string) {
    me.defaulthost = defaultHost
}

// Setting up Default Environment
var Environment = PRODUCTION

// A map of environments and their corresponding servers/baseurls
var EnvironmentsMap = map[Environments](map[Servers]string){

    PRODUCTION : map[Servers]string{
        ENUM_DEFAULT:"http://{defaultHost}",
    },
}
 
// Make and return the map of parameters
func GetBaseURIParameters(config CONFIGURATION) map[string]interface{} {
     kvpMap := map[string]interface{}{
         "defaultHost": config.DefaultHost(),
    }
    return kvpMap;
}

// Gets the URL for a particular alias in the current environment and appends it with template parameters
// return the baseurl
func GetBaseURI(server Servers, config CONFIGURATION) string {
    url := EnvironmentsMap[Environment][server]
    appendedURL, _ := apihelper_pkg.AppendUrlWithTemplateParameters(url, GetBaseURIParameters(config))
    return appendedURL

}
