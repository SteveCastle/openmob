module.exports = client => ({
  ACL: {},
  MailingAddress: {},
  PhoneNumber: {},
  EmailAddress: {},
  Photo: {},
  LayoutType: {
    Layouts: ({ ID }) =>
      client
        .ListLayout()
        .sendMessage({ api: 'v1', filters: [{ LayoutType: ID }], limit: 10 })
        .then(res => res.items)
  },
  LayoutRow: {
    LayoutColumns: ({ ID }) =>
      client
        .ListLayoutColumn()
        .sendMessage({ api: 'v1', filters: [{ LayoutRow: ID }], limit: 10 })
        .then(res => res.items),
    Layout: ({ Layout }) =>
      client
        .GetLayout()
        .sendMessage({ api: 'v1', ID: Layout })
        .then(res => res.item)
  },
  ComponentImplementation: {},
  ComponentType: {
    Components: ({ ID }) =>
      client
        .ListComponent()
        .sendMessage({ api: 'v1', filters: [{ ComponentType: ID }], limit: 10 })
        .then(res => res.items)
  },
  LayoutColumn: {
    Components: ({ ID }) =>
      client
        .ListComponent()
        .sendMessage({ api: 'v1', filters: [{ LayoutColumn: ID }], limit: 10 })
        .then(res => res.items),
    LayoutRow: ({ LayoutRow }) =>
      client
        .GetLayoutRow()
        .sendMessage({ api: 'v1', ID: LayoutRow })
        .then(res => res.item)
  },
  FieldType: {
    Fields: ({ ID }) =>
      client
        .ListField()
        .sendMessage({ api: 'v1', filters: [{ FieldType: ID }], limit: 10 })
        .then(res => res.items)
  },
  Component: {
    Fields: ({ ID }) =>
      client
        .ListField()
        .sendMessage({ api: 'v1', filters: [{ Component: ID }], limit: 10 })
        .then(res => res.items),
    ComponentType: ({ ComponentType }) =>
      client
        .GetComponentType()
        .sendMessage({ api: 'v1', ID: ComponentType })
        .then(res => res.item),
    LayoutColumn: ({ LayoutColumn }) =>
      client
        .GetLayoutColumn()
        .sendMessage({ api: 'v1', ID: LayoutColumn })
        .then(res => res.item)
  },
  Field: {
    FieldType: ({ FieldType }) =>
      client
        .GetFieldType()
        .sendMessage({ api: 'v1', ID: FieldType })
        .then(res => res.item),
    Component: ({ Component }) =>
      client
        .GetComponent()
        .sendMessage({ api: 'v1', ID: Component })
        .then(res => res.item)
  },
  HomePage: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Layout: ({ Layout }) =>
      client
        .GetLayout()
        .sendMessage({ api: 'v1', ID: Layout })
        .then(res => res.item)
  },
  Layout: {
    LayoutRows: ({ ID }) =>
      client
        .ListLayoutRow()
        .sendMessage({ api: 'v1', filters: [{ Layout: ID }], limit: 10 })
        .then(res => res.items),
    HomePages: ({ ID }) =>
      client
        .ListHomePage()
        .sendMessage({ api: 'v1', filters: [{ Layout: ID }], limit: 10 })
        .then(res => res.items),
    LandingPages: ({ ID }) =>
      client
        .ListLandingPage()
        .sendMessage({ api: 'v1', filters: [{ Layout: ID }], limit: 10 })
        .then(res => res.items),
    LayoutType: ({ LayoutType }) =>
      client
        .GetLayoutType()
        .sendMessage({ api: 'v1', ID: LayoutType })
        .then(res => res.item)
  },
  LandingPage: {
    Experiments: ({ ID }) =>
      client
        .ListExperiment()
        .sendMessage({ api: 'v1', filters: [{ LandingPage: ID }], limit: 10 })
        .then(res => res.items),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Layout: ({ Layout }) =>
      client
        .GetLayout()
        .sendMessage({ api: 'v1', ID: Layout })
        .then(res => res.item)
  },
  Experiment: {
    LandingPage: ({ LandingPage }) =>
      client
        .GetLandingPage()
        .sendMessage({ api: 'v1', ID: LandingPage })
        .then(res => res.item)
  },
  Issue: {
    Election: ({ Election }) =>
      client
        .GetElection()
        .sendMessage({ api: 'v1', ID: Election })
        .then(res => res.item)
  },
  Candidate: {
    Election: ({ Election }) =>
      client
        .GetElection()
        .sendMessage({ api: 'v1', ID: Election })
        .then(res => res.item)
  },
  DistrictType: {
    Districts: ({ ID }) =>
      client
        .ListDistrict()
        .sendMessage({ api: 'v1', filters: [{ DistrictType: ID }], limit: 10 })
        .then(res => res.items)
  },
  District: {
    DistrictType: ({ DistrictType }) =>
      client
        .GetDistrictType()
        .sendMessage({ api: 'v1', ID: DistrictType })
        .then(res => res.item)
  },
  Office: {
    Election: ({ Election }) =>
      client
        .GetElection()
        .sendMessage({ api: 'v1', ID: Election })
        .then(res => res.item)
  },
  PollItem: {
    Poll: ({ Poll }) =>
      client
        .GetPoll()
        .sendMessage({ api: 'v1', ID: Poll })
        .then(res => res.item)
  },
  VolunteerOpportunityType: {
    VolunteerOpportunitys: ({ ID }) =>
      client
        .ListVolunteerOpportunity()
        .sendMessage({
          api: 'v1',
          filters: [{ VolunteerOpportunityType: ID }],
          limit: 10
        })
        .then(res => res.items)
  },
  LiveEventType: {
    LiveEvents: ({ ID }) =>
      client
        .ListLiveEvent()
        .sendMessage({ api: 'v1', filters: [{ LiveEventType: ID }], limit: 10 })
        .then(res => res.items)
  },
  Company: {},
  ProductType: {
    Products: ({ ID }) =>
      client
        .ListProduct()
        .sendMessage({ api: 'v1', filters: [{ ProductType: ID }], limit: 10 })
        .then(res => res.items)
  },
  CustomerCart: {
    CustomerOrders: ({ ID }) =>
      client
        .ListCustomerOrder()
        .sendMessage({ api: 'v1', filters: [{ CustomerCart: ID }], limit: 10 })
        .then(res => res.items)
  },
  Payment: {
    CustomerOrder: ({ CustomerOrder }) =>
      client
        .GetCustomerOrder()
        .sendMessage({ api: 'v1', ID: CustomerOrder })
        .then(res => res.item)
  },
  Delivery: {},
  Boycott: {
    BoycottMemberships: ({ ID }) =>
      client
        .ListBoycottMembership()
        .sendMessage({ api: 'v1', filters: [{ Boycott: ID }], limit: 10 })
        .then(res => res.items)
  },
  BoycottMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Boycott: ({ Boycott }) =>
      client
        .GetBoycott()
        .sendMessage({ api: 'v1', ID: Boycott })
        .then(res => res.item)
  },
  Election: {
    Issues: ({ ID }) =>
      client
        .ListIssue()
        .sendMessage({ api: 'v1', filters: [{ Election: ID }], limit: 10 })
        .then(res => res.items),
    Candidates: ({ ID }) =>
      client
        .ListCandidate()
        .sendMessage({ api: 'v1', filters: [{ Election: ID }], limit: 10 })
        .then(res => res.items),
    Offices: ({ ID }) =>
      client
        .ListOffice()
        .sendMessage({ api: 'v1', filters: [{ Election: ID }], limit: 10 })
        .then(res => res.items),
    ElectionMemberships: ({ ID }) =>
      client
        .ListElectionMembership()
        .sendMessage({ api: 'v1', filters: [{ Election: ID }], limit: 10 })
        .then(res => res.items)
  },
  ElectionMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Election: ({ Election }) =>
      client
        .GetElection()
        .sendMessage({ api: 'v1', ID: Election })
        .then(res => res.item)
  },
  PetitionMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Petition: ({ Petition }) =>
      client
        .GetPetition()
        .sendMessage({ api: 'v1', ID: Petition })
        .then(res => res.item)
  },
  PollMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Petition: ({ Petition }) =>
      client
        .GetPetition()
        .sendMessage({ api: 'v1', ID: Petition })
        .then(res => res.item)
  },
  VolunteerOpportunityMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    VolunteerOpportunity: ({ VolunteerOpportunity }) =>
      client
        .GetVolunteerOpportunity()
        .sendMessage({ api: 'v1', ID: VolunteerOpportunity })
        .then(res => res.item)
  },
  LiveEventMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    LiveEvent: ({ LiveEvent }) =>
      client
        .GetLiveEvent()
        .sendMessage({ api: 'v1', ID: LiveEvent })
        .then(res => res.item)
  },
  Product: {
    ProductMemberships: ({ ID }) =>
      client
        .ListProductMembership()
        .sendMessage({ api: 'v1', filters: [{ Product: ID }], limit: 10 })
        .then(res => res.items),
    ProductType: ({ ProductType }) =>
      client
        .GetProductType()
        .sendMessage({ api: 'v1', ID: ProductType })
        .then(res => res.item)
  },
  ProductMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Product: ({ Product }) =>
      client
        .GetProduct()
        .sendMessage({ api: 'v1', ID: Product })
        .then(res => res.item)
  },
  DonationCampaign: {
    DonationCampaignMemberships: ({ ID }) =>
      client
        .ListDonationCampaignMembership()
        .sendMessage({
          api: 'v1',
          filters: [{ DonationCampaign: ID }],
          limit: 10
        })
        .then(res => res.items)
  },
  DonationCampaignMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    DonationCampaign: ({ DonationCampaign }) =>
      client
        .GetDonationCampaign()
        .sendMessage({ api: 'v1', ID: DonationCampaign })
        .then(res => res.item)
  },
  Petition: {
    PetitionMemberships: ({ ID }) =>
      client
        .ListPetitionMembership()
        .sendMessage({ api: 'v1', filters: [{ Petition: ID }], limit: 10 })
        .then(res => res.items),
    PollMemberships: ({ ID }) =>
      client
        .ListPollMembership()
        .sendMessage({ api: 'v1', filters: [{ Petition: ID }], limit: 10 })
        .then(res => res.items),
    PetitionSigners: ({ ID }) =>
      client
        .ListPetitionSigner()
        .sendMessage({ api: 'v1', filters: [{ Petition: ID }], limit: 10 })
        .then(res => res.items)
  },
  PetitionSigner: {
    Petition: ({ Petition }) =>
      client
        .GetPetition()
        .sendMessage({ api: 'v1', ID: Petition })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Poll: {
    PollItems: ({ ID }) =>
      client
        .ListPollItem()
        .sendMessage({ api: 'v1', filters: [{ Poll: ID }], limit: 10 })
        .then(res => res.items),
    PollRespondants: ({ ID }) =>
      client
        .ListPollRespondant()
        .sendMessage({ api: 'v1', filters: [{ Poll: ID }], limit: 10 })
        .then(res => res.items)
  },
  PollRespondant: {
    Poll: ({ Poll }) =>
      client
        .GetPoll()
        .sendMessage({ api: 'v1', ID: Poll })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Purchaser: {
    CustomerOrder: ({ CustomerOrder }) =>
      client
        .GetCustomerOrder()
        .sendMessage({ api: 'v1', ID: CustomerOrder })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  CustomerOrder: {
    Payments: ({ ID }) =>
      client
        .ListPayment()
        .sendMessage({ api: 'v1', filters: [{ CustomerOrder: ID }], limit: 10 })
        .then(res => res.items),
    Purchasers: ({ ID }) =>
      client
        .ListPurchaser()
        .sendMessage({ api: 'v1', filters: [{ CustomerOrder: ID }], limit: 10 })
        .then(res => res.items),
    Donors: ({ ID }) =>
      client
        .ListDonor()
        .sendMessage({ api: 'v1', filters: [{ CustomerOrder: ID }], limit: 10 })
        .then(res => res.items),
    CustomerCart: ({ CustomerCart }) =>
      client
        .GetCustomerCart()
        .sendMessage({ api: 'v1', ID: CustomerCart })
        .then(res => res.item)
  },
  Donor: {
    CustomerOrder: ({ CustomerOrder }) =>
      client
        .GetCustomerOrder()
        .sendMessage({ api: 'v1', ID: CustomerOrder })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  LiveEvent: {
    LiveEventMemberships: ({ ID }) =>
      client
        .ListLiveEventMembership()
        .sendMessage({ api: 'v1', filters: [{ LiveEvent: ID }], limit: 10 })
        .then(res => res.items),
    EventAttendees: ({ ID }) =>
      client
        .ListEventAttendee()
        .sendMessage({ api: 'v1', filters: [{ LiveEvent: ID }], limit: 10 })
        .then(res => res.items),
    LiveEventType: ({ LiveEventType }) =>
      client
        .GetLiveEventType()
        .sendMessage({ api: 'v1', ID: LiveEventType })
        .then(res => res.item)
  },
  EventAttendee: {
    LiveEvent: ({ LiveEvent }) =>
      client
        .GetLiveEvent()
        .sendMessage({ api: 'v1', ID: LiveEvent })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Voter: {
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  VolunteerOpportunity: {
    VolunteerOpportunityMemberships: ({ ID }) =>
      client
        .ListVolunteerOpportunityMembership()
        .sendMessage({
          api: 'v1',
          filters: [{ VolunteerOpportunity: ID }],
          limit: 10
        })
        .then(res => res.items),
    Volunteers: ({ ID }) =>
      client
        .ListVolunteer()
        .sendMessage({
          api: 'v1',
          filters: [{ VolunteerOpportunity: ID }],
          limit: 10
        })
        .then(res => res.items),
    VolunteerOpportunityType: ({ VolunteerOpportunityType }) =>
      client
        .GetVolunteerOpportunityType()
        .sendMessage({ api: 'v1', ID: VolunteerOpportunityType })
        .then(res => res.item)
  },
  Volunteer: {
    VolunteerOpportunity: ({ VolunteerOpportunity }) =>
      client
        .GetVolunteerOpportunity()
        .sendMessage({ api: 'v1', ID: VolunteerOpportunity })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Follower: {
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Territory: {},
  ActivityType: {
    Activitys: ({ ID }) =>
      client
        .ListActivity()
        .sendMessage({ api: 'v1', filters: [{ ActivityType: ID }], limit: 10 })
        .then(res => res.items)
  },
  Activity: {
    ActivityType: ({ ActivityType }) =>
      client
        .GetActivityType()
        .sendMessage({ api: 'v1', ID: ActivityType })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Note: {
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item),
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item)
  },
  Account: {
    OwnerMemberships: ({ ID }) =>
      client
        .ListOwnerMembership()
        .sendMessage({ api: 'v1', filters: [{ Account: ID }], limit: 10 })
        .then(res => res.items),
    Agents: ({ ID }) =>
      client
        .ListAgent()
        .sendMessage({ api: 'v1', filters: [{ Account: ID }], limit: 10 })
        .then(res => res.items)
  },
  OwnerMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Account: ({ Account }) =>
      client
        .GetAccount()
        .sendMessage({ api: 'v1', ID: Account })
        .then(res => res.item)
  },
  Contact: {
    PetitionSigners: ({ ID }) =>
      client
        .ListPetitionSigner()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    PollRespondants: ({ ID }) =>
      client
        .ListPollRespondant()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Purchasers: ({ ID }) =>
      client
        .ListPurchaser()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Donors: ({ ID }) =>
      client
        .ListDonor()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    EventAttendees: ({ ID }) =>
      client
        .ListEventAttendee()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Voters: ({ ID }) =>
      client
        .ListVoter()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Volunteers: ({ ID }) =>
      client
        .ListVolunteer()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Followers: ({ ID }) =>
      client
        .ListFollower()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Activitys: ({ ID }) =>
      client
        .ListActivity()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    Notes: ({ ID }) =>
      client
        .ListNote()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items),
    ContactMemberships: ({ ID }) =>
      client
        .ListContactMembership()
        .sendMessage({ api: 'v1', filters: [{ Contact: ID }], limit: 10 })
        .then(res => res.items)
  },
  ContactMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Contact: ({ Contact }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID: Contact })
        .then(res => res.item)
  },
  Cause: {
    HomePages: ({ ID }) =>
      client
        .ListHomePage()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    LandingPages: ({ ID }) =>
      client
        .ListLandingPage()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    BoycottMemberships: ({ ID }) =>
      client
        .ListBoycottMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    ElectionMemberships: ({ ID }) =>
      client
        .ListElectionMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    PetitionMemberships: ({ ID }) =>
      client
        .ListPetitionMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    PollMemberships: ({ ID }) =>
      client
        .ListPollMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    VolunteerOpportunityMemberships: ({ ID }) =>
      client
        .ListVolunteerOpportunityMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    LiveEventMemberships: ({ ID }) =>
      client
        .ListLiveEventMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    ProductMemberships: ({ ID }) =>
      client
        .ListProductMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    DonationCampaignMemberships: ({ ID }) =>
      client
        .ListDonationCampaignMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    PetitionSigners: ({ ID }) =>
      client
        .ListPetitionSigner()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    PollRespondants: ({ ID }) =>
      client
        .ListPollRespondant()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Purchasers: ({ ID }) =>
      client
        .ListPurchaser()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Donors: ({ ID }) =>
      client
        .ListDonor()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    EventAttendees: ({ ID }) =>
      client
        .ListEventAttendee()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Voters: ({ ID }) =>
      client
        .ListVoter()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Volunteers: ({ ID }) =>
      client
        .ListVolunteer()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Followers: ({ ID }) =>
      client
        .ListFollower()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Activitys: ({ ID }) =>
      client
        .ListActivity()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    Notes: ({ ID }) =>
      client
        .ListNote()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    OwnerMemberships: ({ ID }) =>
      client
        .ListOwnerMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    ContactMemberships: ({ ID }) =>
      client
        .ListContactMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items),
    AgentMemberships: ({ ID }) =>
      client
        .ListAgentMembership()
        .sendMessage({ api: 'v1', filters: [{ Cause: ID }], limit: 10 })
        .then(res => res.items)
  },
  Agent: {
    AgentMemberships: ({ ID }) =>
      client
        .ListAgentMembership()
        .sendMessage({ api: 'v1', filters: [{ Agent: ID }], limit: 10 })
        .then(res => res.items),
    Account: ({ Account }) =>
      client
        .GetAccount()
        .sendMessage({ api: 'v1', ID: Account })
        .then(res => res.item)
  },
  AgentMembership: {
    Cause: ({ Cause }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID: Cause })
        .then(res => res.item),
    Agent: ({ Agent }) =>
      client
        .GetAgent()
        .sendMessage({ api: 'v1', ID: Agent })
        .then(res => res.item)
  },
  Query: {
    listACL: (_, { limit, ordering, filters }) =>
      client
        .ListACL()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getACL: (_, { ID }) =>
      client
        .GetACL()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listMailingAddress: (_, { limit, ordering, filters }) =>
      client
        .ListMailingAddress()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getMailingAddress: (_, { ID }) =>
      client
        .GetMailingAddress()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPhoneNumber: (_, { limit, ordering, filters }) =>
      client
        .ListPhoneNumber()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPhoneNumber: (_, { ID }) =>
      client
        .GetPhoneNumber()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listEmailAddress: (_, { limit, ordering, filters }) =>
      client
        .ListEmailAddress()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getEmailAddress: (_, { ID }) =>
      client
        .GetEmailAddress()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPhoto: (_, { limit, ordering, filters }) =>
      client
        .ListPhoto()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPhoto: (_, { ID }) =>
      client
        .GetPhoto()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLayoutType: (_, { limit, ordering, filters }) =>
      client
        .ListLayoutType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLayoutType: (_, { ID }) =>
      client
        .GetLayoutType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLayoutRow: (_, { limit, ordering, filters }) =>
      client
        .ListLayoutRow()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLayoutRow: (_, { ID }) =>
      client
        .GetLayoutRow()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listComponentImplementation: (_, { limit, ordering, filters }) =>
      client
        .ListComponentImplementation()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getComponentImplementation: (_, { ID }) =>
      client
        .GetComponentImplementation()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listComponentType: (_, { limit, ordering, filters }) =>
      client
        .ListComponentType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getComponentType: (_, { ID }) =>
      client
        .GetComponentType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLayoutColumn: (_, { limit, ordering, filters }) =>
      client
        .ListLayoutColumn()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLayoutColumn: (_, { ID }) =>
      client
        .GetLayoutColumn()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listFieldType: (_, { limit, ordering, filters }) =>
      client
        .ListFieldType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getFieldType: (_, { ID }) =>
      client
        .GetFieldType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listComponent: (_, { limit, ordering, filters }) =>
      client
        .ListComponent()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getComponent: (_, { ID }) =>
      client
        .GetComponent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listField: (_, { limit, ordering, filters }) =>
      client
        .ListField()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getField: (_, { ID }) =>
      client
        .GetField()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listHomePage: (_, { limit, ordering, filters }) =>
      client
        .ListHomePage()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getHomePage: (_, { ID }) =>
      client
        .GetHomePage()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLayout: (_, { limit, ordering, filters }) =>
      client
        .ListLayout()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLayout: (_, { ID }) =>
      client
        .GetLayout()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLandingPage: (_, { limit, ordering, filters }) =>
      client
        .ListLandingPage()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLandingPage: (_, { ID }) =>
      client
        .GetLandingPage()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listExperiment: (_, { limit, ordering, filters }) =>
      client
        .ListExperiment()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getExperiment: (_, { ID }) =>
      client
        .GetExperiment()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listIssue: (_, { limit, ordering, filters }) =>
      client
        .ListIssue()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getIssue: (_, { ID }) =>
      client
        .GetIssue()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listCandidate: (_, { limit, ordering, filters }) =>
      client
        .ListCandidate()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getCandidate: (_, { ID }) =>
      client
        .GetCandidate()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDistrictType: (_, { limit, ordering, filters }) =>
      client
        .ListDistrictType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDistrictType: (_, { ID }) =>
      client
        .GetDistrictType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDistrict: (_, { limit, ordering, filters }) =>
      client
        .ListDistrict()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDistrict: (_, { ID }) =>
      client
        .GetDistrict()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listOffice: (_, { limit, ordering, filters }) =>
      client
        .ListOffice()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getOffice: (_, { ID }) =>
      client
        .GetOffice()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPollItem: (_, { limit, ordering, filters }) =>
      client
        .ListPollItem()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPollItem: (_, { ID }) =>
      client
        .GetPollItem()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listVolunteerOpportunityType: (_, { limit, ordering, filters }) =>
      client
        .ListVolunteerOpportunityType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getVolunteerOpportunityType: (_, { ID }) =>
      client
        .GetVolunteerOpportunityType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLiveEventType: (_, { limit, ordering, filters }) =>
      client
        .ListLiveEventType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLiveEventType: (_, { ID }) =>
      client
        .GetLiveEventType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listCompany: (_, { limit, ordering, filters }) =>
      client
        .ListCompany()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getCompany: (_, { ID }) =>
      client
        .GetCompany()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listProductType: (_, { limit, ordering, filters }) =>
      client
        .ListProductType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getProductType: (_, { ID }) =>
      client
        .GetProductType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listCustomerCart: (_, { limit, ordering, filters }) =>
      client
        .ListCustomerCart()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getCustomerCart: (_, { ID }) =>
      client
        .GetCustomerCart()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPayment: (_, { limit, ordering, filters }) =>
      client
        .ListPayment()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPayment: (_, { ID }) =>
      client
        .GetPayment()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDelivery: (_, { limit, ordering, filters }) =>
      client
        .ListDelivery()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDelivery: (_, { ID }) =>
      client
        .GetDelivery()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listBoycott: (_, { limit, ordering, filters }) =>
      client
        .ListBoycott()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getBoycott: (_, { ID }) =>
      client
        .GetBoycott()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listBoycottMembership: (_, { limit, ordering, filters }) =>
      client
        .ListBoycottMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getBoycottMembership: (_, { ID }) =>
      client
        .GetBoycottMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listElection: (_, { limit, ordering, filters }) =>
      client
        .ListElection()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getElection: (_, { ID }) =>
      client
        .GetElection()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listElectionMembership: (_, { limit, ordering, filters }) =>
      client
        .ListElectionMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getElectionMembership: (_, { ID }) =>
      client
        .GetElectionMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPetitionMembership: (_, { limit, ordering, filters }) =>
      client
        .ListPetitionMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPetitionMembership: (_, { ID }) =>
      client
        .GetPetitionMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPollMembership: (_, { limit, ordering, filters }) =>
      client
        .ListPollMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPollMembership: (_, { ID }) =>
      client
        .GetPollMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listVolunteerOpportunityMembership: (_, { limit, ordering, filters }) =>
      client
        .ListVolunteerOpportunityMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getVolunteerOpportunityMembership: (_, { ID }) =>
      client
        .GetVolunteerOpportunityMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLiveEventMembership: (_, { limit, ordering, filters }) =>
      client
        .ListLiveEventMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLiveEventMembership: (_, { ID }) =>
      client
        .GetLiveEventMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listProduct: (_, { limit, ordering, filters }) =>
      client
        .ListProduct()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getProduct: (_, { ID }) =>
      client
        .GetProduct()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listProductMembership: (_, { limit, ordering, filters }) =>
      client
        .ListProductMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getProductMembership: (_, { ID }) =>
      client
        .GetProductMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDonationCampaign: (_, { limit, ordering, filters }) =>
      client
        .ListDonationCampaign()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDonationCampaign: (_, { ID }) =>
      client
        .GetDonationCampaign()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDonationCampaignMembership: (_, { limit, ordering, filters }) =>
      client
        .ListDonationCampaignMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDonationCampaignMembership: (_, { ID }) =>
      client
        .GetDonationCampaignMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPetition: (_, { limit, ordering, filters }) =>
      client
        .ListPetition()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPetition: (_, { ID }) =>
      client
        .GetPetition()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPetitionSigner: (_, { limit, ordering, filters }) =>
      client
        .ListPetitionSigner()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPetitionSigner: (_, { ID }) =>
      client
        .GetPetitionSigner()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPoll: (_, { limit, ordering, filters }) =>
      client
        .ListPoll()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPoll: (_, { ID }) =>
      client
        .GetPoll()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPollRespondant: (_, { limit, ordering, filters }) =>
      client
        .ListPollRespondant()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPollRespondant: (_, { ID }) =>
      client
        .GetPollRespondant()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listPurchaser: (_, { limit, ordering, filters }) =>
      client
        .ListPurchaser()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getPurchaser: (_, { ID }) =>
      client
        .GetPurchaser()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listCustomerOrder: (_, { limit, ordering, filters }) =>
      client
        .ListCustomerOrder()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getCustomerOrder: (_, { ID }) =>
      client
        .GetCustomerOrder()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listDonor: (_, { limit, ordering, filters }) =>
      client
        .ListDonor()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getDonor: (_, { ID }) =>
      client
        .GetDonor()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listLiveEvent: (_, { limit, ordering, filters }) =>
      client
        .ListLiveEvent()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getLiveEvent: (_, { ID }) =>
      client
        .GetLiveEvent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listEventAttendee: (_, { limit, ordering, filters }) =>
      client
        .ListEventAttendee()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getEventAttendee: (_, { ID }) =>
      client
        .GetEventAttendee()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listVoter: (_, { limit, ordering, filters }) =>
      client
        .ListVoter()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getVoter: (_, { ID }) =>
      client
        .GetVoter()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listVolunteerOpportunity: (_, { limit, ordering, filters }) =>
      client
        .ListVolunteerOpportunity()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getVolunteerOpportunity: (_, { ID }) =>
      client
        .GetVolunteerOpportunity()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listVolunteer: (_, { limit, ordering, filters }) =>
      client
        .ListVolunteer()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getVolunteer: (_, { ID }) =>
      client
        .GetVolunteer()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listFollower: (_, { limit, ordering, filters }) =>
      client
        .ListFollower()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getFollower: (_, { ID }) =>
      client
        .GetFollower()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listTerritory: (_, { limit, ordering, filters }) =>
      client
        .ListTerritory()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getTerritory: (_, { ID }) =>
      client
        .GetTerritory()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listActivityType: (_, { limit, ordering, filters }) =>
      client
        .ListActivityType()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getActivityType: (_, { ID }) =>
      client
        .GetActivityType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listActivity: (_, { limit, ordering, filters }) =>
      client
        .ListActivity()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getActivity: (_, { ID }) =>
      client
        .GetActivity()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listNote: (_, { limit, ordering, filters }) =>
      client
        .ListNote()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getNote: (_, { ID }) =>
      client
        .GetNote()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listAccount: (_, { limit, ordering, filters }) =>
      client
        .ListAccount()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getAccount: (_, { ID }) =>
      client
        .GetAccount()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listOwnerMembership: (_, { limit, ordering, filters }) =>
      client
        .ListOwnerMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getOwnerMembership: (_, { ID }) =>
      client
        .GetOwnerMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listContact: (_, { limit, ordering, filters }) =>
      client
        .ListContact()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getContact: (_, { ID }) =>
      client
        .GetContact()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listContactMembership: (_, { limit, ordering, filters }) =>
      client
        .ListContactMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getContactMembership: (_, { ID }) =>
      client
        .GetContactMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listCause: (_, { limit, ordering, filters }) =>
      client
        .ListCause()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getCause: (_, { ID }) =>
      client
        .GetCause()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listAgent: (_, { limit, ordering, filters }) =>
      client
        .ListAgent()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getAgent: (_, { ID }) =>
      client
        .GetAgent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item),
    listAgentMembership: (_, { limit, ordering, filters }) =>
      client
        .ListAgentMembership()
        .sendMessage({ api: 'v1', limit, ordering, filters })
        .then(res => res.items),
    getAgentMembership: (_, { ID }) =>
      client
        .GetAgentMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.item)
  },
  Mutation: {
    createACL: (_, { acl }) =>
      client
        .CreateACL()
        .sendMessage({ api: 'v1', item: { ...acl } })
        .then(res => ({ ID: res.ID, ...acl })),
    updateACL: (_, { ID, acl }) =>
      client
        .UpdateACL()
        .sendMessage({ api: 'v1', item: { ID, ...acl } })
        .then(res => res.updated),
    deleteACL: (_, { ID }) =>
      client
        .DeleteACL()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createMailingAddress: (_, { mailingAddress }) =>
      client
        .CreateMailingAddress()
        .sendMessage({ api: 'v1', item: { ...mailingAddress } })
        .then(res => ({ ID: res.ID, ...mailingAddress })),
    updateMailingAddress: (_, { ID, mailingAddress }) =>
      client
        .UpdateMailingAddress()
        .sendMessage({ api: 'v1', item: { ID, ...mailingAddress } })
        .then(res => res.updated),
    deleteMailingAddress: (_, { ID }) =>
      client
        .DeleteMailingAddress()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPhoneNumber: (_, { phoneNumber }) =>
      client
        .CreatePhoneNumber()
        .sendMessage({ api: 'v1', item: { ...phoneNumber } })
        .then(res => ({ ID: res.ID, ...phoneNumber })),
    updatePhoneNumber: (_, { ID, phoneNumber }) =>
      client
        .UpdatePhoneNumber()
        .sendMessage({ api: 'v1', item: { ID, ...phoneNumber } })
        .then(res => res.updated),
    deletePhoneNumber: (_, { ID }) =>
      client
        .DeletePhoneNumber()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createEmailAddress: (_, { emailAddress }) =>
      client
        .CreateEmailAddress()
        .sendMessage({ api: 'v1', item: { ...emailAddress } })
        .then(res => ({ ID: res.ID, ...emailAddress })),
    updateEmailAddress: (_, { ID, emailAddress }) =>
      client
        .UpdateEmailAddress()
        .sendMessage({ api: 'v1', item: { ID, ...emailAddress } })
        .then(res => res.updated),
    deleteEmailAddress: (_, { ID }) =>
      client
        .DeleteEmailAddress()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPhoto: (_, { photo }) =>
      client
        .CreatePhoto()
        .sendMessage({ api: 'v1', item: { ...photo } })
        .then(res => ({ ID: res.ID, ...photo })),
    updatePhoto: (_, { ID, photo }) =>
      client
        .UpdatePhoto()
        .sendMessage({ api: 'v1', item: { ID, ...photo } })
        .then(res => res.updated),
    deletePhoto: (_, { ID }) =>
      client
        .DeletePhoto()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLayoutType: (_, { layoutType }) =>
      client
        .CreateLayoutType()
        .sendMessage({ api: 'v1', item: { ...layoutType } })
        .then(res => ({ ID: res.ID, ...layoutType })),
    updateLayoutType: (_, { ID, layoutType }) =>
      client
        .UpdateLayoutType()
        .sendMessage({ api: 'v1', item: { ID, ...layoutType } })
        .then(res => res.updated),
    deleteLayoutType: (_, { ID }) =>
      client
        .DeleteLayoutType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLayoutRow: (_, { layoutRow }) =>
      client
        .CreateLayoutRow()
        .sendMessage({ api: 'v1', item: { ...layoutRow } })
        .then(res => ({ ID: res.ID, ...layoutRow })),
    updateLayoutRow: (_, { ID, layoutRow }) =>
      client
        .UpdateLayoutRow()
        .sendMessage({ api: 'v1', item: { ID, ...layoutRow } })
        .then(res => res.updated),
    deleteLayoutRow: (_, { ID }) =>
      client
        .DeleteLayoutRow()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createComponentImplementation: (_, { componentImplementation }) =>
      client
        .CreateComponentImplementation()
        .sendMessage({ api: 'v1', item: { ...componentImplementation } })
        .then(res => ({ ID: res.ID, ...componentImplementation })),
    updateComponentImplementation: (_, { ID, componentImplementation }) =>
      client
        .UpdateComponentImplementation()
        .sendMessage({ api: 'v1', item: { ID, ...componentImplementation } })
        .then(res => res.updated),
    deleteComponentImplementation: (_, { ID }) =>
      client
        .DeleteComponentImplementation()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createComponentType: (_, { componentType }) =>
      client
        .CreateComponentType()
        .sendMessage({ api: 'v1', item: { ...componentType } })
        .then(res => ({ ID: res.ID, ...componentType })),
    updateComponentType: (_, { ID, componentType }) =>
      client
        .UpdateComponentType()
        .sendMessage({ api: 'v1', item: { ID, ...componentType } })
        .then(res => res.updated),
    deleteComponentType: (_, { ID }) =>
      client
        .DeleteComponentType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLayoutColumn: (_, { layoutColumn }) =>
      client
        .CreateLayoutColumn()
        .sendMessage({ api: 'v1', item: { ...layoutColumn } })
        .then(res => ({ ID: res.ID, ...layoutColumn })),
    updateLayoutColumn: (_, { ID, layoutColumn }) =>
      client
        .UpdateLayoutColumn()
        .sendMessage({ api: 'v1', item: { ID, ...layoutColumn } })
        .then(res => res.updated),
    deleteLayoutColumn: (_, { ID }) =>
      client
        .DeleteLayoutColumn()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createFieldType: (_, { fieldType }) =>
      client
        .CreateFieldType()
        .sendMessage({ api: 'v1', item: { ...fieldType } })
        .then(res => ({ ID: res.ID, ...fieldType })),
    updateFieldType: (_, { ID, fieldType }) =>
      client
        .UpdateFieldType()
        .sendMessage({ api: 'v1', item: { ID, ...fieldType } })
        .then(res => res.updated),
    deleteFieldType: (_, { ID }) =>
      client
        .DeleteFieldType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createComponent: (_, { component }) =>
      client
        .CreateComponent()
        .sendMessage({ api: 'v1', item: { ...component } })
        .then(res => ({ ID: res.ID, ...component })),
    updateComponent: (_, { ID, component }) =>
      client
        .UpdateComponent()
        .sendMessage({ api: 'v1', item: { ID, ...component } })
        .then(res => res.updated),
    deleteComponent: (_, { ID }) =>
      client
        .DeleteComponent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createField: (_, { field }) =>
      client
        .CreateField()
        .sendMessage({ api: 'v1', item: { ...field } })
        .then(res => ({ ID: res.ID, ...field })),
    updateField: (_, { ID, field }) =>
      client
        .UpdateField()
        .sendMessage({ api: 'v1', item: { ID, ...field } })
        .then(res => res.updated),
    deleteField: (_, { ID }) =>
      client
        .DeleteField()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createHomePage: (_, { homePage }) =>
      client
        .CreateHomePage()
        .sendMessage({ api: 'v1', item: { ...homePage } })
        .then(res => ({ ID: res.ID, ...homePage })),
    updateHomePage: (_, { ID, homePage }) =>
      client
        .UpdateHomePage()
        .sendMessage({ api: 'v1', item: { ID, ...homePage } })
        .then(res => res.updated),
    deleteHomePage: (_, { ID }) =>
      client
        .DeleteHomePage()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLayout: (_, { layout }) =>
      client
        .CreateLayout()
        .sendMessage({ api: 'v1', item: { ...layout } })
        .then(res => ({ ID: res.ID, ...layout })),
    updateLayout: (_, { ID, layout }) =>
      client
        .UpdateLayout()
        .sendMessage({ api: 'v1', item: { ID, ...layout } })
        .then(res => res.updated),
    deleteLayout: (_, { ID }) =>
      client
        .DeleteLayout()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLandingPage: (_, { landingPage }) =>
      client
        .CreateLandingPage()
        .sendMessage({ api: 'v1', item: { ...landingPage } })
        .then(res => ({ ID: res.ID, ...landingPage })),
    updateLandingPage: (_, { ID, landingPage }) =>
      client
        .UpdateLandingPage()
        .sendMessage({ api: 'v1', item: { ID, ...landingPage } })
        .then(res => res.updated),
    deleteLandingPage: (_, { ID }) =>
      client
        .DeleteLandingPage()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createExperiment: (_, { experiment }) =>
      client
        .CreateExperiment()
        .sendMessage({ api: 'v1', item: { ...experiment } })
        .then(res => ({ ID: res.ID, ...experiment })),
    updateExperiment: (_, { ID, experiment }) =>
      client
        .UpdateExperiment()
        .sendMessage({ api: 'v1', item: { ID, ...experiment } })
        .then(res => res.updated),
    deleteExperiment: (_, { ID }) =>
      client
        .DeleteExperiment()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createIssue: (_, { issue }) =>
      client
        .CreateIssue()
        .sendMessage({ api: 'v1', item: { ...issue } })
        .then(res => ({ ID: res.ID, ...issue })),
    updateIssue: (_, { ID, issue }) =>
      client
        .UpdateIssue()
        .sendMessage({ api: 'v1', item: { ID, ...issue } })
        .then(res => res.updated),
    deleteIssue: (_, { ID }) =>
      client
        .DeleteIssue()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createCandidate: (_, { candidate }) =>
      client
        .CreateCandidate()
        .sendMessage({ api: 'v1', item: { ...candidate } })
        .then(res => ({ ID: res.ID, ...candidate })),
    updateCandidate: (_, { ID, candidate }) =>
      client
        .UpdateCandidate()
        .sendMessage({ api: 'v1', item: { ID, ...candidate } })
        .then(res => res.updated),
    deleteCandidate: (_, { ID }) =>
      client
        .DeleteCandidate()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDistrictType: (_, { districtType }) =>
      client
        .CreateDistrictType()
        .sendMessage({ api: 'v1', item: { ...districtType } })
        .then(res => ({ ID: res.ID, ...districtType })),
    updateDistrictType: (_, { ID, districtType }) =>
      client
        .UpdateDistrictType()
        .sendMessage({ api: 'v1', item: { ID, ...districtType } })
        .then(res => res.updated),
    deleteDistrictType: (_, { ID }) =>
      client
        .DeleteDistrictType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDistrict: (_, { district }) =>
      client
        .CreateDistrict()
        .sendMessage({ api: 'v1', item: { ...district } })
        .then(res => ({ ID: res.ID, ...district })),
    updateDistrict: (_, { ID, district }) =>
      client
        .UpdateDistrict()
        .sendMessage({ api: 'v1', item: { ID, ...district } })
        .then(res => res.updated),
    deleteDistrict: (_, { ID }) =>
      client
        .DeleteDistrict()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createOffice: (_, { office }) =>
      client
        .CreateOffice()
        .sendMessage({ api: 'v1', item: { ...office } })
        .then(res => ({ ID: res.ID, ...office })),
    updateOffice: (_, { ID, office }) =>
      client
        .UpdateOffice()
        .sendMessage({ api: 'v1', item: { ID, ...office } })
        .then(res => res.updated),
    deleteOffice: (_, { ID }) =>
      client
        .DeleteOffice()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPollItem: (_, { pollItem }) =>
      client
        .CreatePollItem()
        .sendMessage({ api: 'v1', item: { ...pollItem } })
        .then(res => ({ ID: res.ID, ...pollItem })),
    updatePollItem: (_, { ID, pollItem }) =>
      client
        .UpdatePollItem()
        .sendMessage({ api: 'v1', item: { ID, ...pollItem } })
        .then(res => res.updated),
    deletePollItem: (_, { ID }) =>
      client
        .DeletePollItem()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createVolunteerOpportunityType: (_, { volunteerOpportunityType }) =>
      client
        .CreateVolunteerOpportunityType()
        .sendMessage({ api: 'v1', item: { ...volunteerOpportunityType } })
        .then(res => ({ ID: res.ID, ...volunteerOpportunityType })),
    updateVolunteerOpportunityType: (_, { ID, volunteerOpportunityType }) =>
      client
        .UpdateVolunteerOpportunityType()
        .sendMessage({ api: 'v1', item: { ID, ...volunteerOpportunityType } })
        .then(res => res.updated),
    deleteVolunteerOpportunityType: (_, { ID }) =>
      client
        .DeleteVolunteerOpportunityType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLiveEventType: (_, { liveEventType }) =>
      client
        .CreateLiveEventType()
        .sendMessage({ api: 'v1', item: { ...liveEventType } })
        .then(res => ({ ID: res.ID, ...liveEventType })),
    updateLiveEventType: (_, { ID, liveEventType }) =>
      client
        .UpdateLiveEventType()
        .sendMessage({ api: 'v1', item: { ID, ...liveEventType } })
        .then(res => res.updated),
    deleteLiveEventType: (_, { ID }) =>
      client
        .DeleteLiveEventType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createCompany: (_, { company }) =>
      client
        .CreateCompany()
        .sendMessage({ api: 'v1', item: { ...company } })
        .then(res => ({ ID: res.ID, ...company })),
    updateCompany: (_, { ID, company }) =>
      client
        .UpdateCompany()
        .sendMessage({ api: 'v1', item: { ID, ...company } })
        .then(res => res.updated),
    deleteCompany: (_, { ID }) =>
      client
        .DeleteCompany()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createProductType: (_, { productType }) =>
      client
        .CreateProductType()
        .sendMessage({ api: 'v1', item: { ...productType } })
        .then(res => ({ ID: res.ID, ...productType })),
    updateProductType: (_, { ID, productType }) =>
      client
        .UpdateProductType()
        .sendMessage({ api: 'v1', item: { ID, ...productType } })
        .then(res => res.updated),
    deleteProductType: (_, { ID }) =>
      client
        .DeleteProductType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createCustomerCart: (_, { customerCart }) =>
      client
        .CreateCustomerCart()
        .sendMessage({ api: 'v1', item: { ...customerCart } })
        .then(res => ({ ID: res.ID, ...customerCart })),
    updateCustomerCart: (_, { ID, customerCart }) =>
      client
        .UpdateCustomerCart()
        .sendMessage({ api: 'v1', item: { ID, ...customerCart } })
        .then(res => res.updated),
    deleteCustomerCart: (_, { ID }) =>
      client
        .DeleteCustomerCart()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPayment: (_, { payment }) =>
      client
        .CreatePayment()
        .sendMessage({ api: 'v1', item: { ...payment } })
        .then(res => ({ ID: res.ID, ...payment })),
    updatePayment: (_, { ID, payment }) =>
      client
        .UpdatePayment()
        .sendMessage({ api: 'v1', item: { ID, ...payment } })
        .then(res => res.updated),
    deletePayment: (_, { ID }) =>
      client
        .DeletePayment()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDelivery: (_, { delivery }) =>
      client
        .CreateDelivery()
        .sendMessage({ api: 'v1', item: { ...delivery } })
        .then(res => ({ ID: res.ID, ...delivery })),
    updateDelivery: (_, { ID, delivery }) =>
      client
        .UpdateDelivery()
        .sendMessage({ api: 'v1', item: { ID, ...delivery } })
        .then(res => res.updated),
    deleteDelivery: (_, { ID }) =>
      client
        .DeleteDelivery()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createBoycott: (_, { boycott }) =>
      client
        .CreateBoycott()
        .sendMessage({ api: 'v1', item: { ...boycott } })
        .then(res => ({ ID: res.ID, ...boycott })),
    updateBoycott: (_, { ID, boycott }) =>
      client
        .UpdateBoycott()
        .sendMessage({ api: 'v1', item: { ID, ...boycott } })
        .then(res => res.updated),
    deleteBoycott: (_, { ID }) =>
      client
        .DeleteBoycott()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createBoycottMembership: (_, { boycottMembership }) =>
      client
        .CreateBoycottMembership()
        .sendMessage({ api: 'v1', item: { ...boycottMembership } })
        .then(res => ({ ID: res.ID, ...boycottMembership })),
    updateBoycottMembership: (_, { ID, boycottMembership }) =>
      client
        .UpdateBoycottMembership()
        .sendMessage({ api: 'v1', item: { ID, ...boycottMembership } })
        .then(res => res.updated),
    deleteBoycottMembership: (_, { ID }) =>
      client
        .DeleteBoycottMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createElection: (_, { election }) =>
      client
        .CreateElection()
        .sendMessage({ api: 'v1', item: { ...election } })
        .then(res => ({ ID: res.ID, ...election })),
    updateElection: (_, { ID, election }) =>
      client
        .UpdateElection()
        .sendMessage({ api: 'v1', item: { ID, ...election } })
        .then(res => res.updated),
    deleteElection: (_, { ID }) =>
      client
        .DeleteElection()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createElectionMembership: (_, { electionMembership }) =>
      client
        .CreateElectionMembership()
        .sendMessage({ api: 'v1', item: { ...electionMembership } })
        .then(res => ({ ID: res.ID, ...electionMembership })),
    updateElectionMembership: (_, { ID, electionMembership }) =>
      client
        .UpdateElectionMembership()
        .sendMessage({ api: 'v1', item: { ID, ...electionMembership } })
        .then(res => res.updated),
    deleteElectionMembership: (_, { ID }) =>
      client
        .DeleteElectionMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPetitionMembership: (_, { petitionMembership }) =>
      client
        .CreatePetitionMembership()
        .sendMessage({ api: 'v1', item: { ...petitionMembership } })
        .then(res => ({ ID: res.ID, ...petitionMembership })),
    updatePetitionMembership: (_, { ID, petitionMembership }) =>
      client
        .UpdatePetitionMembership()
        .sendMessage({ api: 'v1', item: { ID, ...petitionMembership } })
        .then(res => res.updated),
    deletePetitionMembership: (_, { ID }) =>
      client
        .DeletePetitionMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPollMembership: (_, { pollMembership }) =>
      client
        .CreatePollMembership()
        .sendMessage({ api: 'v1', item: { ...pollMembership } })
        .then(res => ({ ID: res.ID, ...pollMembership })),
    updatePollMembership: (_, { ID, pollMembership }) =>
      client
        .UpdatePollMembership()
        .sendMessage({ api: 'v1', item: { ID, ...pollMembership } })
        .then(res => res.updated),
    deletePollMembership: (_, { ID }) =>
      client
        .DeletePollMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createVolunteerOpportunityMembership: (
      _,
      { volunteerOpportunityMembership }
    ) =>
      client
        .CreateVolunteerOpportunityMembership()
        .sendMessage({ api: 'v1', item: { ...volunteerOpportunityMembership } })
        .then(res => ({ ID: res.ID, ...volunteerOpportunityMembership })),
    updateVolunteerOpportunityMembership: (
      _,
      { ID, volunteerOpportunityMembership }
    ) =>
      client
        .UpdateVolunteerOpportunityMembership()
        .sendMessage({
          api: 'v1',
          item: { ID, ...volunteerOpportunityMembership }
        })
        .then(res => res.updated),
    deleteVolunteerOpportunityMembership: (_, { ID }) =>
      client
        .DeleteVolunteerOpportunityMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLiveEventMembership: (_, { liveEventMembership }) =>
      client
        .CreateLiveEventMembership()
        .sendMessage({ api: 'v1', item: { ...liveEventMembership } })
        .then(res => ({ ID: res.ID, ...liveEventMembership })),
    updateLiveEventMembership: (_, { ID, liveEventMembership }) =>
      client
        .UpdateLiveEventMembership()
        .sendMessage({ api: 'v1', item: { ID, ...liveEventMembership } })
        .then(res => res.updated),
    deleteLiveEventMembership: (_, { ID }) =>
      client
        .DeleteLiveEventMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createProduct: (_, { product }) =>
      client
        .CreateProduct()
        .sendMessage({ api: 'v1', item: { ...product } })
        .then(res => ({ ID: res.ID, ...product })),
    updateProduct: (_, { ID, product }) =>
      client
        .UpdateProduct()
        .sendMessage({ api: 'v1', item: { ID, ...product } })
        .then(res => res.updated),
    deleteProduct: (_, { ID }) =>
      client
        .DeleteProduct()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createProductMembership: (_, { productMembership }) =>
      client
        .CreateProductMembership()
        .sendMessage({ api: 'v1', item: { ...productMembership } })
        .then(res => ({ ID: res.ID, ...productMembership })),
    updateProductMembership: (_, { ID, productMembership }) =>
      client
        .UpdateProductMembership()
        .sendMessage({ api: 'v1', item: { ID, ...productMembership } })
        .then(res => res.updated),
    deleteProductMembership: (_, { ID }) =>
      client
        .DeleteProductMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDonationCampaign: (_, { donationCampaign }) =>
      client
        .CreateDonationCampaign()
        .sendMessage({ api: 'v1', item: { ...donationCampaign } })
        .then(res => ({ ID: res.ID, ...donationCampaign })),
    updateDonationCampaign: (_, { ID, donationCampaign }) =>
      client
        .UpdateDonationCampaign()
        .sendMessage({ api: 'v1', item: { ID, ...donationCampaign } })
        .then(res => res.updated),
    deleteDonationCampaign: (_, { ID }) =>
      client
        .DeleteDonationCampaign()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDonationCampaignMembership: (_, { donationCampaignMembership }) =>
      client
        .CreateDonationCampaignMembership()
        .sendMessage({ api: 'v1', item: { ...donationCampaignMembership } })
        .then(res => ({ ID: res.ID, ...donationCampaignMembership })),
    updateDonationCampaignMembership: (_, { ID, donationCampaignMembership }) =>
      client
        .UpdateDonationCampaignMembership()
        .sendMessage({ api: 'v1', item: { ID, ...donationCampaignMembership } })
        .then(res => res.updated),
    deleteDonationCampaignMembership: (_, { ID }) =>
      client
        .DeleteDonationCampaignMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPetition: (_, { petition }) =>
      client
        .CreatePetition()
        .sendMessage({ api: 'v1', item: { ...petition } })
        .then(res => ({ ID: res.ID, ...petition })),
    updatePetition: (_, { ID, petition }) =>
      client
        .UpdatePetition()
        .sendMessage({ api: 'v1', item: { ID, ...petition } })
        .then(res => res.updated),
    deletePetition: (_, { ID }) =>
      client
        .DeletePetition()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPetitionSigner: (_, { petitionSigner }) =>
      client
        .CreatePetitionSigner()
        .sendMessage({ api: 'v1', item: { ...petitionSigner } })
        .then(res => ({ ID: res.ID, ...petitionSigner })),
    updatePetitionSigner: (_, { ID, petitionSigner }) =>
      client
        .UpdatePetitionSigner()
        .sendMessage({ api: 'v1', item: { ID, ...petitionSigner } })
        .then(res => res.updated),
    deletePetitionSigner: (_, { ID }) =>
      client
        .DeletePetitionSigner()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPoll: (_, { poll }) =>
      client
        .CreatePoll()
        .sendMessage({ api: 'v1', item: { ...poll } })
        .then(res => ({ ID: res.ID, ...poll })),
    updatePoll: (_, { ID, poll }) =>
      client
        .UpdatePoll()
        .sendMessage({ api: 'v1', item: { ID, ...poll } })
        .then(res => res.updated),
    deletePoll: (_, { ID }) =>
      client
        .DeletePoll()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPollRespondant: (_, { pollRespondant }) =>
      client
        .CreatePollRespondant()
        .sendMessage({ api: 'v1', item: { ...pollRespondant } })
        .then(res => ({ ID: res.ID, ...pollRespondant })),
    updatePollRespondant: (_, { ID, pollRespondant }) =>
      client
        .UpdatePollRespondant()
        .sendMessage({ api: 'v1', item: { ID, ...pollRespondant } })
        .then(res => res.updated),
    deletePollRespondant: (_, { ID }) =>
      client
        .DeletePollRespondant()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createPurchaser: (_, { purchaser }) =>
      client
        .CreatePurchaser()
        .sendMessage({ api: 'v1', item: { ...purchaser } })
        .then(res => ({ ID: res.ID, ...purchaser })),
    updatePurchaser: (_, { ID, purchaser }) =>
      client
        .UpdatePurchaser()
        .sendMessage({ api: 'v1', item: { ID, ...purchaser } })
        .then(res => res.updated),
    deletePurchaser: (_, { ID }) =>
      client
        .DeletePurchaser()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createCustomerOrder: (_, { customerOrder }) =>
      client
        .CreateCustomerOrder()
        .sendMessage({ api: 'v1', item: { ...customerOrder } })
        .then(res => ({ ID: res.ID, ...customerOrder })),
    updateCustomerOrder: (_, { ID, customerOrder }) =>
      client
        .UpdateCustomerOrder()
        .sendMessage({ api: 'v1', item: { ID, ...customerOrder } })
        .then(res => res.updated),
    deleteCustomerOrder: (_, { ID }) =>
      client
        .DeleteCustomerOrder()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createDonor: (_, { donor }) =>
      client
        .CreateDonor()
        .sendMessage({ api: 'v1', item: { ...donor } })
        .then(res => ({ ID: res.ID, ...donor })),
    updateDonor: (_, { ID, donor }) =>
      client
        .UpdateDonor()
        .sendMessage({ api: 'v1', item: { ID, ...donor } })
        .then(res => res.updated),
    deleteDonor: (_, { ID }) =>
      client
        .DeleteDonor()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createLiveEvent: (_, { liveEvent }) =>
      client
        .CreateLiveEvent()
        .sendMessage({ api: 'v1', item: { ...liveEvent } })
        .then(res => ({ ID: res.ID, ...liveEvent })),
    updateLiveEvent: (_, { ID, liveEvent }) =>
      client
        .UpdateLiveEvent()
        .sendMessage({ api: 'v1', item: { ID, ...liveEvent } })
        .then(res => res.updated),
    deleteLiveEvent: (_, { ID }) =>
      client
        .DeleteLiveEvent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createEventAttendee: (_, { eventAttendee }) =>
      client
        .CreateEventAttendee()
        .sendMessage({ api: 'v1', item: { ...eventAttendee } })
        .then(res => ({ ID: res.ID, ...eventAttendee })),
    updateEventAttendee: (_, { ID, eventAttendee }) =>
      client
        .UpdateEventAttendee()
        .sendMessage({ api: 'v1', item: { ID, ...eventAttendee } })
        .then(res => res.updated),
    deleteEventAttendee: (_, { ID }) =>
      client
        .DeleteEventAttendee()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createVoter: (_, { voter }) =>
      client
        .CreateVoter()
        .sendMessage({ api: 'v1', item: { ...voter } })
        .then(res => ({ ID: res.ID, ...voter })),
    updateVoter: (_, { ID, voter }) =>
      client
        .UpdateVoter()
        .sendMessage({ api: 'v1', item: { ID, ...voter } })
        .then(res => res.updated),
    deleteVoter: (_, { ID }) =>
      client
        .DeleteVoter()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createVolunteerOpportunity: (_, { volunteerOpportunity }) =>
      client
        .CreateVolunteerOpportunity()
        .sendMessage({ api: 'v1', item: { ...volunteerOpportunity } })
        .then(res => ({ ID: res.ID, ...volunteerOpportunity })),
    updateVolunteerOpportunity: (_, { ID, volunteerOpportunity }) =>
      client
        .UpdateVolunteerOpportunity()
        .sendMessage({ api: 'v1', item: { ID, ...volunteerOpportunity } })
        .then(res => res.updated),
    deleteVolunteerOpportunity: (_, { ID }) =>
      client
        .DeleteVolunteerOpportunity()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createVolunteer: (_, { volunteer }) =>
      client
        .CreateVolunteer()
        .sendMessage({ api: 'v1', item: { ...volunteer } })
        .then(res => ({ ID: res.ID, ...volunteer })),
    updateVolunteer: (_, { ID, volunteer }) =>
      client
        .UpdateVolunteer()
        .sendMessage({ api: 'v1', item: { ID, ...volunteer } })
        .then(res => res.updated),
    deleteVolunteer: (_, { ID }) =>
      client
        .DeleteVolunteer()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createFollower: (_, { follower }) =>
      client
        .CreateFollower()
        .sendMessage({ api: 'v1', item: { ...follower } })
        .then(res => ({ ID: res.ID, ...follower })),
    updateFollower: (_, { ID, follower }) =>
      client
        .UpdateFollower()
        .sendMessage({ api: 'v1', item: { ID, ...follower } })
        .then(res => res.updated),
    deleteFollower: (_, { ID }) =>
      client
        .DeleteFollower()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createTerritory: (_, { territory }) =>
      client
        .CreateTerritory()
        .sendMessage({ api: 'v1', item: { ...territory } })
        .then(res => ({ ID: res.ID, ...territory })),
    updateTerritory: (_, { ID, territory }) =>
      client
        .UpdateTerritory()
        .sendMessage({ api: 'v1', item: { ID, ...territory } })
        .then(res => res.updated),
    deleteTerritory: (_, { ID }) =>
      client
        .DeleteTerritory()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createActivityType: (_, { activityType }) =>
      client
        .CreateActivityType()
        .sendMessage({ api: 'v1', item: { ...activityType } })
        .then(res => ({ ID: res.ID, ...activityType })),
    updateActivityType: (_, { ID, activityType }) =>
      client
        .UpdateActivityType()
        .sendMessage({ api: 'v1', item: { ID, ...activityType } })
        .then(res => res.updated),
    deleteActivityType: (_, { ID }) =>
      client
        .DeleteActivityType()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createActivity: (_, { activity }) =>
      client
        .CreateActivity()
        .sendMessage({ api: 'v1', item: { ...activity } })
        .then(res => ({ ID: res.ID, ...activity })),
    updateActivity: (_, { ID, activity }) =>
      client
        .UpdateActivity()
        .sendMessage({ api: 'v1', item: { ID, ...activity } })
        .then(res => res.updated),
    deleteActivity: (_, { ID }) =>
      client
        .DeleteActivity()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createNote: (_, { note }) =>
      client
        .CreateNote()
        .sendMessage({ api: 'v1', item: { ...note } })
        .then(res => ({ ID: res.ID, ...note })),
    updateNote: (_, { ID, note }) =>
      client
        .UpdateNote()
        .sendMessage({ api: 'v1', item: { ID, ...note } })
        .then(res => res.updated),
    deleteNote: (_, { ID }) =>
      client
        .DeleteNote()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createAccount: (_, { account }) =>
      client
        .CreateAccount()
        .sendMessage({ api: 'v1', item: { ...account } })
        .then(res => ({ ID: res.ID, ...account })),
    updateAccount: (_, { ID, account }) =>
      client
        .UpdateAccount()
        .sendMessage({ api: 'v1', item: { ID, ...account } })
        .then(res => res.updated),
    deleteAccount: (_, { ID }) =>
      client
        .DeleteAccount()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createOwnerMembership: (_, { ownerMembership }) =>
      client
        .CreateOwnerMembership()
        .sendMessage({ api: 'v1', item: { ...ownerMembership } })
        .then(res => ({ ID: res.ID, ...ownerMembership })),
    updateOwnerMembership: (_, { ID, ownerMembership }) =>
      client
        .UpdateOwnerMembership()
        .sendMessage({ api: 'v1', item: { ID, ...ownerMembership } })
        .then(res => res.updated),
    deleteOwnerMembership: (_, { ID }) =>
      client
        .DeleteOwnerMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createContact: (_, { contact }) =>
      client
        .CreateContact()
        .sendMessage({ api: 'v1', item: { ...contact } })
        .then(res => ({ ID: res.ID, ...contact })),
    updateContact: (_, { ID, contact }) =>
      client
        .UpdateContact()
        .sendMessage({ api: 'v1', item: { ID, ...contact } })
        .then(res => res.updated),
    deleteContact: (_, { ID }) =>
      client
        .DeleteContact()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createContactMembership: (_, { contactMembership }) =>
      client
        .CreateContactMembership()
        .sendMessage({ api: 'v1', item: { ...contactMembership } })
        .then(res => ({ ID: res.ID, ...contactMembership })),
    updateContactMembership: (_, { ID, contactMembership }) =>
      client
        .UpdateContactMembership()
        .sendMessage({ api: 'v1', item: { ID, ...contactMembership } })
        .then(res => res.updated),
    deleteContactMembership: (_, { ID }) =>
      client
        .DeleteContactMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createCause: (_, { cause }) =>
      client
        .CreateCause()
        .sendMessage({ api: 'v1', item: { ...cause } })
        .then(res => ({ ID: res.ID, ...cause })),
    updateCause: (_, { ID, cause }) =>
      client
        .UpdateCause()
        .sendMessage({ api: 'v1', item: { ID, ...cause } })
        .then(res => res.updated),
    deleteCause: (_, { ID }) =>
      client
        .DeleteCause()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createAgent: (_, { agent }) =>
      client
        .CreateAgent()
        .sendMessage({ api: 'v1', item: { ...agent } })
        .then(res => ({ ID: res.ID, ...agent })),
    updateAgent: (_, { ID, agent }) =>
      client
        .UpdateAgent()
        .sendMessage({ api: 'v1', item: { ID, ...agent } })
        .then(res => res.updated),
    deleteAgent: (_, { ID }) =>
      client
        .DeleteAgent()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted),
    createAgentMembership: (_, { agentMembership }) =>
      client
        .CreateAgentMembership()
        .sendMessage({ api: 'v1', item: { ...agentMembership } })
        .then(res => ({ ID: res.ID, ...agentMembership })),
    updateAgentMembership: (_, { ID, agentMembership }) =>
      client
        .UpdateAgentMembership()
        .sendMessage({ api: 'v1', item: { ID, ...agentMembership } })
        .then(res => res.updated),
    deleteAgentMembership: (_, { ID }) =>
      client
        .DeleteAgentMembership()
        .sendMessage({ api: 'v1', ID })
        .then(res => res.deleted)
  }
});