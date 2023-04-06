// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersions)
	t.Run("Packages", testPackages)
	t.Run("Tusers", testTusers)
	t.Run("Xnodes", testXnodes)
}

func TestDelete(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsDelete)
	t.Run("Packages", testPackagesDelete)
	t.Run("Tusers", testTusersDelete)
	t.Run("Xnodes", testXnodesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsQueryDeleteAll)
	t.Run("Packages", testPackagesQueryDeleteAll)
	t.Run("Tusers", testTusersQueryDeleteAll)
	t.Run("Xnodes", testXnodesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSliceDeleteAll)
	t.Run("Packages", testPackagesSliceDeleteAll)
	t.Run("Tusers", testTusersSliceDeleteAll)
	t.Run("Xnodes", testXnodesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsExists)
	t.Run("Packages", testPackagesExists)
	t.Run("Tusers", testTusersExists)
	t.Run("Xnodes", testXnodesExists)
}

func TestFind(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsFind)
	t.Run("Packages", testPackagesFind)
	t.Run("Tusers", testTusersFind)
	t.Run("Xnodes", testXnodesFind)
}

func TestBind(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsBind)
	t.Run("Packages", testPackagesBind)
	t.Run("Tusers", testTusersBind)
	t.Run("Xnodes", testXnodesBind)
}

func TestOne(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsOne)
	t.Run("Packages", testPackagesOne)
	t.Run("Tusers", testTusersOne)
	t.Run("Xnodes", testXnodesOne)
}

func TestAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsAll)
	t.Run("Packages", testPackagesAll)
	t.Run("Tusers", testTusersAll)
	t.Run("Xnodes", testXnodesAll)
}

func TestCount(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsCount)
	t.Run("Packages", testPackagesCount)
	t.Run("Tusers", testTusersCount)
	t.Run("Xnodes", testXnodesCount)
}

func TestHooks(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsHooks)
	t.Run("Packages", testPackagesHooks)
	t.Run("Tusers", testTusersHooks)
	t.Run("Xnodes", testXnodesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsInsert)
	t.Run("GooseDBVersions", testGooseDBVersionsInsertWhitelist)
	t.Run("Packages", testPackagesInsert)
	t.Run("Packages", testPackagesInsertWhitelist)
	t.Run("Tusers", testTusersInsert)
	t.Run("Tusers", testTusersInsertWhitelist)
	t.Run("Xnodes", testXnodesInsert)
	t.Run("Xnodes", testXnodesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("TuserToPackageUsingPackage", testTuserToOnePackageUsingPackage)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("PackageToTusers", testPackageToManyTusers)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("TuserToPackageUsingTusers", testTuserToOneSetOpPackageUsingPackage)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("PackageToTusers", testPackageToManyAddOpTusers)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsReload)
	t.Run("Packages", testPackagesReload)
	t.Run("Tusers", testTusersReload)
	t.Run("Xnodes", testXnodesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsReloadAll)
	t.Run("Packages", testPackagesReloadAll)
	t.Run("Tusers", testTusersReloadAll)
	t.Run("Xnodes", testXnodesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSelect)
	t.Run("Packages", testPackagesSelect)
	t.Run("Tusers", testTusersSelect)
	t.Run("Xnodes", testXnodesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsUpdate)
	t.Run("Packages", testPackagesUpdate)
	t.Run("Tusers", testTusersUpdate)
	t.Run("Xnodes", testXnodesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("GooseDBVersions", testGooseDBVersionsSliceUpdateAll)
	t.Run("Packages", testPackagesSliceUpdateAll)
	t.Run("Tusers", testTusersSliceUpdateAll)
	t.Run("Xnodes", testXnodesSliceUpdateAll)
}
