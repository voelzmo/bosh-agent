package fakes

import (
	boshdisk "bosh/platform/disk"
	boshsys "bosh/system"
)

type FakeDiskManager struct {
	FakePartitioner *FakePartitioner
	FakeFormatter   *FakeFormatter
	FakeMounter     *FakeMounter
	FakeFinder      *FakeFinder
}

func NewFakeDiskManager(runner boshsys.CmdRunner) (manager FakeDiskManager) {
	manager.FakePartitioner = &FakePartitioner{}
	manager.FakeFormatter = &FakeFormatter{}
	manager.FakeMounter = &FakeMounter{}
	manager.FakeFinder = &FakeFinder{}
	return
}

func (m FakeDiskManager) GetPartitioner() boshdisk.Partitioner {
	return m.FakePartitioner
}

func (m FakeDiskManager) GetFormatter() boshdisk.Formatter {
	return m.FakeFormatter
}

func (m FakeDiskManager) GetMounter() boshdisk.Mounter {
	return m.FakeMounter
}

func (m FakeDiskManager) GetFinder() boshdisk.Finder {
	return m.FakeFinder
}
