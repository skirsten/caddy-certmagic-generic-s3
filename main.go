package caddycmgs3

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/certmagic"
	cmgs3 "github.com/thomersch/certmagic-generic-s3"
)

func init() {
	caddy.RegisterModule(new(GenericS3))
}

// GenericS3 defines the options for the Caddy module
type GenericS3 struct {
	Storage *cmgs3.S3Storage

	Endpoint        string `json:"endpoint"`
	Bucket          string `json:"bucket"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`

	Prefix string `json:"prefix"`

	// EncryptionKey is optional. If you do not wish to encrypt your certficates and key inside the S3 bucket, leave it empty.
	EncryptionKey string `json:"encryption_key,omitempty"`
}

func (gs3 *GenericS3) Provision(context caddy.Context) (err error) {
	gs3.Storage, err = cmgs3.NewS3Storage(cmgs3.S3Opts{
		Endpoint:        gs3.Endpoint,
		Bucket:          gs3.Bucket,
		AccessKeyID:     gs3.AccessKeyID,
		SecretAccessKey: gs3.SecretAccessKey,
		ObjPrefix:       gs3.Prefix,
		EncryptionKey:   []byte(gs3.EncryptionKey),
	})

	return
}

func (gs3 *GenericS3) CertMagicStorage() (certmagic.Storage, error) {
	return gs3.Storage, nil
}

func (gs3 *GenericS3) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.storage.generic_s3",
		New: func() caddy.Module { return new(GenericS3) },
	}
}
