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

package integration_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	binmock "github.com/pivotal-cf-experimental/go-binmock"

	"testing"
)

func TestDatabaseBackupAndRestore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DatabaseBackupAndRestore Suite")
}

var compiledSDKPath string
var envVars map[string]string

var fakePgDump94 *binmock.Mock
var fakePgDump96 *binmock.Mock
var fakePgDump106 *binmock.Mock
var fakePgDump11 *binmock.Mock
var fakePgRestore94 *binmock.Mock
var fakePgRestore96 *binmock.Mock
var fakePgRestore106 *binmock.Mock
var fakePgRestore11 *binmock.Mock
var fakePgClient *binmock.Mock
var fakeMysqlClient55 *binmock.Mock
var fakeMysqlDump55 *binmock.Mock
var fakeMysqlClient56 *binmock.Mock
var fakeMysqlDump56 *binmock.Mock
var fakeMysqlClient57 *binmock.Mock
var fakeMysqlDump57 *binmock.Mock
var fakeMariaDBClient *binmock.Mock
var fakeMariaDBDump *binmock.Mock

var _ = BeforeSuite(func() {
	var err error
	compiledSDKPath, err = gexec.Build(
		"github.com/cloudfoundry-incubator/database-backup-restore/cmd/database-backup-restore")
	Expect(err).NotTo(HaveOccurred())

	fakePgClient = binmock.NewBinMock(Fail)
	fakePgDump94 = binmock.NewBinMock(Fail)
	fakePgDump96 = binmock.NewBinMock(Fail)
	fakePgDump106 = binmock.NewBinMock(Fail)
	fakePgDump11 = binmock.NewBinMock(Fail)
	fakePgRestore94 = binmock.NewBinMock(Fail)
	fakePgRestore96 = binmock.NewBinMock(Fail)
	fakePgRestore106 = binmock.NewBinMock(Fail)
	fakePgRestore11 = binmock.NewBinMock(Fail)
	fakeMysqlDump55 = binmock.NewBinMock(Fail)
	fakeMysqlClient55 = binmock.NewBinMock(Fail)
	fakeMysqlDump56 = binmock.NewBinMock(Fail)
	fakeMysqlClient56 = binmock.NewBinMock(Fail)
	fakeMysqlDump57 = binmock.NewBinMock(Fail)
	fakeMysqlClient57 = binmock.NewBinMock(Fail)
	fakeMariaDBClient = binmock.NewBinMock(Fail)
	fakeMariaDBDump = binmock.NewBinMock(Fail)
})

var _ = BeforeEach(func() {
	envVars = map[string]string{
		"PG_CLIENT_PATH":        "non-existent",
		"PG_DUMP_9_4_PATH":      "non-existent",
		"PG_DUMP_9_6_PATH":      "non-existent",
		"PG_DUMP_10_6_PATH":     "non-existent",
		"PG_DUMP_11_PATH":       "non-existent",
		"PG_RESTORE_9_4_PATH":   "non-existent",
		"PG_RESTORE_9_6_PATH":   "non-existent",
		"PG_RESTORE_10_6_PATH":  "non-existent",
		"PG_RESTORE_11_PATH":    "non-existent",
		"MARIADB_CLIENT_PATH":   "non-existent",
		"MARIADB_DUMP_PATH":     "non-existent",
		"MYSQL_CLIENT_5_5_PATH": "non-existent",
		"MYSQL_DUMP_5_5_PATH":   "non-existent",
		"MYSQL_CLIENT_5_6_PATH": "non-existent",
		"MYSQL_DUMP_5_6_PATH":   "non-existent",
		"MYSQL_CLIENT_5_7_PATH": "non-existent",
		"MYSQL_DUMP_5_7_PATH":   "non-existent",
	}
})
