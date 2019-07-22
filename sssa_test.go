package sssa

import (
	"encoding/base64"
	"testing"
)

func TestCreateCombine(t *testing.T) {
	// Short, medium, and long tests
	strings := [][]byte{
		[]byte("N17FigASkL6p1EOgJhRaIquQLGvYV0"),
		[]byte("0y10VAfmyH7GLQY6QccCSLKJi8iFgpcSBTLyYOGbiYPqOpStAf1OYuzEBzZR"),
		[]byte("KjRHO1nHmIDidf6fKvsiXWcTqNYo2U9U8juO94EHXVqgearRISTQe0zAjkeUYYBvtcB8VWzZHYm6ktMlhOXXCfRFhbJzBUsXaHb5UDQAvs2GKy6yq0mnp8gCj98ksDlUultqygybYyHvjqR7D7EAWIKPKUVz4of8OzSjZlYg7YtCUMYhwQDryESiYabFID1PKBfKn5WSGgJBIsDw5g2HB2AqC1r3K8GboDN616Swo6qjvSFbseeETCYDB3ikS7uiK67ErIULNqVjf7IKoOaooEhQACmZ5HdWpr34tstg18rO"),
	}

	minimum := []int{4, 6, 20}
	shares := []int{5, 100, 100}

	for i := range strings {
		created, err := Create(minimum[i], shares[i], strings[i])
		if err != nil {
			t.Fatal("Fatal: creating: ", err)
		}
		combined, err := Combine(created)
		//println(len(created), string(combined), err)
		if err != nil {
			t.Fatal("Fatal: combining: ", err)
		}
		if string(combined) != string(strings[i]) {
			t.Fatal("Fatal: combining returned invalid data")
		}
	}
}

func TestLibraryCombine(t *testing.T) {
	shares := []string{
		"U1k9koNN67-og3ZY3Mmikeyj4gEFwK4HXDSglM8i_xc=yA3eU4_XYcJP0ijD63Tvqu1gklhBV32tu8cHPZXP-bk=",
		"O7c_iMBaGmQQE_uU0XRCPQwhfLBdlc6jseTzK_qN-1s=ICDGdloemG50X5GxteWWVZD3EGuxXST4UfZcek_teng=",
		"8qzYpjk7lmB7cRkOl6-7srVTKNYHuqUO2WO31Y0j1Tw=-g6srNoWkZTBqrKA2cMCA-6jxZiZv25rvbrCUWVHb5g=",
		"wGXxa_7FPFSVqdo26VKdgFxqVVWXNfwSDQyFmCh2e5w=8bTrIEs0e5FeiaXcIBaGwtGFxeyNtCG4R883tS3MsZ0=",
		"j8-Y4_7CJvL8aHxc8WMMhP_K2TEsOkxIHb7hBcwIBOo=T5-EOvAlzGMogdPawv3oK88rrygYFza3KSki2q8WEgs=",
	}

	fragments := make([][]byte, len(shares))
	for n := 0; n < len(shares); n++ {
		fragments[n], _ = base64.URLEncoding.DecodeString(shares[n])
	}

	combined, err := Combine(fragments)
	if err != nil {
		t.Fatal("Fatal: combining: ", err)
	}
	if string(combined) != "test-pass" {
		t.Fatal("Failed library cross-language check")
	}
}

func TestIsValidShare(t *testing.T) {
	shares := []string{
		"U1k9koNN67-og3ZY3Mmikeyj4gEFwK4HXDSglM8i_xc=yA3eU4_XYcJP0ijD63Tvqu1gklhBV32tu8cHPZXP-bk=",
		"O7c_iMBaGmQQE_uU0XRCPQwhfLBdlc6jseTzK_qN-1s=ICDGdloemG50X5GxteWWVZD3EGuxXST4UfZcek_teng=",
		"8qzYpjk7lmB7cRkOl6-7srVTKNYHuqUO2WO31Y0j1Tw=-g6srNoWkZTBqrKA2cMCA-6jxZiZv25rvbrCUWVHb5g=",
		"wGXxa_7FPFSVqdo26VKdgFxqVVWXNfwSDQyFmCh2e5w=8bTrIEs0e5FeiaXcIBaGwtGFxeyNtCG4R883tS3MsZ0=",
		"j8-Y4_7CJvL8aHxc8WMMhP_K2TEsOkxIHb7hBcwIBOo=T5-EOvAlzGMogdPawv3oK88rrygYFza3KSki2q8WEgs=",
		"Hello world",
	}

	fragments := make([][]byte, len(shares))
	for n := 0; n < len(shares); n++ {
		fragments[n], _ = base64.URLEncoding.DecodeString(shares[n])
	}

	results := []bool{
		true,
		true,
		true,
		true,
		true,
		false,
	}

	for i := range fragments {
		if IsValidShare(fragments[i]) != results[i] {
			t.Fatal("Checking for valid shares failed:", i)
		}
	}
}
