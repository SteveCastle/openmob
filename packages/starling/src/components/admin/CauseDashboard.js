import React from 'react'
import PropTypes from 'prop-types'
import { Link } from 'gatsby'
import { Router } from '@reach/router'
import App from '@openmob/bluebird/src/components/layout/App'
import SideBar from '@openmob/bluebird/src/components/layout/sidebar/SideBar'
import ContentPanel from '@openmob/bluebird/src/components/layout/ContentPanel'
import SideBarHeader from '@openmob/bluebird/src/components/layout/sidebar/SideBarHeader'
import Menu from '@openmob/bluebird/src/components/menu/Menu'
import MenuItem from '@openmob/bluebird/src/components/menu/MenuItem'
import PageList from './PageList'
import PageEditor from './PageEditor'

const PlaceHolder = () => <div>Placeholder</div>

const CauseDashboard = ({ causeID }) => (
  <App>
    <SideBar>
      <SideBarHeader>
        <Link to="/">Open Mob</Link>
      </SideBarHeader>
      <Menu vertical>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/pages`}>Pages</Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/outreach`}>Outreach</Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/volunteering`}>
            Volunteer Coordination
          </Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/fundraising`}>Fundraising</Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/events`}>Events</Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/elections`}>Elections</Link>
        </MenuItem>
        <MenuItem>
          <Link to={`/app/cause/${causeID}/boycotts`}>Boycotts</Link>
        </MenuItem>
      </Menu>
    </SideBar>
    <ContentPanel>
      <Router>
        <PageList path="/" />
        <PageEditor path="/homepage/:pageID" />
        <PlaceHolder path="/*" />
      </Router>
    </ContentPanel>
  </App>
)

CauseDashboard.propTypes = {
  children: PropTypes.node.isRequired,
}

export default CauseDashboard
