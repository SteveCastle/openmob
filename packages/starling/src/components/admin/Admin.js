import React from "react"
import PropTypes from "prop-types"
import { Router } from "@reach/router"
import { Link } from "gatsby"

import App from "@openmob/bluebird/src/components/layout/App"
import Content from "@openmob/bluebird/src/components/layout/Content"
import SideBar from "@openmob/bluebird/src/components/sidebar/SideBar"
import SideBarHeader from "@openmob/bluebird/src/components/sidebar/SideBarHeader"
import Menu from "@openmob/bluebird/src/components/menu/Menu"
import MenuHeader from "@openmob/bluebird/src/components/menu/MenuHeader"
import Input from "@openmob/bluebird/src/components/forms/Input"
import MenuItem from "@openmob/bluebird/src/components/menu/MenuItem"
import MenuBody from "@openmob/bluebird/src/components/menu/MenuBody"

import EditLayoutType from './edit/EditLayoutType'
import CreateLayoutType from './create/CreateLayoutType'
import ListLayoutType from './list/ListLayoutType'
import EditACL from './edit/EditACL'
import CreateACL from './create/CreateACL'
import ListACL from './list/ListACL'
import EditMailingAddress from './edit/EditMailingAddress'
import CreateMailingAddress from './create/CreateMailingAddress'
import ListMailingAddress from './list/ListMailingAddress'
import EditPhoneNumber from './edit/EditPhoneNumber'
import CreatePhoneNumber from './create/CreatePhoneNumber'
import ListPhoneNumber from './list/ListPhoneNumber'
import EditEmailAddress from './edit/EditEmailAddress'
import CreateEmailAddress from './create/CreateEmailAddress'
import ListEmailAddress from './list/ListEmailAddress'
import EditLayoutRow from './edit/EditLayoutRow'
import CreateLayoutRow from './create/CreateLayoutRow'
import ListLayoutRow from './list/ListLayoutRow'
import EditComponentImplementation from './edit/EditComponentImplementation'
import CreateComponentImplementation from './create/CreateComponentImplementation'
import ListComponentImplementation from './list/ListComponentImplementation'
import EditLayoutColumn from './edit/EditLayoutColumn'
import CreateLayoutColumn from './create/CreateLayoutColumn'
import ListLayoutColumn from './list/ListLayoutColumn'
import EditComponent from './edit/EditComponent'
import CreateComponent from './create/CreateComponent'
import ListComponent from './list/ListComponent'
import EditField from './edit/EditField'
import CreateField from './create/CreateField'
import ListField from './list/ListField'
import EditLayout from './edit/EditLayout'
import CreateLayout from './create/CreateLayout'
import ListLayout from './list/ListLayout'
import EditLandingPage from './edit/EditLandingPage'
import CreateLandingPage from './create/CreateLandingPage'
import ListLandingPage from './list/ListLandingPage'
import EditExperiment from './edit/EditExperiment'
import CreateExperiment from './create/CreateExperiment'
import ListExperiment from './list/ListExperiment'
import EditIssue from './edit/EditIssue'
import CreateIssue from './create/CreateIssue'
import ListIssue from './list/ListIssue'
import EditCandidate from './edit/EditCandidate'
import CreateCandidate from './create/CreateCandidate'
import ListCandidate from './list/ListCandidate'
import EditDistrictType from './edit/EditDistrictType'
import CreateDistrictType from './create/CreateDistrictType'
import ListDistrictType from './list/ListDistrictType'
import EditDistrict from './edit/EditDistrict'
import CreateDistrict from './create/CreateDistrict'
import ListDistrict from './list/ListDistrict'
import EditOffice from './edit/EditOffice'
import CreateOffice from './create/CreateOffice'
import ListOffice from './list/ListOffice'
import EditPollItem from './edit/EditPollItem'
import CreatePollItem from './create/CreatePollItem'
import ListPollItem from './list/ListPollItem'
import EditVolunteerOpportunityType from './edit/EditVolunteerOpportunityType'
import CreateVolunteerOpportunityType from './create/CreateVolunteerOpportunityType'
import ListVolunteerOpportunityType from './list/ListVolunteerOpportunityType'
import EditLiveEventType from './edit/EditLiveEventType'
import CreateLiveEventType from './create/CreateLiveEventType'
import ListLiveEventType from './list/ListLiveEventType'
import EditCompany from './edit/EditCompany'
import CreateCompany from './create/CreateCompany'
import ListCompany from './list/ListCompany'
import EditProductType from './edit/EditProductType'
import CreateProductType from './create/CreateProductType'
import ListProductType from './list/ListProductType'
import EditCustomerCart from './edit/EditCustomerCart'
import CreateCustomerCart from './create/CreateCustomerCart'
import ListCustomerCart from './list/ListCustomerCart'
import EditPayment from './edit/EditPayment'
import CreatePayment from './create/CreatePayment'
import ListPayment from './list/ListPayment'
import EditDelivery from './edit/EditDelivery'
import CreateDelivery from './create/CreateDelivery'
import ListDelivery from './list/ListDelivery'
import EditBoycott from './edit/EditBoycott'
import CreateBoycott from './create/CreateBoycott'
import ListBoycott from './list/ListBoycott'
import EditBoycottMembership from './edit/EditBoycottMembership'
import CreateBoycottMembership from './create/CreateBoycottMembership'
import ListBoycottMembership from './list/ListBoycottMembership'
import EditElection from './edit/EditElection'
import CreateElection from './create/CreateElection'
import ListElection from './list/ListElection'
import EditElectionMembership from './edit/EditElectionMembership'
import CreateElectionMembership from './create/CreateElectionMembership'
import ListElectionMembership from './list/ListElectionMembership'
import EditPetitionMembership from './edit/EditPetitionMembership'
import CreatePetitionMembership from './create/CreatePetitionMembership'
import ListPetitionMembership from './list/ListPetitionMembership'
import EditPollMembership from './edit/EditPollMembership'
import CreatePollMembership from './create/CreatePollMembership'
import ListPollMembership from './list/ListPollMembership'
import EditVolunteerOpportunityMembership from './edit/EditVolunteerOpportunityMembership'
import CreateVolunteerOpportunityMembership from './create/CreateVolunteerOpportunityMembership'
import ListVolunteerOpportunityMembership from './list/ListVolunteerOpportunityMembership'
import EditLiveEventMembership from './edit/EditLiveEventMembership'
import CreateLiveEventMembership from './create/CreateLiveEventMembership'
import ListLiveEventMembership from './list/ListLiveEventMembership'
import EditProduct from './edit/EditProduct'
import CreateProduct from './create/CreateProduct'
import ListProduct from './list/ListProduct'
import EditProductMembership from './edit/EditProductMembership'
import CreateProductMembership from './create/CreateProductMembership'
import ListProductMembership from './list/ListProductMembership'
import EditDonationCampaign from './edit/EditDonationCampaign'
import CreateDonationCampaign from './create/CreateDonationCampaign'
import ListDonationCampaign from './list/ListDonationCampaign'
import EditDonationCampaignMembership from './edit/EditDonationCampaignMembership'
import CreateDonationCampaignMembership from './create/CreateDonationCampaignMembership'
import ListDonationCampaignMembership from './list/ListDonationCampaignMembership'
import EditPetition from './edit/EditPetition'
import CreatePetition from './create/CreatePetition'
import ListPetition from './list/ListPetition'
import EditPetitionSigner from './edit/EditPetitionSigner'
import CreatePetitionSigner from './create/CreatePetitionSigner'
import ListPetitionSigner from './list/ListPetitionSigner'
import EditPoll from './edit/EditPoll'
import CreatePoll from './create/CreatePoll'
import ListPoll from './list/ListPoll'
import EditPollRespondant from './edit/EditPollRespondant'
import CreatePollRespondant from './create/CreatePollRespondant'
import ListPollRespondant from './list/ListPollRespondant'
import EditPurchaser from './edit/EditPurchaser'
import CreatePurchaser from './create/CreatePurchaser'
import ListPurchaser from './list/ListPurchaser'
import EditCustomerOrder from './edit/EditCustomerOrder'
import CreateCustomerOrder from './create/CreateCustomerOrder'
import ListCustomerOrder from './list/ListCustomerOrder'
import EditDonor from './edit/EditDonor'
import CreateDonor from './create/CreateDonor'
import ListDonor from './list/ListDonor'
import EditLiveEvent from './edit/EditLiveEvent'
import CreateLiveEvent from './create/CreateLiveEvent'
import ListLiveEvent from './list/ListLiveEvent'
import EditEventAttendee from './edit/EditEventAttendee'
import CreateEventAttendee from './create/CreateEventAttendee'
import ListEventAttendee from './list/ListEventAttendee'
import EditVoter from './edit/EditVoter'
import CreateVoter from './create/CreateVoter'
import ListVoter from './list/ListVoter'
import EditVolunteerOpportunity from './edit/EditVolunteerOpportunity'
import CreateVolunteerOpportunity from './create/CreateVolunteerOpportunity'
import ListVolunteerOpportunity from './list/ListVolunteerOpportunity'
import EditVolunteer from './edit/EditVolunteer'
import CreateVolunteer from './create/CreateVolunteer'
import ListVolunteer from './list/ListVolunteer'
import EditFollower from './edit/EditFollower'
import CreateFollower from './create/CreateFollower'
import ListFollower from './list/ListFollower'
import EditTerritory from './edit/EditTerritory'
import CreateTerritory from './create/CreateTerritory'
import ListTerritory from './list/ListTerritory'
import EditActivityType from './edit/EditActivityType'
import CreateActivityType from './create/CreateActivityType'
import ListActivityType from './list/ListActivityType'
import EditActivity from './edit/EditActivity'
import CreateActivity from './create/CreateActivity'
import ListActivity from './list/ListActivity'
import EditNote from './edit/EditNote'
import CreateNote from './create/CreateNote'
import ListNote from './list/ListNote'
import EditAccount from './edit/EditAccount'
import CreateAccount from './create/CreateAccount'
import ListAccount from './list/ListAccount'
import EditOwnerMembership from './edit/EditOwnerMembership'
import CreateOwnerMembership from './create/CreateOwnerMembership'
import ListOwnerMembership from './list/ListOwnerMembership'
import EditContact from './edit/EditContact'
import CreateContact from './create/CreateContact'
import ListContact from './list/ListContact'
import EditContactMembership from './edit/EditContactMembership'
import CreateContactMembership from './create/CreateContactMembership'
import ListContactMembership from './list/ListContactMembership'
import EditAgent from './edit/EditAgent'
import CreateAgent from './create/CreateAgent'
import ListAgent from './list/ListAgent'
import EditAgentMembership from './edit/EditAgentMembership'
import CreateAgentMembership from './create/CreateAgentMembership'
import ListAgentMembership from './list/ListAgentMembership'
import EditHomePage from './edit/EditHomePage'
import CreateHomePage from './create/CreateHomePage'
import ListHomePage from './list/ListHomePage'
import EditPhoto from './edit/EditPhoto'
import CreatePhoto from './create/CreatePhoto'
import ListPhoto from './list/ListPhoto'
import EditCause from './edit/EditCause'
import CreateCause from './create/CreateCause'
import ListCause from './list/ListCause'
import EditComponentType from './edit/EditComponentType'
import CreateComponentType from './create/CreateComponentType'
import ListComponentType from './list/ListComponentType'
import EditFieldType from './edit/EditFieldType'
import CreateFieldType from './create/CreateFieldType'
import ListFieldType from './list/ListFieldType'

