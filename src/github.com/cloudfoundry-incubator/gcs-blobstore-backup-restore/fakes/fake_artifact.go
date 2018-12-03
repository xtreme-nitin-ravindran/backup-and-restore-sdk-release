// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	gcs "github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
)

type FakeBackupArtifact struct {
	WriteStub        func(backups map[string]gcs.BackupBucketDirectory) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		backups map[string]gcs.BackupBucketDirectory
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	ReadStub        func() (map[string]gcs.BackupBucketDirectory, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct{}
	readReturns     struct {
		result1 map[string]gcs.BackupBucketDirectory
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 map[string]gcs.BackupBucketDirectory
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackupArtifact) Write(backups map[string]gcs.BackupBucketDirectory) error {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		backups map[string]gcs.BackupBucketDirectory
	}{backups})
	fake.recordInvocation("Write", []interface{}{backups})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(backups)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeReturns.result1
}

func (fake *FakeBackupArtifact) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeBackupArtifact) WriteArgsForCall(i int) map[string]gcs.BackupBucketDirectory {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].backups
}

func (fake *FakeBackupArtifact) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupArtifact) WriteReturnsOnCall(i int, result1 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupArtifact) Read() (map[string]gcs.BackupBucketDirectory, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct{}{})
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *FakeBackupArtifact) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeBackupArtifact) ReadReturns(result1 map[string]gcs.BackupBucketDirectory, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 map[string]gcs.BackupBucketDirectory
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupArtifact) ReadReturnsOnCall(i int, result1 map[string]gcs.BackupBucketDirectory, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 map[string]gcs.BackupBucketDirectory
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 map[string]gcs.BackupBucketDirectory
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackupArtifact) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gcs.BackupArtifact = new(FakeBackupArtifact)