const { gql } = require('apollo-server');
module.exports = `type Time {
    seconds: Int!
    nanos: Int!
  }


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
input PhotoInput {
  secret: Int
  ImgURL: String!
}
​
input LayoutTypeInput {
  secret: Int
  Title: String!
}
​
input LayoutRowInput {
  secret: Int
  Layout: ID!
}
​
input ComponentImplementationInput {
  secret: Int
}
​
input ComponentTypeInput {
  secret: Int
  Title: String!
}
​
input LayoutColumnInput {
  secret: Int
  LayoutRow: ID!
}
​
input FieldTypeInput {
  secret: Int
  Title: String!
}
​
input ComponentInput {
  secret: Int
  ComponentType: ID!
  LayoutColumn: ID
}
​
input FieldInput {
  secret: Int
  FieldType: ID!
  Component: ID
}
​
input HomePageInput {
  secret: Int
  Title: String!
  Cause: ID!
  Layout: ID
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
input CauseInput {
  secret: Int
  Title: String!
  Slug: String!
  Summary: String
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
type Photo {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  ImgURL: String!
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
type LayoutRow {
  ID: ID!
  LayoutColumns: [LayoutColumn]
  CreatedAt: Time!
  UpdatedAt: Time!
  Layout: Layout
}
​
type ComponentImplementation {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type ComponentType {
  ID: ID!
  Components: [Component]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LayoutColumn {
  ID: ID!
  Components: [Component]
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutRow: LayoutRow
}
​
type FieldType {
  ID: ID!
  Fields: [Field]
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type Component {
  ID: ID!
  Fields: [Field]
  CreatedAt: Time!
  UpdatedAt: Time!
  ComponentType: ComponentType
  LayoutColumn: LayoutColumn
}
​
type Field {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  FieldType: FieldType
  Component: Component
}
​
type HomePage {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Cause: Cause
  Layout: Layout
}
​
type Layout {
  ID: ID!
  LayoutRows: [LayoutRow]
  HomePages: [HomePage]
  LandingPages: [LandingPage]
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
type Cause {
  ID: ID!
  HomePages: [HomePage]
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

  type Query {
    
    getACL(ID: ID!): ACL
    listACL(limit: Int): [ACL]
    getMailingAddress(ID: ID!): MailingAddress
    listMailingAddress(limit: Int): [MailingAddress]
    getPhoneNumber(ID: ID!): PhoneNumber
    listPhoneNumber(limit: Int): [PhoneNumber]
    getEmailAddress(ID: ID!): EmailAddress
    listEmailAddress(limit: Int): [EmailAddress]
    getPhoto(ID: ID!): Photo
    listPhoto(limit: Int): [Photo]
    getLayoutType(ID: ID!): LayoutType
    listLayoutType(limit: Int): [LayoutType]
    getLayoutRow(ID: ID!): LayoutRow
    listLayoutRow(limit: Int): [LayoutRow]
    getComponentImplementation(ID: ID!): ComponentImplementation
    listComponentImplementation(limit: Int): [ComponentImplementation]
    getComponentType(ID: ID!): ComponentType
    listComponentType(limit: Int): [ComponentType]
    getLayoutColumn(ID: ID!): LayoutColumn
    listLayoutColumn(limit: Int): [LayoutColumn]
    getFieldType(ID: ID!): FieldType
    listFieldType(limit: Int): [FieldType]
    getComponent(ID: ID!): Component
    listComponent(limit: Int): [Component]
    getField(ID: ID!): Field
    listField(limit: Int): [Field]
    getHomePage(ID: ID!): HomePage
    listHomePage(limit: Int): [HomePage]
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
    getCause(ID: ID!): Cause
    listCause(limit: Int): [Cause]
    getAgent(ID: ID!): Agent
    listAgent(limit: Int): [Agent]
    getAgentMembership(ID: ID!): AgentMembership
    listAgentMembership(limit: Int): [AgentMembership]
  }

  type Mutation {
    
    createACL(acl: ACLInput): ACL
    updateACL(ID: ID!, acl: ACLInput): Int
    deleteACL(ID: ID!): Int
    createMailingAddress(mailingAddress: MailingAddressInput): MailingAddress
    updateMailingAddress(ID: ID!, mailingAddress: MailingAddressInput): Int
    deleteMailingAddress(ID: ID!): Int
    createPhoneNumber(phoneNumber: PhoneNumberInput): PhoneNumber
    updatePhoneNumber(ID: ID!, phoneNumber: PhoneNumberInput): Int
    deletePhoneNumber(ID: ID!): Int
    createEmailAddress(emailAddress: EmailAddressInput): EmailAddress
    updateEmailAddress(ID: ID!, emailAddress: EmailAddressInput): Int
    deleteEmailAddress(ID: ID!): Int
    createPhoto(photo: PhotoInput): Photo
    updatePhoto(ID: ID!, photo: PhotoInput): Int
    deletePhoto(ID: ID!): Int
    createLayoutType(layoutType: LayoutTypeInput): LayoutType
    updateLayoutType(ID: ID!, layoutType: LayoutTypeInput): Int
    deleteLayoutType(ID: ID!): Int
    createLayoutRow(layoutRow: LayoutRowInput): LayoutRow
    updateLayoutRow(ID: ID!, layoutRow: LayoutRowInput): Int
    deleteLayoutRow(ID: ID!): Int
    createComponentImplementation(componentImplementation: ComponentImplementationInput): ComponentImplementation
    updateComponentImplementation(ID: ID!, componentImplementation: ComponentImplementationInput): Int
    deleteComponentImplementation(ID: ID!): Int
    createComponentType(componentType: ComponentTypeInput): ComponentType
    updateComponentType(ID: ID!, componentType: ComponentTypeInput): Int
    deleteComponentType(ID: ID!): Int
    createLayoutColumn(layoutColumn: LayoutColumnInput): LayoutColumn
    updateLayoutColumn(ID: ID!, layoutColumn: LayoutColumnInput): Int
    deleteLayoutColumn(ID: ID!): Int
    createFieldType(fieldType: FieldTypeInput): FieldType
    updateFieldType(ID: ID!, fieldType: FieldTypeInput): Int
    deleteFieldType(ID: ID!): Int
    createComponent(component: ComponentInput): Component
    updateComponent(ID: ID!, component: ComponentInput): Int
    deleteComponent(ID: ID!): Int
    createField(field: FieldInput): Field
    updateField(ID: ID!, field: FieldInput): Int
    deleteField(ID: ID!): Int
    createHomePage(homePage: HomePageInput): HomePage
    updateHomePage(ID: ID!, homePage: HomePageInput): Int
    deleteHomePage(ID: ID!): Int
    createLayout(layout: LayoutInput): Layout
    updateLayout(ID: ID!, layout: LayoutInput): Int
    deleteLayout(ID: ID!): Int
    createLandingPage(landingPage: LandingPageInput): LandingPage
    updateLandingPage(ID: ID!, landingPage: LandingPageInput): Int
    deleteLandingPage(ID: ID!): Int
    createExperiment(experiment: ExperimentInput): Experiment
    updateExperiment(ID: ID!, experiment: ExperimentInput): Int
    deleteExperiment(ID: ID!): Int
    createIssue(issue: IssueInput): Issue
    updateIssue(ID: ID!, issue: IssueInput): Int
    deleteIssue(ID: ID!): Int
    createCandidate(candidate: CandidateInput): Candidate
    updateCandidate(ID: ID!, candidate: CandidateInput): Int
    deleteCandidate(ID: ID!): Int
    createDistrictType(districtType: DistrictTypeInput): DistrictType
    updateDistrictType(ID: ID!, districtType: DistrictTypeInput): Int
    deleteDistrictType(ID: ID!): Int
    createDistrict(district: DistrictInput): District
    updateDistrict(ID: ID!, district: DistrictInput): Int
    deleteDistrict(ID: ID!): Int
    createOffice(office: OfficeInput): Office
    updateOffice(ID: ID!, office: OfficeInput): Int
    deleteOffice(ID: ID!): Int
    createPollItem(pollItem: PollItemInput): PollItem
    updatePollItem(ID: ID!, pollItem: PollItemInput): Int
    deletePollItem(ID: ID!): Int
    createVolunteerOpportunityType(volunteerOpportunityType: VolunteerOpportunityTypeInput): VolunteerOpportunityType
    updateVolunteerOpportunityType(ID: ID!, volunteerOpportunityType: VolunteerOpportunityTypeInput): Int
    deleteVolunteerOpportunityType(ID: ID!): Int
    createLiveEventType(liveEventType: LiveEventTypeInput): LiveEventType
    updateLiveEventType(ID: ID!, liveEventType: LiveEventTypeInput): Int
    deleteLiveEventType(ID: ID!): Int
    createCompany(company: CompanyInput): Company
    updateCompany(ID: ID!, company: CompanyInput): Int
    deleteCompany(ID: ID!): Int
    createProductType(productType: ProductTypeInput): ProductType
    updateProductType(ID: ID!, productType: ProductTypeInput): Int
    deleteProductType(ID: ID!): Int
    createCustomerCart(customerCart: CustomerCartInput): CustomerCart
    updateCustomerCart(ID: ID!, customerCart: CustomerCartInput): Int
    deleteCustomerCart(ID: ID!): Int
    createPayment(payment: PaymentInput): Payment
    updatePayment(ID: ID!, payment: PaymentInput): Int
    deletePayment(ID: ID!): Int
    createDelivery(delivery: DeliveryInput): Delivery
    updateDelivery(ID: ID!, delivery: DeliveryInput): Int
    deleteDelivery(ID: ID!): Int
    createBoycott(boycott: BoycottInput): Boycott
    updateBoycott(ID: ID!, boycott: BoycottInput): Int
    deleteBoycott(ID: ID!): Int
    createBoycottMembership(boycottMembership: BoycottMembershipInput): BoycottMembership
    updateBoycottMembership(ID: ID!, boycottMembership: BoycottMembershipInput): Int
    deleteBoycottMembership(ID: ID!): Int
    createElection(election: ElectionInput): Election
    updateElection(ID: ID!, election: ElectionInput): Int
    deleteElection(ID: ID!): Int
    createElectionMembership(electionMembership: ElectionMembershipInput): ElectionMembership
    updateElectionMembership(ID: ID!, electionMembership: ElectionMembershipInput): Int
    deleteElectionMembership(ID: ID!): Int
    createPetitionMembership(petitionMembership: PetitionMembershipInput): PetitionMembership
    updatePetitionMembership(ID: ID!, petitionMembership: PetitionMembershipInput): Int
    deletePetitionMembership(ID: ID!): Int
    createPollMembership(pollMembership: PollMembershipInput): PollMembership
    updatePollMembership(ID: ID!, pollMembership: PollMembershipInput): Int
    deletePollMembership(ID: ID!): Int
    createVolunteerOpportunityMembership(volunteerOpportunityMembership: VolunteerOpportunityMembershipInput): VolunteerOpportunityMembership
    updateVolunteerOpportunityMembership(ID: ID!, volunteerOpportunityMembership: VolunteerOpportunityMembershipInput): Int
    deleteVolunteerOpportunityMembership(ID: ID!): Int
    createLiveEventMembership(liveEventMembership: LiveEventMembershipInput): LiveEventMembership
    updateLiveEventMembership(ID: ID!, liveEventMembership: LiveEventMembershipInput): Int
    deleteLiveEventMembership(ID: ID!): Int
    createProduct(product: ProductInput): Product
    updateProduct(ID: ID!, product: ProductInput): Int
    deleteProduct(ID: ID!): Int
    createProductMembership(productMembership: ProductMembershipInput): ProductMembership
    updateProductMembership(ID: ID!, productMembership: ProductMembershipInput): Int
    deleteProductMembership(ID: ID!): Int
    createDonationCampaign(donationCampaign: DonationCampaignInput): DonationCampaign
    updateDonationCampaign(ID: ID!, donationCampaign: DonationCampaignInput): Int
    deleteDonationCampaign(ID: ID!): Int
    createDonationCampaignMembership(donationCampaignMembership: DonationCampaignMembershipInput): DonationCampaignMembership
    updateDonationCampaignMembership(ID: ID!, donationCampaignMembership: DonationCampaignMembershipInput): Int
    deleteDonationCampaignMembership(ID: ID!): Int
    createPetition(petition: PetitionInput): Petition
    updatePetition(ID: ID!, petition: PetitionInput): Int
    deletePetition(ID: ID!): Int
    createPetitionSigner(petitionSigner: PetitionSignerInput): PetitionSigner
    updatePetitionSigner(ID: ID!, petitionSigner: PetitionSignerInput): Int
    deletePetitionSigner(ID: ID!): Int
    createPoll(poll: PollInput): Poll
    updatePoll(ID: ID!, poll: PollInput): Int
    deletePoll(ID: ID!): Int
    createPollRespondant(pollRespondant: PollRespondantInput): PollRespondant
    updatePollRespondant(ID: ID!, pollRespondant: PollRespondantInput): Int
    deletePollRespondant(ID: ID!): Int
    createPurchaser(purchaser: PurchaserInput): Purchaser
    updatePurchaser(ID: ID!, purchaser: PurchaserInput): Int
    deletePurchaser(ID: ID!): Int
    createCustomerOrder(customerOrder: CustomerOrderInput): CustomerOrder
    updateCustomerOrder(ID: ID!, customerOrder: CustomerOrderInput): Int
    deleteCustomerOrder(ID: ID!): Int
    createDonor(donor: DonorInput): Donor
    updateDonor(ID: ID!, donor: DonorInput): Int
    deleteDonor(ID: ID!): Int
    createLiveEvent(liveEvent: LiveEventInput): LiveEvent
    updateLiveEvent(ID: ID!, liveEvent: LiveEventInput): Int
    deleteLiveEvent(ID: ID!): Int
    createEventAttendee(eventAttendee: EventAttendeeInput): EventAttendee
    updateEventAttendee(ID: ID!, eventAttendee: EventAttendeeInput): Int
    deleteEventAttendee(ID: ID!): Int
    createVoter(voter: VoterInput): Voter
    updateVoter(ID: ID!, voter: VoterInput): Int
    deleteVoter(ID: ID!): Int
    createVolunteerOpportunity(volunteerOpportunity: VolunteerOpportunityInput): VolunteerOpportunity
    updateVolunteerOpportunity(ID: ID!, volunteerOpportunity: VolunteerOpportunityInput): Int
    deleteVolunteerOpportunity(ID: ID!): Int
    createVolunteer(volunteer: VolunteerInput): Volunteer
    updateVolunteer(ID: ID!, volunteer: VolunteerInput): Int
    deleteVolunteer(ID: ID!): Int
    createFollower(follower: FollowerInput): Follower
    updateFollower(ID: ID!, follower: FollowerInput): Int
    deleteFollower(ID: ID!): Int
    createTerritory(territory: TerritoryInput): Territory
    updateTerritory(ID: ID!, territory: TerritoryInput): Int
    deleteTerritory(ID: ID!): Int
    createActivityType(activityType: ActivityTypeInput): ActivityType
    updateActivityType(ID: ID!, activityType: ActivityTypeInput): Int
    deleteActivityType(ID: ID!): Int
    createActivity(activity: ActivityInput): Activity
    updateActivity(ID: ID!, activity: ActivityInput): Int
    deleteActivity(ID: ID!): Int
    createNote(note: NoteInput): Note
    updateNote(ID: ID!, note: NoteInput): Int
    deleteNote(ID: ID!): Int
    createAccount(account: AccountInput): Account
    updateAccount(ID: ID!, account: AccountInput): Int
    deleteAccount(ID: ID!): Int
    createOwnerMembership(ownerMembership: OwnerMembershipInput): OwnerMembership
    updateOwnerMembership(ID: ID!, ownerMembership: OwnerMembershipInput): Int
    deleteOwnerMembership(ID: ID!): Int
    createContact(contact: ContactInput): Contact
    updateContact(ID: ID!, contact: ContactInput): Int
    deleteContact(ID: ID!): Int
    createContactMembership(contactMembership: ContactMembershipInput): ContactMembership
    updateContactMembership(ID: ID!, contactMembership: ContactMembershipInput): Int
    deleteContactMembership(ID: ID!): Int
    createCause(cause: CauseInput): Cause
    updateCause(ID: ID!, cause: CauseInput): Int
    deleteCause(ID: ID!): Int
    createAgent(agent: AgentInput): Agent
    updateAgent(ID: ID!, agent: AgentInput): Int
    deleteAgent(ID: ID!): Int
    createAgentMembership(agentMembership: AgentMembershipInput): AgentMembership
    updateAgentMembership(ID: ID!, agentMembership: AgentMembershipInput): Int
    deleteAgentMembership(ID: ID!): Int
  }
  `.replace(/[\u200B-\u200D\uFEFF]/g, '');