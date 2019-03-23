const { gql } = require('apollo-server');
module.exports = `
type Time {
    seconds: Int!
    nanos: Int!
  }
input TimeInput {
    seconds: Int!
    nanos: Int!
  }

input LayoutTypeInput {
  secret: Int
  Title: String!
}
​
input ACLInput {
  secret: Int
}
​
input MailingAddressInput {
  secret: Int
  StreetAddress: String!
  City: String!
  State: String!
  ZipCode: String!
}
​
input PhoneNumberInput {
  secret: Int
  PhoneNumber: String!
}
​
input EmailAddressInput {
  secret: Int
  Address: String!
}
​
input LayoutRowInput {
  secret: Int
  Layout: ID!
  Container: Boolean!
}
​
input ComponentImplementationInput {
  secret: Int
  Title: String!
  Path: String!
}
​
input LayoutColumnInput {
  secret: Int
  LayoutRow: ID!
  Width: Int!
}
​
input ComponentInput {
  secret: Int
  ComponentType: ID!
  ComponentImplementation: ID!
  LayoutColumn: ID
}
​
input FieldInput {
  secret: Int
  FieldType: ID!
  StringValue: String
  IntValue: Int
  FloatValue: Float
  BooleanValue: Boolean
  DateTimeValue: TimeInput
  Component: ID
}
​
input LayoutInput {
  secret: Int
  LayoutType: ID
}
​
input LandingPageInput {
  secret: Int
  Title: String!
  Cause: ID!
  Layout: ID
}
​
input ExperimentInput {
  secret: Int
  Title: String!
  LandingPage: ID
}
​
input IssueInput {
  secret: Int
  Title: String!
  Election: ID!
}
​
input CandidateInput {
  secret: Int
  Election: ID!
}
​
input DistrictTypeInput {
  secret: Int
  Title: String!
}
​
input DistrictInput {
  secret: Int
  Geom: Float
  Title: String!
  DistrictType: ID!
}
​
input OfficeInput {
  secret: Int
  Title: String!
  Election: ID
}
​
input PollItemInput {
  secret: Int
  Title: String!
  Poll: ID!
}
​
input VolunteerOpportunityTypeInput {
  secret: Int
  Title: String!
}
​
input LiveEventTypeInput {
  secret: Int
  Title: String!
}
​
input CompanyInput {
  secret: Int
  Title: String!
}
​
input ProductTypeInput {
  secret: Int
  Title: String!
}
​
input CustomerCartInput {
  secret: Int
}
​
input PaymentInput {
  secret: Int
  CustomerOrder: ID!
}
​
input DeliveryInput {
  secret: Int
}
​
input BoycottInput {
  secret: Int
  Title: String!
}
​
input BoycottMembershipInput {
  secret: Int
  Cause: ID!
  Boycott: ID!
}
​
input ElectionInput {
  secret: Int
  Title: String!
}
​
input ElectionMembershipInput {
  secret: Int
  Cause: ID!
  Election: ID!
}
​
input PetitionMembershipInput {
  secret: Int
  Cause: ID!
  Petition: ID!
}
​
input PollMembershipInput {
  secret: Int
  Cause: ID!
  Petition: ID!
}
​
input VolunteerOpportunityMembershipInput {
  secret: Int
  Cause: ID!
  VolunteerOpportunity: ID!
}
​
input LiveEventMembershipInput {
  secret: Int
  Cause: ID!
  LiveEvent: ID!
}
​
input ProductInput {
  secret: Int
  Title: String!
  ProductType: ID!
}
​
input ProductMembershipInput {
  secret: Int
  Cause: ID!
  Product: ID!
}
​
input DonationCampaignInput {
  secret: Int
  Title: String!
}
​
input DonationCampaignMembershipInput {
  secret: Int
  Cause: ID!
  DonationCampaign: ID!
}
​
input PetitionInput {
  secret: Int
  Title: String!
}
​
input PetitionSignerInput {
  secret: Int
  Petition: ID!
  Contact: ID!
  Cause: ID!
}
​
input PollInput {
  secret: Int
  Title: String!
}
​
input PollRespondantInput {
  secret: Int
  Poll: ID!
  Contact: ID!
  Cause: ID!
}
​
input PurchaserInput {
  secret: Int
  CustomerOrder: ID!
  Contact: ID!
  Cause: ID!
}
​
input CustomerOrderInput {
  secret: Int
  CustomerCart: ID!
}
​
input DonorInput {
  secret: Int
  CustomerOrder: ID!
  Contact: ID!
  Cause: ID!
}
​
input LiveEventInput {
  secret: Int
  Title: String!
  LiveEventType: ID!
}
​
input EventAttendeeInput {
  secret: Int
  LiveEvent: ID!
  Contact: ID!
  Cause: ID!
}
​
input VoterInput {
  secret: Int
  Contact: ID!
  Cause: ID!
}
​
input VolunteerOpportunityInput {
  secret: Int
  Title: String!
  VolunteerOpportunityType: ID
}
​
input VolunteerInput {
  secret: Int
  VolunteerOpportunity: ID!
  Contact: ID!
  Cause: ID!
}
​
input FollowerInput {
  secret: Int
  Contact: ID!
  Cause: ID!
}
​
input TerritoryInput {
  secret: Int
  Title: String!
}
​
input ActivityTypeInput {
  secret: Int
  Title: String!
}
​
input ActivityInput {
  secret: Int
  Title: String!
  ActivityType: ID!
  Contact: ID!
  Cause: ID!
}
​
input NoteInput {
  secret: Int
  Contact: ID!
  Cause: ID!
  Body: String
}
​
input AccountInput {
  secret: Int
  Username: String!
}
​
input OwnerMembershipInput {
  secret: Int
  Cause: ID!
  Account: ID!
}
​
input ContactInput {
  secret: Int
  FirstName: String
  MiddleName: String
  LastName: String
  Email: String
  PhoneNumber: String
}
​
input ContactMembershipInput {
  secret: Int
  Cause: ID!
  Contact: ID!
}
​
input AgentInput {
  secret: Int
  Account: ID!
}
​
input AgentMembershipInput {
  secret: Int
  Cause: ID!
  Agent: ID!
}
​
input HomePageInput {
  secret: Int
  Title: String!
  Layout: ID
}
​
input PhotoInput {
  secret: Int
  URI: String!
  Width: Int!
  Height: Int!
}
​
input CauseInput {
  secret: Int
  Title: String!
  Slug: String!
  Summary: String
  HomePage: ID
  Photo: ID
}
​
input ComponentTypeInput {
  secret: Int
  Title: String!
}
​
input FieldTypeInput {
  secret: Int
  Title: String!
  DataType: String!
  PropName: String!
  StringValueDefault: String
  IntValueDefault: Int
  FloatValueDefault: Float
  BooleanValueDefault: Boolean
  DateTimeValueDefault: TimeInput
  ComponentType: ID
}
​


type LayoutType {
  ID: ID!
  Layouts: [Layout]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type ACL {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type MailingAddress {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  StreetAddress: String!
  City: String!
  State: String!
  ZipCode: String!
}
​
type PhoneNumber {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  PhoneNumber: String!
}
​
type EmailAddress {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Address: String!
}
​
type LayoutRow {
  ID: ID!
  LayoutColumns: [LayoutColumn]
  CreatedAt: Time!
  UpdatedAt: Time!
  Layout: Layout
  Container: Boolean!
}
​
type ComponentImplementation {
  ID: ID!
  Components: [Component]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Path: String!
}
​
type LayoutColumn {
  ID: ID!
  Components: [Component]
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutRow: LayoutRow
  Width: Int!
}
​
type Component {
  ID: ID!
  Fields: [Field]
  CreatedAt: Time!
  UpdatedAt: Time!
  ComponentType: ComponentType
  ComponentImplementation: ComponentImplementation
  LayoutColumn: LayoutColumn
}
​
type Field {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  FieldType: FieldType
  StringValue: String
  IntValue: Int
  FloatValue: Float
  BooleanValue: Boolean
  DateTimeValue: Time
  Component: Component
}
​
type Layout {
  ID: ID!
  LayoutRows: [LayoutRow]
  LandingPages: [LandingPage]
  HomePages: [HomePage]
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutType: LayoutType
}
​
type LandingPage {
  ID: ID!
  Experiments: [Experiment]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Cause: Cause
  Layout: Layout
}
​
type Experiment {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LandingPage: LandingPage
}
​
type Issue {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: Election
}
​
type Candidate {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Election: Election
}
​
type DistrictType {
  ID: ID!
  Districts: [District]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type District {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Geom: Float
  Title: String!
  DistrictType: DistrictType
}
​
type Office {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: Election
}
​
type PollItem {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Poll: Poll
}
​
type VolunteerOpportunityType {
  ID: ID!
  VolunteerOpportunitys: [VolunteerOpportunity]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LiveEventType {
  ID: ID!
  LiveEvents: [LiveEvent]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type Company {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type ProductType {
  ID: ID!
  Products: [Product]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type CustomerCart {
  ID: ID!
  CustomerOrders: [CustomerOrder]
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type Payment {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: CustomerOrder
}
​
type Delivery {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type Boycott {
  ID: ID!
  BoycottMemberships: [BoycottMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type BoycottMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Boycott: Boycott
}
​
type Election {
  ID: ID!
  Issues: [Issue]
  Candidates: [Candidate]
  Offices: [Office]
  ElectionMemberships: [ElectionMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type ElectionMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Election: Election
}
​
type PetitionMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Petition: Petition
}
​
type PollMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Petition: Petition
}
​
type VolunteerOpportunityMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  VolunteerOpportunity: VolunteerOpportunity
}
​
type LiveEventMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  LiveEvent: LiveEvent
}
​
type Product {
  ID: ID!
  ProductMemberships: [ProductMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ProductType: ProductType
}
​
type ProductMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Product: Product
}
​
type DonationCampaign {
  ID: ID!
  DonationCampaignMemberships: [DonationCampaignMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type DonationCampaignMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  DonationCampaign: DonationCampaign
}
​
type Petition {
  ID: ID!
  PetitionMemberships: [PetitionMembership]
  PollMemberships: [PollMembership]
  PetitionSigners: [PetitionSigner]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type PetitionSigner {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Petition: Petition
  Contact: Contact
  Cause: Cause
}
​
type Poll {
  ID: ID!
  PollItems: [PollItem]
  PollRespondants: [PollRespondant]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type PollRespondant {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Poll: Poll
  Contact: Contact
  Cause: Cause
}
​
type Purchaser {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: CustomerOrder
  Contact: Contact
  Cause: Cause
}
​
type CustomerOrder {
  ID: ID!
  Payments: [Payment]
  Purchasers: [Purchaser]
  Donors: [Donor]
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerCart: CustomerCart
}
​
type Donor {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: CustomerOrder
  Contact: Contact
  Cause: Cause
}
​
type LiveEvent {
  ID: ID!
  LiveEventMemberships: [LiveEventMembership]
  EventAttendees: [EventAttendee]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LiveEventType: LiveEventType
}
​
type EventAttendee {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LiveEvent: LiveEvent
  Contact: Contact
  Cause: Cause
}
​
type Voter {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Contact
  Cause: Cause
}
​
type VolunteerOpportunity {
  ID: ID!
  VolunteerOpportunityMemberships: [VolunteerOpportunityMembership]
  Volunteers: [Volunteer]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  VolunteerOpportunityType: VolunteerOpportunityType
}
​
type Volunteer {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  VolunteerOpportunity: VolunteerOpportunity
  Contact: Contact
  Cause: Cause
}
​
type Follower {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Contact
  Cause: Cause
}
​
type Territory {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type ActivityType {
  ID: ID!
  Activitys: [Activity]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type Activity {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ActivityType: ActivityType
  Contact: Contact
  Cause: Cause
}
​
type Note {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Contact
  Cause: Cause
  Body: String
}
​
type Account {
  ID: ID!
  OwnerMemberships: [OwnerMembership]
  Agents: [Agent]
  CreatedAt: Time!
  UpdatedAt: Time!
  Username: String!
}
​
type OwnerMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Account: Account
}
​
type Contact {
  ID: ID!
  PetitionSigners: [PetitionSigner]
  PollRespondants: [PollRespondant]
  Purchasers: [Purchaser]
  Donors: [Donor]
  EventAttendees: [EventAttendee]
  Voters: [Voter]
  Volunteers: [Volunteer]
  Followers: [Follower]
  Activitys: [Activity]
  Notes: [Note]
  ContactMemberships: [ContactMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  FirstName: String
  MiddleName: String
  LastName: String
  Email: String
  PhoneNumber: String
}
​
type ContactMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Contact: Contact
}
​
type Agent {
  ID: ID!
  AgentMemberships: [AgentMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Account: Account
}
​
type AgentMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Cause
  Agent: Agent
}
​
type HomePage {
  ID: ID!
  Causes: [Cause]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Layout: Layout
}
​
type Photo {
  ID: ID!
  Causes: [Cause]
  CreatedAt: Time!
  UpdatedAt: Time!
  URI: String!
  Width: Int!
  Height: Int!
}
​
type Cause {
  ID: ID!
  LandingPages: [LandingPage]
  BoycottMemberships: [BoycottMembership]
  ElectionMemberships: [ElectionMembership]
  PetitionMemberships: [PetitionMembership]
  PollMemberships: [PollMembership]
  VolunteerOpportunityMemberships: [VolunteerOpportunityMembership]
  LiveEventMemberships: [LiveEventMembership]
  ProductMemberships: [ProductMembership]
  DonationCampaignMemberships: [DonationCampaignMembership]
  PetitionSigners: [PetitionSigner]
  PollRespondants: [PollRespondant]
  Purchasers: [Purchaser]
  Donors: [Donor]
  EventAttendees: [EventAttendee]
  Voters: [Voter]
  Volunteers: [Volunteer]
  Followers: [Follower]
  Activitys: [Activity]
  Notes: [Note]
  OwnerMemberships: [OwnerMembership]
  ContactMemberships: [ContactMembership]
  AgentMemberships: [AgentMembership]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Slug: String!
  Summary: String
  HomePage: HomePage
  Photo: Photo
}
​
type ComponentType {
  ID: ID!
  Components: [Component]
  FieldTypes: [FieldType]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type FieldType {
  ID: ID!
  Fields: [Field]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  DataType: String!
  PropName: String!
  StringValueDefault: String
  IntValueDefault: Int
  FloatValueDefault: Float
  BooleanValueDefault: Boolean
  DateTimeValueDefault: Time
  ComponentType: ComponentType
}
​

  type Query {
    
    getLayoutType(ID: ID!): LayoutType
    listLayoutType(limit: Int): [LayoutType]
    getACL(ID: ID!): ACL
    listACL(limit: Int): [ACL]
    getMailingAddress(ID: ID!): MailingAddress
    listMailingAddress(limit: Int): [MailingAddress]
    getPhoneNumber(ID: ID!): PhoneNumber
    listPhoneNumber(limit: Int): [PhoneNumber]
    getEmailAddress(ID: ID!): EmailAddress
    listEmailAddress(limit: Int): [EmailAddress]
    getLayoutRow(ID: ID!): LayoutRow
    listLayoutRow(limit: Int): [LayoutRow]
    getComponentImplementation(ID: ID!): ComponentImplementation
    listComponentImplementation(limit: Int): [ComponentImplementation]
    getLayoutColumn(ID: ID!): LayoutColumn
    listLayoutColumn(limit: Int): [LayoutColumn]
    getComponent(ID: ID!): Component
    listComponent(limit: Int): [Component]
    getField(ID: ID!): Field
    listField(limit: Int): [Field]
    getLayout(ID: ID!): Layout
    listLayout(limit: Int): [Layout]
    getLandingPage(ID: ID!): LandingPage
    listLandingPage(limit: Int): [LandingPage]
    getExperiment(ID: ID!): Experiment
    listExperiment(limit: Int): [Experiment]
    getIssue(ID: ID!): Issue
    listIssue(limit: Int): [Issue]
    getCandidate(ID: ID!): Candidate
    listCandidate(limit: Int): [Candidate]
    getDistrictType(ID: ID!): DistrictType
    listDistrictType(limit: Int): [DistrictType]
    getDistrict(ID: ID!): District
    listDistrict(limit: Int): [District]
    getOffice(ID: ID!): Office
    listOffice(limit: Int): [Office]
    getPollItem(ID: ID!): PollItem
    listPollItem(limit: Int): [PollItem]
    getVolunteerOpportunityType(ID: ID!): VolunteerOpportunityType
    listVolunteerOpportunityType(limit: Int): [VolunteerOpportunityType]
    getLiveEventType(ID: ID!): LiveEventType
    listLiveEventType(limit: Int): [LiveEventType]
    getCompany(ID: ID!): Company
    listCompany(limit: Int): [Company]
    getProductType(ID: ID!): ProductType
    listProductType(limit: Int): [ProductType]
    getCustomerCart(ID: ID!): CustomerCart
    listCustomerCart(limit: Int): [CustomerCart]
    getPayment(ID: ID!): Payment
    listPayment(limit: Int): [Payment]
    getDelivery(ID: ID!): Delivery
    listDelivery(limit: Int): [Delivery]
    getBoycott(ID: ID!): Boycott
    listBoycott(limit: Int): [Boycott]
    getBoycottMembership(ID: ID!): BoycottMembership
    listBoycottMembership(limit: Int): [BoycottMembership]
    getElection(ID: ID!): Election
    listElection(limit: Int): [Election]
    getElectionMembership(ID: ID!): ElectionMembership
    listElectionMembership(limit: Int): [ElectionMembership]
    getPetitionMembership(ID: ID!): PetitionMembership
    listPetitionMembership(limit: Int): [PetitionMembership]
    getPollMembership(ID: ID!): PollMembership
    listPollMembership(limit: Int): [PollMembership]
    getVolunteerOpportunityMembership(ID: ID!): VolunteerOpportunityMembership
    listVolunteerOpportunityMembership(limit: Int): [VolunteerOpportunityMembership]
    getLiveEventMembership(ID: ID!): LiveEventMembership
    listLiveEventMembership(limit: Int): [LiveEventMembership]
    getProduct(ID: ID!): Product
    listProduct(limit: Int): [Product]
    getProductMembership(ID: ID!): ProductMembership
    listProductMembership(limit: Int): [ProductMembership]
    getDonationCampaign(ID: ID!): DonationCampaign
    listDonationCampaign(limit: Int): [DonationCampaign]
    getDonationCampaignMembership(ID: ID!): DonationCampaignMembership
    listDonationCampaignMembership(limit: Int): [DonationCampaignMembership]
    getPetition(ID: ID!): Petition
    listPetition(limit: Int): [Petition]
    getPetitionSigner(ID: ID!): PetitionSigner
    listPetitionSigner(limit: Int): [PetitionSigner]
    getPoll(ID: ID!): Poll
    listPoll(limit: Int): [Poll]
    getPollRespondant(ID: ID!): PollRespondant
    listPollRespondant(limit: Int): [PollRespondant]
    getPurchaser(ID: ID!): Purchaser
    listPurchaser(limit: Int): [Purchaser]
    getCustomerOrder(ID: ID!): CustomerOrder
    listCustomerOrder(limit: Int): [CustomerOrder]
    getDonor(ID: ID!): Donor
    listDonor(limit: Int): [Donor]
    getLiveEvent(ID: ID!): LiveEvent
    listLiveEvent(limit: Int): [LiveEvent]
    getEventAttendee(ID: ID!): EventAttendee
    listEventAttendee(limit: Int): [EventAttendee]
    getVoter(ID: ID!): Voter
    listVoter(limit: Int): [Voter]
    getVolunteerOpportunity(ID: ID!): VolunteerOpportunity
    listVolunteerOpportunity(limit: Int): [VolunteerOpportunity]
    getVolunteer(ID: ID!): Volunteer
    listVolunteer(limit: Int): [Volunteer]
    getFollower(ID: ID!): Follower
    listFollower(limit: Int): [Follower]
    getTerritory(ID: ID!): Territory
    listTerritory(limit: Int): [Territory]
    getActivityType(ID: ID!): ActivityType
    listActivityType(limit: Int): [ActivityType]
    getActivity(ID: ID!): Activity
    listActivity(limit: Int): [Activity]
    getNote(ID: ID!): Note
    listNote(limit: Int): [Note]
    getAccount(ID: ID!): Account
    listAccount(limit: Int): [Account]
    getOwnerMembership(ID: ID!): OwnerMembership
    listOwnerMembership(limit: Int): [OwnerMembership]
    getContact(ID: ID!): Contact
    listContact(limit: Int): [Contact]
    getContactMembership(ID: ID!): ContactMembership
    listContactMembership(limit: Int): [ContactMembership]
    getAgent(ID: ID!): Agent
    listAgent(limit: Int): [Agent]
    getAgentMembership(ID: ID!): AgentMembership
    listAgentMembership(limit: Int): [AgentMembership]
    getHomePage(ID: ID!): HomePage
    listHomePage(limit: Int): [HomePage]
    getPhoto(ID: ID!): Photo
    listPhoto(limit: Int): [Photo]
    getCause(ID: ID!): Cause
    listCause(limit: Int): [Cause]
    getComponentType(ID: ID!): ComponentType
    listComponentType(limit: Int): [ComponentType]
    getFieldType(ID: ID!): FieldType
    listFieldType(limit: Int): [FieldType]
  }

  type Mutation {
    
    createLayoutType(layoutType: LayoutTypeInput, buildStatic: Boolean): LayoutType
    updateLayoutType(ID: ID!, layoutType: LayoutTypeInput, buildStatic: Boolean): Int
    deleteLayoutType(ID: ID!, buildStatic: Boolean): Int
    createACL(acl: ACLInput, buildStatic: Boolean): ACL
    updateACL(ID: ID!, acl: ACLInput, buildStatic: Boolean): Int
    deleteACL(ID: ID!, buildStatic: Boolean): Int
    createMailingAddress(mailingAddress: MailingAddressInput, buildStatic: Boolean): MailingAddress
    updateMailingAddress(ID: ID!, mailingAddress: MailingAddressInput, buildStatic: Boolean): Int
    deleteMailingAddress(ID: ID!, buildStatic: Boolean): Int
    createPhoneNumber(phoneNumber: PhoneNumberInput, buildStatic: Boolean): PhoneNumber
    updatePhoneNumber(ID: ID!, phoneNumber: PhoneNumberInput, buildStatic: Boolean): Int
    deletePhoneNumber(ID: ID!, buildStatic: Boolean): Int
    createEmailAddress(emailAddress: EmailAddressInput, buildStatic: Boolean): EmailAddress
    updateEmailAddress(ID: ID!, emailAddress: EmailAddressInput, buildStatic: Boolean): Int
    deleteEmailAddress(ID: ID!, buildStatic: Boolean): Int
    createLayoutRow(layoutRow: LayoutRowInput, buildStatic: Boolean): LayoutRow
    updateLayoutRow(ID: ID!, layoutRow: LayoutRowInput, buildStatic: Boolean): Int
    deleteLayoutRow(ID: ID!, buildStatic: Boolean): Int
    createComponentImplementation(componentImplementation: ComponentImplementationInput, buildStatic: Boolean): ComponentImplementation
    updateComponentImplementation(ID: ID!, componentImplementation: ComponentImplementationInput, buildStatic: Boolean): Int
    deleteComponentImplementation(ID: ID!, buildStatic: Boolean): Int
    createLayoutColumn(layoutColumn: LayoutColumnInput, buildStatic: Boolean): LayoutColumn
    updateLayoutColumn(ID: ID!, layoutColumn: LayoutColumnInput, buildStatic: Boolean): Int
    deleteLayoutColumn(ID: ID!, buildStatic: Boolean): Int
    createComponent(component: ComponentInput, buildStatic: Boolean): Component
    updateComponent(ID: ID!, component: ComponentInput, buildStatic: Boolean): Int
    deleteComponent(ID: ID!, buildStatic: Boolean): Int
    createField(field: FieldInput, buildStatic: Boolean): Field
    updateField(ID: ID!, field: FieldInput, buildStatic: Boolean): Int
    deleteField(ID: ID!, buildStatic: Boolean): Int
    createLayout(layout: LayoutInput, buildStatic: Boolean): Layout
    updateLayout(ID: ID!, layout: LayoutInput, buildStatic: Boolean): Int
    deleteLayout(ID: ID!, buildStatic: Boolean): Int
    createLandingPage(landingPage: LandingPageInput, buildStatic: Boolean): LandingPage
    updateLandingPage(ID: ID!, landingPage: LandingPageInput, buildStatic: Boolean): Int
    deleteLandingPage(ID: ID!, buildStatic: Boolean): Int
    createExperiment(experiment: ExperimentInput, buildStatic: Boolean): Experiment
    updateExperiment(ID: ID!, experiment: ExperimentInput, buildStatic: Boolean): Int
    deleteExperiment(ID: ID!, buildStatic: Boolean): Int
    createIssue(issue: IssueInput, buildStatic: Boolean): Issue
    updateIssue(ID: ID!, issue: IssueInput, buildStatic: Boolean): Int
    deleteIssue(ID: ID!, buildStatic: Boolean): Int
    createCandidate(candidate: CandidateInput, buildStatic: Boolean): Candidate
    updateCandidate(ID: ID!, candidate: CandidateInput, buildStatic: Boolean): Int
    deleteCandidate(ID: ID!, buildStatic: Boolean): Int
    createDistrictType(districtType: DistrictTypeInput, buildStatic: Boolean): DistrictType
    updateDistrictType(ID: ID!, districtType: DistrictTypeInput, buildStatic: Boolean): Int
    deleteDistrictType(ID: ID!, buildStatic: Boolean): Int
    createDistrict(district: DistrictInput, buildStatic: Boolean): District
    updateDistrict(ID: ID!, district: DistrictInput, buildStatic: Boolean): Int
    deleteDistrict(ID: ID!, buildStatic: Boolean): Int
    createOffice(office: OfficeInput, buildStatic: Boolean): Office
    updateOffice(ID: ID!, office: OfficeInput, buildStatic: Boolean): Int
    deleteOffice(ID: ID!, buildStatic: Boolean): Int
    createPollItem(pollItem: PollItemInput, buildStatic: Boolean): PollItem
    updatePollItem(ID: ID!, pollItem: PollItemInput, buildStatic: Boolean): Int
    deletePollItem(ID: ID!, buildStatic: Boolean): Int
    createVolunteerOpportunityType(volunteerOpportunityType: VolunteerOpportunityTypeInput, buildStatic: Boolean): VolunteerOpportunityType
    updateVolunteerOpportunityType(ID: ID!, volunteerOpportunityType: VolunteerOpportunityTypeInput, buildStatic: Boolean): Int
    deleteVolunteerOpportunityType(ID: ID!, buildStatic: Boolean): Int
    createLiveEventType(liveEventType: LiveEventTypeInput, buildStatic: Boolean): LiveEventType
    updateLiveEventType(ID: ID!, liveEventType: LiveEventTypeInput, buildStatic: Boolean): Int
    deleteLiveEventType(ID: ID!, buildStatic: Boolean): Int
    createCompany(company: CompanyInput, buildStatic: Boolean): Company
    updateCompany(ID: ID!, company: CompanyInput, buildStatic: Boolean): Int
    deleteCompany(ID: ID!, buildStatic: Boolean): Int
    createProductType(productType: ProductTypeInput, buildStatic: Boolean): ProductType
    updateProductType(ID: ID!, productType: ProductTypeInput, buildStatic: Boolean): Int
    deleteProductType(ID: ID!, buildStatic: Boolean): Int
    createCustomerCart(customerCart: CustomerCartInput, buildStatic: Boolean): CustomerCart
    updateCustomerCart(ID: ID!, customerCart: CustomerCartInput, buildStatic: Boolean): Int
    deleteCustomerCart(ID: ID!, buildStatic: Boolean): Int
    createPayment(payment: PaymentInput, buildStatic: Boolean): Payment
    updatePayment(ID: ID!, payment: PaymentInput, buildStatic: Boolean): Int
    deletePayment(ID: ID!, buildStatic: Boolean): Int
    createDelivery(delivery: DeliveryInput, buildStatic: Boolean): Delivery
    updateDelivery(ID: ID!, delivery: DeliveryInput, buildStatic: Boolean): Int
    deleteDelivery(ID: ID!, buildStatic: Boolean): Int
    createBoycott(boycott: BoycottInput, buildStatic: Boolean): Boycott
    updateBoycott(ID: ID!, boycott: BoycottInput, buildStatic: Boolean): Int
    deleteBoycott(ID: ID!, buildStatic: Boolean): Int
    createBoycottMembership(boycottMembership: BoycottMembershipInput, buildStatic: Boolean): BoycottMembership
    updateBoycottMembership(ID: ID!, boycottMembership: BoycottMembershipInput, buildStatic: Boolean): Int
    deleteBoycottMembership(ID: ID!, buildStatic: Boolean): Int
    createElection(election: ElectionInput, buildStatic: Boolean): Election
    updateElection(ID: ID!, election: ElectionInput, buildStatic: Boolean): Int
    deleteElection(ID: ID!, buildStatic: Boolean): Int
    createElectionMembership(electionMembership: ElectionMembershipInput, buildStatic: Boolean): ElectionMembership
    updateElectionMembership(ID: ID!, electionMembership: ElectionMembershipInput, buildStatic: Boolean): Int
    deleteElectionMembership(ID: ID!, buildStatic: Boolean): Int
    createPetitionMembership(petitionMembership: PetitionMembershipInput, buildStatic: Boolean): PetitionMembership
    updatePetitionMembership(ID: ID!, petitionMembership: PetitionMembershipInput, buildStatic: Boolean): Int
    deletePetitionMembership(ID: ID!, buildStatic: Boolean): Int
    createPollMembership(pollMembership: PollMembershipInput, buildStatic: Boolean): PollMembership
    updatePollMembership(ID: ID!, pollMembership: PollMembershipInput, buildStatic: Boolean): Int
    deletePollMembership(ID: ID!, buildStatic: Boolean): Int
    createVolunteerOpportunityMembership(volunteerOpportunityMembership: VolunteerOpportunityMembershipInput, buildStatic: Boolean): VolunteerOpportunityMembership
    updateVolunteerOpportunityMembership(ID: ID!, volunteerOpportunityMembership: VolunteerOpportunityMembershipInput, buildStatic: Boolean): Int
    deleteVolunteerOpportunityMembership(ID: ID!, buildStatic: Boolean): Int
    createLiveEventMembership(liveEventMembership: LiveEventMembershipInput, buildStatic: Boolean): LiveEventMembership
    updateLiveEventMembership(ID: ID!, liveEventMembership: LiveEventMembershipInput, buildStatic: Boolean): Int
    deleteLiveEventMembership(ID: ID!, buildStatic: Boolean): Int
    createProduct(product: ProductInput, buildStatic: Boolean): Product
    updateProduct(ID: ID!, product: ProductInput, buildStatic: Boolean): Int
    deleteProduct(ID: ID!, buildStatic: Boolean): Int
    createProductMembership(productMembership: ProductMembershipInput, buildStatic: Boolean): ProductMembership
    updateProductMembership(ID: ID!, productMembership: ProductMembershipInput, buildStatic: Boolean): Int
    deleteProductMembership(ID: ID!, buildStatic: Boolean): Int
    createDonationCampaign(donationCampaign: DonationCampaignInput, buildStatic: Boolean): DonationCampaign
    updateDonationCampaign(ID: ID!, donationCampaign: DonationCampaignInput, buildStatic: Boolean): Int
    deleteDonationCampaign(ID: ID!, buildStatic: Boolean): Int
    createDonationCampaignMembership(donationCampaignMembership: DonationCampaignMembershipInput, buildStatic: Boolean): DonationCampaignMembership
    updateDonationCampaignMembership(ID: ID!, donationCampaignMembership: DonationCampaignMembershipInput, buildStatic: Boolean): Int
    deleteDonationCampaignMembership(ID: ID!, buildStatic: Boolean): Int
    createPetition(petition: PetitionInput, buildStatic: Boolean): Petition
    updatePetition(ID: ID!, petition: PetitionInput, buildStatic: Boolean): Int
    deletePetition(ID: ID!, buildStatic: Boolean): Int
    createPetitionSigner(petitionSigner: PetitionSignerInput, buildStatic: Boolean): PetitionSigner
    updatePetitionSigner(ID: ID!, petitionSigner: PetitionSignerInput, buildStatic: Boolean): Int
    deletePetitionSigner(ID: ID!, buildStatic: Boolean): Int
    createPoll(poll: PollInput, buildStatic: Boolean): Poll
    updatePoll(ID: ID!, poll: PollInput, buildStatic: Boolean): Int
    deletePoll(ID: ID!, buildStatic: Boolean): Int
    createPollRespondant(pollRespondant: PollRespondantInput, buildStatic: Boolean): PollRespondant
    updatePollRespondant(ID: ID!, pollRespondant: PollRespondantInput, buildStatic: Boolean): Int
    deletePollRespondant(ID: ID!, buildStatic: Boolean): Int
    createPurchaser(purchaser: PurchaserInput, buildStatic: Boolean): Purchaser
    updatePurchaser(ID: ID!, purchaser: PurchaserInput, buildStatic: Boolean): Int
    deletePurchaser(ID: ID!, buildStatic: Boolean): Int
    createCustomerOrder(customerOrder: CustomerOrderInput, buildStatic: Boolean): CustomerOrder
    updateCustomerOrder(ID: ID!, customerOrder: CustomerOrderInput, buildStatic: Boolean): Int
    deleteCustomerOrder(ID: ID!, buildStatic: Boolean): Int
    createDonor(donor: DonorInput, buildStatic: Boolean): Donor
    updateDonor(ID: ID!, donor: DonorInput, buildStatic: Boolean): Int
    deleteDonor(ID: ID!, buildStatic: Boolean): Int
    createLiveEvent(liveEvent: LiveEventInput, buildStatic: Boolean): LiveEvent
    updateLiveEvent(ID: ID!, liveEvent: LiveEventInput, buildStatic: Boolean): Int
    deleteLiveEvent(ID: ID!, buildStatic: Boolean): Int
    createEventAttendee(eventAttendee: EventAttendeeInput, buildStatic: Boolean): EventAttendee
    updateEventAttendee(ID: ID!, eventAttendee: EventAttendeeInput, buildStatic: Boolean): Int
    deleteEventAttendee(ID: ID!, buildStatic: Boolean): Int
    createVoter(voter: VoterInput, buildStatic: Boolean): Voter
    updateVoter(ID: ID!, voter: VoterInput, buildStatic: Boolean): Int
    deleteVoter(ID: ID!, buildStatic: Boolean): Int
    createVolunteerOpportunity(volunteerOpportunity: VolunteerOpportunityInput, buildStatic: Boolean): VolunteerOpportunity
    updateVolunteerOpportunity(ID: ID!, volunteerOpportunity: VolunteerOpportunityInput, buildStatic: Boolean): Int
    deleteVolunteerOpportunity(ID: ID!, buildStatic: Boolean): Int
    createVolunteer(volunteer: VolunteerInput, buildStatic: Boolean): Volunteer
    updateVolunteer(ID: ID!, volunteer: VolunteerInput, buildStatic: Boolean): Int
    deleteVolunteer(ID: ID!, buildStatic: Boolean): Int
    createFollower(follower: FollowerInput, buildStatic: Boolean): Follower
    updateFollower(ID: ID!, follower: FollowerInput, buildStatic: Boolean): Int
    deleteFollower(ID: ID!, buildStatic: Boolean): Int
    createTerritory(territory: TerritoryInput, buildStatic: Boolean): Territory
    updateTerritory(ID: ID!, territory: TerritoryInput, buildStatic: Boolean): Int
    deleteTerritory(ID: ID!, buildStatic: Boolean): Int
    createActivityType(activityType: ActivityTypeInput, buildStatic: Boolean): ActivityType
    updateActivityType(ID: ID!, activityType: ActivityTypeInput, buildStatic: Boolean): Int
    deleteActivityType(ID: ID!, buildStatic: Boolean): Int
    createActivity(activity: ActivityInput, buildStatic: Boolean): Activity
    updateActivity(ID: ID!, activity: ActivityInput, buildStatic: Boolean): Int
    deleteActivity(ID: ID!, buildStatic: Boolean): Int
    createNote(note: NoteInput, buildStatic: Boolean): Note
    updateNote(ID: ID!, note: NoteInput, buildStatic: Boolean): Int
    deleteNote(ID: ID!, buildStatic: Boolean): Int
    createAccount(account: AccountInput, buildStatic: Boolean): Account
    updateAccount(ID: ID!, account: AccountInput, buildStatic: Boolean): Int
    deleteAccount(ID: ID!, buildStatic: Boolean): Int
    createOwnerMembership(ownerMembership: OwnerMembershipInput, buildStatic: Boolean): OwnerMembership
    updateOwnerMembership(ID: ID!, ownerMembership: OwnerMembershipInput, buildStatic: Boolean): Int
    deleteOwnerMembership(ID: ID!, buildStatic: Boolean): Int
    createContact(contact: ContactInput, buildStatic: Boolean): Contact
    updateContact(ID: ID!, contact: ContactInput, buildStatic: Boolean): Int
    deleteContact(ID: ID!, buildStatic: Boolean): Int
    createContactMembership(contactMembership: ContactMembershipInput, buildStatic: Boolean): ContactMembership
    updateContactMembership(ID: ID!, contactMembership: ContactMembershipInput, buildStatic: Boolean): Int
    deleteContactMembership(ID: ID!, buildStatic: Boolean): Int
    createAgent(agent: AgentInput, buildStatic: Boolean): Agent
    updateAgent(ID: ID!, agent: AgentInput, buildStatic: Boolean): Int
    deleteAgent(ID: ID!, buildStatic: Boolean): Int
    createAgentMembership(agentMembership: AgentMembershipInput, buildStatic: Boolean): AgentMembership
    updateAgentMembership(ID: ID!, agentMembership: AgentMembershipInput, buildStatic: Boolean): Int
    deleteAgentMembership(ID: ID!, buildStatic: Boolean): Int
    createHomePage(homePage: HomePageInput, buildStatic: Boolean): HomePage
    updateHomePage(ID: ID!, homePage: HomePageInput, buildStatic: Boolean): Int
    deleteHomePage(ID: ID!, buildStatic: Boolean): Int
    createPhoto(photo: PhotoInput, buildStatic: Boolean): Photo
    updatePhoto(ID: ID!, photo: PhotoInput, buildStatic: Boolean): Int
    deletePhoto(ID: ID!, buildStatic: Boolean): Int
    createCause(cause: CauseInput, buildStatic: Boolean): Cause
    updateCause(ID: ID!, cause: CauseInput, buildStatic: Boolean): Int
    deleteCause(ID: ID!, buildStatic: Boolean): Int
    createComponentType(componentType: ComponentTypeInput, buildStatic: Boolean): ComponentType
    updateComponentType(ID: ID!, componentType: ComponentTypeInput, buildStatic: Boolean): Int
    deleteComponentType(ID: ID!, buildStatic: Boolean): Int
    createFieldType(fieldType: FieldTypeInput, buildStatic: Boolean): FieldType
    updateFieldType(ID: ID!, fieldType: FieldTypeInput, buildStatic: Boolean): Int
    deleteFieldType(ID: ID!, buildStatic: Boolean): Int
  }
  `.replace(/[\u200B-\u200D\uFEFF]/g, '');