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
  overflow: hidden;
`;
function SideBar({ children }) {
  return (
    <StyledSideBar width={256} bg="admin.light.sidebarBg" height="5">
      {children}
    </StyledSideBar>
  );
}

SideBar.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element)
};

export default SideBar;
