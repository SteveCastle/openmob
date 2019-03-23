import React from "react"
import PropTypes from "prop-types"
import { Link } from 'gatsby';

import App from "@openmob/bluebird/src/components/layout/App"
import SideBar from "@openmob/bluebird/src/components/sidebar/SideBar"
import SideBarHeader from "@openmob/bluebird/src/components/layout/SideBarHeader"
import Menu from "@openmob/bluebird/src/components/menu/Menu"
import MenuHeader from "@openmob/bluebird/src/components/menu/MenuHeader"
import Input from "@openmob/bluebird/src/components/forms/Input"
import MenuItem from "@openmob/bluebird/src/components/menu/MenuItem";
import "./layout.css"

const Layout = ({ children, title, id, summary }) => (
  <App>
    <SideBar>
    <SideBarHeader><Link to="/" style={{color:"white", textDecoration:"none"}}>Open Mob</Link></SideBarHeader>
      <Menu vertical>
      <MenuHeader><Input/></MenuHeader>
        <MenuItem title="Causes"/>
        <MenuItem title="Layouts"/>
        </Menu>
    </SideBar>
  </App>
)

Layout.propTypes = {
  children: PropTypes.node.isRequired,
}

export default Layout
