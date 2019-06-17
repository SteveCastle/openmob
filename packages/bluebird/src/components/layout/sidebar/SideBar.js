import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap
} from 'styled-system';

const StyledSideBar = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  flex: 1 1 20rem;
  overflow: hidden;
`;
function SideBar({ children }) {
  return <StyledSideBar bg="admin.light.sidebarBg">{children}</StyledSideBar>;
}

SideBar.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ])
};

export default SideBar;
