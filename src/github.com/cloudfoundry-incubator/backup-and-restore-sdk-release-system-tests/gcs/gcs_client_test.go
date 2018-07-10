package gcs_test

import (
	"io/ioutil"

	"fmt"

	. "github.com/onsi/gomega"
)

type GCSClient struct{}

func (c GCSClient) WriteBlobToBucket(bucket, blobName, body string) {
	file, err := ioutil.TempFile("", "bbr-sdk-gcs-system-tests")
	Expect(err).NotTo(HaveOccurred())

	_, err = file.WriteString(body)
	Expect(err).NotTo(HaveOccurred())

	MustRunSuccessfully("gsutil", "cp", file.Name(), fmt.Sprintf("gs://%s/%s", bucket, blobName))
}

func (c GCSClient) DeleteBlobInBucket(bucket, blobName string) {
	MustRunSuccessfully("gsutil", "rm", fmt.Sprintf("gs://%s/%s", bucket, blobName))
}