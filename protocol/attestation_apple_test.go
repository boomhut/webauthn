package protocol

import (
	"crypto/sha256"
	"testing"

	"github.com/boomhut/webauthn/metadata"
)

func Test_verifyAppleFormat(t *testing.T) {
	type args struct {
		att            AttestationObject
		clientDataHash []byte
	}

	successAttResponse := attestationTestUnpackResponse(t, appleTestResponse["success"]).Response.AttestationObject
	successClientDataHash := sha256.Sum256(attestationTestUnpackResponse(t, appleTestResponse["success"]).Raw.AttestationResponse.ClientDataJSON)

	tests := []struct {
		name    string
		args    args
		want    string
		want1   []any
		wantErr bool
	}{
		{
			"success",
			args{
				successAttResponse,
				successClientDataHash[:],
			},
			string(metadata.AnonCA),
			nil,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := verifyAppleFormat(tt.args.att, tt.args.clientDataHash, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyAppleFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("verifyAppleFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var appleTestResponse = map[string]string{
	`success`: `{
		"rawId": "U5cxFNxLbU9-SAi1K7k9atYwXhghkAMbxpL__VPtBlw",
		"id": "U5cxFNxLbU9-SAi1K7k9atYwXhghkAMbxpL__VPtBlw",
		"response": {
		  "clientDataJSON": "eyJ0eXBlIjoid2ViYXV0aG4uY3JlYXRlIiwiY2hhbGxlbmdlIjoia093TXZFMm1RTzZvdTBCMGpqRDBWQSIsIm9yaWdpbiI6Imh0dHBzOi8vNmNjM2M5ZTc5NjdhLm5ncm9rLmlvIn0",
		  "attestationObject": "o2NmbXRlYXBwbGVnYXR0U3RtdKJjYWxnJmN4NWOCWQJIMIICRDCCAcmgAwIBAgIGAXUCfWGDMAoGCCqGSM49BAMCMEgxHDAaBgNVBAMME0FwcGxlIFdlYkF1dGhuIENBIDExEzARBgNVBAoMCkFwcGxlIEluYy4xEzARBgNVBAgMCkNhbGlmb3JuaWEwHhcNMjAxMDA3MDk0NjEyWhcNMjAxMDA4MDk1NjEyWjCBkTFJMEcGA1UEAwxANjEyNzZmYzAyZDNmZThkMTZiMzNiNTU0OWQ4MTkyMzZjODE3NDZhODNmMmU5NGE2ZTRiZWUxYzcwZjgxYjViYzEaMBgGA1UECwwRQUFBIENlcnRpZmljYXRpb24xEzARBgNVBAoMCkFwcGxlIEluYy4xEzARBgNVBAgMCkNhbGlmb3JuaWEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR5_lkIu1EpyAk4t1TATSs0DvpmFbmHaYv1naTlPqPm_vsD2qEnDVgE6KthwVqsokNcfb82nXHKFcUjsABKG3W3o1UwUzAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB_wQEAwIE8DAzBgkqhkiG92NkCAIEJjAkoSIEIJxgAhVAs-GYNN_jfsYkRcieGylPeSzka5QTwyMO84aBMAoGCCqGSM49BAMCA2kAMGYCMQDaHBjrI75xAF7SXzyF5zSQB_Lg9PjTdyye-w7stiqy84K6lmo8d3fIptYjLQx81bsCMQCvC8MSN-aewiaU0bMsdxRbdDerCJJj3xJb3KZwloevJ3daCmCcrZrAPYfLp2kDOshZAjgwggI0MIIBuqADAgECAhBWJVOVx6f7QOviKNgmCFO2MAoGCCqGSM49BAMDMEsxHzAdBgNVBAMMFkFwcGxlIFdlYkF1dGhuIFJvb3QgQ0ExEzARBgNVBAoMCkFwcGxlIEluYy4xEzARBgNVBAgMCkNhbGlmb3JuaWEwHhcNMjAwMzE4MTgzODAxWhcNMzAwMzEzMDAwMDAwWjBIMRwwGgYDVQQDDBNBcHBsZSBXZWJBdXRobiBDQSAxMRMwEQYDVQQKDApBcHBsZSBJbmMuMRMwEQYDVQQIDApDYWxpZm9ybmlhMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEgy6HLyYUkYECJbn1_Na7Y3i19V8_ywRbxzWZNHX9VJBE35v-GSEXZcaaHdoFCzjUUINAGkNPsk0RLVbD4c-_y5iR_sBpYIG--Wy8d8iN3a9Gpa7h3VFbWvqrk76cCyaRo2YwZDASBgNVHRMBAf8ECDAGAQH_AgEAMB8GA1UdIwQYMBaAFCbXZNnFeMJaZ9Gn3msS0Btj8cbXMB0GA1UdDgQWBBTrroLE_6GsW1HUzyRhBQC-Y713iDAOBgNVHQ8BAf8EBAMCAQYwCgYIKoZIzj0EAwMDaAAwZQIxAN2LGjSBpfrZ27TnZXuEHhRMJ7dbh2pBhsKxR1dQM3In7-VURX72SJUMYy5cSD5wwQIwLIpgRNwgH8_lm8NNKTDBSHhR2WDtanXx60rKvjjNJbiX0MgFvvDH94sHpXHG6A4HaGF1dGhEYXRhWJhWHo8_bWPQzAMKYRIrGXu__PkMUfuqHM4RH7Jea4WDgkUAAAAAAAAAAAAAAAAAAAAAAAAAAAAUomGfdaNI-cYgWrq2klNk97zkcg-lAQIDJiABIVggef5ZCLtRKcgJOLdUwE0rNA76ZhW5h2mL9Z2k5T6j5v4iWCD7A9qhJw1YBOirYcFarKJDXH2_Np1xyhXFI7AASht1tw"},
		"type": "public-key"
	  }`,
}
