// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package system_tests

import (
	"time"

	"strconv"

	. "github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/system_tests/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("S3 versioned backup and restore", func() {
	var region string
	var cloneRegion string
	var unversionedRegion string
	var bucket string
	var cloneBucket string
	var unversionedBucket string
	var fileName1 string
	var fileName2 string
	var fileName3 string
	var artifactDirPath string

	var backuperInstance JobInstance
	var backuperInstanceWithClonedBucket JobInstance
	var backuperInstanceWithUnversionedBucket JobInstance
	var backuperInstanceWithCustomCaCertBlobstore JobInstance

	BeforeEach(func() {
		backuperInstance = JobInstance{
			Environment: MustHaveEnv("BOSH_ENVIRONMENT"),
			Deployment:  MustHaveEnv("BOSH_DEPLOYMENT"),
			Name:        "backuper",
			Index:       "0",
		}
		backuperInstanceWithClonedBucket = JobInstance{
			Environment: MustHaveEnv("BOSH_ENVIRONMENT"),
			Deployment:  MustHaveEnv("BOSH_DEPLOYMENT"),
			Name:        "clone-backuper",
			Index:       "0",
		}
		backuperInstanceWithUnversionedBucket = JobInstance{
			Environment: MustHaveEnv("BOSH_ENVIRONMENT"),
			Deployment:  MustHaveEnv("BOSH_DEPLOYMENT"),
			Name:        "versioned-backuper-unversioned-bucket",
			Index:       "0",
		}
		backuperInstanceWithCustomCaCertBlobstore = JobInstance{
			Environment: MustHaveEnv("BOSH_ENVIRONMENT"),
			Deployment:  MustHaveEnv("BOSH_DEPLOYMENT"),
			Name:        "unversioned-custom-ca-cert-backuper",
			Index:       "0",
		}

		region = MustHaveEnv("AWS_TEST_BUCKET_REGION")
		bucket = MustHaveEnv("AWS_TEST_BUCKET_NAME")
		cloneRegion = MustHaveEnv("AWS_TEST_CLONE_BUCKET_REGION")
		cloneBucket = MustHaveEnv("AWS_TEST_CLONE_BUCKET_NAME")
		unversionedRegion = MustHaveEnv("AWS_TEST_UNVERSIONED_BUCKET_REGION")
		unversionedBucket = MustHaveEnv("AWS_TEST_UNVERSIONED_BUCKET_NAME")

		DeleteAllVersionsFromBucket(region, bucket)
		DeleteAllVersionsFromBucket(cloneRegion, cloneBucket)
		DeleteAllVersionsFromBucket(unversionedRegion, unversionedBucket) // will it work?

		artifactDirPath = "/tmp/s3-versioned-blobstore-backup-restorer" + strconv.FormatInt(time.Now().Unix(), 10)
		backuperInstance.RunAndSucceed("mkdir -p " + artifactDirPath)
		backuperInstanceWithClonedBucket.RunAndSucceed("mkdir -p " + artifactDirPath)
		backuperInstanceWithUnversionedBucket.RunAndSucceed("mkdir -p " + artifactDirPath)
	})

	AfterEach(func() {
		DeleteAllVersionsFromBucket(region, bucket)
		DeleteAllVersionsFromBucket(cloneRegion, cloneBucket)
		DeleteAllVersionsFromBucket(unversionedRegion, unversionedBucket)
		backuperInstance.RunAndSucceed("rm -rf " + artifactDirPath)
		backuperInstanceWithClonedBucket.RunAndSucceed("rm -rf " + artifactDirPath)
		backuperInstanceWithUnversionedBucket.RunAndSucceed("rm -rf " + artifactDirPath)
	})

	It("backs up and restores in-place", func() {
		fileName1 = UploadTimestampedFileToBucket(region, bucket, "file1", "FILE1")
		fileName2 = UploadTimestampedFileToBucket(region, bucket, "file2", "FILE2")

		backuperInstance.RunAndSucceed("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/backup")

		DeleteFileFromBucket(region, bucket, fileName1)
		WriteFileInBucket(region, bucket, fileName2, "FILE2_NEW")
		fileName3 = UploadTimestampedFileToBucket(region, bucket, "file3", "FILE3")

		backuperInstance.RunAndSucceed("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/restore")

		filesList := ListFilesFromBucket(region, bucket)
		Expect(filesList).To(ConsistOf(fileName1, fileName2, fileName3))

		Expect(GetFileContentsFromBucket(region, bucket, fileName1)).To(Equal("FILE1"))
		Expect(GetFileContentsFromBucket(region, bucket, fileName2)).To(Equal("FILE2"))
		Expect(GetFileContentsFromBucket(region, bucket, fileName3)).To(Equal("FILE3"))
	})

	It("backs up and restores to a different bucket", func() {
		fileName1 = UploadTimestampedFileToBucket(region, bucket, "file1", "FILE1")
		fileName2 = UploadTimestampedFileToBucket(region, bucket, "file2", "FILE2")

		backuperInstance.RunAndSucceed("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/backup")

		backuperInstance.Download(artifactDirPath+"/blobstore.json", "/tmp/blobstore.json")
		backuperInstanceWithClonedBucket.Upload("/tmp/blobstore.json", artifactDirPath+"/blobstore.json")

		backuperInstanceWithClonedBucket.RunAndSucceed("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/restore")

		filesList := ListFilesFromBucket(cloneRegion, cloneBucket)
		Expect(filesList).To(ConsistOf(fileName1, fileName2))

		Expect(GetFileContentsFromBucket(cloneRegion, cloneBucket, fileName1)).To(Equal("FILE1"))
		Expect(GetFileContentsFromBucket(cloneRegion, cloneBucket, fileName2)).To(Equal("FILE2"))
	})

	It("fails when the bucket is not versioned", func() {
		session := backuperInstanceWithUnversionedBucket.Run("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/backup")

		Expect(session).To(gexec.Exit(1))
		Expect(session.Out).To(gbytes.Say("is not versioned"))
		Expect(backuperInstanceWithUnversionedBucket.Run("stat " + artifactDirPath + "/blobstore.json")).To(gexec.Exit(1))
	})

	It("connects with a blobstore with custom CA cert", func() {
		session := backuperInstanceWithCustomCaCertBlobstore.Run("BBR_ARTIFACT_DIRECTORY=" + artifactDirPath +
			" /var/vcap/jobs/s3-versioned-blobstore-backup-restorer/bin/bbr/backup")

		Expect(session).To(gexec.Exit(1))
		Expect(session.Out).NotTo(gbytes.Say("CERTIFICATE_VERIFY_FAILED"))
		Expect(session.Out).NotTo(gbytes.Say("no such host"))
		Expect(session.Out).To(gbytes.Say("A header you provided implies functionality that is not implemented"))
	})
})