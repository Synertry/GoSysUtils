/*
 *           Copyright Synertry 2022.
 *  Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE_1_0.txt or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Source: https://gosamples.dev/pretty-print-json/

package JSON

import (
	"bytes"
	"encoding/json"
	"io"
)

/*
	Function json.MarshalIndent generates JSON encoding of the value with indentation.
	You can specify a prefix of each JSON line and indent copied one or more times according to the indentation level.
	In our example, we pretty-print JSON using four spaces for indentation.
*/
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

/*
	If you use json.Encode, you can set indentation through Encoder.
	SetIndent method similarly to as in marshaling, by defining a prefix and indent.
*/
func PrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "\t")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

/*
	Package encoding/json also has a useful function json.Indent to beautify JSON string without indentation to JSON with indentation.
	The function needs the source JSON, output buffer, prefix, and indent.
*/
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "\t"); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
