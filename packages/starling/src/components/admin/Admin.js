import React from "react"
import PropTypes from "prop-types"
import { Link } from "gatsby"

import App from "@openmob/bluebird/src/components/layout/App"
import SideBar from "@openmob/bluebird/src/components/sidebar/SideBar"
import SideBarHeader from "@openmob/bluebird/src/components/sidebar/SideBarHeader"
import Menu from "@openmob/bluebird/src/components/menu/Menu"
import MenuHeader from "@openmob/bluebird/src/components/menu/MenuHeader"
import Input from "@openmob/bluebird/src/components/forms/Input"
import MenuItem from "@openmob/bluebird/src/components/menu/MenuItem"
import MenuBody from "@openmob/bluebird/src/components/menu/MenuBody"

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
  </App>
)

Layout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Layout
