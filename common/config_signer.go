package common

import (
	"crypto/sha256"
	"fmt"

	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/ecc"
)

// ConfigSigner singer with use config keys
type ConfigSigner struct {
	Keys []*ecc.PrivateKey `json:"keys"`
	api  *API
}

// NewConfigSigner create config signer with keys in config in api
func NewConfigSigner(api *API) *ConfigSigner {
	res := &ConfigSigner{
		Keys: make([]*ecc.PrivateKey, 0),
		api:  api,
	}
	for _, k := range api.Cfg.Keys {
		res.Keys = append(res.Keys, &k.PriKey)
	}
	for _, k := range api.Cfg.Prikeys {
		res.Keys = append(res.Keys, &k)
	}
	return res
}

func (c *ConfigSigner) add(wifKey string) error {
	privKey, err := ecc.NewPrivateKey(wifKey)
	if err != nil {
		return err
	}
	c.Keys = append(c.Keys, privKey)
	return nil
}

// AvailableKeys imp Signer interface
func (c *ConfigSigner) AvailableKeys() (out []ecc.PublicKey, err error) {
	for _, k := range c.Keys {
		out = append(out, k.PublicKey())
	}
	return
}

// ImportPrivateKey imp Signer interface
func (c *ConfigSigner) ImportPrivateKey(wifPrivKey string) (err error) {
	return c.add(wifPrivKey)
}

// Sign imp Signer interface
func (c *ConfigSigner) Sign(tx *eos.SignedTransaction, chainID []byte, requiredKeys ...ecc.PublicKey) (*eos.SignedTransaction, error) {
	txdata, cfd, err := tx.PackedTransactionAndCFD()
	if err != nil {
		return nil, err
	}

	sigDigest := sigDigest(chainID, txdata, cfd)

	keyMap := c.keyMap()
	for _, key := range requiredKeys {
		privKey := keyMap[key.String()]
		if privKey == nil {
			return nil, fmt.Errorf("private key for %q not in keybag", key)
		}

		sig, err := privKey.Sign(sigDigest)
		if err != nil {
			return nil, err
		}

		tx.Signatures = append(tx.Signatures, sig)
	}

	return tx, nil
}

func (c *ConfigSigner) keyMap() map[string]*ecc.PrivateKey {
	out := map[string]*ecc.PrivateKey{}
	for _, key := range c.Keys {
		out[key.PublicKey().String()] = key
	}
	return out
}

func sigDigest(chainID, payload, contextFreeData []byte) []byte {
	h := sha256.New()
	if len(chainID) == 0 {
		_, _ = h.Write(make([]byte, 32, 32))
	} else {
		_, _ = h.Write(chainID)
	}
	_, _ = h.Write(payload)

	if len(contextFreeData) > 0 {
		h2 := sha256.New()
		_, _ = h2.Write(contextFreeData)
		_, _ = h.Write(h2.Sum(nil))
	} else {
		_, _ = h.Write(make([]byte, 32, 32))
	}
	return h.Sum(nil)
}
