package vault

import (
	"errors"
	"regexp"
	"strings"
)

func validateUrl(url string) error {
	if strings.HasPrefix(url, "http://") {
		return errors.New("URL must be HTTPS")
	}

	regex := regexp.MustCompile(`[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	if !regex.MatchString(url) {
		return errors.New("Invalid URL")
	}

	return nil
}

func (vault *Vault) Login(serverUrl, username, password string) (err error) {
	if serverUrl != vault.ServerUrl && serverUrl != "" {
		err = validateUrl(serverUrl)
		if err != nil {
			return err
		}
		_, err = bw("config server " + serverUrl)
		if err != nil {
			return
		}
		vault.ServerUrl = serverUrl
	}

	res, err := bw("login " + username + " " + password)
	if err != nil {
		return
	}

	// get the BW_SESSION token
	token := res["message"].(string)
	token = strings.Split(token, "\n")[2]
	token = strings.Split(token, "BW_SESSION=")[1]
	token = strings.TrimPrefix(token, "\"")
	token = strings.TrimSuffix(token, "\"")
	vault.SessionKey = token

	// TODO: Save token to env var

	err = vault.reload()
	return
}

func (vault *Vault) Logout() (err error) {
	_, err = bw("logout")
	if err != nil {
		return
	}

	vault.Status = "unauthenticated"
	return
}
