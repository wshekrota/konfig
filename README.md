# konfig - input your parameters
Layered approach to config value usage.

We usually set simple config values in our progrrams but would be nice if they could have default underlying values.
Just like a birthday layer cake.

![Alt text](image.png)

The storage is a simple json file distributed multiline for readability.

Read the file from disk and process it.

*func GetConf(string) map[string]interface{}*

Returns map with key/values. Values need to be asserted.
Add routine to check for required settings.

*func RequiredKeys(map[string]interface{}, []string) bool*
