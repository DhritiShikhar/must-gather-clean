package obfuscator

import (
	"testing"

	"github.com/openshift/must-gather-clean/pkg/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSSHObfuscatorStatic(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		output string
		report map[string]string
	}{
		{
			name:   "valid SSH key",
			input:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8xw==",
			output: "ssh-rsa",
		},
		{
			name:   "string and SSH key",
			input:  "some string ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8xw==",
			output: "some string ssh-rsa",
		},
		{
			name:   "invalid SSH key",
			input:  "invalid SSH key",
			output: "invalid SSH key",
		},
		{
			name:   "empty SSH key",
			input:  "",
			output: "",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			o, err := NewSSHObfuscator(schema.ObfuscateReplacementTypeStatic, NewSimpleTracker())
			require.NoError(t, err)
			assert.Equal(t, tc.output, o.Contents(tc.input))
		})
	}
}

func TestSSHObfuscatorConsistent(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  []string
		output []string
		report map[string]string
	}{
		{
			name:   "valid SSH key",
			input:  []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8xw=="},
			output: []string{"x-0000000001-x"},
		},
		{
			name:   "multiple ssh",
			input:  []string{"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu9xw== ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8x2== ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8x3=="},
			output: []string{"x-0000000001-x x-0000000002-x x-0000000003-x"},
		},
		{
			name:   "string and SSH key",
			input:  []string{"some string ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCbXCby9r69mn+lGn7/mjZRkr+ShGWmVcXT4pbwA8IJBkjJg/EtXFuL1VjP5QbbWvjakQ1ZpMEYkL4V1Gm1etzkoDuMV+VhvvL8uW59XezLH1My9RQ5vtXY7GpB3t4qbTX2AQ5abAlTAoRgOxr5mKT62m3uUpU6HBWkcqwhNGRNPQOhUBybbpxMyakJ/TyS5F7GOajsCWdhx3ErldXrtUgbArPwR16Nh0lA3jO81QJnKzbkcaVlCNd8A3to0Dx1g5cel2HDK37Ri6xYZssh1qGN+fecc7Gf4lqvp1gGMtKMyZw8t54/cJrSeVhzi+mq8aeTIaOAwpoa8C4H80HE35wog1tsS0WALlPdNZ8IyPZRfhH3iG12X0WttB5x2hHngQaYzSWzs1TvEGwrci1Y8EFE1xXG6ArAPG5Iy79tmXlOZM/R/D1K6oVRrVB6T4fWKtHFHJExlRI6HWT+Qxye96RPWxEdKEhWzOLRrBiWPSXYCtT4SCbBirP4C/htnDNcMGlT/HIETVf0R+ixjnsqeYYQn15cXvWSSDQ4LTnW9vBrDLsWVFV8hJ4outZ67Ztf/tBuGKfUFzLkTCFhWJER1bbH7Zhxn5xCplI4REr2+PKnhRaPCrz6W2TRO94pACkJG3M4eP3OyCbVfC1N1c0+MPwwJ0R7TAllli94t5jQthu8xw=="},
			output: []string{"some string x-0000000001-x"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			o, err := NewSSHObfuscator(schema.ObfuscateReplacementTypeConsistent, NewSimpleTracker())
			require.NoError(t, err)
			for i := 0; i < len(tc.input); i++ {
				assert.Equal(t, tc.output[i], o.Contents(tc.input[i]))
			}
		})
	}
}