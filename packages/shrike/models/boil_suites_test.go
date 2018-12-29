// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Accounts", testAccounts)
	t.Run("Acls", testAcls)
	t.Run("Activities", testActivities)
	t.Run("ActivityTypes", testActivityTypes)
	t.Run("Agents", testAgents)
	t.Run("Boycotts", testBoycotts)
	t.Run("Candidates", testCandidates)
	t.Run("Causes", testCauses)
	t.Run("Companies", testCompanies)
	t.Run("Contacts", testContacts)
	t.Run("CustomerCarts", testCustomerCarts)
	t.Run("CustomerOrders", testCustomerOrders)
	t.Run("Deliveries", testDeliveries)
	t.Run("Districts", testDistricts)
	t.Run("Donations", testDonations)
	t.Run("Elections", testElections)
	t.Run("EmailAddresses", testEmailAddresses)
	t.Run("Experiments", testExperiments)
	t.Run("LandingPages", testLandingPages)
	t.Run("MailingAddresses", testMailingAddresses)
	t.Run("Notes", testNotes)
	t.Run("Offices", testOffices)
	t.Run("Payments", testPayments)
	t.Run("PhoneNumbers", testPhoneNumbers)
	t.Run("Products", testProducts)
	t.Run("ProductTypes", testProductTypes)
	t.Run("Territories", testTerritories)
	t.Run("Voters", testVoters)
}

