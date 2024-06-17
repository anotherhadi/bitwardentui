package vault

type VaultStatus uint8

const (
	VaultStatusUnauthenticated VaultStatus = 0
	VaultStatusLocked          VaultStatus = 1
	VaultStatusUnlocked        VaultStatus = 1
)

type Vault struct {
	ServerUrl  string
	Status     string // unauthenticated, locked, unlocked
	LastSync   string
	Search     string
	UserEmail  string
	UserId     string
	SessionKey string
}

func LoadVault() (Vault, error) {
	vault := Vault{}

	// Default:
	vault.ServerUrl = "bitwarden.com"
	vault.Status = "unauthenticated"

	// Load status
	status, err := bw("status")
	if err != nil {
		return vault, err
	}

	// check if key exists before asigning
	if _, ok := status["serverUrl"]; ok {
		vault.ServerUrl = status["serverUrl"].(string)
	}
	if _, ok := status["lastSync"]; ok {
		vault.LastSync = status["lastSync"].(string)
	}
	if _, ok := status["userEmail"]; ok {
		vault.UserEmail = status["userEmail"].(string)
	}
	if _, ok := status["userId"]; ok {
		vault.UserId = status["userId"].(string)
	}
	if _, ok := status["status"]; ok {
		vault.Status = status["status"].(string)
	}

	return vault, nil
}

func (vault *Vault) reload() (err error) {
	newVault, err := LoadVault()
	if err != nil {
		return
	}
	newVault.SessionKey = vault.SessionKey
	*vault = newVault
	return
}
