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
  Layout: Int!
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
  LayoutRow: Int!
}
​
input FieldTypeInput {
  secret: Int
  Title: String!
}
​
input ComponentInput {
  secret: Int
  ComponentType: Int!
  LayoutColumn: Int
}
​
input FieldInput {
  secret: Int
  FieldType: Int!
  Component: Int
}
​
input LayoutInput {
  secret: Int
  LayoutType: Int
}
​
input LandingPageInput {
  secret: Int
  Title: String!
  Cause: Int!
  Layout: Int
}
​
input ExperimentInput {
  secret: Int
  Title: String!
  LandingPage: Int
}
​
input IssueInput {
  secret: Int
  Title: String!
  Election: Int!
}
​
input CandidateInput {
  secret: Int
  Election: Int!
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
  DistrictType: Int!
}
​
input OfficeInput {
  secret: Int
  Title: String!
  Election: Int
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
  CustomerOrder: Int!
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
  Cause: Int!
  Boycott: Int!
}
​
input ElectionInput {
  secret: Int
  Title: String!
}
​
input ElectionMembershipInput {
  secret: Int
  Cause: Int!
  Election: Int!
}
​
input PetitionMembershipInput {
  secret: Int
  Cause: Int!
  Petition: Int!
}
​
input PollMembershipInput {
  secret: Int
  Cause: Int!
  Petition: Int!
}
​
input VolunteerOpportunityMembershipInput {
  secret: Int
  Cause: Int!
  VolunteerOpportunity: Int!
}
​
input LiveEventMembershipInput {
  secret: Int
  Cause: Int!
  LiveEvent: Int!
}
​
input ProductInput {
  secret: Int
  Title: String!
  ProductType: Int!
}
​
input ProductMembershipInput {
  secret: Int
  Cause: Int!
  Product: Int!
}
​
input DonationCampaignInput {
  secret: Int
  Title: String!
}
​
input DonationCampaignMembershipInput {
  secret: Int
  Cause: Int!
  DonationCampaign: Int!
}
​
input PetitionInput {
  secret: Int
  Title: String!
}
​
input PetitionSignerInput {
  secret: Int
  Petition: Int!
  Contact: Int!
  Cause: Int!
}
​
input PollInput {
  secret: Int
  Title: String!
}
​
input PollRespondantInput {
  secret: Int
  Poll: Int!
  Contact: Int!
  Cause: Int!
}
​
input PurchaserInput {
  secret: Int
  CustomerOrder: Int!
  Contact: Int!
  Cause: Int!
}
​
input CustomerOrderInput {
  secret: Int
  CustomerCart: Int!
}
​
input DonorInput {
  secret: Int
  CustomerOrder: Int!
  Contact: Int!
  Cause: Int!
}
​
input LiveEventInput {
  secret: Int
  Title: String!
  LiveEventType: Int!
}
​
input EventAttendeeInput {
  secret: Int
  LiveEvent: Int!
  Contact: Int!
  Cause: Int!
}
​
input VoterInput {
  secret: Int
  Contact: Int!
  Cause: Int!
}
​
input VolunteerOpportunityInput {
  secret: Int
  Title: String!
  ElectionType: Int
}
​
input VolunteerInput {
  secret: Int
  VolunteerOpportunity: Int!
  Contact: Int!
  Cause: Int!
}
​
input TerritoryInput {
  secret: Int
  Title: String!
}
​
input FollowerInput {
  secret: Int
  Contact: Int!
  Cause: Int!
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
  ActivityType: Int!
  Contact: Int!
  Cause: Int!
}
​
input NoteInput {
  secret: Int
  Contact: Int!
  Cause: Int!
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
  Cause: Int!
  Account: Int!
}
​
input ContactInput {
  secret: Int
}
​
input ContactMembershipInput {
  secret: Int
  Cause: Int!
  Contact: Int!
}
​
input CauseInput {
  secret: Int
  Title: String!
  Summary: String
}
​
input AgentInput {
  secret: Int
  Account: Int!
}
​
input AgentMembershipInput {
  secret: Int
  Cause: Int!
  Agent: Int!
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
  Layout: Int!
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
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type LayoutColumn {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutRow: Int!
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
  ComponentType: Int!
  LayoutColumn: Int
}
​
type Field {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  FieldType: Int!
  Component: Int
}
​
type Layout {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LayoutType: Int
}
​
type LandingPage {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Cause: Int!
  Layout: Int
}
​
type Experiment {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LandingPage: Int
}
​
type Issue {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: Int!
}
​
type Candidate {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Election: Int!
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
  DistrictType: Int!
}
​
type Office {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Election: Int
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
  CustomerOrder: Int!
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
  Cause: Int!
  Boycott: Int!
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
  Cause: Int!
  Election: Int!
}
​
type PetitionMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  Petition: Int!
}
​
type PollMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  Petition: Int!
}
​
type VolunteerOpportunityMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  VolunteerOpportunity: Int!
}
​
type LiveEventMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  LiveEvent: Int!
}
​
type Product {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ProductType: Int!
}
​
type ProductMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  Product: Int!
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
  Cause: Int!
  DonationCampaign: Int!
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
  Petition: Int!
  Contact: Int!
  Cause: Int!
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
  Poll: Int!
  Contact: Int!
  Cause: Int!
}
​
type Purchaser {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: Int!
  Contact: Int!
  Cause: Int!
}
​
type CustomerOrder {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerCart: Int!
}
​
type Donor {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  CustomerOrder: Int!
  Contact: Int!
  Cause: Int!
}
​
type LiveEvent {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  LiveEventType: Int!
}
​
type EventAttendee {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  LiveEvent: Int!
  Contact: Int!
  Cause: Int!
}
​
type Voter {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Int!
  Cause: Int!
}
​
type VolunteerOpportunity {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  ElectionType: Int
}
​
type Volunteer {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  VolunteerOpportunity: Int!
  Contact: Int!
  Cause: Int!
}
​
type Territory {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
}
​
type Follower {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Int!
  Cause: Int!
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
  ActivityType: Int!
  Contact: Int!
  Cause: Int!
}
​
type Note {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Contact: Int!
  Cause: Int!
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
  Cause: Int!
  Account: Int!
}
​
type Contact {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
}
​
type ContactMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  Contact: Int!
}
​
type Cause {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Title: String!
  Summary: String
}
​
type Agent {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Account: Int!
}
​
type AgentMembership {
  ID: ID!
  CreatedAt: Time!
  UpdatedAt: Time!
  Cause: Int!
  Agent: Int!
}
​

  type Query {
    
    getACL(ID: Int): ACL
    listACL: [ACL]
    getMailingAddress(ID: Int): MailingAddress
    listMailingAddress: [MailingAddress]
    getPhoneNumber(ID: Int): PhoneNumber
    listPhoneNumber: [PhoneNumber]
    getEmailAddress(ID: Int): EmailAddress
    listEmailAddress: [EmailAddress]
    getPhoto(ID: Int): Photo
    listPhoto: [Photo]
    getLayoutType(ID: Int): LayoutType
    listLayoutType: [LayoutType]
    getLayoutRow(ID: Int): LayoutRow
    listLayoutRow: [LayoutRow]
    getComponentImplementation(ID: Int): ComponentImplementation
    listComponentImplementation: [ComponentImplementation]
    getComponentType(ID: Int): ComponentType
    listComponentType: [ComponentType]
    getLayoutColumn(ID: Int): LayoutColumn
    listLayoutColumn: [LayoutColumn]
    getFieldType(ID: Int): FieldType
    listFieldType: [FieldType]
    getComponent(ID: Int): Component
    listComponent: [Component]
    getField(ID: Int): Field
    listField: [Field]
    getLayout(ID: Int): Layout
    listLayout: [Layout]
    getLandingPage(ID: Int): LandingPage
    listLandingPage: [LandingPage]
    getExperiment(ID: Int): Experiment
    listExperiment: [Experiment]
    getIssue(ID: Int): Issue
    listIssue: [Issue]
    getCandidate(ID: Int): Candidate
    listCandidate: [Candidate]
    getDistrictType(ID: Int): DistrictType
    listDistrictType: [DistrictType]
    getDistrict(ID: Int): District
    listDistrict: [District]
    getOffice(ID: Int): Office
    listOffice: [Office]
    getVolunteerOpportunityType(ID: Int): VolunteerOpportunityType
    listVolunteerOpportunityType: [VolunteerOpportunityType]
    getLiveEventType(ID: Int): LiveEventType
    listLiveEventType: [LiveEventType]
    getCompany(ID: Int): Company
    listCompany: [Company]
    getProductType(ID: Int): ProductType
    listProductType: [ProductType]
    getCustomerCart(ID: Int): CustomerCart
    listCustomerCart: [CustomerCart]
    getPayment(ID: Int): Payment
    listPayment: [Payment]
    getDelivery(ID: Int): Delivery
    listDelivery: [Delivery]
    getBoycott(ID: Int): Boycott
    listBoycott: [Boycott]
    getBoycottMembership(ID: Int): BoycottMembership
    listBoycottMembership: [BoycottMembership]
    getElection(ID: Int): Election
    listElection: [Election]
    getElectionMembership(ID: Int): ElectionMembership
    listElectionMembership: [ElectionMembership]
    getPetitionMembership(ID: Int): PetitionMembership
    listPetitionMembership: [PetitionMembership]
    getPollMembership(ID: Int): PollMembership
    listPollMembership: [PollMembership]
    getVolunteerOpportunityMembership(ID: Int): VolunteerOpportunityMembership
    listVolunteerOpportunityMembership: [VolunteerOpportunityMembership]
    getLiveEventMembership(ID: Int): LiveEventMembership
    listLiveEventMembership: [LiveEventMembership]
    getProduct(ID: Int): Product
    listProduct: [Product]
    getProductMembership(ID: Int): ProductMembership
    listProductMembership: [ProductMembership]
    getDonationCampaign(ID: Int): DonationCampaign
    listDonationCampaign: [DonationCampaign]
    getDonationCampaignMembership(ID: Int): DonationCampaignMembership
    listDonationCampaignMembership: [DonationCampaignMembership]
    getPetition(ID: Int): Petition
    listPetition: [Petition]
    getPetitionSigner(ID: Int): PetitionSigner
    listPetitionSigner: [PetitionSigner]
    getPoll(ID: Int): Poll
    listPoll: [Poll]
    getPollRespondant(ID: Int): PollRespondant
    listPollRespondant: [PollRespondant]
    getPurchaser(ID: Int): Purchaser
    listPurchaser: [Purchaser]
    getCustomerOrder(ID: Int): CustomerOrder
    listCustomerOrder: [CustomerOrder]
    getDonor(ID: Int): Donor
    listDonor: [Donor]
    getLiveEvent(ID: Int): LiveEvent
    listLiveEvent: [LiveEvent]
    getEventAttendee(ID: Int): EventAttendee
    listEventAttendee: [EventAttendee]
    getVoter(ID: Int): Voter
    listVoter: [Voter]
    getVolunteerOpportunity(ID: Int): VolunteerOpportunity
    listVolunteerOpportunity: [VolunteerOpportunity]
    getVolunteer(ID: Int): Volunteer
    listVolunteer: [Volunteer]
    getTerritory(ID: Int): Territory
    listTerritory: [Territory]
    getFollower(ID: Int): Follower
    listFollower: [Follower]
    getActivityType(ID: Int): ActivityType
    listActivityType: [ActivityType]
    getActivity(ID: Int): Activity
    listActivity: [Activity]
    getNote(ID: Int): Note
    listNote: [Note]
    getAccount(ID: Int): Account
    listAccount: [Account]
    getOwnerMembership(ID: Int): OwnerMembership
    listOwnerMembership: [OwnerMembership]
    getContact(ID: Int): Contact
    listContact: [Contact]
    getContactMembership(ID: Int): ContactMembership
    listContactMembership: [ContactMembership]
    getCause(ID: Int): Cause
    listCause: [Cause]
    getAgent(ID: Int): Agent
    listAgent: [Agent]
    getAgentMembership(ID: Int): AgentMembership
    listAgentMembership: [AgentMembership]
  }

  type Mutation {
    
    createACL(acl: ACLInput): ACL
    updateACL(ID: Int, acl: ACLInput): Int
    deleteACL(ID: Int): Int
    createMailingAddress(mailingAddress: MailingAddressInput): MailingAddress
    updateMailingAddress(ID: Int, mailingAddress: MailingAddressInput): Int
    deleteMailingAddress(ID: Int): Int
    createPhoneNumber(phoneNumber: PhoneNumberInput): PhoneNumber
    updatePhoneNumber(ID: Int, phoneNumber: PhoneNumberInput): Int
    deletePhoneNumber(ID: Int): Int
    createEmailAddress(emailAddress: EmailAddressInput): EmailAddress
    updateEmailAddress(ID: Int, emailAddress: EmailAddressInput): Int
    deleteEmailAddress(ID: Int): Int
    createPhoto(photo: PhotoInput): Photo
    updatePhoto(ID: Int, photo: PhotoInput): Int
    deletePhoto(ID: Int): Int
    createLayoutType(layoutType: LayoutTypeInput): LayoutType
    updateLayoutType(ID: Int, layoutType: LayoutTypeInput): Int
    deleteLayoutType(ID: Int): Int
    createLayoutRow(layoutRow: LayoutRowInput): LayoutRow
    updateLayoutRow(ID: Int, layoutRow: LayoutRowInput): Int
    deleteLayoutRow(ID: Int): Int
    createComponentImplementation(componentImplementation: ComponentImplementationInput): ComponentImplementation
    updateComponentImplementation(ID: Int, componentImplementation: ComponentImplementationInput): Int
    deleteComponentImplementation(ID: Int): Int
    createComponentType(componentType: ComponentTypeInput): ComponentType
    updateComponentType(ID: Int, componentType: ComponentTypeInput): Int
    deleteComponentType(ID: Int): Int
    createLayoutColumn(layoutColumn: LayoutColumnInput): LayoutColumn
    updateLayoutColumn(ID: Int, layoutColumn: LayoutColumnInput): Int
    deleteLayoutColumn(ID: Int): Int
    createFieldType(fieldType: FieldTypeInput): FieldType
    updateFieldType(ID: Int, fieldType: FieldTypeInput): Int
    deleteFieldType(ID: Int): Int
    createComponent(component: ComponentInput): Component
    updateComponent(ID: Int, component: ComponentInput): Int
    deleteComponent(ID: Int): Int
    createField(field: FieldInput): Field
    updateField(ID: Int, field: FieldInput): Int
    deleteField(ID: Int): Int
    createLayout(layout: LayoutInput): Layout
    updateLayout(ID: Int, layout: LayoutInput): Int
    deleteLayout(ID: Int): Int
    createLandingPage(landingPage: LandingPageInput): LandingPage
    updateLandingPage(ID: Int, landingPage: LandingPageInput): Int
    deleteLandingPage(ID: Int): Int
    createExperiment(experiment: ExperimentInput): Experiment
    updateExperiment(ID: Int, experiment: ExperimentInput): Int
    deleteExperiment(ID: Int): Int
    createIssue(issue: IssueInput): Issue
    updateIssue(ID: Int, issue: IssueInput): Int
    deleteIssue(ID: Int): Int
    createCandidate(candidate: CandidateInput): Candidate
    updateCandidate(ID: Int, candidate: CandidateInput): Int
    deleteCandidate(ID: Int): Int
    createDistrictType(districtType: DistrictTypeInput): DistrictType
    updateDistrictType(ID: Int, districtType: DistrictTypeInput): Int
    deleteDistrictType(ID: Int): Int
    createDistrict(district: DistrictInput): District
    updateDistrict(ID: Int, district: DistrictInput): Int
    deleteDistrict(ID: Int): Int
    createOffice(office: OfficeInput): Office
    updateOffice(ID: Int, office: OfficeInput): Int
    deleteOffice(ID: Int): Int
    createVolunteerOpportunityType(volunteerOpportunityType: VolunteerOpportunityTypeInput): VolunteerOpportunityType
    updateVolunteerOpportunityType(ID: Int, volunteerOpportunityType: VolunteerOpportunityTypeInput): Int
    deleteVolunteerOpportunityType(ID: Int): Int
    createLiveEventType(liveEventType: LiveEventTypeInput): LiveEventType
    updateLiveEventType(ID: Int, liveEventType: LiveEventTypeInput): Int
    deleteLiveEventType(ID: Int): Int
    createCompany(company: CompanyInput): Company
    updateCompany(ID: Int, company: CompanyInput): Int
    deleteCompany(ID: Int): Int
    createProductType(productType: ProductTypeInput): ProductType
    updateProductType(ID: Int, productType: ProductTypeInput): Int
    deleteProductType(ID: Int): Int
    createCustomerCart(customerCart: CustomerCartInput): CustomerCart
    updateCustomerCart(ID: Int, customerCart: CustomerCartInput): Int
    deleteCustomerCart(ID: Int): Int
    createPayment(payment: PaymentInput): Payment
    updatePayment(ID: Int, payment: PaymentInput): Int
    deletePayment(ID: Int): Int
    createDelivery(delivery: DeliveryInput): Delivery
    updateDelivery(ID: Int, delivery: DeliveryInput): Int
    deleteDelivery(ID: Int): Int
    createBoycott(boycott: BoycottInput): Boycott
    updateBoycott(ID: Int, boycott: BoycottInput): Int
    deleteBoycott(ID: Int): Int
    createBoycottMembership(boycottMembership: BoycottMembershipInput): BoycottMembership
    updateBoycottMembership(ID: Int, boycottMembership: BoycottMembershipInput): Int
    deleteBoycottMembership(ID: Int): Int
    createElection(election: ElectionInput): Election
    updateElection(ID: Int, election: ElectionInput): Int
    deleteElection(ID: Int): Int
    createElectionMembership(electionMembership: ElectionMembershipInput): ElectionMembership
    updateElectionMembership(ID: Int, electionMembership: ElectionMembershipInput): Int
    deleteElectionMembership(ID: Int): Int
    createPetitionMembership(petitionMembership: PetitionMembershipInput): PetitionMembership
    updatePetitionMembership(ID: Int, petitionMembership: PetitionMembershipInput): Int
    deletePetitionMembership(ID: Int): Int
    createPollMembership(pollMembership: PollMembershipInput): PollMembership
    updatePollMembership(ID: Int, pollMembership: PollMembershipInput): Int
    deletePollMembership(ID: Int): Int
    createVolunteerOpportunityMembership(volunteerOpportunityMembership: VolunteerOpportunityMembershipInput): VolunteerOpportunityMembership
    updateVolunteerOpportunityMembership(ID: Int, volunteerOpportunityMembership: VolunteerOpportunityMembershipInput): Int
    deleteVolunteerOpportunityMembership(ID: Int): Int
    createLiveEventMembership(liveEventMembership: LiveEventMembershipInput): LiveEventMembership
    updateLiveEventMembership(ID: Int, liveEventMembership: LiveEventMembershipInput): Int
    deleteLiveEventMembership(ID: Int): Int
    createProduct(product: ProductInput): Product
    updateProduct(ID: Int, product: ProductInput): Int
    deleteProduct(ID: Int): Int
    createProductMembership(productMembership: ProductMembershipInput): ProductMembership
    updateProductMembership(ID: Int, productMembership: ProductMembershipInput): Int
    deleteProductMembership(ID: Int): Int
    createDonationCampaign(donationCampaign: DonationCampaignInput): DonationCampaign
    updateDonationCampaign(ID: Int, donationCampaign: DonationCampaignInput): Int
    deleteDonationCampaign(ID: Int): Int
    createDonationCampaignMembership(donationCampaignMembership: DonationCampaignMembershipInput): DonationCampaignMembership
    updateDonationCampaignMembership(ID: Int, donationCampaignMembership: DonationCampaignMembershipInput): Int
    deleteDonationCampaignMembership(ID: Int): Int
    createPetition(petition: PetitionInput): Petition
    updatePetition(ID: Int, petition: PetitionInput): Int
    deletePetition(ID: Int): Int
    createPetitionSigner(petitionSigner: PetitionSignerInput): PetitionSigner
    updatePetitionSigner(ID: Int, petitionSigner: PetitionSignerInput): Int
    deletePetitionSigner(ID: Int): Int
    createPoll(poll: PollInput): Poll
    updatePoll(ID: Int, poll: PollInput): Int
    deletePoll(ID: Int): Int
    createPollRespondant(pollRespondant: PollRespondantInput): PollRespondant
    updatePollRespondant(ID: Int, pollRespondant: PollRespondantInput): Int
    deletePollRespondant(ID: Int): Int
    createPurchaser(purchaser: PurchaserInput): Purchaser
    updatePurchaser(ID: Int, purchaser: PurchaserInput): Int
    deletePurchaser(ID: Int): Int
    createCustomerOrder(customerOrder: CustomerOrderInput): CustomerOrder
    updateCustomerOrder(ID: Int, customerOrder: CustomerOrderInput): Int
    deleteCustomerOrder(ID: Int): Int
    createDonor(donor: DonorInput): Donor
    updateDonor(ID: Int, donor: DonorInput): Int
    deleteDonor(ID: Int): Int
    createLiveEvent(liveEvent: LiveEventInput): LiveEvent
    updateLiveEvent(ID: Int, liveEvent: LiveEventInput): Int
    deleteLiveEvent(ID: Int): Int
    createEventAttendee(eventAttendee: EventAttendeeInput): EventAttendee
    updateEventAttendee(ID: Int, eventAttendee: EventAttendeeInput): Int
    deleteEventAttendee(ID: Int): Int
    createVoter(voter: VoterInput): Voter
    updateVoter(ID: Int, voter: VoterInput): Int
    deleteVoter(ID: Int): Int
    createVolunteerOpportunity(volunteerOpportunity: VolunteerOpportunityInput): VolunteerOpportunity
    updateVolunteerOpportunity(ID: Int, volunteerOpportunity: VolunteerOpportunityInput): Int
    deleteVolunteerOpportunity(ID: Int): Int
    createVolunteer(volunteer: VolunteerInput): Volunteer
    updateVolunteer(ID: Int, volunteer: VolunteerInput): Int
    deleteVolunteer(ID: Int): Int
    createTerritory(territory: TerritoryInput): Territory
    updateTerritory(ID: Int, territory: TerritoryInput): Int
    deleteTerritory(ID: Int): Int
    createFollower(follower: FollowerInput): Follower
    updateFollower(ID: Int, follower: FollowerInput): Int
    deleteFollower(ID: Int): Int
    createActivityType(activityType: ActivityTypeInput): ActivityType
    updateActivityType(ID: Int, activityType: ActivityTypeInput): Int
    deleteActivityType(ID: Int): Int
    createActivity(activity: ActivityInput): Activity
    updateActivity(ID: Int, activity: ActivityInput): Int
    deleteActivity(ID: Int): Int
    createNote(note: NoteInput): Note
    updateNote(ID: Int, note: NoteInput): Int
    deleteNote(ID: Int): Int
    createAccount(account: AccountInput): Account
    updateAccount(ID: Int, account: AccountInput): Int
    deleteAccount(ID: Int): Int
    createOwnerMembership(ownerMembership: OwnerMembershipInput): OwnerMembership
    updateOwnerMembership(ID: Int, ownerMembership: OwnerMembershipInput): Int
    deleteOwnerMembership(ID: Int): Int
    createContact(contact: ContactInput): Contact
    updateContact(ID: Int, contact: ContactInput): Int
    deleteContact(ID: Int): Int
    createContactMembership(contactMembership: ContactMembershipInput): ContactMembership
    updateContactMembership(ID: Int, contactMembership: ContactMembershipInput): Int
    deleteContactMembership(ID: Int): Int
    createCause(cause: CauseInput): Cause
    updateCause(ID: Int, cause: CauseInput): Int
    deleteCause(ID: Int): Int
    createAgent(agent: AgentInput): Agent
    updateAgent(ID: Int, agent: AgentInput): Int
    deleteAgent(ID: Int): Int
    createAgentMembership(agentMembership: AgentMembershipInput): AgentMembership
    updateAgentMembership(ID: Int, agentMembership: AgentMembershipInput): Int
    deleteAgentMembership(ID: Int): Int
  }
  `.replace(/[\u200B-\u200D\uFEFF]/g, '');