func TestDelete(t *testing.T) {
	t.Run("Accounts", testAccountsDelete)
	t.Run("Acls", testAclsDelete)
	t.Run("Activities", testActivitiesDelete)
	t.Run("ActivityTypes", testActivityTypesDelete)
	t.Run("Agents", testAgentsDelete)
	t.Run("Boycotts", testBoycottsDelete)
	t.Run("Candidates", testCandidatesDelete)
	t.Run("Causes", testCausesDelete)
	t.Run("Companies", testCompaniesDelete)
	t.Run("Contacts", testContactsDelete)
	t.Run("CustomerCarts", testCustomerCartsDelete)
	t.Run("CustomerOrders", testCustomerOrdersDelete)
	t.Run("Deliveries", testDeliveriesDelete)
	t.Run("Districts", testDistrictsDelete)
	t.Run("Donations", testDonationsDelete)
	t.Run("Elections", testElectionsDelete)
	t.Run("EmailAddresses", testEmailAddressesDelete)
	t.Run("Experiments", testExperimentsDelete)
	t.Run("LandingPages", testLandingPagesDelete)
	t.Run("MailingAddresses", testMailingAddressesDelete)
	t.Run("Notes", testNotesDelete)
	t.Run("Offices", testOfficesDelete)
	t.Run("Payments", testPaymentsDelete)
	t.Run("PhoneNumbers", testPhoneNumbersDelete)
	t.Run("Products", testProductsDelete)
	t.Run("ProductTypes", testProductTypesDelete)
	t.Run("Territories", testTerritoriesDelete)
	t.Run("Voters", testVotersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsQueryDeleteAll)
	t.Run("Acls", testAclsQueryDeleteAll)
	t.Run("Activities", testActivitiesQueryDeleteAll)
	t.Run("ActivityTypes", testActivityTypesQueryDeleteAll)
	t.Run("Agents", testAgentsQueryDeleteAll)
	t.Run("Boycotts", testBoycottsQueryDeleteAll)
	t.Run("Candidates", testCandidatesQueryDeleteAll)
	t.Run("Causes", testCausesQueryDeleteAll)
	t.Run("Companies", testCompaniesQueryDeleteAll)
	t.Run("Contacts", testContactsQueryDeleteAll)
	t.Run("CustomerCarts", testCustomerCartsQueryDeleteAll)
	t.Run("CustomerOrders", testCustomerOrdersQueryDeleteAll)
	t.Run("Deliveries", testDeliveriesQueryDeleteAll)
	t.Run("Districts", testDistrictsQueryDeleteAll)
	t.Run("Donations", testDonationsQueryDeleteAll)
	t.Run("Elections", testElectionsQueryDeleteAll)
	t.Run("EmailAddresses", testEmailAddressesQueryDeleteAll)
	t.Run("Experiments", testExperimentsQueryDeleteAll)
	t.Run("LandingPages", testLandingPagesQueryDeleteAll)
	t.Run("MailingAddresses", testMailingAddressesQueryDeleteAll)
	t.Run("Notes", testNotesQueryDeleteAll)
	t.Run("Offices", testOfficesQueryDeleteAll)
	t.Run("Payments", testPaymentsQueryDeleteAll)
	t.Run("PhoneNumbers", testPhoneNumbersQueryDeleteAll)
	t.Run("Products", testProductsQueryDeleteAll)
	t.Run("ProductTypes", testProductTypesQueryDeleteAll)
	t.Run("Territories", testTerritoriesQueryDeleteAll)
	t.Run("Voters", testVotersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceDeleteAll)
	t.Run("Acls", testAclsSliceDeleteAll)
	t.Run("Activities", testActivitiesSliceDeleteAll)
	t.Run("ActivityTypes", testActivityTypesSliceDeleteAll)
	t.Run("Agents", testAgentsSliceDeleteAll)
	t.Run("Boycotts", testBoycottsSliceDeleteAll)
	t.Run("Candidates", testCandidatesSliceDeleteAll)
	t.Run("Causes", testCausesSliceDeleteAll)
	t.Run("Companies", testCompaniesSliceDeleteAll)
	t.Run("Contacts", testContactsSliceDeleteAll)
	t.Run("CustomerCarts", testCustomerCartsSliceDeleteAll)
	t.Run("CustomerOrders", testCustomerOrdersSliceDeleteAll)
	t.Run("Deliveries", testDeliveriesSliceDeleteAll)
	t.Run("Districts", testDistrictsSliceDeleteAll)
	t.Run("Donations", testDonationsSliceDeleteAll)
	t.Run("Elections", testElectionsSliceDeleteAll)
	t.Run("EmailAddresses", testEmailAddressesSliceDeleteAll)
	t.Run("Experiments", testExperimentsSliceDeleteAll)
	t.Run("LandingPages", testLandingPagesSliceDeleteAll)
	t.Run("MailingAddresses", testMailingAddressesSliceDeleteAll)
	t.Run("Notes", testNotesSliceDeleteAll)
	t.Run("Offices", testOfficesSliceDeleteAll)
	t.Run("Payments", testPaymentsSliceDeleteAll)
	t.Run("PhoneNumbers", testPhoneNumbersSliceDeleteAll)
	t.Run("Products", testProductsSliceDeleteAll)
	t.Run("ProductTypes", testProductTypesSliceDeleteAll)
	t.Run("Territories", testTerritoriesSliceDeleteAll)
	t.Run("Voters", testVotersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Accounts", testAccountsExists)
	t.Run("Acls", testAclsExists)
	t.Run("Activities", testActivitiesExists)
	t.Run("ActivityTypes", testActivityTypesExists)
	t.Run("Agents", testAgentsExists)
	t.Run("Boycotts", testBoycottsExists)
	t.Run("Candidates", testCandidatesExists)
	t.Run("Causes", testCausesExists)
	t.Run("Companies", testCompaniesExists)
	t.Run("Contacts", testContactsExists)
	t.Run("CustomerCarts", testCustomerCartsExists)
	t.Run("CustomerOrders", testCustomerOrdersExists)
	t.Run("Deliveries", testDeliveriesExists)
	t.Run("Districts", testDistrictsExists)
	t.Run("Donations", testDonationsExists)
	t.Run("Elections", testElectionsExists)
	t.Run("EmailAddresses", testEmailAddressesExists)
	t.Run("Experiments", testExperimentsExists)
	t.Run("LandingPages", testLandingPagesExists)
	t.Run("MailingAddresses", testMailingAddressesExists)
	t.Run("Notes", testNotesExists)
	t.Run("Offices", testOfficesExists)
	t.Run("Payments", testPaymentsExists)
	t.Run("PhoneNumbers", testPhoneNumbersExists)
	t.Run("Products", testProductsExists)
	t.Run("ProductTypes", testProductTypesExists)
	t.Run("Territories", testTerritoriesExists)
	t.Run("Voters", testVotersExists)
}

func TestFind(t *testing.T) {
	t.Run("Accounts", testAccountsFind)
	t.Run("Acls", testAclsFind)
	t.Run("Activities", testActivitiesFind)
	t.Run("ActivityTypes", testActivityTypesFind)
	t.Run("Agents", testAgentsFind)
	t.Run("Boycotts", testBoycottsFind)
	t.Run("Candidates", testCandidatesFind)
	t.Run("Causes", testCausesFind)
	t.Run("Companies", testCompaniesFind)
	t.Run("Contacts", testContactsFind)
	t.Run("CustomerCarts", testCustomerCartsFind)
	t.Run("CustomerOrders", testCustomerOrdersFind)
	t.Run("Deliveries", testDeliveriesFind)
	t.Run("Districts", testDistrictsFind)
	t.Run("Donations", testDonationsFind)
	t.Run("Elections", testElectionsFind)
	t.Run("EmailAddresses", testEmailAddressesFind)
	t.Run("Experiments", testExperimentsFind)
	t.Run("LandingPages", testLandingPagesFind)
	t.Run("MailingAddresses", testMailingAddressesFind)
	t.Run("Notes", testNotesFind)
	t.Run("Offices", testOfficesFind)
	t.Run("Payments", testPaymentsFind)
	t.Run("PhoneNumbers", testPhoneNumbersFind)
	t.Run("Products", testProductsFind)
	t.Run("ProductTypes", testProductTypesFind)
	t.Run("Territories", testTerritoriesFind)
	t.Run("Voters", testVotersFind)
}

func TestBind(t *testing.T) {
	t.Run("Accounts", testAccountsBind)
	t.Run("Acls", testAclsBind)
	t.Run("Activities", testActivitiesBind)
	t.Run("ActivityTypes", testActivityTypesBind)
	t.Run("Agents", testAgentsBind)
	t.Run("Boycotts", testBoycottsBind)
	t.Run("Candidates", testCandidatesBind)
	t.Run("Causes", testCausesBind)
	t.Run("Companies", testCompaniesBind)
	t.Run("Contacts", testContactsBind)
	t.Run("CustomerCarts", testCustomerCartsBind)
	t.Run("CustomerOrders", testCustomerOrdersBind)
	t.Run("Deliveries", testDeliveriesBind)
	t.Run("Districts", testDistrictsBind)
	t.Run("Donations", testDonationsBind)
	t.Run("Elections", testElectionsBind)
	t.Run("EmailAddresses", testEmailAddressesBind)
	t.Run("Experiments", testExperimentsBind)
	t.Run("LandingPages", testLandingPagesBind)
	t.Run("MailingAddresses", testMailingAddressesBind)
	t.Run("Notes", testNotesBind)
	t.Run("Offices", testOfficesBind)
	t.Run("Payments", testPaymentsBind)
	t.Run("PhoneNumbers", testPhoneNumbersBind)
	t.Run("Products", testProductsBind)
	t.Run("ProductTypes", testProductTypesBind)
	t.Run("Territories", testTerritoriesBind)
	t.Run("Voters", testVotersBind)
}

func TestOne(t *testing.T) {
	t.Run("Accounts", testAccountsOne)
	t.Run("Acls", testAclsOne)
	t.Run("Activities", testActivitiesOne)
	t.Run("ActivityTypes", testActivityTypesOne)
	t.Run("Agents", testAgentsOne)
	t.Run("Boycotts", testBoycottsOne)
	t.Run("Candidates", testCandidatesOne)
	t.Run("Causes", testCausesOne)
	t.Run("Companies", testCompaniesOne)
	t.Run("Contacts", testContactsOne)
	t.Run("CustomerCarts", testCustomerCartsOne)
	t.Run("CustomerOrders", testCustomerOrdersOne)
	t.Run("Deliveries", testDeliveriesOne)
	t.Run("Districts", testDistrictsOne)
	t.Run("Donations", testDonationsOne)
	t.Run("Elections", testElectionsOne)
	t.Run("EmailAddresses", testEmailAddressesOne)
	t.Run("Experiments", testExperimentsOne)
	t.Run("LandingPages", testLandingPagesOne)
	t.Run("MailingAddresses", testMailingAddressesOne)
	t.Run("Notes", testNotesOne)
	t.Run("Offices", testOfficesOne)
	t.Run("Payments", testPaymentsOne)
	t.Run("PhoneNumbers", testPhoneNumbersOne)
	t.Run("Products", testProductsOne)
	t.Run("ProductTypes", testProductTypesOne)
	t.Run("Territories", testTerritoriesOne)
	t.Run("Voters", testVotersOne)
}

func TestAll(t *testing.T) {
	t.Run("Accounts", testAccountsAll)
	t.Run("Acls", testAclsAll)
	t.Run("Activities", testActivitiesAll)
	t.Run("ActivityTypes", testActivityTypesAll)
	t.Run("Agents", testAgentsAll)
	t.Run("Boycotts", testBoycottsAll)
	t.Run("Candidates", testCandidatesAll)
	t.Run("Causes", testCausesAll)
	t.Run("Companies", testCompaniesAll)
	t.Run("Contacts", testContactsAll)
	t.Run("CustomerCarts", testCustomerCartsAll)
	t.Run("CustomerOrders", testCustomerOrdersAll)
	t.Run("Deliveries", testDeliveriesAll)
	t.Run("Districts", testDistrictsAll)
	t.Run("Donations", testDonationsAll)
	t.Run("Elections", testElectionsAll)
	t.Run("EmailAddresses", testEmailAddressesAll)
	t.Run("Experiments", testExperimentsAll)
	t.Run("LandingPages", testLandingPagesAll)
	t.Run("MailingAddresses", testMailingAddressesAll)
	t.Run("Notes", testNotesAll)
	t.Run("Offices", testOfficesAll)
	t.Run("Payments", testPaymentsAll)
	t.Run("PhoneNumbers", testPhoneNumbersAll)
	t.Run("Products", testProductsAll)
	t.Run("ProductTypes", testProductTypesAll)
	t.Run("Territories", testTerritoriesAll)
	t.Run("Voters", testVotersAll)
}

func TestCount(t *testing.T) {
	t.Run("Accounts", testAccountsCount)
	t.Run("Acls", testAclsCount)
	t.Run("Activities", testActivitiesCount)
	t.Run("ActivityTypes", testActivityTypesCount)
	t.Run("Agents", testAgentsCount)
	t.Run("Boycotts", testBoycottsCount)
	t.Run("Candidates", testCandidatesCount)
	t.Run("Causes", testCausesCount)
	t.Run("Companies", testCompaniesCount)
	t.Run("Contacts", testContactsCount)
	t.Run("CustomerCarts", testCustomerCartsCount)
	t.Run("CustomerOrders", testCustomerOrdersCount)
	t.Run("Deliveries", testDeliveriesCount)
	t.Run("Districts", testDistrictsCount)
	t.Run("Donations", testDonationsCount)
	t.Run("Elections", testElectionsCount)
	t.Run("EmailAddresses", testEmailAddressesCount)
	t.Run("Experiments", testExperimentsCount)
	t.Run("LandingPages", testLandingPagesCount)
	t.Run("MailingAddresses", testMailingAddressesCount)
	t.Run("Notes", testNotesCount)
	t.Run("Offices", testOfficesCount)
	t.Run("Payments", testPaymentsCount)
	t.Run("PhoneNumbers", testPhoneNumbersCount)
	t.Run("Products", testProductsCount)
	t.Run("ProductTypes", testProductTypesCount)
	t.Run("Territories", testTerritoriesCount)
	t.Run("Voters", testVotersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Accounts", testAccountsHooks)
	t.Run("Acls", testAclsHooks)
	t.Run("Activities", testActivitiesHooks)
	t.Run("ActivityTypes", testActivityTypesHooks)
	t.Run("Agents", testAgentsHooks)
	t.Run("Boycotts", testBoycottsHooks)
	t.Run("Candidates", testCandidatesHooks)
	t.Run("Causes", testCausesHooks)
	t.Run("Companies", testCompaniesHooks)
	t.Run("Contacts", testContactsHooks)
	t.Run("CustomerCarts", testCustomerCartsHooks)
	t.Run("CustomerOrders", testCustomerOrdersHooks)
	t.Run("Deliveries", testDeliveriesHooks)
	t.Run("Districts", testDistrictsHooks)
	t.Run("Donations", testDonationsHooks)
	t.Run("Elections", testElectionsHooks)
	t.Run("EmailAddresses", testEmailAddressesHooks)
	t.Run("Experiments", testExperimentsHooks)
	t.Run("LandingPages", testLandingPagesHooks)
	t.Run("MailingAddresses", testMailingAddressesHooks)
	t.Run("Notes", testNotesHooks)
	t.Run("Offices", testOfficesHooks)
	t.Run("Payments", testPaymentsHooks)
	t.Run("PhoneNumbers", testPhoneNumbersHooks)
	t.Run("Products", testProductsHooks)
	t.Run("ProductTypes", testProductTypesHooks)
	t.Run("Territories", testTerritoriesHooks)
	t.Run("Voters", testVotersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Accounts", testAccountsInsert)
	t.Run("Accounts", testAccountsInsertWhitelist)
	t.Run("Acls", testAclsInsert)
	t.Run("Acls", testAclsInsertWhitelist)
	t.Run("Activities", testActivitiesInsert)
	t.Run("Activities", testActivitiesInsertWhitelist)
	t.Run("ActivityTypes", testActivityTypesInsert)
	t.Run("ActivityTypes", testActivityTypesInsertWhitelist)
	t.Run("Agents", testAgentsInsert)
	t.Run("Agents", testAgentsInsertWhitelist)
	t.Run("Boycotts", testBoycottsInsert)
	t.Run("Boycotts", testBoycottsInsertWhitelist)
	t.Run("Candidates", testCandidatesInsert)
	t.Run("Candidates", testCandidatesInsertWhitelist)
	t.Run("Causes", testCausesInsert)
	t.Run("Causes", testCausesInsertWhitelist)
	t.Run("Companies", testCompaniesInsert)
	t.Run("Companies", testCompaniesInsertWhitelist)
	t.Run("Contacts", testContactsInsert)
	t.Run("Contacts", testContactsInsertWhitelist)
	t.Run("CustomerCarts", testCustomerCartsInsert)
	t.Run("CustomerCarts", testCustomerCartsInsertWhitelist)
	t.Run("CustomerOrders", testCustomerOrdersInsert)
	t.Run("CustomerOrders", testCustomerOrdersInsertWhitelist)
	t.Run("Deliveries", testDeliveriesInsert)
	t.Run("Deliveries", testDeliveriesInsertWhitelist)
	t.Run("Districts", testDistrictsInsert)
	t.Run("Districts", testDistrictsInsertWhitelist)
	t.Run("Donations", testDonationsInsert)
	t.Run("Donations", testDonationsInsertWhitelist)
	t.Run("Elections", testElectionsInsert)
	t.Run("Elections", testElectionsInsertWhitelist)
	t.Run("EmailAddresses", testEmailAddressesInsert)
	t.Run("EmailAddresses", testEmailAddressesInsertWhitelist)
	t.Run("Experiments", testExperimentsInsert)
	t.Run("Experiments", testExperimentsInsertWhitelist)
	t.Run("LandingPages", testLandingPagesInsert)
	t.Run("LandingPages", testLandingPagesInsertWhitelist)
	t.Run("MailingAddresses", testMailingAddressesInsert)
	t.Run("MailingAddresses", testMailingAddressesInsertWhitelist)
	t.Run("Notes", testNotesInsert)
	t.Run("Notes", testNotesInsertWhitelist)
	t.Run("Offices", testOfficesInsert)
	t.Run("Offices", testOfficesInsertWhitelist)
	t.Run("Payments", testPaymentsInsert)
	t.Run("Payments", testPaymentsInsertWhitelist)
	t.Run("PhoneNumbers", testPhoneNumbersInsert)
	t.Run("PhoneNumbers", testPhoneNumbersInsertWhitelist)
	t.Run("Products", testProductsInsert)
	t.Run("Products", testProductsInsertWhitelist)
	t.Run("ProductTypes", testProductTypesInsert)
	t.Run("ProductTypes", testProductTypesInsertWhitelist)
	t.Run("Territories", testTerritoriesInsert)
	t.Run("Territories", testTerritoriesInsertWhitelist)
	t.Run("Voters", testVotersInsert)
	t.Run("Voters", testVotersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

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
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Accounts", testAccountsReload)
	t.Run("Acls", testAclsReload)
	t.Run("Activities", testActivitiesReload)
	t.Run("ActivityTypes", testActivityTypesReload)
	t.Run("Agents", testAgentsReload)
	t.Run("Boycotts", testBoycottsReload)
	t.Run("Candidates", testCandidatesReload)
	t.Run("Causes", testCausesReload)
	t.Run("Companies", testCompaniesReload)
	t.Run("Contacts", testContactsReload)
	t.Run("CustomerCarts", testCustomerCartsReload)
	t.Run("CustomerOrders", testCustomerOrdersReload)
	t.Run("Deliveries", testDeliveriesReload)
	t.Run("Districts", testDistrictsReload)
	t.Run("Donations", testDonationsReload)
	t.Run("Elections", testElectionsReload)
	t.Run("EmailAddresses", testEmailAddressesReload)
	t.Run("Experiments", testExperimentsReload)
	t.Run("LandingPages", testLandingPagesReload)
	t.Run("MailingAddresses", testMailingAddressesReload)
	t.Run("Notes", testNotesReload)
	t.Run("Offices", testOfficesReload)
	t.Run("Payments", testPaymentsReload)
	t.Run("PhoneNumbers", testPhoneNumbersReload)
	t.Run("Products", testProductsReload)
	t.Run("ProductTypes", testProductTypesReload)
	t.Run("Territories", testTerritoriesReload)
	t.Run("Voters", testVotersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Accounts", testAccountsReloadAll)
	t.Run("Acls", testAclsReloadAll)
	t.Run("Activities", testActivitiesReloadAll)
	t.Run("ActivityTypes", testActivityTypesReloadAll)
	t.Run("Agents", testAgentsReloadAll)
	t.Run("Boycotts", testBoycottsReloadAll)
	t.Run("Candidates", testCandidatesReloadAll)
	t.Run("Causes", testCausesReloadAll)
	t.Run("Companies", testCompaniesReloadAll)
	t.Run("Contacts", testContactsReloadAll)
	t.Run("CustomerCarts", testCustomerCartsReloadAll)
	t.Run("CustomerOrders", testCustomerOrdersReloadAll)
	t.Run("Deliveries", testDeliveriesReloadAll)
	t.Run("Districts", testDistrictsReloadAll)
	t.Run("Donations", testDonationsReloadAll)
	t.Run("Elections", testElectionsReloadAll)
	t.Run("EmailAddresses", testEmailAddressesReloadAll)
	t.Run("Experiments", testExperimentsReloadAll)
	t.Run("LandingPages", testLandingPagesReloadAll)
	t.Run("MailingAddresses", testMailingAddressesReloadAll)
	t.Run("Notes", testNotesReloadAll)
	t.Run("Offices", testOfficesReloadAll)
	t.Run("Payments", testPaymentsReloadAll)
	t.Run("PhoneNumbers", testPhoneNumbersReloadAll)
	t.Run("Products", testProductsReloadAll)
	t.Run("ProductTypes", testProductTypesReloadAll)
	t.Run("Territories", testTerritoriesReloadAll)
	t.Run("Voters", testVotersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Accounts", testAccountsSelect)
	t.Run("Acls", testAclsSelect)
	t.Run("Activities", testActivitiesSelect)
	t.Run("ActivityTypes", testActivityTypesSelect)
	t.Run("Agents", testAgentsSelect)
	t.Run("Boycotts", testBoycottsSelect)
	t.Run("Candidates", testCandidatesSelect)
	t.Run("Causes", testCausesSelect)
	t.Run("Companies", testCompaniesSelect)
	t.Run("Contacts", testContactsSelect)
	t.Run("CustomerCarts", testCustomerCartsSelect)
	t.Run("CustomerOrders", testCustomerOrdersSelect)
	t.Run("Deliveries", testDeliveriesSelect)
	t.Run("Districts", testDistrictsSelect)
	t.Run("Donations", testDonationsSelect)
	t.Run("Elections", testElectionsSelect)
	t.Run("EmailAddresses", testEmailAddressesSelect)
	t.Run("Experiments", testExperimentsSelect)
	t.Run("LandingPages", testLandingPagesSelect)
	t.Run("MailingAddresses", testMailingAddressesSelect)
	t.Run("Notes", testNotesSelect)
	t.Run("Offices", testOfficesSelect)
	t.Run("Payments", testPaymentsSelect)
	t.Run("PhoneNumbers", testPhoneNumbersSelect)
	t.Run("Products", testProductsSelect)
	t.Run("ProductTypes", testProductTypesSelect)
	t.Run("Territories", testTerritoriesSelect)
	t.Run("Voters", testVotersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Accounts", testAccountsUpdate)
	t.Run("Acls", testAclsUpdate)
	t.Run("Activities", testActivitiesUpdate)
	t.Run("ActivityTypes", testActivityTypesUpdate)
	t.Run("Agents", testAgentsUpdate)
	t.Run("Boycotts", testBoycottsUpdate)
	t.Run("Candidates", testCandidatesUpdate)
	t.Run("Causes", testCausesUpdate)
	t.Run("Companies", testCompaniesUpdate)
	t.Run("Contacts", testContactsUpdate)
	t.Run("CustomerCarts", testCustomerCartsUpdate)
	t.Run("CustomerOrders", testCustomerOrdersUpdate)
	t.Run("Deliveries", testDeliveriesUpdate)
	t.Run("Districts", testDistrictsUpdate)
	t.Run("Donations", testDonationsUpdate)
	t.Run("Elections", testElectionsUpdate)
	t.Run("EmailAddresses", testEmailAddressesUpdate)
	t.Run("Experiments", testExperimentsUpdate)
	t.Run("LandingPages", testLandingPagesUpdate)
	t.Run("MailingAddresses", testMailingAddressesUpdate)
	t.Run("Notes", testNotesUpdate)
	t.Run("Offices", testOfficesUpdate)
	t.Run("Payments", testPaymentsUpdate)
	t.Run("PhoneNumbers", testPhoneNumbersUpdate)
	t.Run("Products", testProductsUpdate)
	t.Run("ProductTypes", testProductTypesUpdate)
	t.Run("Territories", testTerritoriesUpdate)
	t.Run("Voters", testVotersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Accounts", testAccountsSliceUpdateAll)
	t.Run("Acls", testAclsSliceUpdateAll)
	t.Run("Activities", testActivitiesSliceUpdateAll)
	t.Run("ActivityTypes", testActivityTypesSliceUpdateAll)
	t.Run("Agents", testAgentsSliceUpdateAll)
	t.Run("Boycotts", testBoycottsSliceUpdateAll)
	t.Run("Candidates", testCandidatesSliceUpdateAll)
	t.Run("Causes", testCausesSliceUpdateAll)
	t.Run("Companies", testCompaniesSliceUpdateAll)
	t.Run("Contacts", testContactsSliceUpdateAll)
	t.Run("CustomerCarts", testCustomerCartsSliceUpdateAll)
	t.Run("CustomerOrders", testCustomerOrdersSliceUpdateAll)
	t.Run("Deliveries", testDeliveriesSliceUpdateAll)
	t.Run("Districts", testDistrictsSliceUpdateAll)
	t.Run("Donations", testDonationsSliceUpdateAll)
	t.Run("Elections", testElectionsSliceUpdateAll)
	t.Run("EmailAddresses", testEmailAddressesSliceUpdateAll)
	t.Run("Experiments", testExperimentsSliceUpdateAll)
	t.Run("LandingPages", testLandingPagesSliceUpdateAll)
	t.Run("MailingAddresses", testMailingAddressesSliceUpdateAll)
	t.Run("Notes", testNotesSliceUpdateAll)
	t.Run("Offices", testOfficesSliceUpdateAll)
	t.Run("Payments", testPaymentsSliceUpdateAll)
	t.Run("PhoneNumbers", testPhoneNumbersSliceUpdateAll)
	t.Run("Products", testProductsSliceUpdateAll)
	t.Run("ProductTypes", testProductTypesSliceUpdateAll)
	t.Run("Territories", testTerritoriesSliceUpdateAll)
	t.Run("Voters", testVotersSliceUpdateAll)
}