const Layout = ({ children, title, id, summary }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">
          Open Mob
        </Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuHeader>
          <Input />
        </MenuHeader>
        <MenuBody>
<MenuItem>
          <Link to="/app/admin/layout-types">LayoutType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/acls">ACL</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/mailing-addresss">MailingAddress</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/phone-numbers">PhoneNumber</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/email-addresss">EmailAddress</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/layout-rows">LayoutRow</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/component-implementations">ComponentImplementation</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/layout-columns">LayoutColumn</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/components">Component</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/fields">Field</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/layouts">Layout</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/landing-pages">LandingPage</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/experiments">Experiment</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/issues">Issue</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/candidates">Candidate</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/district-types">DistrictType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/districts">District</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/offices">Office</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/poll-items">PollItem</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/volunteer-opportunity-types">VolunteerOpportunityType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/live-event-types">LiveEventType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/companys">Company</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/product-types">ProductType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/customer-carts">CustomerCart</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/payments">Payment</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/deliverys">Delivery</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/boycotts">Boycott</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/boycott-memberships">BoycottMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/elections">Election</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/election-memberships">ElectionMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/petition-memberships">PetitionMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/poll-memberships">PollMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/volunteer-opportunity-memberships">VolunteerOpportunityMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/live-event-memberships">LiveEventMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/products">Product</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/product-memberships">ProductMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/donation-campaigns">DonationCampaign</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/donation-campaign-memberships">DonationCampaignMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/petitions">Petition</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/petition-signers">PetitionSigner</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/polls">Poll</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/poll-respondants">PollRespondant</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/purchasers">Purchaser</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/customer-orders">CustomerOrder</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/donors">Donor</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/live-events">LiveEvent</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/event-attendees">EventAttendee</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/voters">Voter</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/volunteer-opportunitys">VolunteerOpportunity</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/volunteers">Volunteer</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/followers">Follower</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/territorys">Territory</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/activity-types">ActivityType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/activitys">Activity</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/notes">Note</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/accounts">Account</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/owner-memberships">OwnerMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/contacts">Contact</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/contact-memberships">ContactMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/agents">Agent</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/agent-memberships">AgentMembership</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/home-pages">HomePage</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/photos">Photo</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/causes">Cause</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/component-types">ComponentType</Link>
        </MenuItem>
<MenuItem>
          <Link to="/app/admin/field-types">FieldType</Link>
        </MenuItem>

        </MenuBody>
      </Menu>
    </SideBar>
    <Content>
    <Router>
<ListLayoutType path="/layout-types"/>
<CreateLayoutType path="/layout-types/create"/>
<EditLayoutType path="/layout-types/:id"/>
<ListACL path="/acls"/>
<CreateACL path="/acls/create"/>
<EditACL path="/acls/:id"/>
<ListMailingAddress path="/mailing-addresss"/>
<CreateMailingAddress path="/mailing-addresss/create"/>
<EditMailingAddress path="/mailing-addresss/:id"/>
<ListPhoneNumber path="/phone-numbers"/>
<CreatePhoneNumber path="/phone-numbers/create"/>
<EditPhoneNumber path="/phone-numbers/:id"/>
<ListEmailAddress path="/email-addresss"/>
<CreateEmailAddress path="/email-addresss/create"/>
<EditEmailAddress path="/email-addresss/:id"/>
<ListLayoutRow path="/layout-rows"/>
<CreateLayoutRow path="/layout-rows/create"/>
<EditLayoutRow path="/layout-rows/:id"/>
<ListComponentImplementation path="/component-implementations"/>
<CreateComponentImplementation path="/component-implementations/create"/>
<EditComponentImplementation path="/component-implementations/:id"/>
<ListLayoutColumn path="/layout-columns"/>
<CreateLayoutColumn path="/layout-columns/create"/>
<EditLayoutColumn path="/layout-columns/:id"/>
<ListComponent path="/components"/>
<CreateComponent path="/components/create"/>
<EditComponent path="/components/:id"/>
<ListField path="/fields"/>
<CreateField path="/fields/create"/>
<EditField path="/fields/:id"/>
<ListLayout path="/layouts"/>
<CreateLayout path="/layouts/create"/>
<EditLayout path="/layouts/:id"/>
<ListLandingPage path="/landing-pages"/>
<CreateLandingPage path="/landing-pages/create"/>
<EditLandingPage path="/landing-pages/:id"/>
<ListExperiment path="/experiments"/>
<CreateExperiment path="/experiments/create"/>
<EditExperiment path="/experiments/:id"/>
<ListIssue path="/issues"/>
<CreateIssue path="/issues/create"/>
<EditIssue path="/issues/:id"/>
<ListCandidate path="/candidates"/>
<CreateCandidate path="/candidates/create"/>
<EditCandidate path="/candidates/:id"/>
<ListDistrictType path="/district-types"/>
<CreateDistrictType path="/district-types/create"/>
<EditDistrictType path="/district-types/:id"/>
<ListDistrict path="/districts"/>
<CreateDistrict path="/districts/create"/>
<EditDistrict path="/districts/:id"/>
<ListOffice path="/offices"/>
<CreateOffice path="/offices/create"/>
<EditOffice path="/offices/:id"/>
<ListPollItem path="/poll-items"/>
<CreatePollItem path="/poll-items/create"/>
<EditPollItem path="/poll-items/:id"/>
<ListVolunteerOpportunityType path="/volunteer-opportunity-types"/>
<CreateVolunteerOpportunityType path="/volunteer-opportunity-types/create"/>
<EditVolunteerOpportunityType path="/volunteer-opportunity-types/:id"/>
<ListLiveEventType path="/live-event-types"/>
<CreateLiveEventType path="/live-event-types/create"/>
<EditLiveEventType path="/live-event-types/:id"/>
<ListCompany path="/companys"/>
<CreateCompany path="/companys/create"/>
<EditCompany path="/companys/:id"/>
<ListProductType path="/product-types"/>
<CreateProductType path="/product-types/create"/>
<EditProductType path="/product-types/:id"/>
<ListCustomerCart path="/customer-carts"/>
<CreateCustomerCart path="/customer-carts/create"/>
<EditCustomerCart path="/customer-carts/:id"/>
<ListPayment path="/payments"/>
<CreatePayment path="/payments/create"/>
<EditPayment path="/payments/:id"/>
<ListDelivery path="/deliverys"/>
<CreateDelivery path="/deliverys/create"/>
<EditDelivery path="/deliverys/:id"/>
<ListBoycott path="/boycotts"/>
<CreateBoycott path="/boycotts/create"/>
<EditBoycott path="/boycotts/:id"/>
<ListBoycottMembership path="/boycott-memberships"/>
<CreateBoycottMembership path="/boycott-memberships/create"/>
<EditBoycottMembership path="/boycott-memberships/:id"/>
<ListElection path="/elections"/>
<CreateElection path="/elections/create"/>
<EditElection path="/elections/:id"/>
<ListElectionMembership path="/election-memberships"/>
<CreateElectionMembership path="/election-memberships/create"/>
<EditElectionMembership path="/election-memberships/:id"/>
<ListPetitionMembership path="/petition-memberships"/>
<CreatePetitionMembership path="/petition-memberships/create"/>
<EditPetitionMembership path="/petition-memberships/:id"/>
<ListPollMembership path="/poll-memberships"/>
<CreatePollMembership path="/poll-memberships/create"/>
<EditPollMembership path="/poll-memberships/:id"/>
<ListVolunteerOpportunityMembership path="/volunteer-opportunity-memberships"/>
<CreateVolunteerOpportunityMembership path="/volunteer-opportunity-memberships/create"/>
<EditVolunteerOpportunityMembership path="/volunteer-opportunity-memberships/:id"/>
<ListLiveEventMembership path="/live-event-memberships"/>
<CreateLiveEventMembership path="/live-event-memberships/create"/>
<EditLiveEventMembership path="/live-event-memberships/:id"/>
<ListProduct path="/products"/>
<CreateProduct path="/products/create"/>
<EditProduct path="/products/:id"/>
<ListProductMembership path="/product-memberships"/>
<CreateProductMembership path="/product-memberships/create"/>
<EditProductMembership path="/product-memberships/:id"/>
<ListDonationCampaign path="/donation-campaigns"/>
<CreateDonationCampaign path="/donation-campaigns/create"/>
<EditDonationCampaign path="/donation-campaigns/:id"/>
<ListDonationCampaignMembership path="/donation-campaign-memberships"/>
<CreateDonationCampaignMembership path="/donation-campaign-memberships/create"/>
<EditDonationCampaignMembership path="/donation-campaign-memberships/:id"/>
<ListPetition path="/petitions"/>
<CreatePetition path="/petitions/create"/>
<EditPetition path="/petitions/:id"/>
<ListPetitionSigner path="/petition-signers"/>
<CreatePetitionSigner path="/petition-signers/create"/>
<EditPetitionSigner path="/petition-signers/:id"/>
<ListPoll path="/polls"/>
<CreatePoll path="/polls/create"/>
<EditPoll path="/polls/:id"/>
<ListPollRespondant path="/poll-respondants"/>
<CreatePollRespondant path="/poll-respondants/create"/>
<EditPollRespondant path="/poll-respondants/:id"/>
<ListPurchaser path="/purchasers"/>
<CreatePurchaser path="/purchasers/create"/>
<EditPurchaser path="/purchasers/:id"/>
<ListCustomerOrder path="/customer-orders"/>
<CreateCustomerOrder path="/customer-orders/create"/>
<EditCustomerOrder path="/customer-orders/:id"/>
<ListDonor path="/donors"/>
<CreateDonor path="/donors/create"/>
<EditDonor path="/donors/:id"/>
<ListLiveEvent path="/live-events"/>
<CreateLiveEvent path="/live-events/create"/>
<EditLiveEvent path="/live-events/:id"/>
<ListEventAttendee path="/event-attendees"/>
<CreateEventAttendee path="/event-attendees/create"/>
<EditEventAttendee path="/event-attendees/:id"/>
<ListVoter path="/voters"/>
<CreateVoter path="/voters/create"/>
<EditVoter path="/voters/:id"/>
<ListVolunteerOpportunity path="/volunteer-opportunitys"/>
<CreateVolunteerOpportunity path="/volunteer-opportunitys/create"/>
<EditVolunteerOpportunity path="/volunteer-opportunitys/:id"/>
<ListVolunteer path="/volunteers"/>
<CreateVolunteer path="/volunteers/create"/>
<EditVolunteer path="/volunteers/:id"/>
<ListFollower path="/followers"/>
<CreateFollower path="/followers/create"/>
<EditFollower path="/followers/:id"/>
<ListTerritory path="/territorys"/>
<CreateTerritory path="/territorys/create"/>
<EditTerritory path="/territorys/:id"/>
<ListActivityType path="/activity-types"/>
<CreateActivityType path="/activity-types/create"/>
<EditActivityType path="/activity-types/:id"/>
<ListActivity path="/activitys"/>
<CreateActivity path="/activitys/create"/>
<EditActivity path="/activitys/:id"/>
<ListNote path="/notes"/>
<CreateNote path="/notes/create"/>
<EditNote path="/notes/:id"/>
<ListAccount path="/accounts"/>
<CreateAccount path="/accounts/create"/>
<EditAccount path="/accounts/:id"/>
<ListOwnerMembership path="/owner-memberships"/>
<CreateOwnerMembership path="/owner-memberships/create"/>
<EditOwnerMembership path="/owner-memberships/:id"/>
<ListContact path="/contacts"/>
<CreateContact path="/contacts/create"/>
<EditContact path="/contacts/:id"/>
<ListContactMembership path="/contact-memberships"/>
<CreateContactMembership path="/contact-memberships/create"/>
<EditContactMembership path="/contact-memberships/:id"/>
<ListAgent path="/agents"/>
<CreateAgent path="/agents/create"/>
<EditAgent path="/agents/:id"/>
<ListAgentMembership path="/agent-memberships"/>
<CreateAgentMembership path="/agent-memberships/create"/>
<EditAgentMembership path="/agent-memberships/:id"/>
<ListHomePage path="/home-pages"/>
<CreateHomePage path="/home-pages/create"/>
<EditHomePage path="/home-pages/:id"/>
<ListPhoto path="/photos"/>
<CreatePhoto path="/photos/create"/>
<EditPhoto path="/photos/:id"/>
<ListCause path="/causes"/>
<CreateCause path="/causes/create"/>
<EditCause path="/causes/:id"/>
<ListComponentType path="/component-types"/>
<CreateComponentType path="/component-types/create"/>
<EditComponentType path="/component-types/:id"/>
<ListFieldType path="/field-types"/>
<CreateFieldType path="/field-types/create"/>
<EditFieldType path="/field-types/:id"/>

    </Router>
    </Content>
  </App>
)

Layout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Layout
