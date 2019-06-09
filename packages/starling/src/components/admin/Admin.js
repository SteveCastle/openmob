import React from 'react';
import PropTypes from 'prop-types';
import { Router } from '@reach/router';
import { Link } from 'gatsby';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';

import App from '@openmob/bluebird/src/components/layout/App';
import SideBar from '@openmob/bluebird/src/components/layout/sidebar/SideBar';
import SideBarHeader from '@openmob/bluebird/src/components/layout/sidebar/SideBarHeader';
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel';
import Menu from '@openmob/bluebird/src/components/menu/Menu';
import MenuHeader from '@openmob/bluebird/src/components/menu/MenuHeader';
import Input from '@openmob/bluebird/src/components/forms/Input';
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem';
import MenuBody from '@openmob/bluebird/src/components/menu/MenuBody';

import EditACL from './edit/EditACL';
import CreateACL from './create/CreateACL';
import ListACL from './list/ListACL';
import EditMailingAddress from './edit/EditMailingAddress';
import CreateMailingAddress from './create/CreateMailingAddress';
import ListMailingAddress from './list/ListMailingAddress';
import EditPhoneNumber from './edit/EditPhoneNumber';
import CreatePhoneNumber from './create/CreatePhoneNumber';
import ListPhoneNumber from './list/ListPhoneNumber';
import EditEmailAddress from './edit/EditEmailAddress';
import CreateEmailAddress from './create/CreateEmailAddress';
import ListEmailAddress from './list/ListEmailAddress';
import EditLayoutType from './edit/EditLayoutType';
import CreateLayoutType from './create/CreateLayoutType';
import ListLayoutType from './list/ListLayoutType';
import EditLayoutRow from './edit/EditLayoutRow';
import CreateLayoutRow from './create/CreateLayoutRow';
import ListLayoutRow from './list/ListLayoutRow';
import EditComponentTypeFields from './edit/EditComponentTypeFields';
import CreateComponentTypeFields from './create/CreateComponentTypeFields';
import ListComponentTypeFields from './list/ListComponentTypeFields';
import EditComponentType from './edit/EditComponentType';
import CreateComponentType from './create/CreateComponentType';
import ListComponentType from './list/ListComponentType';
import EditComponentImplementation from './edit/EditComponentImplementation';
import CreateComponentImplementation from './create/CreateComponentImplementation';
import ListComponentImplementation from './list/ListComponentImplementation';
import EditLayoutColumn from './edit/EditLayoutColumn';
import CreateLayoutColumn from './create/CreateLayoutColumn';
import ListLayoutColumn from './list/ListLayoutColumn';
import EditFieldType from './edit/EditFieldType';
import CreateFieldType from './create/CreateFieldType';
import ListFieldType from './list/ListFieldType';
import EditLayout from './edit/EditLayout';
import CreateLayout from './create/CreateLayout';
import ListLayout from './list/ListLayout';
import EditExperimentVariant from './edit/EditExperimentVariant';
import CreateExperimentVariant from './create/CreateExperimentVariant';
import ListExperimentVariant from './list/ListExperimentVariant';
import EditExperiment from './edit/EditExperiment';
import CreateExperiment from './create/CreateExperiment';
import ListExperiment from './list/ListExperiment';
import EditLandingPage from './edit/EditLandingPage';
import CreateLandingPage from './create/CreateLandingPage';
import ListLandingPage from './list/ListLandingPage';
import EditIssue from './edit/EditIssue';
import CreateIssue from './create/CreateIssue';
import ListIssue from './list/ListIssue';
import EditField from './edit/EditField';
import CreateField from './create/CreateField';
import ListField from './list/ListField';
import EditComponent from './edit/EditComponent';
import CreateComponent from './create/CreateComponent';
import ListComponent from './list/ListComponent';
import EditCandidate from './edit/EditCandidate';
import CreateCandidate from './create/CreateCandidate';
import ListCandidate from './list/ListCandidate';
import EditDistrictType from './edit/EditDistrictType';
import CreateDistrictType from './create/CreateDistrictType';
import ListDistrictType from './list/ListDistrictType';
import EditDistrict from './edit/EditDistrict';
import CreateDistrict from './create/CreateDistrict';
import ListDistrict from './list/ListDistrict';
import EditVolunteerOpportunityType from './edit/EditVolunteerOpportunityType';
import CreateVolunteerOpportunityType from './create/CreateVolunteerOpportunityType';
import ListVolunteerOpportunityType from './list/ListVolunteerOpportunityType';
import EditOffice from './edit/EditOffice';
import CreateOffice from './create/CreateOffice';
import ListOffice from './list/ListOffice';
import EditPollItem from './edit/EditPollItem';
import CreatePollItem from './create/CreatePollItem';
import ListPollItem from './list/ListPollItem';
import EditLiveEventType from './edit/EditLiveEventType';
import CreateLiveEventType from './create/CreateLiveEventType';
import ListLiveEventType from './list/ListLiveEventType';
import EditCompany from './edit/EditCompany';
import CreateCompany from './create/CreateCompany';
import ListCompany from './list/ListCompany';
import EditProductType from './edit/EditProductType';
import CreateProductType from './create/CreateProductType';
import ListProductType from './list/ListProductType';
import EditCustomerCart from './edit/EditCustomerCart';
import CreateCustomerCart from './create/CreateCustomerCart';
import ListCustomerCart from './list/ListCustomerCart';
import EditPayment from './edit/EditPayment';
import CreatePayment from './create/CreatePayment';
import ListPayment from './list/ListPayment';
import EditDelivery from './edit/EditDelivery';
import CreateDelivery from './create/CreateDelivery';
import ListDelivery from './list/ListDelivery';
import EditElection from './edit/EditElection';
import CreateElection from './create/CreateElection';
import ListElection from './list/ListElection';
import EditBoycott from './edit/EditBoycott';
import CreateBoycott from './create/CreateBoycott';
import ListBoycott from './list/ListBoycott';
import EditBoycottMembership from './edit/EditBoycottMembership';
import CreateBoycottMembership from './create/CreateBoycottMembership';
import ListBoycottMembership from './list/ListBoycottMembership';
import EditElectionMembership from './edit/EditElectionMembership';
import CreateElectionMembership from './create/CreateElectionMembership';
import ListElectionMembership from './list/ListElectionMembership';
import EditPetitionMembership from './edit/EditPetitionMembership';
import CreatePetitionMembership from './create/CreatePetitionMembership';
import ListPetitionMembership from './list/ListPetitionMembership';
import EditPollMembership from './edit/EditPollMembership';
import CreatePollMembership from './create/CreatePollMembership';
import ListPollMembership from './list/ListPollMembership';
import EditVolunteerOpportunityMembership from './edit/EditVolunteerOpportunityMembership';
import CreateVolunteerOpportunityMembership from './create/CreateVolunteerOpportunityMembership';
import ListVolunteerOpportunityMembership from './list/ListVolunteerOpportunityMembership';
import EditLiveEventMembership from './edit/EditLiveEventMembership';
import CreateLiveEventMembership from './create/CreateLiveEventMembership';
import ListLiveEventMembership from './list/ListLiveEventMembership';
import EditProduct from './edit/EditProduct';
import CreateProduct from './create/CreateProduct';
import ListProduct from './list/ListProduct';
import EditProductMembership from './edit/EditProductMembership';
import CreateProductMembership from './create/CreateProductMembership';
import ListProductMembership from './list/ListProductMembership';
import EditDonationCampaign from './edit/EditDonationCampaign';
import CreateDonationCampaign from './create/CreateDonationCampaign';
import ListDonationCampaign from './list/ListDonationCampaign';
import EditDonationCampaignMembership from './edit/EditDonationCampaignMembership';
import CreateDonationCampaignMembership from './create/CreateDonationCampaignMembership';
import ListDonationCampaignMembership from './list/ListDonationCampaignMembership';
import EditPetition from './edit/EditPetition';
import CreatePetition from './create/CreatePetition';
import ListPetition from './list/ListPetition';
import EditPetitionSigner from './edit/EditPetitionSigner';
import CreatePetitionSigner from './create/CreatePetitionSigner';
import ListPetitionSigner from './list/ListPetitionSigner';
import EditPoll from './edit/EditPoll';
import CreatePoll from './create/CreatePoll';
import ListPoll from './list/ListPoll';
import EditPollRespondant from './edit/EditPollRespondant';
import CreatePollRespondant from './create/CreatePollRespondant';
import ListPollRespondant from './list/ListPollRespondant';
import EditPurchaser from './edit/EditPurchaser';
import CreatePurchaser from './create/CreatePurchaser';
import ListPurchaser from './list/ListPurchaser';
import EditCustomerOrder from './edit/EditCustomerOrder';
import CreateCustomerOrder from './create/CreateCustomerOrder';
import ListCustomerOrder from './list/ListCustomerOrder';
import EditDonor from './edit/EditDonor';
import CreateDonor from './create/CreateDonor';
import ListDonor from './list/ListDonor';
import EditLiveEvent from './edit/EditLiveEvent';
import CreateLiveEvent from './create/CreateLiveEvent';
import ListLiveEvent from './list/ListLiveEvent';
import EditEventAttendee from './edit/EditEventAttendee';
import CreateEventAttendee from './create/CreateEventAttendee';
import ListEventAttendee from './list/ListEventAttendee';
import EditVoter from './edit/EditVoter';
import CreateVoter from './create/CreateVoter';
import ListVoter from './list/ListVoter';
import EditVolunteerOpportunity from './edit/EditVolunteerOpportunity';
import CreateVolunteerOpportunity from './create/CreateVolunteerOpportunity';
import ListVolunteerOpportunity from './list/ListVolunteerOpportunity';
import EditVolunteer from './edit/EditVolunteer';
import CreateVolunteer from './create/CreateVolunteer';
import ListVolunteer from './list/ListVolunteer';
import EditActivityType from './edit/EditActivityType';
import CreateActivityType from './create/CreateActivityType';
import ListActivityType from './list/ListActivityType';
import EditFollower from './edit/EditFollower';
import CreateFollower from './create/CreateFollower';
import ListFollower from './list/ListFollower';
import EditTerritory from './edit/EditTerritory';
import CreateTerritory from './create/CreateTerritory';
import ListTerritory from './list/ListTerritory';
import EditActivity from './edit/EditActivity';
import CreateActivity from './create/CreateActivity';
import ListActivity from './list/ListActivity';
import EditNote from './edit/EditNote';
import CreateNote from './create/CreateNote';
import ListNote from './list/ListNote';
import EditAgent from './edit/EditAgent';
import CreateAgent from './create/CreateAgent';
import ListAgent from './list/ListAgent';
import EditAccount from './edit/EditAccount';
import CreateAccount from './create/CreateAccount';
import ListAccount from './list/ListAccount';
import EditOwnerMembership from './edit/EditOwnerMembership';
import CreateOwnerMembership from './create/CreateOwnerMembership';
import ListOwnerMembership from './list/ListOwnerMembership';
import EditAgentMembership from './edit/EditAgentMembership';
import CreateAgentMembership from './create/CreateAgentMembership';
import ListAgentMembership from './list/ListAgentMembership';
import EditContact from './edit/EditContact';
import CreateContact from './create/CreateContact';
import ListContact from './list/ListContact';
import EditContactMembership from './edit/EditContactMembership';
import CreateContactMembership from './create/CreateContactMembership';
import ListContactMembership from './list/ListContactMembership';
import EditHomePage from './edit/EditHomePage';
import CreateHomePage from './create/CreateHomePage';
import ListHomePage from './list/ListHomePage';
import EditPhoto from './edit/EditPhoto';
import CreatePhoto from './create/CreatePhoto';
import ListPhoto from './list/ListPhoto';
import EditCause from './edit/EditCause';
import CreateCause from './create/CreateCause';
import ListCause from './list/ListCause';

