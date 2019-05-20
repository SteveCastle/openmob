import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  height,
  fontSize,
  color,
  borderBottom,
  borderRadius,
  display,
  flexWrap,
} from 'styled-system';

const StyledMenubar = styled.div`
  ${borderBottom}
  ${borderRadius}
  ${space}
  ${width}
  ${height}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  overflow: hidden;
  & a {
      text-decoration: none;
      color: #717171;
      font-size: 1.2rem;
      font-weight:300;
      padding: 0 20px;
  }
`;
function Menubar({ children }) {
  return (
    <StyledMenubar
      bg="admin.light.menubarBg"
      display="flex"
      p={4}
      borderBottom="#cacaca solid 1px"
    >
      {children}
    </StyledMenubar>
  );
}

Menubar.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};

export default Menubar;
