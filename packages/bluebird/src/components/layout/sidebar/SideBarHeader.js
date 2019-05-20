import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  alignItems,
  justifyContent,
  flexWrap,
} from 'styled-system';

const StyledSideBarHeader = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${alignItems}
  ${justifyContent}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  a {
    text-decoration: none;
    color: white;
  }
`;

function SideBarHeader({ children }) {
  return (
    <StyledSideBarHeader
      color="gray.0"
      display="flex"
      p={2}
      m={2}
      fontSize={4}
      justifyContent="center"
      alignItems="center"
    >
      {children}
    </StyledSideBarHeader>
  );
}

StyledSideBarHeader.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]),
};

export default SideBarHeader;
