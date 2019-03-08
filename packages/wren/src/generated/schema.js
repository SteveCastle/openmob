const { gql } = require('apollo-server');
module.exports = `type Time {
    seconds: Int!
    nanos: Int!
  }


input ComponentImplementationInput {
  secret: Int
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
  ElectionType: ID
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


type ComponentImplementation {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
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
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LayoutRow {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Layout: ID!
}
​
type ComponentType {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LayoutColumn {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutRow: ID!
}
​
type FieldType {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type Component {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  ComponentType: ID!
  LayoutColumn: ID
}
​
type Field {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  FieldType: ID!
  Component: ID
}
​
type HomePage {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Cause: ID!
  Layout: ID
}
​
type Layout {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutType: ID
}
​
type LandingPage {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Cause: ID!
  Layout: ID
}
​
type Experiment {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LandingPage: ID
}
​
type Issue {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: ID!
}
​
type Candidate {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Election: ID!
}
​
type DistrictType {
  ID: ID!
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
  DistrictType: ID!
}
​
type Office {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: ID
}
​
type PollItem {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Poll: ID!
}
​
type VolunteerOpportunityType {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LiveEventType {
  ID: ID!
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
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type CustomerCart {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type Payment {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: ID!
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
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type BoycottMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Boycott: ID!
}
​
type Election {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type ElectionMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Election: ID!
}
​
type PetitionMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Petition: ID!
}
​
type PollMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Petition: ID!
}
​
type VolunteerOpportunityMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  VolunteerOpportunity: ID!
}
​
type LiveEventMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  LiveEvent: ID!
}
​
type Product {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ProductType: ID!
}
​
type ProductMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Product: ID!
}
​
type DonationCampaign {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type DonationCampaignMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  DonationCampaign: ID!
}
​
type Petition {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type PetitionSigner {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Petition: ID!
  Contact: ID!
  Cause: ID!
}
​
type Poll {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type PollRespondant {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Poll: ID!
  Contact: ID!
  Cause: ID!
}
​
type Purchaser {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: ID!
  Contact: ID!
  Cause: ID!
}
​
type CustomerOrder {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerCart: ID!
}
​
type Donor {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: ID!
  Contact: ID!
  Cause: ID!
}
​
type LiveEvent {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LiveEventType: ID!
}
​
type EventAttendee {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LiveEvent: ID!
  Contact: ID!
  Cause: ID!
}
​
type Voter {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: ID!
  Cause: ID!
}
​
type VolunteerOpportunity {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ElectionType: ID
}
​
type Volunteer {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  VolunteerOpportunity: ID!
  Contact: ID!
  Cause: ID!
}
​
type Follower {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: ID!
  Cause: ID!
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
  ActivityType: ID!
  Contact: ID!
  Cause: ID!
}
​
type Note {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: ID!
  Cause: ID!
  Body: String
}
​
type Account {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Username: String!
}
​
type OwnerMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Account: ID!
}
​
type Contact {
  ID: ID!
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
  Cause: ID!
  Contact: ID!
}
​
type Cause {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Slug: String!
  Summary: String
}
​
type Agent {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Account: ID!
}
​
type AgentMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: ID!
  Agent: ID!
}
​

  type Query {
    
    getComponentImplementation(ID: ID!): ComponentImplementation
    listComponentImplementation: [ComponentImplementation]
    getACL(ID: ID!): ACL
    listACL: [ACL]
    getMailingAddress(ID: ID!): MailingAddress
    listMailingAddress: [MailingAddress]
    getPhoneNumber(ID: ID!): PhoneNumber
    listPhoneNumber: [PhoneNumber]
    getEmailAddress(ID: ID!): EmailAddress
    listEmailAddress: [EmailAddress]
    getPhoto(ID: ID!): Photo
    listPhoto: [Photo]
    getLayoutType(ID: ID!): LayoutType
    listLayoutType: [LayoutType]
    getLayoutRow(ID: ID!): LayoutRow
    listLayoutRow: [LayoutRow]
    getComponentType(ID: ID!): ComponentType
    listComponentType: [ComponentType]
    getLayoutColumn(ID: ID!): LayoutColumn
    listLayoutColumn: [LayoutColumn]
    getFieldType(ID: ID!): FieldType
    listFieldType: [FieldType]
    getComponent(ID: ID!): Component
    listComponent: [Component]
    getField(ID: ID!): Field
    listField: [Field]
    getHomePage(ID: ID!): HomePage
    listHomePage: [HomePage]
    getLayout(ID: ID!): Layout
    listLayout: [Layout]
    getLandingPage(ID: ID!): LandingPage
    listLandingPage: [LandingPage]
    getExperiment(ID: ID!): Experiment
    listExperiment: [Experiment]
    getIssue(ID: ID!): Issue
    listIssue: [Issue]
    getCandidate(ID: ID!): Candidate
    listCandidate: [Candidate]
    getDistrictType(ID: ID!): DistrictType
    listDistrictType: [DistrictType]
    getDistrict(ID: ID!): District
    listDistrict: [District]
    getOffice(ID: ID!): Office
    listOffice: [Office]
    getPollItem(ID: ID!): PollItem
    listPollItem: [PollItem]
    getVolunteerOpportunityType(ID: ID!): VolunteerOpportunityType
    listVolunteerOpportunityType: [VolunteerOpportunityType]
    getLiveEventType(ID: ID!): LiveEventType
    listLiveEventType: [LiveEventType]
    getCompany(ID: ID!): Company
    listCompany: [Company]
    getProductType(ID: ID!): ProductType
    listProductType: [ProductType]
    getCustomerCart(ID: ID!): CustomerCart
    listCustomerCart: [CustomerCart]
    getPayment(ID: ID!): Payment
    listPayment: [Payment]
    getDelivery(ID: ID!): Delivery
    listDelivery: [Delivery]
    getBoycott(ID: ID!): Boycott
    listBoycott: [Boycott]
    getBoycottMembership(ID: ID!): BoycottMembership
    listBoycottMembership: [BoycottMembership]
    getElection(ID: ID!): Election
    listElection: [Election]
    getElectionMembership(ID: ID!): ElectionMembership
    listElectionMembership: [ElectionMembership]
    getPetitionMembership(ID: ID!): PetitionMembership
    listPetitionMembership: [PetitionMembership]
    getPollMembership(ID: ID!): PollMembership
    listPollMembership: [PollMembership]
    getVolunteerOpportunityMembership(ID: ID!): VolunteerOpportunityMembership
    listVolunteerOpportunityMembership: [VolunteerOpportunityMembership]
    getLiveEventMembership(ID: ID!): LiveEventMembership
    listLiveEventMembership: [LiveEventMembership]
    getProduct(ID: ID!): Product
    listProduct: [Product]
    getProductMembership(ID: ID!): ProductMembership
    listProductMembership: [ProductMembership]
    getDonationCampaign(ID: ID!): DonationCampaign
    listDonationCampaign: [DonationCampaign]
    getDonationCampaignMembership(ID: ID!): DonationCampaignMembership
    listDonationCampaignMembership: [DonationCampaignMembership]
    getPetition(ID: ID!): Petition
    listPetition: [Petition]
    getPetitionSigner(ID: ID!): PetitionSigner
    listPetitionSigner: [PetitionSigner]
    getPoll(ID: ID!): Poll
    listPoll: [Poll]
    getPollRespondant(ID: ID!): PollRespondant
    listPollRespondant: [PollRespondant]
    getPurchaser(ID: ID!): Purchaser
    listPurchaser: [Purchaser]
    getCustomerOrder(ID: ID!): CustomerOrder
    listCustomerOrder: [CustomerOrder]
    getDonor(ID: ID!): Donor
    listDonor: [Donor]
    getLiveEvent(ID: ID!): LiveEvent
    listLiveEvent: [LiveEvent]
    getEventAttendee(ID: ID!): EventAttendee
    listEventAttendee: [EventAttendee]
    getVoter(ID: ID!): Voter
    listVoter: [Voter]
    getVolunteerOpportunity(ID: ID!): VolunteerOpportunity
    listVolunteerOpportunity: [VolunteerOpportunity]
    getVolunteer(ID: ID!): Volunteer
    listVolunteer: [Volunteer]
    getFollower(ID: ID!): Follower
    listFollower: [Follower]
    getTerritory(ID: ID!): Territory
    listTerritory: [Territory]
    getActivityType(ID: ID!): ActivityType
    listActivityType: [ActivityType]
    getActivity(ID: ID!): Activity
    listActivity: [Activity]
    getNote(ID: ID!): Note
    listNote: [Note]
    getAccount(ID: ID!): Account
    listAccount: [Account]
    getOwnerMembership(ID: ID!): OwnerMembership
    listOwnerMembership: [OwnerMembership]
    getContact(ID: ID!): Contact
    listContact: [Contact]
    getContactMembership(ID: ID!): ContactMembership
    listContactMembership: [ContactMembership]
    getCause(ID: ID!): Cause
    listCause: [Cause]
    getAgent(ID: ID!): Agent
    listAgent: [Agent]
    getAgentMembership(ID: ID!): AgentMembership
    listAgentMembership: [AgentMembership]
  }

  type Mutation {
    
    createComponentImplementation(componentImplementation: ComponentImplementationInput): ComponentImplementation
    updateComponentImplementation(ID: ID!, componentImplementation: ComponentImplementationInput): Int
    deleteComponentImplementation(ID: ID!): Int
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