package vault

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"
)

func printMap(mapr interface{}) {
	b, err := json.MarshalIndent(mapr, "", "  ")
	if err != nil {
		panic(err)
	}
	println(string(b))
}

// Use the bitwarden-cli utility to execute commands with the --response and --nointeraction flags
func bw(args string) (map[string]interface{}, error) {
	args += " --response --nointeraction"
	cmd := exec.Command("bw", strings.Split(args, " ")...)

	out, err := cmd.Output()
	if err != nil && out == nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(out), &result)
	if err != nil {
		return nil, err
	}

	if result["success"] == false {
		return result, errors.New(result["message"].(string))
	}

	data := result["data"].(map[string]interface{})
	// check if the "template" key is present in the data
	if _, ok := data["template"]; ok {
		template := data["template"].(map[string]interface{})
		return template, nil
	}
	return data, nil
}
