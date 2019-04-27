import { Link } from 'gatsby'
import PropTypes from 'prop-types'
import React from 'react'
import Navbar from '@openmob/bluebird/src/components/navbar/Navbar'
import Container from '@openmob/bluebird/src/components/layout/Container'

const Header = ({ siteTitle, id, summary }) => (
  <Navbar>
    <Container>
      <Link to="/">{siteTitle}</Link>
      <Link to="/app">Admin</Link>
    </Container>
  </Navbar>
)

Header.propTypes = {
  siteTitle: PropTypes.string,
}

Header.defaultProps = {
  siteTitle: ``,
}

export default Header