const Layout = ({ children, title, id, summary }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuHeader>
          <Input block />
          <MenuItem noBorder>
            <Link to={`/app`}>
              <FontAwesomeIcon icon={faArrowLeft} />
              Back to Dashboard
            </Link>
          </MenuItem>
        </MenuHeader>
        <MenuBody>
          <MenuItem>
            <Link to="/app/admin/acl">ACL</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/mailing-address">MailingAddress</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/phone-number">PhoneNumber</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/email-address">EmailAddress</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/layout-type">LayoutType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/layout-row">LayoutRow</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/component-type-fields">
              ComponentTypeFields
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/component-type">ComponentType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/component-implementation">
              ComponentImplementation
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/layout-column">LayoutColumn</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/field-type">FieldType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/layout">Layout</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/experiment-variant">ExperimentVariant</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/experiment">Experiment</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/landing-page">LandingPage</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/issue">Issue</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/field">Field</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/component">Component</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/candidate">Candidate</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/district-type">DistrictType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/district">District</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/volunteer-opportunity-type">
              VolunteerOpportunityType
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/office">Office</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/poll-item">PollItem</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/live-event-type">LiveEventType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/company">Company</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/product-type">ProductType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/customer-cart">CustomerCart</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/payment">Payment</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/delivery">Delivery</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/election">Election</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/boycott">Boycott</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/boycott-membership">BoycottMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/election-membership">ElectionMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/petition-membership">PetitionMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/poll-membership">PollMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/volunteer-opportunity-membership">
              VolunteerOpportunityMembership
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/live-event-membership">
              LiveEventMembership
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/product">Product</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/product-membership">ProductMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/donation-campaign">DonationCampaign</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/donation-campaign-membership">
              DonationCampaignMembership
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/petition">Petition</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/petition-signer">PetitionSigner</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/poll">Poll</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/poll-respondant">PollRespondant</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/purchaser">Purchaser</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/customer-order">CustomerOrder</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/donor">Donor</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/live-event">LiveEvent</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/event-attendee">EventAttendee</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/voter">Voter</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/volunteer-opportunity">
              VolunteerOpportunity
            </Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/volunteer">Volunteer</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/activity-type">ActivityType</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/follower">Follower</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/territory">Territory</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/activity">Activity</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/note">Note</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/agent">Agent</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/account">Account</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/owner-membership">OwnerMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/agent-membership">AgentMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/contact">Contact</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/contact-membership">ContactMembership</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/home-page">HomePage</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/photo">Photo</Link>
          </MenuItem>
          <MenuItem>
            <Link to="/app/admin/cause">Cause</Link>
          </MenuItem>
        </MenuBody>
      </Menu>
    </SideBar>
    <ContentPanel>
      <Router>
        <ListACL path="/acl" />
        <CreateACL path="/acl/create" />
        <EditACL path="/acl/:id" />
        <ListMailingAddress path="/mailing-address" />
        <CreateMailingAddress path="/mailing-address/create" />
        <EditMailingAddress path="/mailing-address/:id" />
        <ListPhoneNumber path="/phone-number" />
        <CreatePhoneNumber path="/phone-number/create" />
        <EditPhoneNumber path="/phone-number/:id" />
        <ListEmailAddress path="/email-address" />
        <CreateEmailAddress path="/email-address/create" />
        <EditEmailAddress path="/email-address/:id" />
        <ListLayoutType path="/layout-type" />
        <CreateLayoutType path="/layout-type/create" />
        <EditLayoutType path="/layout-type/:id" />
        <ListLayoutRow path="/layout-row" />
        <CreateLayoutRow path="/layout-row/create" />
        <EditLayoutRow path="/layout-row/:id" />
        <ListComponentTypeFields path="/component-type-fields" />
        <CreateComponentTypeFields path="/component-type-fields/create" />
        <EditComponentTypeFields path="/component-type-fields/:id" />
        <ListComponentType path="/component-type" />
        <CreateComponentType path="/component-type/create" />
        <EditComponentType path="/component-type/:id" />
        <ListComponentImplementation path="/component-implementation" />
        <CreateComponentImplementation path="/component-implementation/create" />
        <EditComponentImplementation path="/component-implementation/:id" />
        <ListLayoutColumn path="/layout-column" />
        <CreateLayoutColumn path="/layout-column/create" />
        <EditLayoutColumn path="/layout-column/:id" />
        <ListFieldType path="/field-type" />
        <CreateFieldType path="/field-type/create" />
        <EditFieldType path="/field-type/:id" />
        <ListLayout path="/layout" />
        <CreateLayout path="/layout/create" />
        <EditLayout path="/layout/:id" />
        <ListExperimentVariant path="/experiment-variant" />
        <CreateExperimentVariant path="/experiment-variant/create" />
        <EditExperimentVariant path="/experiment-variant/:id" />
        <ListExperiment path="/experiment" />
        <CreateExperiment path="/experiment/create" />
        <EditExperiment path="/experiment/:id" />
        <ListLandingPage path="/landing-page" />
        <CreateLandingPage path="/landing-page/create" />
        <EditLandingPage path="/landing-page/:id" />
        <ListIssue path="/issue" />
        <CreateIssue path="/issue/create" />
        <EditIssue path="/issue/:id" />
        <ListField path="/field" />
        <CreateField path="/field/create" />
        <EditField path="/field/:id" />
        <ListComponent path="/component" />
        <CreateComponent path="/component/create" />
        <EditComponent path="/component/:id" />
        <ListCandidate path="/candidate" />
        <CreateCandidate path="/candidate/create" />
        <EditCandidate path="/candidate/:id" />
        <ListDistrictType path="/district-type" />
        <CreateDistrictType path="/district-type/create" />
        <EditDistrictType path="/district-type/:id" />
        <ListDistrict path="/district" />
        <CreateDistrict path="/district/create" />
        <EditDistrict path="/district/:id" />
        <ListVolunteerOpportunityType path="/volunteer-opportunity-type" />
        <CreateVolunteerOpportunityType path="/volunteer-opportunity-type/create" />
        <EditVolunteerOpportunityType path="/volunteer-opportunity-type/:id" />
        <ListOffice path="/office" />
        <CreateOffice path="/office/create" />
        <EditOffice path="/office/:id" />
        <ListPollItem path="/poll-item" />
        <CreatePollItem path="/poll-item/create" />
        <EditPollItem path="/poll-item/:id" />
        <ListLiveEventType path="/live-event-type" />
        <CreateLiveEventType path="/live-event-type/create" />
        <EditLiveEventType path="/live-event-type/:id" />
        <ListCompany path="/company" />
        <CreateCompany path="/company/create" />
        <EditCompany path="/company/:id" />
        <ListProductType path="/product-type" />
        <CreateProductType path="/product-type/create" />
        <EditProductType path="/product-type/:id" />
        <ListCustomerCart path="/customer-cart" />
        <CreateCustomerCart path="/customer-cart/create" />
        <EditCustomerCart path="/customer-cart/:id" />
        <ListPayment path="/payment" />
        <CreatePayment path="/payment/create" />
        <EditPayment path="/payment/:id" />
        <ListDelivery path="/delivery" />
        <CreateDelivery path="/delivery/create" />
        <EditDelivery path="/delivery/:id" />
        <ListElection path="/election" />
        <CreateElection path="/election/create" />
        <EditElection path="/election/:id" />
        <ListBoycott path="/boycott" />
        <CreateBoycott path="/boycott/create" />
        <EditBoycott path="/boycott/:id" />
        <ListBoycottMembership path="/boycott-membership" />
        <CreateBoycottMembership path="/boycott-membership/create" />
        <EditBoycottMembership path="/boycott-membership/:id" />
        <ListElectionMembership path="/election-membership" />
        <CreateElectionMembership path="/election-membership/create" />
        <EditElectionMembership path="/election-membership/:id" />
        <ListPetitionMembership path="/petition-membership" />
        <CreatePetitionMembership path="/petition-membership/create" />
        <EditPetitionMembership path="/petition-membership/:id" />
        <ListPollMembership path="/poll-membership" />
        <CreatePollMembership path="/poll-membership/create" />
        <EditPollMembership path="/poll-membership/:id" />
        <ListVolunteerOpportunityMembership path="/volunteer-opportunity-membership" />
        <CreateVolunteerOpportunityMembership path="/volunteer-opportunity-membership/create" />
        <EditVolunteerOpportunityMembership path="/volunteer-opportunity-membership/:id" />
        <ListLiveEventMembership path="/live-event-membership" />
        <CreateLiveEventMembership path="/live-event-membership/create" />
        <EditLiveEventMembership path="/live-event-membership/:id" />
        <ListProduct path="/product" />
        <CreateProduct path="/product/create" />
        <EditProduct path="/product/:id" />
        <ListProductMembership path="/product-membership" />
        <CreateProductMembership path="/product-membership/create" />
        <EditProductMembership path="/product-membership/:id" />
        <ListDonationCampaign path="/donation-campaign" />
        <CreateDonationCampaign path="/donation-campaign/create" />
        <EditDonationCampaign path="/donation-campaign/:id" />
        <ListDonationCampaignMembership path="/donation-campaign-membership" />
        <CreateDonationCampaignMembership path="/donation-campaign-membership/create" />
        <EditDonationCampaignMembership path="/donation-campaign-membership/:id" />
        <ListPetition path="/petition" />
        <CreatePetition path="/petition/create" />
        <EditPetition path="/petition/:id" />
        <ListPetitionSigner path="/petition-signer" />
        <CreatePetitionSigner path="/petition-signer/create" />
        <EditPetitionSigner path="/petition-signer/:id" />
        <ListPoll path="/poll" />
        <CreatePoll path="/poll/create" />
        <EditPoll path="/poll/:id" />
        <ListPollRespondant path="/poll-respondant" />
        <CreatePollRespondant path="/poll-respondant/create" />
        <EditPollRespondant path="/poll-respondant/:id" />
        <ListPurchaser path="/purchaser" />
        <CreatePurchaser path="/purchaser/create" />
        <EditPurchaser path="/purchaser/:id" />
        <ListCustomerOrder path="/customer-order" />
        <CreateCustomerOrder path="/customer-order/create" />
        <EditCustomerOrder path="/customer-order/:id" />
        <ListDonor path="/donor" />
        <CreateDonor path="/donor/create" />
        <EditDonor path="/donor/:id" />
        <ListLiveEvent path="/live-event" />
        <CreateLiveEvent path="/live-event/create" />
        <EditLiveEvent path="/live-event/:id" />
        <ListEventAttendee path="/event-attendee" />
        <CreateEventAttendee path="/event-attendee/create" />
        <EditEventAttendee path="/event-attendee/:id" />
        <ListVoter path="/voter" />
        <CreateVoter path="/voter/create" />
        <EditVoter path="/voter/:id" />
        <ListVolunteerOpportunity path="/volunteer-opportunity" />
        <CreateVolunteerOpportunity path="/volunteer-opportunity/create" />
        <EditVolunteerOpportunity path="/volunteer-opportunity/:id" />
        <ListVolunteer path="/volunteer" />
        <CreateVolunteer path="/volunteer/create" />
        <EditVolunteer path="/volunteer/:id" />
        <ListActivityType path="/activity-type" />
        <CreateActivityType path="/activity-type/create" />
        <EditActivityType path="/activity-type/:id" />
        <ListFollower path="/follower" />
        <CreateFollower path="/follower/create" />
        <EditFollower path="/follower/:id" />
        <ListTerritory path="/territory" />
        <CreateTerritory path="/territory/create" />
        <EditTerritory path="/territory/:id" />
        <ListActivity path="/activity" />
        <CreateActivity path="/activity/create" />
        <EditActivity path="/activity/:id" />
        <ListNote path="/note" />
        <CreateNote path="/note/create" />
        <EditNote path="/note/:id" />
        <ListAgent path="/agent" />
        <CreateAgent path="/agent/create" />
        <EditAgent path="/agent/:id" />
        <ListAccount path="/account" />
        <CreateAccount path="/account/create" />
        <EditAccount path="/account/:id" />
        <ListOwnerMembership path="/owner-membership" />
        <CreateOwnerMembership path="/owner-membership/create" />
        <EditOwnerMembership path="/owner-membership/:id" />
        <ListAgentMembership path="/agent-membership" />
        <CreateAgentMembership path="/agent-membership/create" />
        <EditAgentMembership path="/agent-membership/:id" />
        <ListContact path="/contact" />
        <CreateContact path="/contact/create" />
        <EditContact path="/contact/:id" />
        <ListContactMembership path="/contact-membership" />
        <CreateContactMembership path="/contact-membership/create" />
        <EditContactMembership path="/contact-membership/:id" />
        <ListHomePage path="/home-page" />
        <CreateHomePage path="/home-page/create" />
        <EditHomePage path="/home-page/:id" />
        <ListPhoto path="/photo" />
        <CreatePhoto path="/photo/create" />
        <EditPhoto path="/photo/:id" />
        <ListCause path="/cause" />
        <CreateCause path="/cause/create" />
        <EditCause path="/cause/:id" />
      </Router>
    </ContentPanel>
  </App>
);

Layout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default Layout;
